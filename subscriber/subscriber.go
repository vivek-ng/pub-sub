package subscriber

import (
	"fmt"

	"github.com/vivek-ng/pub-sub/pubsub"
)

type Subscriber struct {
	pubsub *pubsub.PubSub
}

func NewSubscriber(pb *pubsub.PubSub) *Subscriber {
	return &Subscriber{
		pubsub: pb,
	}
}

func (s *Subscriber) Subscribe() {
	// for i := 0; i < 500; i++ {
	// 	s.pubsub.Publish("topic1", "i love pizza")
	// }
	ch := s.pubsub.Subscribe("topic1")
	count := 0
	for val := range ch {
		count += 1
		fmt.Println(val)
	}
	fmt.Println(count, "vivek")
}
