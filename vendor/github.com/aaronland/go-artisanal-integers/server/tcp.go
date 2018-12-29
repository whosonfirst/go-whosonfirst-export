package server

// EXPERIMENTAL

import (
	"bufio"
	"github.com/aaronland/go-artisanal-integers"
	"log"
	"net"
	"net/url"
	"strconv"
)

type TCPServer struct {
	artisanalinteger.Server
	url *url.URL
}

func NewTCPServer(u *url.URL, args ...interface{}) (*TCPServer, error) {

	server := TCPServer{
		url: u,
	}

	return &server, nil
}

func (s *TCPServer) Address() string {
	return s.url.Host
}

func (s *TCPServer) ListenAndServe(service artisanalinteger.Service) error {

	listener, err := net.Listen("tcp", s.url.Host)

	if err != nil {
		return err
	}

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		// log.Println(conn.RemoteAddr().String())

		go func() {

			defer conn.Close()

			i, err := service.NextInt()

			if err != nil {
				return
			}

			str_i := strconv.FormatInt(i, 10)

			bufout := bufio.NewWriter(conn)
			bufout.WriteString(str_i + "\n")
			bufout.Flush()
		}()
	}
}
