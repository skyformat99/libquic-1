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
	"github.com/goiiot/libquic/io"
)

type packet interface {
	EncodeTo(w io.BufferedWriter) error
	Bytes() []byte
}

// Type defines quic long header packet type
type Type = byte

// Types for different packet
//
// the packet field of Version Negotiation packet is unused, which can be
// randomly selected by the server
const (
	TypeInitial     Type = 0x7F
	TypeRetry       Type = 0x7E
	TypeHandshake   Type = 0x7D
	Type0RTTProtect Type = 0x7C
)
