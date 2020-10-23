package api

import (
	"net/http"
	"strconv"
	"time"
)

type queue struct {
	WaitList      []chan bool
	NextBlock     *time.Timer
	NextChan      chan bool
	StopChan      chan bool
	Running       int
	SecondRRemain int
}

func newQueue() *queue {
	queue := &queue{
		WaitList:  make([]chan bool, 0),
		NextBlock: nil,
		NextChan:  make(chan bool),
		StopChan:  make(chan bool),
		Running:   0,
	}
	go queue.manage()
	return queue
}

func (q *queue) manage() {
	for {
		select {
		case <-q.NextChan:
			q.runNext()
		case <-q.StopChan:
			break
		}
	}
}

func (q *queue) setLimits(header http.Header) {
	if header["X-Secondly-RateLimit-Remaining"] == nil {
		return
	}
	n, _ := strconv.Atoi(header["X-Secondly-RateLimit-Remaining"][0])
	q.SecondRRemain = n
	if n == 0 {
		q.NextBlock = time.NewTimer(time.Second * 2)
	}
}

func (q *queue) waitTurn() {
	if q.NextBlock != nil {
		<-q.NextBlock.C
		q.NextBlock = nil
	}
	if q.Running < 2 {
		q.Running++
		return
	}
	newChan := make(chan bool)
	q.WaitList = append(q.WaitList, newChan)
	<-newChan
}

func (q *queue) next() {
	q.NextChan <- true
}

func (q *queue) stop() {
	q.StopChan <- true
}

func (q *queue) runNext() {
	if len(q.WaitList) < 1 {
		if q.Running > 0 {
			q.Running--
		}
		return
	}
	toClose := q.WaitList[0]
	q.drop(0)
	toClose <- true
}

func (q *queue) drop(i int) {
	q.WaitList = append(q.WaitList[:i], q.WaitList[i+1:]...)
}
