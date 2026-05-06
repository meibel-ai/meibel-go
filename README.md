# Meibel Go SDK

The official Go SDK for the [Meibel API](https://docs.meibel.ai). Provides document parsing, datasource management, and AI agent orchestration.

## Installation

Install from Git (v2):

```bash
go get github.com/meibel-ai/meibel-go/v2@v2.0.0
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"os"

	meibel "github.com/meibel-ai/meibel-go/v2"
)

func main() {
	client := meibel.NewMeibelClient(
		meibel.WithAPIKey("your-api-key"),
	)

	// Parse a document
	file, _ := os.Open("document.pdf")
	defer file.Close()

	job, err := client.Documents.ParseDocument(context.Background(), file, "document.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(job.JobID)

	// List datasources
	datasources, err := client.Datasources.ListDatasources(context.Background())
	if err != nil {
		panic(err)
	}
	for _, ds := range datasources.Items {
		fmt.Println(ds.Name)
	}
}
```

## Nested Resources

Resources are organized hierarchically. Content, downloads, data elements, and table descriptions are accessed through `Datasources`:

```go
// Upload content to a datasource
result, err := client.Datasources.Content.UploadContent(ctx, file, "data.csv", nil)

// List data elements
elements, err := client.Datasources.DataElements.ListDataElements(ctx, "ds-123", nil)
```

Agent sessions (chat) are accessed through `Agents`:

```go
session, err := client.Agents.Sessions.CreateSession(ctx, &meibel.CreateSessionRequest{
    AgentID: "agent-123",
})
response, err := client.Agents.Sessions.SendChatMessage(ctx, &meibel.SendChatMessageRequest{
    SessionID: session.ID,
    Message:   "Hello",
})
```

## Documentation

- [API Reference](https://docs.meibel.ai/api-reference/overview)
- [SDK Guide](https://docs.meibel.ai/sdk/go)

## License

MIT
