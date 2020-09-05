package publisher

import "github.com/vivek-ng/pub-sub/pubsub"

type Publisher struct {
	pubsub *pubsub.PubSub
}

func NewPublisher(pb *pubsub.PubSub) *Publisher {
	return &Publisher{
		pubsub: pb,
	}
}

func (p *Publisher) Publish() {
	for i := 0; i < 500; i++ {
		p.pubsub.Publish("topic1", "i love pizza")
	}
}
