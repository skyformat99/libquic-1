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
)

// Initial packet
//
// An Initial packet uses long headers with a type value of 0x7F.
//
// It carries the first CRYPTO frames sent by the client and server to perform
// key exchange, and carries ACKs in either direction.
type Initial struct {
	LongHeader
	Token []byte
}

func (i *Initial) EncodeTo(w *bytes.Buffer) error {
	return nil
}

func (i *Initial) Payload() []byte {
	return nil
}

func (i *Initial) Bytes() []byte {
	return nil
}

func (i *Initial) Type() Type {
	return TypeInitial
}
