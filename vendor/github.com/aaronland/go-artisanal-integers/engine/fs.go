package engine

import (
	"errors"
	"fmt"
	"github.com/aaronland/go-artisanal-integers"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type FSEngine struct {
	artisanalinteger.Engine
	key       string
	offset    int64
	increment int64
	mu        *sync.Mutex
}

func NewFSEngine(dsn string) (*FSEngine, error) {

	abs_path, err := filepath.Abs(dsn)

	if err != nil {
		return nil, err
	}

	root := filepath.Dir(abs_path)

	_, err = os.Stat(root)

	if os.IsNotExist(err) {

		err := os.MkdirAll(root, 0755)

		if err != nil {
			return nil, err
		}
	}

	_, err = os.Stat(abs_path)

	if os.IsNotExist(err) {

		err := write_int(abs_path, 0)

		if err != nil {
			return nil, err
		}
	}

	mu := new(sync.Mutex)

	eng := FSEngine{
		key:       abs_path,
		increment: 2,
		offset:    1,
		mu:        mu,
	}

	return &eng, nil
}

func (eng *FSEngine) SetLastInt(i int64) error {

	last, err := eng.LastInt()

	if err != nil {
		return err
	}

	if i < last {
		return errors.New("integer value too small")
	}

	eng.mu.Lock()
	defer eng.mu.Unlock()

	return write_int(eng.key, i)
}

func (eng *FSEngine) SetKey(k string) error {
	// FIX ME
	return nil
}

func (eng *FSEngine) SetOffset(i int64) error {
	eng.offset = i
	return nil
}

func (eng *FSEngine) SetIncrement(i int64) error {
	eng.increment = i
	return nil
}

func (eng *FSEngine) LastInt() (int64, error) {

	eng.mu.Lock()
	defer eng.mu.Unlock()

	return read_int(eng.key)
}

func (eng *FSEngine) NextInt() (int64, error) {

	eng.mu.Lock()
	defer eng.mu.Unlock()

	i, err := read_int(eng.key)

	if err != nil {
		return -1, err
	}

	i = i + eng.increment

	err = write_int(eng.key, i)

	if err != nil {
		return -1, err
	}

	return i, nil
}

func (eng *FSEngine) Close() error {
	return nil
}

func read_int(path string) (int64, error) {

	fh, err := os.Open(path)

	if err != nil {
		return -1, err
	}

	defer fh.Close()

	b, err := ioutil.ReadAll(fh)

	if err != nil {
		return -1, err
	}

	i, err := strconv.ParseInt(string(b), 10, 64)

	if err != nil {
		return -1, err
	}

	return i, nil
}

func write_int(path string, i int64) error {

	fh, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer fh.Close()

	body := fmt.Sprintf("%d", i)

	_, err = fh.Write([]byte(body))
	return err
}
