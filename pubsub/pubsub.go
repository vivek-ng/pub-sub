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
		once        sync.Once
		closed      bool
	}
)

func NewPubSub(capacity int) *PubSub {
	return &PubSub{
		capacity:    capacity,
		subscribers: map[string][]DataChan{},
	}
}

func (pb *PubSub) Subscribe(topic string) DataChan {
	pb.mux.Lock()
	defer pb.mux.Unlock()
	ch := make(DataChan, pb.capacity)
	tp, ok := pb.subscribers[topic]
	if ok {
		pb.subscribers[topic] = append(tp, ch)
	} else {
		pb.subscribers[topic] = append([]DataChan{}, ch)
	}
	return ch
}

func (pb *PubSub) Publish(topic string, message interface{}) {
	pb.mux.RLock()
	defer pb.mux.RUnlock()
	if pb.closed {
		return
	}
	dt := Data{
		Message: message,
	}
	for _, sub := range pb.subscribers[topic] {
		go func(data Data, sub DataChan) {
			sub <- dt
		}(dt, sub)
	}
}

func (pb *PubSub) Close() {
	pb.mux.Lock()
	defer pb.mux.Unlock()
	pb.once.Do(func() {
		pb.closed = true
		for _, v := range pb.subscribers {
			for _, ch := range v {
				close(ch)
			}
		}
	})
}
