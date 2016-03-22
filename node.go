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

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"net"
	"time"

	"golang.org/x/crypto/sha3"
)

// NodeID is the unique identifier for a given node on a Kademlia network.
type NodeID string

// Bytes constructs a NodeID for a given hex string.
func (nid NodeID) Bytes() []byte {
	decoded, err := hex.DecodeString(string(nid))
	if err != nil {
		panic("kademlia: failed to decode hex NodeID string: " + err.Error())
	}
	return decoded
}

// NewNodeID constructs a NodeID of size b.
func NewNodeID(b int, data []byte) NodeID {
	hash := make([]byte, b)
	sha3.ShakeSum256(hash, data)
	return NodeID(hex.EncodeToString(hash))
}

// Node is the representation of a client participating in a Kademlia network.
type Node struct {
	ID   NodeID
	IP   net.IP
	Port uint32
}

// NewNode creates a new Node with a NodeID of size b.
func NewNode(b int, ip net.IP, port uint32) *Node {
	buf := bytes.NewBuffer([]byte(ip))
	binary.Write(buf, binary.LittleEndian, port)
	return &Node{
		ID:   NewNodeID(b, buf.Bytes()),
		IP:   ip,
		Port: port,
	}
}
