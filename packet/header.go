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

// LongHeader for quic invariant long header
type LongHeader struct {
	PacketType           Type
	Version              uint32
	dcil, scil           byte
	DstConnID, SrcConnID []byte
}

func (lh *LongHeader) EncodeTo(w io.BufferedWriter) error {
	if err := w.WriteByte(lh.PacketType | 0x80); err != nil {
		return err
	}

	// if err := w.Write()
	return nil
}

// ShortHeader for quic invariant short header
type ShortHeader struct {
	DstConnID []byte
}

func (sh *ShortHeader) EncodeTo(w io.BufferedWriter) error {
	if err := w.WriteByte(0x7F); err != nil {
		return err
	}

	_, err := w.Write(sh.DstConnID)
	return err
}
