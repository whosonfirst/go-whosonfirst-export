package engine

import (
	"errors"
	"github.com/aaronland/go-artisanal-integers"
	"sync"
)

type MemoryEngine struct {
	artisanalinteger.Engine
	key       string
	increment int64
	offset    int64
	mu        *sync.Mutex
	last      int64
}

func NewMemoryEngine(dsn string) (*MemoryEngine, error) {

	mu := new(sync.Mutex)

	eng := MemoryEngine{
		key:       "integers",
		increment: 2,
		offset:    1,
		mu:        mu,
		last:      0,
	}

	// PLEASE WRITE ME: check to see if we should read a value persisted to disk

	return &eng, nil
}

func (eng *MemoryEngine) SetLastInt(i int64) error {

	eng.mu.Lock()
	defer eng.mu.Unlock()

	last, err := eng.LastInt()

	if err != nil {
		return err
	}

	if last > i {
		return errors.New("integer is too small")
	}

	eng.last = i
	return nil
}

func (eng *MemoryEngine) SetKey(k string) error {
	return nil
}

func (eng *MemoryEngine) SetOffset(i int64) error {

	eng.mu.Lock()
	defer eng.mu.Unlock()

	eng.offset = i
	return nil
}

func (eng *MemoryEngine) SetIncrement(i int64) error {

	eng.mu.Lock()
	defer eng.mu.Unlock()

	eng.increment = i
	return nil
}

func (eng *MemoryEngine) NextInt() (int64, error) {

	eng.mu.Lock()
	defer eng.mu.Unlock()

	last := eng.last
	next := last + (eng.increment * eng.offset)

	eng.last = next

	go eng.persist(eng.last)
	return next, nil
}

func (eng *MemoryEngine) LastInt() (int64, error) {

	eng.mu.Lock()
	defer eng.mu.Unlock()

	return eng.last, nil
}

func (eng *MemoryEngine) Close() error {
	return nil
}

// PLEASE WRITE ME

func (eng *MemoryEngine) persist(i int64) error {
	return nil
}
