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
	"net"
	"net/http"
	"sync"
	"time"

	gxsync "github.com/dubbogo/gost/sync"

	"github.com/gorilla/websocket"

	perrors "github.com/pkg/errors"

	uatomic "go.uber.org/atomic"
)

var (
	errSelfConnect        = perrors.New("connect self!")
	serverFastFailTimeout = time.Second * 1

	serverID uatomic.Int32
)

// Server interface
type Server interface {
	EndPoint
}

// StreamServer is like tcp/websocket/wss server
type StreamServer interface {
	Server
	// Listener get the network listener
	Listener() net.Listener
}

// PacketServer is like udp listen endpoint
type PacketServer interface {
	Server
	// PacketConn get the network listener
	PacketConn() net.PacketConn
}

type server struct {
	ServerOptions

	// endpoint ID
	endPointID EndPointID

	// net
	pktListener    net.PacketConn
	streamListener net.Listener
	lock           sync.Mutex // for server
	endPointType   EndPointType
	server         *http.Server // for ws or wss server
	sync.Once
	done chan struct{}
	wg   sync.WaitGroup
}

func (s *server) init(opts ...ServerOption) {
	for _, opt := range opts {
		opt(&(s.ServerOptions))
	}
}

func newServer(t EndPointType, opts ...ServerOption) *server { _ = "STUB: not implemented"; return nil }

// NewTCPServer builds a tcp server.
func NewTCPServer(opts ...ServerOption) Server { _ = "STUB: not implemented"; return *new(Server) }

// NewUDPEndPoint builds a unconnected udp server.
func NewUDPEndPoint(opts ...ServerOption) Server { _ = "STUB: not implemented"; return *new(Server) }

// NewWSServer builds a websocket server.
func NewWSServer(opts ...ServerOption) Server { _ = "STUB: not implemented"; return *new(Server) }

// NewWSSServer builds a secure websocket server.
func NewWSSServer(opts ...ServerOption) Server { _ = "STUB: not implemented"; return *new(Server) }

func (s *server) ID() int32 { _ = "STUB: not implemented"; return 0 }

func (s *server) EndPointType() EndPointType { _ = "STUB: not implemented"; return *new(EndPointType) }

func (s *server) stop() { _ = "STUB: not implemented"; return }

// if the log output is "shutdown ctx: context deadline exceeded"， it means that
// there are still some active connections.

// let the server exit asap when got error from RunEventLoop.

func (s *server) GetTaskPool() gxsync.GenericTaskPool {
	_ = "STUB: not implemented"
	return *new(gxsync.GenericTaskPool)
}

func (s *server) IsClosed() bool { _ = "STUB: not implemented"; return false }

// net.ipv4.tcp_max_syn_backlog
// net.ipv4.tcp_timestamps
// net.ipv4.tcp_tw_recycle
func (s *server) listenTCP() error { _ = "STUB: not implemented"; return nil }

func (s *server) listenUDP() error { _ = "STUB: not implemented"; return nil }

// Listen announces on the local network address.
func (s *server) listen() error { _ = "STUB: not implemented"; return nil }

func (s *server) accept(newSession NewSessionCallback) (Session, error) {
	_ = "STUB: not implemented"
	return *new(Session), nil
}

func (s *server) runTCPEventLoop(newSession NewSessionCallback) { _ = "STUB: not implemented"; return }

//	change the error checking from "netErr.Temporary()" to "netErr.Timeout()".
//  as per https://github.com/golang/go/issues/45729,
//  Timeout() correctly captures subset of Temporary() errors that could be retried.
//  The rest of Temporary() errors should not be retried anyway (like syscall errors, out of file descriptors)

func (s *server) runUDPEventLoop(newSession NewSessionCallback) { _ = "STUB: not implemented"; return }

type wsHandler struct {
	http.ServeMux
	server     *server
	newSession NewSessionCallback
	upgrader   websocket.Upgrader
}

func newWSHandler(server *server, newSession NewSessionCallback) *wsHandler {
	_ = "STUB: not implemented"
	return nil
}

// in default, ReadBufferSize & WriteBufferSize is 4k
// HandshakeTimeout: server.HTTPTimeout,
// allow connections from any origin

func (s *wsHandler) serveWSRequest(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return

	// w.WriteHeader(http.StatusMethodNotAllowed)
}

// conn.SetReadLimit(int64(handler.maxMsgLen))

// runWSEventLoop serve websocket client request
// @newSession: new websocket connection callback
func (s *server) runWSEventLoop(newSession NewSessionCallback) { _ = "STUB: not implemented"; return }

// ReadTimeout:    server.HTTPTimeout,
// WriteTimeout:   server.HTTPTimeout,

// serve websocket client request
// RunWSSEventLoop serve websocket client request
func (s *server) runWSSEventLoop(newSession NewSessionCallback) { _ = "STUB: not implemented"; return }

// do not verify peer certs

// ReadTimeout:    server.HTTPTimeout,
// WriteTimeout:   server.HTTPTimeout,

// RunEventLoop serves client request.
// @newSession: new connection callback
func (s *server) RunEventLoop(newSession NewSessionCallback) { _ = "STUB: not implemented"; return }

func (s *server) Listener() net.Listener { _ = "STUB: not implemented"; return *new(net.Listener) }

func (s *server) PacketConn() net.PacketConn {
	_ = "STUB: not implemented"
	return *new(net.PacketConn)
}

func (s *server) Close() { _ = "STUB: not implemented"; return }
