package pool

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type workableUnit func(args ...interface{})

type work struct {
	f         workableUnit
	createdAt time.Time
	args      []interface{}
}

var p *pool

type pool struct {
	procs         uint8
	maxGoRoutines uint32
	queueSize     uint32
	workChan      chan work
	ttl           time.Duration
	wg            sync.WaitGroup
}

func GetPool(cpu uint8, goroutines uint32, queueSize uint32, expireTime int) *pool {
	if p == nil {
		p = &pool{
			procs:         cpu,
			maxGoRoutines: goroutines,
			queueSize:     queueSize,
			ttl:           time.Duration(expireTime) * time.Second,
		}
	}
	return p
}

func (p *pool) enableMaxProcs() {
	runtime.GOMAXPROCS(int(p.procs))
}

func (p *pool) enableQueue() {
	p.workChan = make(chan work, p.queueSize)
}

func worker(id int, jobs <-chan work, p *pool) {
	for w := range jobs {
		if time.Since(w.createdAt) > p.ttl {
			continue
		}
		w.f(w.args...)
	}
	p.wg.Done()
}

func (p *pool) enableMaxGoRoutines() {
	for i := 0; i < int(p.maxGoRoutines); i++ {
		p.wg.Add(1)
		go worker(i, p.workChan, p)
	}
}

func (p *pool) Stop() {
	close(p.workChan)
	p.wg.Wait()
}

func (p *pool) Start() {
	p.enableMaxProcs()
	p.enableQueue()
	p.enableMaxGoRoutines()
}

func (p *pool) Submit(f workableUnit, args ...interface{}) error {
	if p.workChan == nil {
		return fmt.Errorf("channel is not initialized")
	}
	p.workChan <- work{f, time.Now(), args}
	return nil
}
