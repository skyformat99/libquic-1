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
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// QuicVersion the version type
type QuicVersion = uint32

const (
	// V1 represent quic version 1, which uses TLS as a cryptographic
	// handshake protocol
	V1 QuicVersion = 0x00000001
)

var (
	// ErrInvalidConnectionID raised when connection's length not valid
	ErrInvalidConnectionID = errors.New("invalid connection id, length should be in range(0, 4 - 18)")
)

// ConnectionID for large number manipulation
type ConnectionID struct {
	value [18]byte
	len   byte
}

// SetByRand
func (c *ConnectionID) SetByRand() error {
	n, err := rand.Read(c.value[:])
	c.len = byte(n)
	return err
}

// SetByReader read from a buffer and set value
func (c *ConnectionID) SetByReader(r *bytes.Buffer, len int) (err error) {
	_, err = io.ReadFull(r, c.value[:len])
	return
}

// SetByBytes read from bytes and set value
func (c *ConnectionID) SetByBytes(data []byte) error {
	length := len(data)

	if length == 0 {
		c.len = 0
	}

	if length < 4 || length > 18 {
		return ErrInvalidConnectionID
	}

	for i, v := range data {
		c.value[i] = v
	}
	c.len = byte(length)

	return nil
}

// SetByValue set connection id via multiple uint64 value
func (c *ConnectionID) SetByValue(value ...uint64) error {
	if len(value) > 3 || (len(value) == 3 && value[0] > (1<<(144-128)-1)) {
		return ErrVarLenIntTooLarge
	}

	return nil
}

func (c *ConnectionID) Equal(other *ConnectionID) bool {
	return bytes.Equal(c.value[:c.len], other.value[:other.len])
}

func (c *ConnectionID) Bytes() []byte {
	return c.value[:c.len]
}

func (c *ConnectionID) Len() int {
	return int(c.len)
}

func (c *ConnectionID) String() string {
	return fmt.Sprintf("%#x", c.Bytes())
}
