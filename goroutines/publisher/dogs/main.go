package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type (
	subscriber chan interface{}         // Subscriber is a pipeline
	topicFunc  func(v interface{}) bool // Subject is a filter
)

// Publisher object
type Publisher struct {
	m           sync.RWMutex             // read-write lock
	buffer      int                      // The cache size of the subscription queue
	timeout     time.Duration            // Release timeout
	subscribers map[subscriber]topicFunc // All subscriber information
}

// Build a publisher object, you can set the publishing timeout and the length of the cache queue
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Add a new subscriber, subscribe to the filtered topics
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// Subscribe to all topics, no filtering
func (p *Publisher) Subscriber() chan interface{} {
	return p.SubscribeTopic(nil)
}

// Exit subscription
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// Send the subject, can tolerate a certain timeout
func (p *Publisher) sendTopic(
	sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	// did not subscribe to the topic
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

// Post a topic
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// Close the publisher object and close all subscriber pipelines
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func main() {
	// instantiate the publisher object
	pub := NewPublisher(100*time.Millisecond, 10)
	// Close the publisher object and close all subscriber pipelines
	defer pub.Close()
	// Add a new subscriber, subscribe to all topics, no filtering
	all := pub.Subscriber()
	// Add a new subscriber, subscribe to the filtered topics
	golang := pub.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	// Post a topic
	pub.Publish("hello, world!")
	// Post a topic
	pub.Publish("hello, golang!")
	go func() {
		// Accept the subject of the corresponding subscription
		for msg := range all {
			fmt.Println("all: ", msg)
		}
	}()
	go func() {
		// Accept the subject of the corresponding subscription
		for msg := range golang {
			fmt.Println("golang: ", msg)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

