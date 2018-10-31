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
	"math/rand"

	"github.com/goiiot/libquic/common"
)

type VersionNegotiation struct {
	DstConnID, SrcConnID common.ConnectionID
	SupportedVersions    []common.QuicVersion
}

func (v *VersionNegotiation) EncodeTo(w *bytes.Buffer) (err error) {
	// v.longHeaderPacket.encodeTo(w)

	// version MUST be 0x00000000
	w.Write([]byte{0x00, 0x00, 0x00, 0x00})
	if err = w.WriteByte(byte(v.DstConnID.Len()-3)<<4 | byte(v.DstConnID.Len()-3)); err != nil {
		return
	}

	if _, err = w.Write(v.DstConnID.Bytes()); err != nil {
		return
	}

	if _, err = w.Write(v.SrcConnID.Bytes()); err != nil {
		return
	}

	// destination
	for _, v := range v.SupportedVersions {
		_, err = w.Write([]byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
		if err != nil {
			return
		}
	}
	return nil
}

func (v *VersionNegotiation) Payload() []byte {
	return nil
}

func (v *VersionNegotiation) Bytes() []byte {
	return nil
}

// Type is not specified for VersionNegotiation packet, use random number
func (v *VersionNegotiation) Type() Type {
	return byte(rand.Int())
}
