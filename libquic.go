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

// TransParamType transport parameter types
type TransParamType = uint16

// TransParams definition
const (
	TransParamInitialMaxStreamDataBidiLocal  TransParamType = 0x0000
	TransParamInitialMaxData                 TransParamType = 0x0001
	TransParamInitialMaxBidiStreams          TransParamType = 0x0002
	TransParamIdleTimeout                    TransParamType = 0x0003
	TransParamPreferredAddress               TransParamType = 0x0004
	TransParamMaxPacketSize                  TransParamType = 0x0005
	TransParamStatelessResetToken            TransParamType = 0x0006
	TransParamAckDelayExponent               TransParamType = 0x0007
	TransParamInitialMaxUniStreams           TransParamType = 0x0008
	TransParamDisableMigration               TransParamType = 0x0009
	TransParamInitialMaxStreamDataBidiRemote TransParamType = 0x000A
	TransParamInitialMaxStreamDataUni        TransParamType = 0x000B
	TransParamMaxAckDelay                    TransParamType = 0x000C
	TransParamOriginalConnectionID           TransParamType = 0x000D
)
