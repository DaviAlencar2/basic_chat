package models

import "net"

type ChatGroup struct{
	Connections []net.Conn
	Name string
}
