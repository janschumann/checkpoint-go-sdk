# Checkpoint SDK for Go

checkpoint-go-sdk is the [Checkpoint](https://www.checkpoint.com/de/products/next-generation-firewall/) 
SDK for the Go programming language. 

API Documentation: https://sc1.checkpoint.com/documents/latest/APIs/#introduction~v1.5

The core parts of this SDK are borrowed from the [AWS SDK for Go](https://github.com/aws/aws-sdk-go),
because I like their approach to handle API Errors, retry logic, pagination, configuration etc
and I did not want to reinvent the wheel.

This is in a very early development stage. There are still some core decisions to make. E.g. I am 
currently thinking about generating the API components via the 
[checkpoint API definition](https://sc1.checkpoint.com/documents/latest/APIs/data/v1.5/dynamic/apis.json)

## Getting started

The best way to get started working with the SDK is to use `go get` to add the SDK to your Go Workspace

```sh
go get github.com/janschumann/checkpoint-go-sdk
```

### Hello Host

This example shows how you can use the SDK to make an API request to create a host:

```go
package main

import (
	"fmt"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"github.com/janschumann/checkpoint-go-sdk/service/host"
)

func main() {
	// configure the client
	c := client.Must(client.New(checkpoint.NewConfig().
		WithCredentials(credentials.NewStaticCredentials("admin", "<pw>")).
		WithApiHost("<api-host>")))

    // create host client
    h := host.New(c)

    // add a host
    res, _ := h.AddHost(&host.AddHostInput{
        Name:      "host1",
        IpAddress: "192.168.10.1",
    })
    fmt.Println(fmt.Sprintf("Add Host Response: %v", res))

    // publish the change
    pr, _ := c.Publish()
    fmt.Println(fmt.Sprintf("Publish Response: %s", pr.TaskId))
    
    // logout
    _, _ = c.Logout()
}
```

## License

TBD
