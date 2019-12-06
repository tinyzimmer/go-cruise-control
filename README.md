# go-cruise-control

A go REST-client library for LinkedIn's Cruise Control.

[![](https://godoc.org/github.com/tinyzimmer/go-cruise-control?status.svg)](http://godoc.org/github.com/tinyzimmer/go-cruise-control)


Will improve docs in a little bit

## Basic Usage

```go
package main

import (
	"fmt"

	gocc "github.com/tinyzimmer/go-cruise-control"
	"github.com/tinyzimmer/go-cruise-control/pkg/types"
)

func main() {
	client, err := gocc.New(&types.ClientOptions{
		URL: "http://localhost:8090",
	})
	if err != nil {
		panic(err)
	}
	res, err := client.GetState(types.GetStateDefaults())
	if err != nil {
		panic(err)
	}
	fmt.Println(res.MonitorState.State)
}

```
