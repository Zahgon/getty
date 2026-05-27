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

package main

import (
	"errors"
	"sync"

	getty "github.com/AlexStocks/getty/transport"
)

const (
	WritePkgTimeout = 1e8
	WritePkgASAP    = 0e9
)

var (
	errTooManySessions = errors.New("too many echo sessions")
	hbHandler          = &HeartbeatHandler{}
	msgHandler         = &MessageHandler{}
	echoMsgHandler     = newEchoMessageHandler()
)

type PackageHandler interface {
	Handle(getty.Session, getty.UDPContext) error
}

////////////////////////////////////////////
// heartbeat handler
////////////////////////////////////////////

type HeartbeatHandler struct{}

func (h *HeartbeatHandler) Handle(session getty.Session, ctx getty.UDPContext) error {
	_ = "STUB: not implemented"
	return nil
}

// return session.WritePkg(getty.UDPContext{Pkg: &rspPkg, PeerAddr: ctx.PeerAddr}, WritePkgTimeout)

////////////////////////////////////////////
// message handler
////////////////////////////////////////////

type MessageHandler struct{}

func (h *MessageHandler) Handle(session getty.Session, ctx getty.UDPContext) error {
	_ = "STUB: not implemented"
	return nil
}

// write echo message handle logic here.
// return session.WritePkg(ctx, WritePkgTimeout)

////////////////////////////////////////////
// EchoMessageHandler
////////////////////////////////////////////

type clientEchoSession struct {
	session getty.Session
	reqNum  int32
}

type EchoMessageHandler struct {
	handlers map[uint32]PackageHandler

	rwlock     sync.RWMutex
	sessionMap map[getty.Session]*clientEchoSession
}

func newEchoMessageHandler() *EchoMessageHandler { _ = "STUB: not implemented"; return nil }

func (h *EchoMessageHandler) OnOpen(session getty.Session) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *EchoMessageHandler) OnError(session getty.Session, err error) {
	_ = "STUB: not implemented"
	return
}

func (h *EchoMessageHandler) OnClose(session getty.Session) { _ = "STUB: not implemented"; return }

func (h *EchoMessageHandler) OnMessage(session getty.Session, udpCtx any) {
	_ = "STUB: not implemented"
	return
}

func (h *EchoMessageHandler) OnCron(session getty.Session) {
	_ = "STUB: not implemented"
	// flag   bool
	return
}

// flag = true

// udp session是根据本地udp socket fd生成的，如果关闭则连同socket也一同关闭了
//if flag {
//	h.rwlock.Lock()
//	delete(h.sessionMap, session)
//	h.rwlock.Unlock()
//	session.Close()
//}
