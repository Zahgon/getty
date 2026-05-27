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
	"math/rand"
	"net"
	"sync"
	"time"

	getty "github.com/AlexStocks/getty/transport"
)

var (
	reqID uint32
	r     *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

////////////////////////////////////////////////////////////////////
// echo client
////////////////////////////////////////////////////////////////////

type EchoClient struct {
	lock        sync.RWMutex
	sessions    []*clientEchoSession
	gettyClient getty.Client
	serverAddr  net.UDPAddr
}

func (c *EchoClient) isAvailable() bool { _ = "STUB: not implemented"; return false }

func (c *EchoClient) close() { _ = "STUB: not implemented"; return }

func (c *EchoClient) selectSession() getty.Session {
	_ = "STUB: not implemented"
	// get route server session
	return *new(getty.Session)
}

func (c *EchoClient) addSession(session getty.Session) { _ = "STUB: not implemented"; return }

func (c *EchoClient) removeSession(session getty.Session) { _ = "STUB: not implemented"; return }

func (c *EchoClient) updateSession(session getty.Session) { _ = "STUB: not implemented"; return }

func (c *EchoClient) getClientEchoSession(session getty.Session) (clientEchoSession, error) {
	_ = "STUB: not implemented"
	return *new(clientEchoSession), nil
}

func (c *EchoClient) heartbeat(session getty.Session) { _ = "STUB: not implemented"; return }

// pkg.H.ServiceID = 0

// if err := session.WritePkg(ctx, WritePkgTimeout); err != nil {
