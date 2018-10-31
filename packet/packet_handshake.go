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

// Handshake packet
type Handshake struct {
}

// EncodeTo encode Handshake to buffer
func (h *Handshake) EncodeTo(w *bytes.Buffer) error {
	return nil
}

// Payload get payload of Handshake packet
func (h *Handshake) Payload() []byte {
	return nil
}

// Bytes encoded for Handshake packet
func (h *Handshake) Bytes() []byte {
	return nil
}

// Type of Handshake packet
func (h *Handshake) Type() Type {
	return TypeHandshake
}
