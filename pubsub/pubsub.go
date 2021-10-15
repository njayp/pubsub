package pubsub


import "sync"


type PubSub struct {
	mu   sync.Mutex
	subs[T any] map[string][]chan T
}

func NewPubsub() *PubSub {
	ps := &PubSub{}
	ps.subs = make(map[string][]chan string)
	return ps
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic string, msg string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, ch := range ps.subs[topic] {
		go func(ch chan string) {
			ch <- msg
		}(ch)
	}
}
