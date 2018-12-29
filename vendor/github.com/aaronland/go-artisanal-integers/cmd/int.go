package main

import (
	"bufio"
	"flag"
	"github.com/aaronland/go-artisanal-integers/engine"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	var dsn = flag.String("dsn", "", "The data source name (dsn) for connecting to the artisanal integer engine.")
	var last = flag.Int("set-last-int", 0, "Set the last known integer.")
	var offset = flag.Int("set-offset", 0, "Set the offset used to mint integers.")
	var increment = flag.Int("set-increment", 0, "Set the increment used to mint integers.")
	var continuous = flag.Bool("continuous", false, "Continuously mint integers. This is mostly only useful for debugging.")

	flag.Parse()

	eng, err := engine.NewMemoryEngine(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	if *last != 0 {

		err = eng.SetLastInt(int64(*last))

		if err != nil {
			log.Fatal(err)
		}
	}

	if *increment != 0 {

		err = eng.SetIncrement(int64(*increment))

		if err != nil {
			log.Fatal(err)
		}
	}

	if *offset != 0 {

		err = eng.SetOffset(int64(*offset))

		if err != nil {
			log.Fatal(err)
		}
	}

	writers := []io.Writer{
		os.Stdout,
	}

	multi := io.MultiWriter(writers...)
	writer := bufio.NewWriter(multi)

	for {

		next, err := eng.NextInt()

		if err != nil {
			log.Fatal(err)
		}

		str_next := strconv.FormatInt(next, 10)
		writer.WriteString(str_next + "\n")
		writer.Flush()

		if !*continuous {
			break
		}

	}

	os.Exit(0)
}
