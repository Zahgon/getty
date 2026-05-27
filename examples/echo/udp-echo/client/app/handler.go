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

	getty "github.com/AlexStocks/getty/transport"
)

var errSessionNotExist = errors.New("session not exist")

////////////////////////////////////////////
// EchoMessageHandler
////////////////////////////////////////////

type clientEchoSession struct {
	session getty.Session
	reqNum  int32
}

type EchoMessageHandler struct {
	client *EchoClient
}

func newEchoMessageHandler(client *EchoClient) *EchoMessageHandler {
	_ = "STUB: not implemented"
	return nil
}

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

func (h *EchoMessageHandler) OnCron(session getty.Session) { _ = "STUB: not implemented"; return }

// UDP_ENDPOINT session should be long live.
