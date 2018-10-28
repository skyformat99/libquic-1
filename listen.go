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

import "net"

// listener implements net.Listener
type listener struct {
	conn *net.UDPConn
}

// Accept waits for and returns the next connection to the listener.
func (l *listener) Accept() (net.Conn, error) {
	return nil, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (l *listener) Close() error {
	return nil
}

// Addr returns the listener's network address.
func (l *listener) Addr() net.Addr {
	return nil
}

// start quic logic
func (l *listener) start() error {
	go func() {
	}()
	return nil
}

// Listen announces on the local network address.
//
// The network must be "quic", "quic4", "quic6"
func Listen(network, address string) (net.Listener, error) {
	udpNetwork := getUDPNetwork(network)
	pktConn, err := net.ListenPacket(udpNetwork, address)
	if err != nil {
		return nil, err
	}
	l := &listener{conn: pktConn.(*net.UDPConn)}

	return l, l.start()
}
