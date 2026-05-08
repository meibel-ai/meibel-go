package v2

import (
	"context"
)

// PageIterator provides iteration over paginated API results.
// It implements a pull-based iterator pattern that fetches pages on demand.
type PageIterator[T any] struct {
	// fetch is the function that fetches the next page of results.
	fetch func(ctx context.Context, cursor string) (*Page[T], error)
	// current holds the items from the current page.
	current []T
	// index is the current position within the current page.
	index int
	// cursor is the cursor for fetching the next page.
	cursor string
	// done indicates whether all pages have been consumed.
	done bool
	// err holds any error that occurred during fetching.
	err error
}

// Page represents a single page of results.
type Page[T any] struct {
	// Items contains the items in this page.
	Items []T
	// NextCursor is the cursor for the next page, or empty if this is the last page.
	NextCursor string
}

// NewPageIterator creates a new PageIterator with the given fetch function.
func NewPageIterator[T any](fetch func(ctx context.Context, cursor string) (*Page[T], error)) *PageIterator[T] {
	return &PageIterator[T]{
		fetch: fetch,
	}
}

// Next advances the iterator to the next item.
// It returns true if there is a next item, false otherwise.
// Call Item() to get the current item after Next() returns true.
func (p *PageIterator[T]) Next(ctx context.Context) bool {
	if p.err != nil || p.done {
		return false
	}

	// Try to advance within the current page
	p.index++
	if p.index < len(p.current) {
		return true
	}

	// Need to fetch the next page
	if p.cursor == "" && p.current != nil {
		// We've exhausted all pages
		p.done = true
		return false
	}

	page, err := p.fetch(ctx, p.cursor)
	if err != nil {
		p.err = err
		return false
	}

	p.current = page.Items
	p.cursor = page.NextCursor
	p.index = 0

	if len(p.current) == 0 {
		p.done = true
		return false
	}

	return true
}

// Item returns the current item.
// Call this after Next() returns true.
func (p *PageIterator[T]) Item() T {
	return p.current[p.index]
}

// Err returns any error that occurred during iteration.
func (p *PageIterator[T]) Err() error {
	return p.err
}

// Collect fetches all remaining items into a slice.
// Use with caution on large datasets.
func (p *PageIterator[T]) Collect(ctx context.Context) ([]T, error) {
	var items []T
	for p.Next(ctx) {
		items = append(items, p.Item())
	}
	if err := p.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// Take fetches up to n items from the iterator.
func (p *PageIterator[T]) Take(ctx context.Context, n int) ([]T, error) {
	var items []T
	for p.Next(ctx) && len(items) < n {
		items = append(items, p.Item())
	}
	if err := p.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// ForEach calls fn for each item in the iterator.
// If fn returns an error, iteration stops and the error is returned.
func (p *PageIterator[T]) ForEach(ctx context.Context, fn func(item T) error) error {
	for p.Next(ctx) {
		if err := fn(p.Item()); err != nil {
			return err
		}
	}
	return p.Err()
}

// OffsetPageIterator provides iteration over offset-based paginated API results.
type OffsetPageIterator[T any] struct {
	// fetch is the function that fetches a page of results at the given offset.
	fetch func(ctx context.Context, offset, limit int) (*OffsetPage[T], error)
	// limit is the page size.
	limit int
	// offset is the current offset.
	offset int
	// current holds the items from the current page.
	current []T
	// index is the current position within the current page.
	index int
	// total is the total number of items (if known).
	total int
	// totalKnown indicates whether total is known.
	totalKnown bool
	// done indicates whether all pages have been consumed.
	done bool
	// err holds any error that occurred during fetching.
	err error
}

// OffsetPage represents a single page of offset-based results.
type OffsetPage[T any] struct {
	// Items contains the items in this page.
	Items []T
	// Total is the total number of items (optional).
	Total *int
}

// NewOffsetPageIterator creates a new OffsetPageIterator with the given fetch function and limit.
func NewOffsetPageIterator[T any](limit int, fetch func(ctx context.Context, offset, limit int) (*OffsetPage[T], error)) *OffsetPageIterator[T] {
	return &OffsetPageIterator[T]{
		fetch: fetch,
		limit: limit,
	}
}

// Next advances the iterator to the next item.
func (p *OffsetPageIterator[T]) Next(ctx context.Context) bool {
	if p.err != nil || p.done {
		return false
	}

	// Try to advance within the current page
	p.index++
	if p.index < len(p.current) {
		return true
	}

	// Check if we've reached the known total
	if p.totalKnown && p.offset >= p.total {
		p.done = true
		return false
	}

	// Need to fetch the next page
	page, err := p.fetch(ctx, p.offset, p.limit)
	if err != nil {
		p.err = err
		return false
	}

	p.current = page.Items
	p.index = 0
	p.offset += len(page.Items)

	if page.Total != nil {
		p.total = *page.Total
		p.totalKnown = true
	}

	if len(p.current) == 0 || len(p.current) < p.limit {
		p.done = true
	}

	if len(p.current) == 0 {
		return false
	}

	return true
}

// Item returns the current item.
func (p *OffsetPageIterator[T]) Item() T {
	return p.current[p.index]
}

// Err returns any error that occurred during iteration.
func (p *OffsetPageIterator[T]) Err() error {
	return p.err
}

// Collect fetches all remaining items into a slice.
func (p *OffsetPageIterator[T]) Collect(ctx context.Context) ([]T, error) {
	var items []T
	for p.Next(ctx) {
		items = append(items, p.Item())
	}
	if err := p.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// Take fetches up to n items from the iterator.
func (p *OffsetPageIterator[T]) Take(ctx context.Context, n int) ([]T, error) {
	var items []T
	for p.Next(ctx) && len(items) < n {
		items = append(items, p.Item())
	}
	if err := p.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
