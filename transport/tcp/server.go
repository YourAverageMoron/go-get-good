package tcp

import (
	"fmt"
	"net"
)

type TCPServer struct {
	Port string
}

func (s *TCPServer) Start() error {
    fmt.Printf("listening on port %s\n", s.Port)
	ln, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go s.handleConn(conn)
	}
}

func (s *TCPServer) handleConn(conn net.Conn) {
    defer conn.Close()
    // conn.Read
}
