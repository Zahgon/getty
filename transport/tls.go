/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package getty

import (
	"crypto/tls"
)

// TlsConfigBuilder  tls config builder interface
type TlsConfigBuilder interface {
	BuildTlsConfig() (*tls.Config, error)
}

// ServerTlsConfigBuilder impl TlsConfigBuilder for server
type ServerTlsConfigBuilder struct {
	ServerKeyCertChainPath        string
	ServerPrivateKeyPath          string
	ServerKeyPassword             string
	ServerTrustCertCollectionPath string
}

// BuildTlsConfig impl TlsConfigBuilder method
func (s *ServerTlsConfigBuilder) BuildTlsConfig() (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// do not verify peer certs

// ClientTlsConfigBuilder impl TlsConfigBuilder for client
type ClientTlsConfigBuilder struct {
	ClientKeyCertChainPath        string
	ClientPrivateKeyPath          string
	ClientKeyPassword             string
	ClientTrustCertCollectionPath string
}

// BuildTlsConfig impl TlsConfigBuilder method
func (c *ClientTlsConfigBuilder) BuildTlsConfig() (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
