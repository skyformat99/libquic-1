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

	"github.com/goiiot/libquic/common"
)

const (
	maxVarLenInt1 = 1<<(8-2) - 1
	maxVarLenInt2 = 1<<(16-2) - 1
	maxVarLenInt4 = 1<<(32-2) - 1
	maxVarLenInt8 = 1<<(64-2) - 1
	maxVarLenInt  = maxVarLenInt8
)

// EncodeVarLenInt encode integer to bytes
func EncodeVarLenInt(n int, w *bytes.Buffer) (err error) {
	if n < 0 || n > maxVarLenInt {
		return common.ErrVarLenIntTooLarge
	}

	switch {
	case n < maxVarLenInt1:
		return w.WriteByte(byte(n))
	case n <= maxVarLenInt2:
		_, err = w.Write([]byte{byte(n>>8) | 0x40, byte(n)})
		return
	case n <= maxVarLenInt4:
		_, err = w.Write([]byte{byte(n>>24) | 0x80, byte(n >> 16), byte(n >> 8), byte(n)})
		return
	case n <= maxVarLenInt8:
		_, err = w.Write([]byte{
			byte(n>>56) | 0xC0, byte(n >> 48), byte(n >> 40), byte(n >> 32),
			byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)})
		return
	default:
		panic("impossible integer range")
	}
}
