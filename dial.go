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

import (
	"context"
	"net"
	"time"

	"github.com/goiiot/libquic/utils"
)

// Dialer wraps net.Dialer for quic dial
type Dialer struct {
	net.Dialer
}

// Dial connects to the address on the named network.
//
// Known networks are "quic", "quic4" (IPv4-only), "quic6" (IPv6-only)
//
// Examples:
// 	Dial("quic", "golang.org:http")
// 	Dial("quic4", "192.0.2.1:http")
// 	Dial("quic", "198.51.100.1:80")
// 	Dial("quic6", "[2001:db8::1]:domain")
// 	Dial("quic", "[fe80::1%lo0]:53")
// 	Dial("quic", ":80")
//
func Dial(network, address string) (net.Conn, error) {
	var d Dialer
	return d.dial(network, address)
}

// DialTimeout acts like Dial but takes a timeout.
//
// The timeout includes name resolution, if required.
// When using TCP, and the host in the address parameter resolves to
// multiple IP addresses, the timeout is spread over each consecutive
// dial, such that each is given an appropriate fraction of the time
// to connect.
//
// See func Dial for a description of the network and address
// parameters.
func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := Dialer{Dialer: net.Dialer{Timeout: timeout}}
	return d.dial(network, address)
}

func (d *Dialer) dial(network, address string) (net.Conn, error) {
	return d.dialContext(context.Background(), network, address)
}

func (d *Dialer) dialContext(ctx context.Context, network, address string) (net.Conn, error) {
	udpProto := utils.GetUDPNetwork(network)

	if udpProto == "" {
		return nil, &net.OpError{
			Op:     "dial",
			Net:    network,
			Source: nil,
			Addr:   nil,
			Err:    net.UnknownNetworkError(network),
		}
	}

	udpConn, err := d.DialContext(ctx, network, address)
	if err != nil {
		return nil, err
	}

	quicConn := &QUICConn{conn: udpConn.(*net.UDPConn)}
	return quicConn, quicConn.start()
}
