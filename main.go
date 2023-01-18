package main

import (
	"log"
	"net"
)

func maini() {
	s := newServer()()
	go s.run()

	listener, err := net.Listen(("tcp",":8888"))
	if err != nil {
		log.Printf("unable to accept connection: %s", err.Error())
		continue
	}

	defer listener.Close()
	log.Printf("server started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unabke to accept connection: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}

