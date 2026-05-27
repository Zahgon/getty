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
	"time"
)

const (
	APP_CONF_FILE     = "APP_CONF_FILE"
	APP_LOG_CONF_FILE = "APP_LOG_CONF_FILE"
)

var conf *Config

type (
	GettySessionParam struct {
		CompressEncoding bool   `default:"false"` // Accept-Encoding: gzip, deflate, sdch
		TcpNoDelay       bool   `default:"true"`
		TcpKeepAlive     bool   `default:"true"`
		TcpRBufSize      int    `default:"262144"`
		TcpWBufSize      int    `default:"65536"`
		PkgWQSize        int    `default:"1024"`
		TcpReadTimeout   string `default:"1s"`
		tcpReadTimeout   time.Duration
		TcpWriteTimeout  string `default:"5s"`
		tcpWriteTimeout  time.Duration
		WaitTimeout      string `default:"7s"`
		waitTimeout      time.Duration
		MaxMsgLen        int    `default:"1024"`
		SessionName      string `default:"echo-server"`
	}

	// Config holds supported types by the multiconfig package
	Config struct {
		// local address
		AppName     string   `default:"echo-server"`
		Host        string   `default:"127.0.0.1"`
		Ports       []string `default:"[\"10000\"]"`
		Paths       []string `default:"[\"/echo\"]"`
		ProfilePort int      `default:"10086"`

		// session
		HeartbeatPeriod string `default:"30s"`
		heartbeatPeriod time.Duration
		SessionTimeout  string `default:"60s"`
		sessionTimeout  time.Duration
		SessionNumber   int `default:"1000"`

		// app
		FailFastTimeout string `default:"5s"`
		failFastTimeout time.Duration

		// session tcp parameters
		GettySessionParam GettySessionParam `required:"true"`
	}
)

func initConf() { _ = "STUB: not implemented"; return }

// configure

// gxlog.CInfo("config{%#v}\n", conf)

// log
