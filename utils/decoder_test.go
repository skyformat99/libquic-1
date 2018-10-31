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

package utils

import (
	"bytes"
	"testing"
)

func BenchmarkDecodeVarLenInt(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.Write([]byte{0xC2, 0x19, 0x7C, 0x5E, 0xFF, 0x14, 0xE8, 0x8C})
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DecodeVarLenInt(buf)
	}
}

func TestDecodeVarLenInt(t *testing.T) {
	for decoded, encoded := range decodedEncodedMap {
		buf := bytes.NewBuffer(encoded)
		result, err := DecodeVarLenInt(buf)
		if err != nil || result != decoded {
			t.Errorf("result: %v, target: %v", result, decoded)
		}
	}
}
