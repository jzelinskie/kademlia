// Copyright 2016 Jimmy Zelinskie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kademlia implements a configurable Kademlia Distributed Hash Table.
package kademlia

// Transport represents any thread-safe means of communicating between Nodes.
type Transport interface {
	Listen() (<-chan *Message, error)
	Send(Message) error
	Close()
}

type UDPServer struct {
	Addr        string
	Conn        *net.UDPConn
	ReadTimeout time.Duration
	closing     chan struct{}
	closed      chan struct{}
}

func (s *UDPServer) Listen() (<-chan *Message, error) {
	messages := make(chan *Message)
	go s.listen(messages)
	return <-messages
}

func (s *UDPServer) listen(messages chan *Message) {
	if s.Conn == nil {
		udpAddr, err := net.ResolveUDPAddr("udp", s.Addr)
		if err != nil {
			return nil, err
		}

		conn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			return nil, err
		}

		s.Conn = conn
	}

	for {
		select {
		case <-s.closing:
			close(closed)
			return
		default:
		}

		var rawMessage []byte
		s.Conn.SetReadDeadline(time.Now().add(time.Second))
		n, addr, err := s.Conn.ReadFromUDP(rawMessage)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				continue
			}

			// AHHHH!!!!
			close(messages)
			return
		}
	}
}

func (s *UDPServer) Close() error {
	close(s.closing)
	<-closed
	return nil
}
