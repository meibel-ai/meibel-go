# Meibel Go SDK

The official Go SDK for the [Meibel API](https://docs.meibel.ai). Provides document parsing, datasource management, and AI agent orchestration.

## Installation

```bash
go get github.com/meibel-ai/meibel-go@v0.1.0-beta.1
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"os"

	meibel "github.com/meibel-ai/meibel-go"
)

func main() {
	client := meibel.NewMeibelClient(
		meibel.WithAPIKey("your-api-key"),
		meibel.WithBaseURL("https://api.meibel.ai/v2"),
	)

	// Parse a document
	file, _ := os.Open("document.pdf")
	defer file.Close()

	job, err := client.Documents.ParseDocument(context.Background(), file)
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

## Documentation

- [API Reference](https://docs.meibel.ai/api-reference/overview)
- [SDK Guide](https://docs.meibel.ai/sdk/go)

## License

MIT