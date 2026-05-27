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
	getty "github.com/AlexStocks/getty/transport"
)

var echoPkgHandler = NewEchoPackageHandler()

type EchoPackageHandler struct{}

func NewEchoPackageHandler() *EchoPackageHandler { _ = "STUB: not implemented"; return nil }

func (h *EchoPackageHandler) Read(ss getty.Session, data []byte) (any, int, error) {
	_ = "STUB: not implemented"
	// log.Debug("get client package:%s", gxstrings.String(data))
	return *new(any), 0, nil
}

// return data, len(data), nil

func (h *EchoPackageHandler) Write(ss getty.Session, pkg any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
