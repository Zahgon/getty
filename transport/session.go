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
	"sync"
	"time"

	gxcontext "github.com/dubbogo/gost/context"

	gxtime "github.com/dubbogo/gost/time"
	"github.com/gorilla/websocket"

	uatomic "go.uber.org/atomic"
)

const (
	maxReadBufLen   = 4 * 1024
	netIOTimeout    = 1e9      // 1s
	period          = 60 * 1e9 // 1 minute
	pendingDuration = 3e9
	// MaxWheelTimeSpan 900s, 15 minute
	MaxWheelTimeSpan = 900e9
	maxPacketLen     = 16 * 1024

	defaultTLSHandshakeTimeout = time.Second * 3

	defaultSessionName    = "session"
	defaultTCPSessionName = "tcp-session"
	defaultUDPSessionName = "udp-session"
	defaultWSSessionName  = "ws-session"
	defaultWSSSessionName = "wss-session"
	outputFormat          = "session %s, Read Bytes: %d, Write Bytes: %d, Read Pkgs: %d, Write Pkgs: %d"
)

var defaultTimerWheel *gxtime.TimerWheel

func init() {
	gxtime.InitDefaultTimerWheel()
	defaultTimerWheel = gxtime.GetDefaultTimerWheel()
}

// Session wrap connection between the server and the client
type Session interface {
	Connection
	Reset()
	Conn() net.Conn
	Stat() string
	IsClosed() bool
	// EndPoint get endpoint type
	EndPoint() EndPoint
	SetMaxMsgLen(int)
	SetName(string)
	SetEventListener(EventListener)
	SetPkgHandler(ReadWriter)
	SetReader(Reader)
	SetWriter(Writer)
	SetCronPeriod(int)
	SetWaitTime(time.Duration)
	GetAttribute(any) any
	SetAttribute(any, any)
	RemoveAttribute(any)

	// WritePkg the Writer will invoke this function. Pls attention that if timeout is less than 0, WritePkg will send @pkg asap.
	// for udp session, the first parameter should be UDPContext.
	// totalBytesLength: @pkg stream bytes length after encoding @pkg.
	// sendBytesLength: stream bytes length that sent out successfully.
	// err: maybe it has illegal data, encoding error, or write out system error.
	WritePkg(pkg any, timeout time.Duration) (totalBytesLength int, sendBytesLength int, err error)
	WriteBytes([]byte) (int, error)
	WriteBytesArray(...[]byte) (int, error)
	Close()

	AddCloseCallback(handler, key any, callback CallBackFunc)
	RemoveCloseCallback(handler, key any)
}

// getty base session
type session struct {
	name     string
	endPoint EndPoint

	// net read Write
	Connection

	listener EventListener

	// codec
	reader Reader // @reader should be nil when @conn is a gettyWSConn object.
	writer Writer

	// handle logic
	maxMsgLen int32

	// heartbeat
	period time.Duration

	// done
	wait time.Duration
	once *sync.Once
	done chan struct{}

	// attribute
	attrs *gxcontext.ValuesContext

	// goroutines sync
	grNum      uatomic.Int32
	lock       sync.RWMutex
	packetLock sync.RWMutex

	// callbacks
	closeCallback      callbacks
	closeCallbackMutex sync.RWMutex
}

func newSession(endPoint EndPoint, conn Connection) *session { _ = "STUB: not implemented"; return nil }

func newTCPSession(conn net.Conn, endPoint EndPoint) Session {
	_ = "STUB: not implemented"
	return *new(Session)
}

func newUDPSession(conn *net.UDPConn, endPoint EndPoint) Session {
	_ = "STUB: not implemented"
	return *new(Session)
}

func newWSSession(conn *websocket.Conn, endPoint EndPoint) Session {
	_ = "STUB: not implemented"
	return *new(Session)
}

func (s *session) Reset() { _ = "STUB: not implemented"; return }

func (s *session) Conn() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

func (s *session) EndPoint() EndPoint { _ = "STUB: not implemented"; return *new(EndPoint) }

func (s *session) gettyConn() *gettyConn { _ = "STUB: not implemented"; return nil }

// Stat get the connect statistic data
func (s *session) Stat() string { _ = "STUB: not implemented"; return "" }

// IsClosed check whether the session has been closed.
func (s *session) IsClosed() bool { _ = "STUB: not implemented"; return false }

// SetMaxMsgLen set maximum package length of every package in (EventListener)OnMessage(@pkgs)
func (s *session) SetMaxMsgLen(length int) { _ = "STUB: not implemented"; return }

// SetName set session name
func (s *session) SetName(name string) { _ = "STUB: not implemented"; return }

// SetEventListener set event listener
func (s *session) SetEventListener(listener EventListener) { _ = "STUB: not implemented"; return }

// SetPkgHandler set package handler
func (s *session) SetPkgHandler(handler ReadWriter) { _ = "STUB: not implemented"; return }

func (s *session) SetReader(reader Reader) { _ = "STUB: not implemented"; return }

func (s *session) SetWriter(writer Writer) { _ = "STUB: not implemented"; return }

// SetCronPeriod period is in millisecond. Websocket session will send ping frame automatically every peroid.
func (s *session) SetCronPeriod(period int) { _ = "STUB: not implemented"; return }

// SetWaitTime set maximum wait time when session got error or got exit signal
func (s *session) SetWaitTime(waitTime time.Duration) { _ = "STUB: not implemented"; return }

// GetAttribute get attribute of key @session:key
func (s *session) GetAttribute(key any) any { _ = "STUB: not implemented"; return *new(any) }

// SetAttribute set attribute of key @session:key
func (s *session) SetAttribute(key any, value any) { _ = "STUB: not implemented"; return }

// RemoveAttribute remove attribute of key @session:key
func (s *session) RemoveAttribute(key any) { _ = "STUB: not implemented"; return }

func (s *session) sessionToken() string { _ = "STUB: not implemented"; return "" }

func (s *session) WritePkg(pkg any, timeout time.Duration) (pkgBytesLenth int, successCount int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// WriteBytes for codecs
func (s *session) WriteBytes(pkg []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteBytesArray Write multiple packages at once. so we invoke write sys.call just one time.
func (s *session) WriteBytesArray(pkgs ...[]byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// reduce syscall and memcopy for multiple packages

// get len

// merge the pkgs

func heartbeat(_ gxtime.TimerID, _ time.Time, arg any) error { _ = "STUB: not implemented"; return nil }

// if enable task pool, run @f asynchronously.

// func (s *session) RunEventLoop() {
func (s *session) run() { _ = "STUB: not implemented"; return }

// call session opened

// start read gr

func (s *session) addTask(pkg any) {
	_ = "STUB: not implemented"

	// If the session is closed, there is no need to perform CPU-intensive operations.
	return
}

func (s *session) handlePackage() { _ = "STUB: not implemented"; return }

// get package from tcp stream(packet)
func (s *session) handleTCPPackage() error { _ = "STUB: not implemented"; return nil }

// do not handle the left stream in pktBuf and exit asap.
// it is impossible packing a package by the left stream.

// for clause for the network timeout condition check
// s.conn.SetReadTimeout(time.Now().Add(s.rTimeout))

//when read EOF, means that the peer has closed the connection, stop to reconnect to maintain the connection pool.

// as https://github.com/apache/dubbo-getty/issues/77#issuecomment-939652203
// this branch is impossible. Even if it happens, the bufLen will be zero and the error
// is io.EOF when getty continues to read the socket.

// for case 3/case 4

// handle case 1

// handle case 2/case 3

// handle case 4

// continue to handle case 5

// get package from udp packet
func (s *session) handleUDPPackage() error { _ = "STUB: not implemented"; return nil }

// get package from websocket stream
func (s *session) handleWSPackage() error { _ = "STUB: not implemented"; return nil }

func (s *session) stop() { _ = "STUB: not implemented"; return }

// s.done is a blocked channel. if it has not been closed, the default branch will be invoked.

// let read/Write timeout asap

func (s *session) gc() { _ = "STUB: not implemented"; return }

// Close will be invoked by NewSessionCallback(if return error is not nil)
// or (session)handleLoop automatically. It's thread safe.
func (s *session) Close() { _ = "STUB: not implemented"; return }

// GetActive return connection's time
func (s *session) GetActive() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// UpdateActive update connection's active time
func (s *session) UpdateActive() { _ = "STUB: not implemented"; return }

func (s *session) ID() uint32 { _ = "STUB: not implemented"; return 0 }

func (s *session) LocalAddr() string { _ = "STUB: not implemented"; return "" }

func (s *session) RemoteAddr() string { _ = "STUB: not implemented"; return "" }

func (s *session) IncReadPkgNum() { _ = "STUB: not implemented"; return }

func (s *session) IncWritePkgNum() { _ = "STUB: not implemented"; return }

func (s *session) Send(pkg any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *session) ReadTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (s *session) SetSession(ss Session) { _ = "STUB: not implemented"; return }
