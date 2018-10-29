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
	"testing"
)

func TestPutUint16(t *testing.T) {
	buf := make([]byte, 2)
	PutUint16(buf, 0xAABB)
	if buf[0] != 0xAA || buf[1] != 0xBB {
		t.Fail()
	}
}

func TestPutUint32(t *testing.T) {
	buf := make([]byte, 4)
	PutUint32(buf, 0xAABBCCDD)
	if buf[0] != 0xAA || buf[1] != 0xBB || buf[2] != 0xCC || buf[3] != 0xDD {
		t.Fail()
	}
}
