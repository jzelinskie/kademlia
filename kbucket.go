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
