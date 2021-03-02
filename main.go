package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to start server: %v", err.Error())
	}

	defer lis.Close()
	log.Printf("Started server on :8888")

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %v", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
