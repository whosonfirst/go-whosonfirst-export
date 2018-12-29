package engine

import (
	"errors"
	"github.com/aaronland/go-artisanal-integers"
	"sync"
)

type ExampleEngine struct {
	artisanalinteger.Engine
	key       string
	increment int64
	offset    int64
	mu        *sync.Mutex
}

func NewExampleEngine(dsn string) (*ExampleEngine, error) {

	mu := new(sync.Mutex)

	eng := ExampleEngine{
		key:       "integers",
		increment: 2,
		offset:    1,
		mu:        mu,
	}

	return &eng, nil
}

func (eng *ExampleEngine) SetLastInt(i int64) error {
	return errors.New("Please implement me")
}

func (eng *ExampleEngine) SetKey(k string) error {
	return errors.New("Please implement me")
}

func (eng *ExampleEngine) SetOffset(i int64) error {
	return errors.New("Please implement me")
}

func (eng *ExampleEngine) SetIncrement(i int64) error {
	return errors.New("Please implement me")
}

func (eng *ExampleEngine) NextInt() (int64, error) {
	return -1, errors.New("Please implement me")
}

func (eng *ExampleEngine) LastInt() (int64, error) {
	return -1, errors.New("Please implement me")
}

func (eng *ExampleEngine) Close() error {
	return errors.New("Please implement me")
}
