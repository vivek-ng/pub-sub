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
	ch := s.pubsub.Subscribe("topic1")
	count := 0
	for val := range ch {
		count += 1
		fmt.Println(val)
	}
	fmt.Println(count, "vivek")
}
