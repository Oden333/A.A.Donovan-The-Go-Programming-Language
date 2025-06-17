package memo

import "sync"

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready

	//* Each entry contains the memoized result of a call to the
	//* function f, as before, but it additionally contains a channel called ready. Just after
	//* the entry’s result has been set, this channel will be closed, to broadcast
	//* to any other goroutines that it is now safe for them to read the result from the entry.
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New4(f Func) *Memo4 {
	return &Memo4{f: f, cache: make(map[string]*entry)}
}

type Memo4 struct {
	f     Func
	mu    sync.Mutex //* guards cache map
	cache map[string]*entry
}

func (memo *Memo4) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		<-e.ready // wait for ready ( operation blocks until the channel is closed )
	}
	return e.res.value, e.res.err
	//? These variables are shared among multiple goroutines because
	//? closing of the ready channel happens before any other goroutine receives the broadcast event
}
