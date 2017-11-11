// Copyright 2017 Tim Oster. All rights reserved.
// Use of this source code is governed by the MIT license.
// More information can be found in the LICENSE file.

package rmnp

import (
	"net"
	"fmt"
)

type Server struct {
	protocolImpl
}

func NewServer(address string) *Server {
	s := new(Server)

	s.readFunc = func(conn *net.UDPConn, buffer []byte) (int, *net.UDPAddr, bool) {
		length, addr, err := conn.ReadFromUDP(buffer)

		if err != nil{
			return 0, nil, false
		}

		return length, addr, true
	}

	s.writeFunc = func(c *Connection, buffer []byte) {
		c.conn.WriteToUDP(buffer, c.addr)
	}

	AddConnectionCallback(&s.onConnect, func(connection *Connection) {
		fmt.Println("new client")
	})

	s.init(address)
	return s
}

func (s *Server) Start() {
	socket, err := net.ListenUDP("udp", s.address)
	checkError("Cannot listen on udp", err)
	s.socket = socket
	s.listen()
}

func (s *Server) Stop() {
	s.socket.Close()
}