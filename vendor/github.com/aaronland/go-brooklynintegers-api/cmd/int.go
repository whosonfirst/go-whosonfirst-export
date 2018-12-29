package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/aaronland/go-brooklynintegers-api"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var count = flag.Int("count", 1, "The number of Brooklyn Integers to mint.")
	var clients = flag.Int("clients", 10, "...")
	var timings = flag.Bool("timings", false, "...")

	flag.Parse()

	writers := []io.Writer{
		os.Stdout,
	}

	multi := io.MultiWriter(writers...)
	writer := bufio.NewWriter(multi)

	client := api.NewAPIClient()

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	t := *clients
	throttle := make(chan bool, t)

	for i := 0; i < t; i++ {
		throttle <- true
	}

	t1 := time.Now()

	i := int32(0)

	for atomic.LoadInt32(&i) < int32(*count) {

		<-throttle
		wg.Add(1)

		go func(throttle chan bool) {

			defer func() {
				throttle <- true
				wg.Done()
			}()

			if atomic.LoadInt32(&i) >= int32(*count) {
				return
			}

			bi, err := client.NextInt()

			if err != nil {
				log.Println(err)
				return
			}

			mu.Lock()

			str_i := strconv.Itoa(int(bi))

			writer.WriteString(str_i + "\n")
			writer.Flush()

			mu.Unlock()

			atomic.AddInt32(&i, 1)

		}(throttle)

	}

	wg.Wait()

	t2 := time.Since(t1)

	if *timings {
		t := fmt.Sprintf("time to mint %d integers: %v\n", *count, t2)
		os.Stdout.Write([]byte(t))
	}

	os.Exit(0)
}
