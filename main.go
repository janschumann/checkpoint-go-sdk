package main

import (
	"fmt"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"github.com/janschumann/checkpoint-go-sdk/service/host"
	"github.com/janschumann/checkpoint-go-sdk/service/session"
	"github.com/janschumann/checkpoint-go-sdk/service/task"
)

func main() {

	c := client.New(checkpoint.NewConfig().
		WithCredentials(credentials.NewStaticCredentials("admin", "dkPL.N5zqAJKad7yFoZAN(Pm")).
		WithApiHost("management-wbo7vw5btkaac.westeurope.cloudapp.azure.com").
		WithInsecure(true))

	h := host.New(c)
	//	_, _, err := h.AddHost(&host.AddHostInput{
	//		Name:       "foo3",
	//		Ip4Address: "192.168.20.11",
	//	})
	//
	//if err != nil {
	//	res, _, _ := h.ShowHost(&host.ShowHostInput{
	//		Uid:       res.Uid,
	//	})
	//	fmt.Printf("%v", res)
	//	_, _, _ = h.DeleteHost(&host.ShowHostInput{
	//		Uid:       res.Uid,
	//	})
	//}
	hs, err := h.ShowHosts()
	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		for _, o := range *hs.Objects {
			_, err = h.DeleteHost(&host.ShowHostInput{
				Uid: o.Uid,
			})
			if err != nil {
				fmt.Printf("%s", err.Error())
			}
		}
		fmt.Printf("%v", hs)
	}

	fmt.Printf("%v", hs)

	s := session.New(c)
	sr, err := s.ShowSessions()
	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		for _, o := range *sr.Objects {
			_, err = s.Disconnect(&session.DisconnectInput{
				Uid:     o.Uid,
				Discard: true,
			})
			if err != nil {
				fmt.Printf("%s", err.Error())
			}
		}
		fmt.Printf("%v", sr)
	}

	pr, err := s.Publish()

	fmt.Printf("%v ... %v ... %v", pr, err)

	t := task.New(c)
	tr, err := t.ShowTasks()

	fmt.Printf("%v ... %v ... %v", tr, err)

}
