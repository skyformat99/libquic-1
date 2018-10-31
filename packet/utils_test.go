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
	"testing"
)

var decodedEncodedMap = map[int][]byte{
	7023993: {0x6B, 0x2D, 0x79}, // 30 bit
	7061767: {0x6B, 0xC1, 0x07}, // 30 bit
}

func TestDecodePacketID(t *testing.T) {

}

func TestEncodePacketID(t *testing.T) {
}
