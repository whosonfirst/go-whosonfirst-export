package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	var roots flags.MultiString

	flag.Var(&roots, "root", "...")
	debug := flag.Bool("debug", false, "...")

	flag.Parse()

	readers := make([]reader.Reader, 0)

	for _, root := range roots {

		r, err := reader.NewFSReader(root)

		if err != nil {
			log.Fatal(err)
		}

		readers = append(readers, r)
	}

	mr, err := reader.NewMultiReader(readers...)

	if err != nil {
		log.Fatal(err)
	}

	for _, str_id := range flag.Args() {

		id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		rel_path, err := uri.Id2RelPath(id)

		if err != nil {
			log.Fatal(err)
		}

		fh, err := mr.Read(rel_path)

		if *debug {

			if err == nil {
				fh, err = mr.Read(rel_path)
			}
		}

		if err != nil {
			log.Fatal(err)
		}

		io.Copy(os.Stdout, fh)
	}

}
