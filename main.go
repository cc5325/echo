package main

import (
	"errors"
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
			if errors.Is(err, os.ErrDeadlineExceeded) {
				conn.Write([]byte("too slow!\n"))
				conn.Write([]byte("bye!\n"))
			} else {
				log.Printf("error with connection %s: %s", conn.RemoteAddr(), err)
			}
			conn.Close()
			return
		}
		msg = msg[:n]
		conn.Write(msg)
	}
}

func main() {
	addr := os.Getenv("ECHO_ADDRESS")
	if len(addr) == 0 {
		addr = "0.0.0.0:5325"
	}
	s, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening in %s", addr)
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
		}
		go Handle(conn)
	}
}
