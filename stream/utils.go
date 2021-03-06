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

package stream

import (
	"bytes"

	"github.com/goiiot/libquic/common"
	"github.com/goiiot/libquic/utils"
)

// EncodeStreamID encode stream with its type
func EncodeStreamID(id uint64, typ Type, w *bytes.Buffer) error {
	if id > MaxStreamID {
		return common.ErrVarLenIntTooLarge
	}

	streamID := uint64((id << 2) | uint64(typ))
	return utils.EncodeVarLenInt(streamID, w)
}
