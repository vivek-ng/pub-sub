package main

import (
	"fmt"
	"time"

	"github.com/vivek-ng/pub-sub/pubsub"
)

func main() {
	pb := pubsub.NewPubSub(5)
	ch1 := pb.Subscribe("test")
	go func() {
		//for {
		pb.Publish("test", "vivek is awesome!")
		//}
	}()

	go func() {
		for {
			select {
			case d := <-ch1:
				go func(data interface{}) {
					fmt.Println(data)
				}(d)
			}
		}
	}()
	time.Sleep(10 * time.Second)
}
