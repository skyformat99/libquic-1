/*
 * Copyright Go-IIoT (https://github.com/goiiot)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package packet

import (
	"bytes"
	"testing"
)

var decodedEncodedMap = map[int][]byte{
	7023993: {0xC0, 0x6B, 0x2D, 0x79}, // 30 bit
	7061767: {0xC0, 0x6B, 0xC1, 0x07}, // 30 bit
}

func TestDecodePacketID(t *testing.T) {
	for decoded, encoded := range decodedEncodedMap {
		buf := bytes.NewBuffer(encoded)
		if result, err := DecodePacketID(buf); err != nil {
			t.Errorf("failed decode packet id: %v", err)
		} else if result != decoded {
			t.Errorf("result: %v, target: %v", result, decoded)
		}
	}
}

func TestEncodePacketID(t *testing.T) {
	for decoded, encoded := range decodedEncodedMap {
		buf := &bytes.Buffer{}
		if err := EncodePacketID(decoded, buf); err != nil {
			t.Errorf("failed decode packet id: %v", err)
		} else if !bytes.Equal(buf.Bytes(), encoded) {
			t.Errorf("result: %v, target: %v", buf.Bytes(), encoded)
		}
	}
}
