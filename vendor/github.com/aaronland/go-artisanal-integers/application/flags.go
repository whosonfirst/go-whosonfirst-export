package application

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func NewFlagSet(name string) *flag.FlagSet {

	fs := flag.NewFlagSet(name, flag.ExitOnError)

	fs.Usage = func() {
		fs.PrintDefaults()
	}

	return fs
}

func AssignCommonFlags(fs *flag.FlagSet) {

	fs.String("protocol", "http", "The protocol for the server to implement. Valid options are: http,tcp.")
	fs.String("host", "localhost", "The hostname to listen for requests on")
	fs.Int("port", 8080, "The port number to listen for requests on")
	fs.String("path", "/", "The path to listen for requests on")
}

func ParseFlags(fs *flag.FlagSet) {

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "-h" {
		fs.Usage()
		os.Exit(0)
	}

	fs.Parse(args)

	fs.VisitAll(func(fl *flag.Flag) {

		name := fl.Name
		env := name

		env = strings.ToUpper(env)
		env = strings.Replace(env, "-", "_", -1)
		env = fmt.Sprintf("ARTISANAL_%s", env)

		val, ok := os.LookupEnv(env)

		if ok {
			log.Printf("set -%s flag from %s environment variable\n", name, env)
			fs.Set(name, val)
		}
	})

}

func Lookup(fl *flag.FlagSet, k string) (interface{}, error) {

	v := fl.Lookup(k)

	if v == nil {
		msg := fmt.Sprintf("Unknown flag '%s'", k)
		return nil, errors.New(msg)
	}

	// Go is weird...
	return v.Value.(flag.Getter).Get(), nil
}

func StringVar(fl *flag.FlagSet, k string) (string, error) {

	i, err := Lookup(fl, k)

	if err != nil {
		return "", err
	}

	return i.(string), nil
}

func IntVar(fl *flag.FlagSet, k string) (int, error) {

	i, err := Lookup(fl, k)

	if err != nil {
		return 0, err
	}

	return i.(int), nil
}

func BoolVar(fl *flag.FlagSet, k string) (bool, error) {

	i, err := Lookup(fl, k)

	if err != nil {
		return false, err
	}

	return i.(bool), nil
}
