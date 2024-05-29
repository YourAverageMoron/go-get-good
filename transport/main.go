package main

import (
	"ryan-jones.io/transport/tcp"
)

func main() {
	s := &tcp.TCPServer{
        Port: ":3000",
	}

    go s.Start()
    select {}
}
