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

package common

import (
	"errors"
)

var (
	// ErrVarLenIntTooLarge happens when variable length integer is too large
	ErrVarLenIntTooLarge = errors.New("variable length integer value too large ")
)

// ErrorType defines different error code
type ErrorType = byte

// ErrorTypes definition
const (
	ErrNone               ErrorType = 0x0
	ErrInternal           ErrorType = 0x1
	ErrServerBusy         ErrorType = 0x2
	ErrFlowControl        ErrorType = 0x3
	ErrStreamID           ErrorType = 0x4
	ErrStreamState        ErrorType = 0x5
	ErrFinalOffset        ErrorType = 0x6
	ErrFrameEncoding      ErrorType = 0x7
	ErrTransportParameter ErrorType = 0x8
	ErrVersionNegotiation ErrorType = 0x9
	ErrProtocolViolation  ErrorType = 0xA
	ErrInvalidMigration   ErrorType = 0xC
)
