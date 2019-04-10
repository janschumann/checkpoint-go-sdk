package main

import (
	"fmt"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"github.com/janschumann/checkpoint-go-sdk/service/host"
	"github.com/janschumann/checkpoint-go-sdk/service/session"
)

func main() {

	c := client.New(checkpoint.NewConfig().
		WithCredentials(credentials.NewStaticCredentials("admin", "dkPL.N5zqAJKad7yFoZAN(Pm")).
		WithApiHost("management-wbo7vw5btkaac.westeurope.cloudapp.azure.com").
		WithDisableSSL(true))

	h := host.New(c)
	_, _, err := h.AddHost(&host.AddHostInput{
		Name:       "foo",
		Ip4Address: "192.168.20.10",
	})

	s := session.New(c)
	pr, pe, err := s.Publish()

	fmt.Printf("%v ... %v ... %v", pr, pe, err)
}
