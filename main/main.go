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
		for val := range ch1 {
			fmt.Println(val)
		}
		fmt.Println("channel closed !!")
	}()
	time.Sleep(3 * time.Second)
	pb.Close()
	time.Sleep(10 * time.Second)

}
