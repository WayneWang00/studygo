package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = (1 << workerBits) - 1
	numberMax   int64 = (1 << numberBits) - 1
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000
)

type Worker struct {
	mu        sync.Mutex
	timeStamp int64
	workerId  int64
	numberId  int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker id failed")
	}

	return &Worker{
		timeStamp: 0,
		workerId:  workerId,
		numberId:  0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6

	if now == w.timeStamp {
		w.numberId++
		if w.numberId > numberMax {
			for now <= w.timeStamp {
				now = time.Now().UnixNano() / 1e6
			}
			w.workerId = 0
		}
	} else {
		w.timeStamp = now
		w.numberId = 0
	}

	n := int64((now-startTime)<<timeShift | w.workerId<<workerShift | w.numberId)
	return n
}

func main() {
	worker, err := NewWorker(1)
	if err != nil {
		fmt.Println("new worker failed:", err)
		return
	}
	fmt.Printf("new worker:%+v\n", worker)

	id := worker.GetId()
	fmt.Println("get id:", id)
}
