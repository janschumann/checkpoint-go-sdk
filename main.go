package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {

	sess := httpsession.Must(httpsession.NewSession(&checkpoint.Config{
		Credentials: credentials.NewStaticCredentials("admin", "dkPL.N5zqAJKad7yFoZAN(Pm", ""),
		DisableSSL: aws.Bool(true),
	}))
	sess.Auth()

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
