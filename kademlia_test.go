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
	"net"
	"testing"
)

func TestNodeID(t *testing.T) {
	var table = []struct {
		ip       string
		port     uint32
		expected string
	}{
		{"127.0.0.1", 0, "5fbda2308f8be2c4f2e58a8469f1e2b43e5b1f37"},
		{"192.168.99.100", 5000, "d61f08a1ee4b32c8d5846f9c7feff6148ffb9017"},
	}

	for _, tt := range table {
		expected, err := NewNodeID(tt.expected)
		if err != nil {
			t.Errorf("failed to parse NodeID from %s: %s", tt.expected, err.Error())
		}

		if node := NewNode(net.ParseIP(tt.ip), tt.port); node.NodeID() != expected {
			t.Errorf("calculated NodeID (%s) did not match expected (%#v)", node.NodeID(), tt.expected)
		}
	}
}
