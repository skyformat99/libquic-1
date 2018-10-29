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

package frame

// Type defines different frame types
type Type = byte

// Types definition
const (
	FramePadding            Type = 0x00
	FrameRstStream          Type = 0x01
	FrameConnectionClose    Type = 0x02
	FrameApplicationClose   Type = 0x03
	FrameMaxData            Type = 0x04
	FrameMaxStreamData      Type = 0x05
	FrameMaxStreamID        Type = 0x06
	FramePing               Type = 0x07
	FrameBlocked            Type = 0x08
	FrameStreamBlocked      Type = 0x09
	FrameStreamIDBlocked    Type = 0x0a
	FrameNewConnectionID    Type = 0x0b
	FrameStopSending        Type = 0x0c
	FrameRetireConnectionID Type = 0x0d
	FramePathChallenge      Type = 0x0e
	FramePathResponse       Type = 0x0f
	FrameStream             Type = 0x10 // 0x10 - 0x17
	FrameCrypto             Type = 0x18
	FrameNewToken           Type = 0x19
	FrameAck                Type = 0x1a // 0x1a - 0x1b
)
