package main

import (
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

// accepts all incoming commands
func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NAME:
			s.name(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MESSAGE:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	// prints remote address
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		name:     "anonymous",
		commands: s.commands,
	}

	c.readInput()

}

func (s *server) name(c *client, args []string) {

}

func (s *server) join(c *client, args []string) {

}

func (s *server) listRooms(c *client, args []string) {

}

func (s *server) msg(c *client, args []string) {

}

func (s *server) quit(c *client, args []string) {

}
