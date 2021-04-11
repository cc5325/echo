package main

import (
	"log"
	"net"
	"os"
	"time"
)

func Handle(conn net.Conn) {
	for {
		conn.SetDeadline(time.Now().Add(240 * time.Second))
		msg := make([]byte, 4096)
		n, err := conn.Read(msg)
		if err != nil {
			log.Printf("error with connection %s: %s", conn.RemoteAddr(), err)
			conn.Close()
			return
		}
		msg = msg[:n]
		conn.Write(msg)
	}
}

func main() {
	s, err := net.Listen("tcp", os.Getenv("ECHO_ADDRESS"))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Printf("error accepting connection: %s", err)
		}
		go Handle(conn)
	}
}
