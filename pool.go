package overclock

import (
	"sync"
)

type Executor interface {
	Start() error
	Execute(func() error) error
	Close() error
}

type executor struct {
	poolSize     int
	queueSize    int
	works        sync.WaitGroup
	queue        chan func() error
	panicHandler func(err error)
}

func NewExecutor(poolSize int, queueSize int, panicHandler func(err error)) Executor {
	return &executor{
		poolSize:     poolSize,
		queueSize:    queueSize,
		works:        sync.WaitGroup{},
		queue:        make(chan func() error, queueSize),
		panicHandler: panicHandler,
	}
}

func (e *executor) Start() error {
	for i := 0; i < e.poolSize; i++ {
		go func() {
			for work := range e.queue {
				//e.works.Add(1)
				if err := work(); err != nil {
					e.panicHandler(err)
				}
				e.works.Done()
			}
		}()
	}
	return nil
}

func (e *executor) Execute(f func() error) error {
	e.works.Add(1)
	e.queue <- f
	return nil
}

func (e *executor) Close() error {
	close(e.queue)
	e.works.Wait()
	return nil
}
