// Copyright Go-IIoT (https://github.com/goiiot)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libquic

import (
	"net"
	"time"
)

// QUICConn the quic implementation of net.Conn
type QUICConn struct {
	conn *net.UDPConn
}

// ReadByte read one byte form quic connection
func ReadByte() (byte, error) {
	return 0, nil
}

// WriteByte write one byte to quic connection
func WriteByte(c byte) error {
	return nil
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *QUICConn) Read(b []byte) (n int, err error) {
	return 0, nil
}

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *QUICConn) Write(b []byte) (n int, err error) {
	return 0, nil
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c *QUICConn) Close() error {
	return nil
}

// LocalAddr returns the local network address.
func (c *QUICConn) LocalAddr() net.Addr {
	return nil
}

// RemoteAddr returns the remote network address.
func (c *QUICConn) RemoteAddr() net.Addr {
	return nil
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail with a timeout (see type Error) instead of
// blocking. The deadline applies to all future and pending
// I/O, not just the immediately following call to Read or
// Write. After a deadline has been exceeded, the connection
// can be refreshed by setting a deadline in the future.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
func (c *QUICConn) SetDeadline(t time.Time) error {
	return nil
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (c *QUICConn) SetReadDeadline(t time.Time) error {
	return nil
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *QUICConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// start quic connection logic
// this will spawn new goroutines to handle underlying udp connection
func (c *QUICConn) start() error {
	go func() {
	}()
	return nil
}
