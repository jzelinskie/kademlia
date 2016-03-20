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
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

// B is the size in bits of the keys used to identify nodes and store and
// retrieve data.
const B = 160

// NodeID is the unique identifier for a given node on a Kademlia network.
type NodeID [B / 8]byte

// NewNodeID constructs a NodeID for a given hex string.
func NewNodeID(id string) (NodeID, error) {
	var nodeID [B / 8]byte
	decoded, err := hex.DecodeString(id)
	copy(nodeID[:], decoded)
	return NodeID(nodeID), err
}

// String constructs a hex string for a given NodeID.
func (id NodeID) String() string {
	return fmt.Sprintf("%x", [B / 8]byte(id))
}

// Node is the representation of a client participating in a Kademlia network.
type Node struct {
	ID   NodeID
	IP   net.IP
	Port uint32
}

// NewNode creates a new Node to be stored in a KBucket.
func NewNode(ip net.IP, port uint32) *Node {
	return &Node{
		IP:   ip,
		Port: port,
	}
}

// NodeID returns the unique identifier for a Node.
func (n *Node) NodeID() NodeID {
	var bottom NodeID
	if n.ID == bottom {
		buf := bytes.NewBuffer([]byte(n.IP))
		binary.Write(buf, binary.LittleEndian, n.Port)
		return NodeID(sha1.Sum(buf.Bytes()))
	}

	return n.ID
}

// KBucket is a collection of long-living nodes with a similar prefix.
type KBucket struct {
	Nodes        map[NodeID]Node
	Size         int
	LastModified time.Time
}

// NewKBucket allocates a new KBucket of size k.
func NewKBucket(k int) *KBucket {
	return &KBucket{
		Nodes:        make(map[NodeID]Node),
		Size:         k,
		LastModified: time.Now(),
	}
}

func (kb *KBucket) ping() error                  { return nil }
func (kb *KBucket) store() error                 { return nil }
func (kb *KBucket) findNode(n NodeID) error      { return nil }
func (kb *KBucket) findValue(value []byte) error { return nil }
