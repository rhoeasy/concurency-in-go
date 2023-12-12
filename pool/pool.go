package pool

import (
	"context"
	"hash/crc32"
	"log"
)

type (
	Pool struct {
		works []chan string
	}
)

func NewPool(size int, queueLen int) *Pool {
	p := &Pool{
		works: make([]chan string, size),
	}
	for i := 0; i < size; i++ {
		p.works[i] = make(chan string, queueLen)
	}
	return p
}

func (p Pool) poolSize() int {
	return len(p.works)
}

func (p Pool) Run(ctxt context.Context) {
	for i, work := range p.works {
		go func(ctx context.Context, i int, w chan string) {
			for {
				select {
				case <-ctx.Done():
					log.Printf("worker[%v] is destroyed.", i)
					return
				case anything := <-w:
					log.Printf("worker[%v] is working with task[%v]", i, anything)
				}
			}

		}(ctxt, i, work)
	}

}

func (p *Pool) SubmitTask(task string) {
	crc := crc32.ChecksumIEEE([]byte(task))
	mod := int(crc) % p.poolSize()
	log.Printf("worker[%v] will received task[%v:%v]", mod, task, crc)
	p.works[mod] <- task
}
