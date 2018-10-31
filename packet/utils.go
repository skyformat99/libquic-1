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

	"github.com/goiiot/libquic/common"
)

func boolToByte(val bool) byte {
	if val {
		return 1
	}
	return 0
}

const (
	maxVarLenID1 = 1<<(8-1) - 1
	maxVarLenID2 = 1<<(16-2) - 1
	maxVarLenID4 = 1<<(32-2) - 1
	maxVarLenID  = maxVarLenID4
)

// EncodePacketID use different but similar method as EncodeVarLenInt
func EncodePacketID(n int, w *bytes.Buffer) (err error) {
	if n < 0 || n > maxVarLenID {
		return common.ErrVarLenIntTooLarge
	}

	switch {
	case n <= maxVarLenID1:
		return w.WriteByte(byte(n))
	case n <= maxVarLenID2:
		_, err = w.Write([]byte{byte(n>>8) | 0x80, byte(n)})
		return
	case n <= maxVarLenID4:
		_, err = w.Write([]byte{byte(n>>24) | 0xC0, byte(n >> 16), byte(n >> 8), byte(n)})
	}
	return nil
}

// DecodePacketID use different but similar method as DecodeVarLenInt
func DecodePacketID(r *bytes.Buffer) (int, error) {
	var (
		b1, b2, b3, b4 byte
		err            error
	)

	if b1, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if (b1|0x80)>>6 == 0 {
		return int(b1), nil
	}

	length := b1 >> 6
	b1 &= 0xC0

	if b2, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if length == 2 {
		return int(uint64(b2) | uint64(b1)<<8), nil
	}

	if b3, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if b4, err = r.ReadByte(); err != nil {
		return 0, err
	}

	return int(uint64(b4) | uint64(b3)<<8 | uint64(b2)<<16 | uint64(b1)<<24), nil
}
