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
	"sync"

	gxsync "github.com/dubbogo/gost/sync"
)

const (
	defaultReconnectInterval    = 3e8 // 300ms
	connectInterval             = 5e8 // 500ms
	connectTimeout              = 3e9
	defaultMaxReconnectAttempts = 50
	maxBackOffTimes             = 10
)

var (
	sessionClientKey   = "session-client-owner"
	connectPingPackage = []byte("connect-ping")

	clientID           = EndPointID(0)
	ignoreReconnectKey = "ignore-reconnect"
)

type Client interface {
	EndPoint
}

type client struct {
	ClientOptions

	// endpoint ID
	endPointID EndPointID

	// net
	sync.Mutex
	endPointType EndPointType

	newSession NewSessionCallback
	ssMap      map[Session]struct{}

	sync.Once
	done chan struct{}
	wg   sync.WaitGroup
}

func (c *client) init(opts ...ClientOption) {
	for _, opt := range opts {
		opt(&(c.ClientOptions))
	}
}

func newClient(t EndPointType, opts ...ClientOption) *client { _ = "STUB: not implemented"; return nil }

// NewTCPClient builds a tcp client.
func NewTCPClient(opts ...ClientOption) Client { _ = "STUB: not implemented"; return *new(Client) }

// NewUDPClient builds a connected udp client
func NewUDPClient(opts ...ClientOption) Client { _ = "STUB: not implemented"; return *new(Client) }

// NewWSClient builds a ws client.
func NewWSClient(opts ...ClientOption) Client { _ = "STUB: not implemented"; return *new(Client) }

// NewWSSClient function builds a wss client.
func NewWSSClient(opts ...ClientOption) Client { _ = "STUB: not implemented"; return *new(Client) }

func (c *client) ID() EndPointID { _ = "STUB: not implemented"; return *new(EndPointID) }

func (c *client) EndPointType() EndPointType { _ = "STUB: not implemented"; return *new(EndPointType) }

func (c *client) dialTCP() Session { _ = "STUB: not implemented"; return *new(Session) }

func (c *client) dialUDP() Session { _ = "STUB: not implemented"; return *new(Session) }

// check connection alive by write/read action

func (c *client) dialWS() Session { _ = "STUB: not implemented"; return *new(Session) }

func (c *client) dialWSS() Session { _ = "STUB: not implemented"; return *new(Session) }

// dialer.EnableCompression = true

func (c *client) dial() Session { _ = "STUB: not implemented"; return *new(Session) }

func (c *client) GetTaskPool() gxsync.GenericTaskPool {
	_ = "STUB: not implemented"
	return *new(gxsync.GenericTaskPool)
}

func (c *client) sessionNum() int { _ = "STUB: not implemented"; return 0 }

func (c *client) connect() { _ = "STUB: not implemented"; return }

// client has been closed

// don't distinguish between tcp connection and websocket connection. Because
// gorilla/websocket/conn.go:(Conn)Close also invoke net.Conn.Close()

// there are two methods to keep connection pool. the first approach is like
// redigo's lazy connection pool(https://github.com/gomodule/redigo/blob/master/redis/pool.go:),
// in which you should apply testOnBorrow to check alive of the connection.
// the second way is as follows. @RunEventLoop detects the aliveness of the connection
// in regular time interval.
// the active method maybe overburden the cpu slightly.
// however, you can get a active tcp connection very quickly.
func (c *client) RunEventLoop(newSession NewSessionCallback) { _ = "STUB: not implemented"; return }

// a for-loop connect to make sure the connection pool is valid
func (c *client) reConnect() { _ = "STUB: not implemented"; return }

//exit reconnect when the number of connection pools is sufficient or the current reconnection attempts exceeds the max reconnection attempts.

func (c *client) stop() { _ = "STUB: not implemented"; return }

func (c *client) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (c *client) Close() { _ = "STUB: not implemented"; return }
