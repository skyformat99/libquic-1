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

import (
	"github.com/goiiot/libquic/common"
)

type Frame interface {
	FrameType() Type
}

// Type defines different frame types
type Type = byte

// Types definition
const (
	TypePadding            Type = 0x00
	TypeRstStream          Type = 0x01
	TypeConnectionClose    Type = 0x02
	TypeApplicationClose   Type = 0x03
	TypeMaxData            Type = 0x04
	TypeMaxStreamData      Type = 0x05
	TypeMaxStreamID        Type = 0x06
	TypePing               Type = 0x07
	TypeBlocked            Type = 0x08
	TypeStreamBlocked      Type = 0x09
	TypeStreamIDBlocked    Type = 0x0A
	TypeNewConnectionID    Type = 0x0B
	TypeStopSending        Type = 0x0C
	TypeRetireConnectionID Type = 0x0D
	TypePathChallenge      Type = 0x0E
	TypePathResponse       Type = 0x0F
	TypeStream             Type = 0x10 // 0x10 - 0x17
	TypeCrypto             Type = 0x18
	TypeNewToken           Type = 0x19
	TypeAck                Type = 0x1A // 0x1A - 0x1B
)

type Padding struct {
}

func (p *Padding) FrameType() Type {
	return TypePadding
}

type ResetStream struct {
	StreamID    uint64
	AppErrCode  uint16
	FinalOffset uint64
}

func (r *ResetStream) FrameType() Type {
	return TypeRstStream
}

type ConnectionClose struct {
	ErrCode      uint16
	ReasonPhrase []byte
}

func (c *ConnectionClose) FrameType() Type {
	return TypeConnectionClose
}

type ApplicationClose struct {
	ErrCode      uint16
	ReasonPhrase []byte
}

func (a *ApplicationClose) FrameType() Type {
	return TypeApplicationClose
}

type MaxData struct {
	MaxData uint64
}

func (a *MaxData) FrameType() Type {
	return TypeMaxData
}

type MaxStreamData struct {
	StreamID      uint64
	MaxStreamData uint64
}

func (a *MaxStreamData) FrameType() Type {
	return TypeMaxStreamData
}

type MaxStreamID struct {
}

func (a *MaxStreamID) FrameType() Type {
	return TypeMaxStreamID
}

type Ping struct {
}

func (a *Ping) FrameType() Type {
	return TypePing
}

type Blocked struct {
	Offset uint64
}

func (a *Blocked) FrameType() Type {
	return TypeBlocked
}

type StreamBlocked struct {
	StreamID uint64
	Offset   uint64
}

func (a *StreamBlocked) FrameType() Type {
	return TypeStreamBlocked
}

type StreamIDBlocked struct {
	StreamID uint64
}

func (a *StreamIDBlocked) FrameType() Type {
	return TypeStreamIDBlocked
}

type NewConnectionID struct {
	SeqNum              uint64
	ConnectionID        common.ConnectionID
	StatelessResetToken [128]byte
}

func (a *NewConnectionID) FrameType() Type {
	return TypeNewConnectionID
}

type RetireConnectionID struct {
	SeqNum uint64
}

func (a *RetireConnectionID) FrameType() Type {
	return TypeRetireConnectionID
}

type StopSending struct {
	StreamID   uint64
	AppErrCode uint16
}

func (a *StopSending) FrameType() Type {
	return TypeStopSending
}

type AckBlock struct {
	Gap uint64
}

type ECNSection struct {
	ECT0Count  uint64
	ECT1Count  uint64
	ECNCECount uint64
}

type Ack struct {
	LargestAcked  uint64
	AckDelay      uint64
	AckBlockCount uint64
	AckBlocks     []AckBlock
	ECN           *ECNSection
}

func (a *Ack) FrameType() Type {
	return TypeAck
}

type PathChallenge struct {
	Data [8]byte
}

func (a *PathChallenge) FrameType() Type {
	return TypePathChallenge
}

type PathResponse struct {
	Data [8]byte
}

func (a *PathResponse) FrameType() Type {
	return TypePathResponse
}

type NewToken struct {
	// Token An opaque blob that the client may use with a future Initial packet
	Token []byte
}

func (a *NewToken) FrameType() Type {
	return TypeNewToken
}

type Stream struct {
	StreamID   uint64
	Offset     *uint64
	StreamData []byte
}

func (a *Stream) FrameType() Type {
	return TypeStream
}

type Crypto struct {
	Offset     uint64
	CryptoData []byte
}

func (a *Crypto) FrameType() Type {
	return TypeCrypto
}
