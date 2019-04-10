package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
)

func main() {

	client := client.New(checkpoint.NewConfig().
		WithCredentials(credentials.NewStaticCredentials("admin", "dkPL.N5zqAJKad7yFoZAN(Pm", "")).
		WithApiHost("management-wbo7vw5btkaac.westeurope.cloudapp.azure.com").
		WithDisableSSL(true))

	h := host.New(sess)
	_, _, err := h.AddHost(&host.AddHostInput{
		Name: "foo",
		Ip4Address: "192.168.20.10",
	})

	sess.Auth()

	h = host.New(sess)
	_, _, err = h.AddHost(&host.AddHostInput{
		Name: "bar",
		Ip4Address: "192.168.20.11",
	})

	sess.Auth()

	svc := session.New(sess)
	pr, pe, err := svc.Publish()

	fmt.Sprintf("%v ... %v ... %v", pr, pe, err)
}
