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
	"net"
	"net/url"
	"time"
)

// ErrInvalidBootstrap is the error returned when an invalid address is
// attempted to be used to bootstrap a client into a DHT.
var ErrInvalidBootstrap = errors.New("location to bootstrap node was invalid")

// DHT represents the methods by which a library consumer interacts with a DHT.
type DHT interface {
	Get(Key) ([]byte, error)
	Set([]byte) (Key, error)
}

// Kademlia represents a concrete implementation of the Kademlia Distributed
// Hash Table that adheres to the DHT interface.
type Kademlia struct {
	// Alpha is a small number representing the degree of parallelism in network
	// calls.
	Alpha int

	// KT is used for the encoding and decoding of keys in the Kademlia DHT.
	KT KeyTranscoder

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

	// kbuckets are the means of storing contacts (other nodes).
	kbuckets []KBucket
}

// Get fetches a value with the specified key.
func (k *Kademlia) Get(key Key) (value []byte, err error) {
	return []byte{}, nil
}

// Set stores a value with the specified key.
func (k *Kademlia) Set(value []byte) (Key, error) {
	return "", nil
}

// Transport represents the RPC layer of a Kademlia DHT.
type Transport interface {
	Ping() error
	Store() error
	FindNode(n RandomID) error
	FindValue(value []byte) error
}

// Contact is the representation of a client participating in a Kademlia
// network.
type Contact struct {
	ID   RandomID
	IP   net.IP
	Port uint32
}

// NewContact creates a new Node with a NodeID of size b bytes.
func NewContact(b int, ip net.IP, port uint32) *Contact {
	return &Contact{
		ID:   NewRandomID(b),
		IP:   ip,
		Port: port,
	}
}

// NewMainlineDHT returns a DHT that connects to the Mainline BitTorrent DHT.
func NewMainlineDHT(bootstrap []url.URL) (DHT, error) {
	if bootstrap == nil {
		return nil, ErrInvalidBootstrap
	}

	// TODO(jzelinskie): actually bootstrap here

	return &Kademlia{
		Alpha:     3,
		KT:        SHA1Transcoder{},
		Expire:    time.Hour * 24,
		Refresh:   time.Hour,
		Replicate: time.Hour,
		Republish: time.Hour * 23,
	}, nil
}
