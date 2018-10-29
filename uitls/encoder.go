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
	"io"

	. "github.com/goiiot/libquic/common"
)

func encodePktLenInt(n int, w io.ByteWriter) error {
	if n < 0 || n > MaxVarLenInt {
		return ErrVarLenIntTooLarge
	}
	if n == 0 {
		w.WriteByte(0)
		return nil
	}

	// TODO: implement variable length integer encoding for packet number
	return nil
}

func encodeVarLenInt(n int, w io.ByteWriter) error {
	if n < 0 || n > MaxVarLenInt {
		return ErrVarLenIntTooLarge
	}

	if n == 0 {
		w.WriteByte(0)
		return nil
	}

	// TODO: implement variable length integer encoding
	return nil
}
