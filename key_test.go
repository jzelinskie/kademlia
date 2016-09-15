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

import "testing"

func TestSHA3Shake256KeyTranscoder(t *testing.T) {
	var table = []struct {
		outputSize int
		data       []byte
		expected   string
	}{
		{1, []byte{0x0}, "b8"},
		{10, []byte{0x0}, "b8d01df855f7075882c6"},
	}

	for _, tt := range table {
		transcoder := SHA3Shake256Transcoder{tt.outputSize}
		if key := transcoder.Encode(tt.data); string(key) != tt.expected {
			t.Errorf("calculated key (%s) did not match expected (%s)", key, tt.expected)
		}
	}
}
