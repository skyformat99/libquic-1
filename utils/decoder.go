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
)

// DecodeVarLenInt decode variable length integer from io.Reader
func DecodeVarLenInt(r *bytes.Buffer) (uint64, error) {
	var (
		b1, b2, b3, b4, b5, b6, b7, b8 byte
		err                            error
	)

	if b1, err = r.ReadByte(); err != nil {
		return 0, err
	}

	length := 1 << (b1 >> 6)
	b1 &= 0x3F
	if length == 1 {
		return uint64(b1), nil
	}

	if b2, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if length == 2 {
		return uint64(b2) | uint64(b1)<<8, nil
	}

	if b3, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if b4, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if length == 4 {
		return uint64(b4) | uint64(b3)<<8 | uint64(b2)<<16 | uint64(b1)<<24, nil
	}

	if b5, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if b6, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if b7, err = r.ReadByte(); err != nil {
		return 0, err
	}

	if b8, err = r.ReadByte(); err != nil {
		return 0, err
	}

	return uint64(b8) | uint64(b7)<<8 | uint64(b6)<<16 | uint64(b5)<<24 |
		uint64(b4)<<32 | uint64(b3)<<40 | uint64(b2)<<48 | uint64(b1)<<56, nil
}
