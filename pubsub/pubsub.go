package pubsub

import "sync"

type (
	Data struct {
		Message interface{}
	}

	DataChan chan Data

	PubSub struct {
		capacity    int
		subscribers map[string][]DataChan
		mux         sync.RWMutex
	}
)

func NewPubSub(capacity int) *PubSub {
	return &PubSub{
		capacity: capacity,
	}
}

func (pb *PubSub) Subscribe(topic string) DataChan {
	pb.mux.Lock()
	defer pb.mux.Unlock()
	ch := make(DataChan)
	_, ok := pb.subscribers[topic]
	if ok {
		pb.subscribers[topic] = append(pb.subscribers[topic], ch)
	} else {
		pb.subscribers[topic] = append([]DataChan{}, ch)
	}
	return ch
}

func (pb *PubSub) Publish(topic string, message interface{}) {
	pb.mux.RLock()
	defer pb.mux.RUnlock()
	dt := Data{
		Message: message,
	}

	for _, subs := range pb.subscribers[topic] {
		subs <- dt
	}
}
