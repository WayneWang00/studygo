package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type process struct {
	rw   sync.RWMutex
	mu   sync.Mutex
	pass string
}

var pro = process{pass: "Pass"}

func Change(p *process, pass string) {
	p.rw.Lock()
	defer p.rw.Unlock()

	fmt.Println("change:", pass, time.Now().Second())
	p.pass = pass
	time.Sleep(3 * time.Second)
}

func RwMutexShow(p *process) string {
	p.rw.RLock()
	defer p.rw.RUnlock()

	fmt.Println("rwShow:", p.pass, time.Now().Second())
	time.Sleep(1 * time.Second)

	return p.pass
}

func MutexShow(p *process) string {
	p.mu.Lock()
	defer p.mu.Unlock()

	fmt.Println("muShow:", p.pass, time.Now().Second())
	time.Sleep(1 * time.Second)

	return p.pass
}

func main() {
	var f = func(*process) string { return "" }
	var wg sync.WaitGroup

	if len(os.Args) < 2 {
		fmt.Println("rw mutex", time.Now().Second())
		f = RwMutexShow
	} else {
		fmt.Println("mutex", time.Now().Second())
		f = MutexShow
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("show:", f(&pro), time.Now().Second())
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		Change(&pro, "12345")
	}()

	wg.Wait()
}
