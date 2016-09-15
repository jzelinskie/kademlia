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
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// Key represents the unique identifier for a value stored in a Kademlia DHT.
//
// Keys are encoded as hex strings in order to function as variable-length
// keys in Go maps.
type Key string

// KeyTranscoder is a function that can encode data as a Key capable of being
// used in a Kademlia DHT.
type KeyTranscoder interface {
	Encode([]byte) Key
	Decode(Key) []byte
}

// SHA3Shake256Transcoder implements a variable length KeyTranscoder using the
// output of a variable-length SHA3-256 ShakeHash represented in hex.
type SHA3Shake256Transcoder struct {
	OutputSize int
}

// Encode constructs a new Key.
func (t SHA3Shake256Transcoder) Encode(data []byte) Key {
	hash := make([]byte, t.OutputSize)
	sha3.ShakeSum256(hash, data)
	return Key(hex.EncodeToString(hash))
}

// Decode constructs the original bytes out of a Key.
func (t SHA3Shake256Transcoder) Decode(k Key) []byte {
	decoded, err := hex.DecodeString(string(k))
	if err != nil {
		panic("failed to decode hexadecimal form of Key: " + err.Error())
	}
	return decoded
}
