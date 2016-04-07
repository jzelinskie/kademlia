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
	"crypto/rand"
	"encoding/hex"
	"net"
)

func randomBytes(length int) []byte {
	bytes := make([]byte, b)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("kademlia: failed to generate random bytes for NodeID: " + err.Error())
	}
	return bytes
}

// RandomID is a randomly-generated, unique identifier for something on a
// Kademlia network.
//
// RandomIDs are encoded as hex strings in order to function as variable-length
// keys in Go maps.
type RandomID string

// Bytes returns the byte-interpretation of the hex string.
func (r RandomID) Bytes() []byte {
	decoded, err := hex.DecodeString(string(r))
	if err != nil {
		panic("kademlia: failed to decode RandomID: " + err.Error())
	}
	return decoded
}

// NewRandomID constructs a RandomID of the provided length using the Go
// standard library's PRNG.
func NewRandomID(length int) NodeID { return NodeID(hex.EncodeToString(randomBytes(b))) }
