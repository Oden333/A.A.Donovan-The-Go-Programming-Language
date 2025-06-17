package memo

import (
	"errors"
	"sync"
)

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// Func is the type of the function to memoize.
type Func struct {
	fn   func(key string) (interface{}, error)
	done chan struct{}
}

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex //* guards cache map
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		if memo.f.done != nil {
			res := make(chan result)
			go func() {
				r, err := memo.f.fn(key)
				res <- result{r, err}
			}()
			select {
			case <-memo.f.done:
				e.res.value, e.res.err = nil, errors.New("Done signal received")
			case r := <-res:
				e.res.value, e.res.err = r.value, r.err
			}
		} else {
			e.res.value, e.res.err = memo.f.fn(key)
		}

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		<-e.ready // wait for ready ( operation blocks until the channel is closed )
	}
	return e.res.value, e.res.err
}
