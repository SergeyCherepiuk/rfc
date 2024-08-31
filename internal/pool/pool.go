package pool

import (
	"context"
	"sync"
)

type Pool[In, Out any] struct {
	callback func(In) (Out, error)
	in       chan In
	out      chan Out
	err      chan error
	wg       sync.WaitGroup
}

func NewPool[In, Out any](ctx context.Context, workers int, callback func(In) (Out, error)) *Pool[In, Out] {
	pool := Pool[In, Out]{
		callback: callback,
		in:       make(chan In),
		out:      make(chan Out),
		err:      make(chan error),
	}

	pool.wg.Add(workers)
	for range workers {
		go pool.worker(ctx)
	}

	return &pool
}

func (p *Pool[In, Out]) worker(ctx context.Context) {
	defer p.wg.Done()

	for {
		select {
		case in, ok := <-p.in:
			if !ok {
				return
			}

			result, err := p.callback(in)
			if err != nil {
				select {
				case p.err <- err:
				case <-ctx.Done():
					return
				}
			} else {
				select {
				case p.out <- result:
				case <-ctx.Done():
					return
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pool[In, Out]) In() chan<- In {
	return p.in
}

func (p *Pool[In, Out]) Out() <-chan Out {
	return p.out
}

func (p *Pool[In, Out]) Err() <-chan error {
	return p.err
}

func (p *Pool[In, Out]) Close() {
	close(p.in)
	p.wg.Wait()
	close(p.out)
	close(p.err)
}
