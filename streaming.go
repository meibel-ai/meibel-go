package meibelgo

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// SSEEvent represents a Server-Sent Event.
type SSEEvent struct {
	// Event is the event type (from "event:" field).
	Event string
	// Data is the event data (from "data:" field).
	Data string
	// ID is the event ID (from "id:" field).
	ID string
	// Retry is the retry interval in milliseconds (from "retry:" field).
	Retry int
}

// EventStream provides streaming access to Server-Sent Events.
// Use Events() and Errors() to receive parsed events and errors.
type EventStream[T any] struct {
	eventsCh <-chan T
	errorsCh <-chan error
	events   chan T
	errors   chan error
	response *http.Response
	cancel   context.CancelFunc
	parser   func(data string) (T, error)
}

// Events returns the channel for receiving parsed events.
func (s *EventStream[T]) Events() <-chan T { return s.eventsCh }

// Errors returns the channel for receiving stream errors.
func (s *EventStream[T]) Errors() <-chan error { return s.errorsCh }

// NewEventStream creates a new EventStream from an HTTP response.
// The parser function is used to parse the event data into the target type.
func NewEventStream[T any](resp *http.Response, parser func(data string) (T, error)) *EventStream[T] {
	events := make(chan T, 10)
	errors := make(chan error, 1)

	ctx, cancel := context.WithCancel(context.Background())

	stream := &EventStream[T]{
		eventsCh: events,
		errorsCh: errors,
		events:   events,
		errors:   errors,
		response: resp,
		cancel:   cancel,
		parser:   parser,
	}

	go stream.run(ctx)

	return stream
}

// Close closes the event stream and releases resources.
func (s *EventStream[T]) Close() error {
	s.cancel()
	if s.response != nil && s.response.Body != nil {
		return s.response.Body.Close()
	}
	return nil
}

func (s *EventStream[T]) run(ctx context.Context) {
	defer close(s.events)
	defer close(s.errors)
	defer s.response.Body.Close()

	reader := bufio.NewReader(s.response.Body)
	var eventBuilder strings.Builder
	var currentEvent SSEEvent

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				select {
				case s.errors <- err:
				case <-ctx.Done():
				}
			}
			return
		}

		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")

		// Empty line signals end of event
		if line == "" {
			if eventBuilder.Len() > 0 || currentEvent.Event != "" {
				currentEvent.Data = eventBuilder.String()
				eventBuilder.Reset()

				if currentEvent.Data != "" {
					parsed, err := s.parser(currentEvent.Data)
					if err != nil {
						select {
						case s.errors <- fmt.Errorf("failed to parse event: %w", err):
						case <-ctx.Done():
							return
						}
					} else {
						select {
						case s.events <- parsed:
						case <-ctx.Done():
							return
						}
					}
				}
				currentEvent = SSEEvent{}
			}
			continue
		}

		// Comment line
		if strings.HasPrefix(line, ":") {
			continue
		}

		// Parse field
		colonIdx := strings.Index(line, ":")
		if colonIdx == -1 {
			continue
		}

		field := line[:colonIdx]
		value := line[colonIdx+1:]
		if strings.HasPrefix(value, " ") {
			value = value[1:]
		}

		switch field {
		case "event":
			currentEvent.Event = value
		case "data":
			if eventBuilder.Len() > 0 {
				eventBuilder.WriteString("\n")
			}
			eventBuilder.WriteString(value)
		case "id":
			currentEvent.ID = value
		case "retry":
			if n, err := strconv.Atoi(value); err == nil {
				currentEvent.Retry = n
			}
		}
	}
}

// JSONEventStream creates an EventStream that parses JSON events.
func JSONEventStream[T any](resp *http.Response) *EventStream[T] {
	return NewEventStream[T](resp, func(data string) (T, error) {
		var result T
		err := json.Unmarshal([]byte(data), &result)
		return result, err
	})
}

// Range allows iterating over events using a range loop (Go 1.23+).
// Note: This is a helper for manual iteration; for production use,
// prefer reading directly from the Events channel.
func (s *EventStream[T]) Range(ctx context.Context) <-chan T {
	return s.eventsCh
}

// Recv receives the next event from the stream.
// It returns the event, a boolean indicating if more events are available,
// and any error that occurred.
func (s *EventStream[T]) Recv(ctx context.Context) (T, bool, error) {
	var zero T
	select {
	case <-ctx.Done():
		return zero, false, ctx.Err()
	case err := <-s.errorsCh:
		return zero, false, err
	case event, ok := <-s.eventsCh:
		return event, ok, nil
	}
}

// Collect reads all events from the stream into a slice.
// Use with caution on long-running streams.
func (s *EventStream[T]) Collect(ctx context.Context) ([]T, error) {
	var events []T
	for {
		select {
		case <-ctx.Done():
			return events, ctx.Err()
		case err := <-s.errorsCh:
			return events, err
		case event, ok := <-s.eventsCh:
			if !ok {
				return events, nil
			}
			events = append(events, event)
		}
	}
}
