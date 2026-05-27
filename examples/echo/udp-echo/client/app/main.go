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
	_ "net/http/pprof"

	gxlog "github.com/AlexStocks/goext/log"

	getty "github.com/AlexStocks/getty/transport"

	log "github.com/AlexStocks/getty/util"
)

const (
	pprofPath = "/debug/pprof/"
)

const (
	WritePkgTimeout = 1e9
	WritePkgASAP    = 0e9
)

var (
	connectedClient   EchoClient
	unconnectedClient EchoClient
)

////////////////////////////////////////////////////////////////////
// main
////////////////////////////////////////////////////////////////////

func main() {
	initConf()

	initProfiling()

	initClient()
	log.Info("%s starts successfull! its version=%s\n", conf.AppName, getty.Version)
	gxlog.CInfo("%s starts successfull! its version=%s\n", conf.AppName, getty.Version)

	go test()

	initSignal()
}

func initProfiling() { _ = "STUB: not implemented"; return }

func newSession(session getty.Session) error { _ = "STUB: not implemented"; return nil }

func initClient() { _ = "STUB: not implemented"; return }

func uninitClient() { _ = "STUB: not implemented"; return }

func initSignal() {
	_ = "STUB: not implemented"
	// signal.Notify的ch信道是阻塞的(signal.Notify不会阻塞发送信号), 需要设置缓冲
	return
}

// It is not possible to block SIGKILL or syscall.SIGSTOP

// reload()

// log.Warn("app exit now by force...")
// os.Exit(1)

// 要么fastFailTimeout时间内执行完毕下面的逻辑然后程序退出，要么执行上面的超时函数程序强行退出

// fmt.Println("app exit now...")

func echo(client *EchoClient) { _ = "STUB: not implemented"; return }

// pkg.H.ServiceID = 0

// err := session.WritePkg(ctx, WritePkgTimeout)

func testEchoClient(client *EchoClient) { _ = "STUB: not implemented"; return }

func test() { _ = "STUB: not implemented"; return }
