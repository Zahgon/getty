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
	"bytes"
	"errors"
	"unsafe"
)

// //////////////////////////////////////////
//
//	echo command
//
// //////////////////////////////////////////
type echoCommand uint32

const (
	heartbeatCmd = iota
	echoCmd
)

var echoCommandStrings = [...]string{
	"heartbeat",
	"echo",
}

func (c echoCommand) String() string { _ = "STUB: not implemented"; return "" }

////////////////////////////////////////////
// EchoPkgHandler
////////////////////////////////////////////

const (
	echoPkgMagic     = 0x20160905
	maxEchoStringLen = 0xff

	echoHeartbeatRequestString  = "ping"
	echoHeartbeatResponseString = "pong"
)

var (
	ErrNotEnoughStream = errors.New("packet stream is not enough")
	ErrTooLargePackage = errors.New("package length is exceed the echo package's legal maximum length")
	ErrIllegalMagic    = errors.New("package magic is not right")
)

var echoPkgHeaderLen int

func init() {
	echoPkgHeaderLen = (int)((uint)(unsafe.Sizeof(EchoPkgHeader{})))
}

type EchoPkgHeader struct {
	Magic uint32
	LogID uint32 // log id

	Sequence  uint32 // request/response sequence
	ServiceID uint32 // service id

	Command uint32 // operation command code
	Code    int32  // error code

	Len uint16 // body length
	_   uint16
	_   int32 // reserved, maybe used as package md5 checksum
}

type EchoPackage struct {
	H EchoPkgHeader
	B string
}

func (p EchoPackage) String() string { _ = "STUB: not implemented"; return "" }

func (p EchoPackage) Marshal() (*bytes.Buffer, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *EchoPackage) Unmarshal(buf *bytes.Buffer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// header
