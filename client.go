package chat

import "net"

type client struct {
	net.Conn
	net.Addr
	name     string
	room     *room
	commands chan<- command
}
