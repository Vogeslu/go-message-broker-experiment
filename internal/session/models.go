package session

import (
	"net"
)

type Session struct {
	Conn    net.Conn
	Greeted bool
	Name    *string
}
