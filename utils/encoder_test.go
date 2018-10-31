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

func BenchmarkEncodeVarLenInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		EncodeVarLenInt(151288809941952652, buf)
	}
}

func TestEncodeVarLenInt(t *testing.T) {
	for decoded, encoded := range decodedEncodedMap {
		buf := &bytes.Buffer{}
		if err := EncodeVarLenInt(decoded, buf); err != nil {
			t.Errorf("encode var int failed")
		} else {
			if bytes.Compare(buf.Bytes(), encoded) != 0 {
				t.Errorf("result: %v, target: %v", buf.Bytes(), encoded)
			}
		}
	}
}
