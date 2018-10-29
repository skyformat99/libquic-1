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
	"encoding/binary"
)

// PutUint16 wraps (network byte order) BigEndian.PutUint16
func PutUint16(d []byte, v uint16) {
	binary.BigEndian.PutUint16(d[:], v)
}

// PutUint32 wraps (network byte order) BigEndian.PutUint32
func PutUint32(d []byte, v uint32) {
	binary.BigEndian.PutUint32(d[:], v)
}
