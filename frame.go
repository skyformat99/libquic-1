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

package libquic

// FrameType defines different frame types
type FrameType = byte

// FrameTypes definition
const (
	FramePadding            FrameType = 0x00
	FrameRstStream          FrameType = 0x01
	FrameConnectionClose    FrameType = 0x02
	FrameApplicationClose   FrameType = 0x03
	FrameMaxData            FrameType = 0x04
	FrameMaxStreamData      FrameType = 0x05
	FrameMaxStreamID        FrameType = 0x06
	FramePing               FrameType = 0x07
	FrameBlocked            FrameType = 0x08
	FrameStreamBlocked      FrameType = 0x09
	FrameStreamIDBlocked    FrameType = 0x0a
	FrameNewConnectionID    FrameType = 0x0b
	FrameStopSending        FrameType = 0x0c
	FrameRetireConnectionID FrameType = 0x0d
	FramePathChallenge      FrameType = 0x0e
	FramePathResponse       FrameType = 0x0f
	FrameStream             FrameType = 0x10 // 0x10 - 0x17
	FrameCrypto             FrameType = 0x18
	FrameNewToken           FrameType = 0x19
	FrameAck                FrameType = 0x1a // 0x1a - 0x1b
)
