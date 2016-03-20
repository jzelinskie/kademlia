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

package kademlia

import "time"

// Server represents a Kademlia Distributed Hash Table server.
type Server struct {
	// Addr is the listen address of the server.
	Addr string

	// Alpha is a small number representing the degree of parallelism in network
	// calls.
	Alpha int

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
	Republish time.Duration
}

// NewServer returns a new Kademlia server that
func NewServer() *Server {
	return &Server{
		Addr:      ":9043",
		Alpha:     3,
		K:         20,
		Expire:    time.Second * 86400,
		Refresh:   time.Second * 3600,
		Replicate: time.Second * 3600,
		Republish: time.Second * 86400,
	}
}

// Get fetches a value with the specified key.
func (s *Server) Get(key []byte) (value []byte) { return []byte{} }

// Set stores a value with the specified key.
func (s *Server) Set(key []byte) (value []byte) { return []byte{} }
