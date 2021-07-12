package worker

import (
	"fmt"
)

type Query struct {
	k string
	v string
}

var QueryQueue chan Query

type Worker struct {
	WP chan chan Query
	QChan chan Query
	quit chan bool
}

func NewWorker(wp chan chan Query) Worker {
	return Worker{
		WP: wp,
		QChan: make(chan Query),
		quit: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WP <- w.QChan

			select {
			case job := <- w.QChan:
				// lookup the key
				fmt.Printf("%+v\n", job)
				continue
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}