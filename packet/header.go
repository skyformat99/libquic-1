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

// LongHeader for quic invariant long header
type LongHeader struct {
	PacketType           Type
	Version              common.QuicVersion
	DstConnID, SrcConnID common.ConnectionID
}

// encodeTo encode LongHeader to bytes
func (lh *LongHeader) encodeTo(w *bytes.Buffer) (err error) {
	if err = w.WriteByte(lh.PacketType | 0x80); err != nil {
		return
	}

	_, err = w.Write([]byte{byte(lh.Version >> 24), byte(lh.Version >> 16), byte(lh.Version >> 8), byte(lh.Version)})
	if err != nil {
		return
	}

	if err = w.WriteByte(byte(lh.DstConnID.Len()-3)<<4 | byte(lh.DstConnID.Len()-3)); err != nil {
		return
	}

	if _, err = w.Write(lh.DstConnID.Bytes()); err != nil {
		return
	}

	if _, err = w.Write(lh.SrcConnID.Bytes()); err != nil {
		return
	}

	return
}

// ShortHeader for quic invariant short header
type ShortHeader struct {
	KeyPhase  bool
	DstConnID common.ConnectionID
}

// encodeTo encode ShortHeader to bytes
func (sh *ShortHeader) encodeTo(w *bytes.Buffer) (err error) {
	flag := boolToByte(sh.KeyPhase) << 6
	if err := w.WriteByte(0x07F & (0x30 | flag)); err != nil {
		return err
	}

	_, err = w.Write(sh.DstConnID.Bytes())

	return
}
