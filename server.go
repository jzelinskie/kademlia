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

import (
	"errors"
	"time"
)

// Config represents a Kademlia Distributed Hash Table server.
type Config struct {
	// Alpha is a small number representing the degree of parallelism in network
	// calls.
	Alpha int

	// B is the size in bytes of the keys used to identify nodes and store and
	// retrieve data.
	B int

	// K is the maximum number of contacts stored in a KBucket.
	K int

	// Expire is the time after which a key/value pair expires; this is a TTL
	// from the original publication.
	Expire time.Duration

	// Refresh is the time after which an otherwise unaccessed KBucket must be
	// refreshed.
	Refresh time.Duration

	// Replicate is the interval between Kademlia replication events, when a node
	// is required to publish its entire database.
	Replicate time.Duration

	// Republish is the time after which the original publisher must republish a
	// key/value pair.
	//
	// This value should be smaller than Expire in order to prevent the network
	// from racing to delete active data.
	Republish time.Duration
}

// DefaultConfig is a Kademlia server configuration with sane defaults.
var DefaultConfig = &Config{
	Alpha:     3,
	B:         20,
	K:         20,
	Expire:    time.Hour * 24,
	Refresh:   time.Hour,
	Replicate: time.Hour,
	Republish: time.Hour * 23,
}

type Server struct {
	// Transport is the means by which Nodes will communicate with each other.
	Transport Transport

	// KBuckets are the means of storing contacts (other nodes).
	KBuckets []KBucket

	// Inbox is the means of routing messages to the proper goroutine.
	Inbox Inbox
}

// Get fetches a value with the specified key.
func (s *Server) Get(k Key) (value []byte, err error) {
	if len(k.Bytes()) != s.B {
		return nil, errors.New("kademlia: attempted to get key with length not equal to B")
	}
	return []byte{}, nil
}

// Set stores a value with the specified key.
func (s *Server) Set(k Key) (value []byte, err error) {
	if len(k.Bytes()) != s.B {
		return nil, errors.New("kademlia: attempted to set key with length not equal to B")
	}
	return []byte{}, nil
}

// ping is the PING RPC from the Kademlia paper.
func (s *Server) ping() error                  { return nil }
func (s *Server) store() error                 { return nil }
func (s *Server) findNode(n NodeID) error      { return nil }
func (s *Server) findValue(value []byte) error { return nil }

// Node is the representation of a client participating in a Kademlia network.
type Node struct {
	ID   RandomID
	IP   net.IP
	Port uint32
}

// NewNode creates a new Node with a NodeID of size b bytes.
func NewNode(b int, ip net.IP, port uint32) *Node {
	return &Node{
		ID:   NewRandomID(b),
		IP:   ip,
		Port: port,
	}
}
