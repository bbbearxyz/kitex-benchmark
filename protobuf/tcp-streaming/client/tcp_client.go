/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
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
	"encoding/binary"
	"github.com/bbbearxyz/kitex-benchmark/runner"
	"net"
)

func NewPBGrpcClient(opt *runner.Options) runner.Client {
	cli := &tcpClient{}
	cli.address = opt.Address
	return cli
}

type tcpClient struct {
	address   string
}

func (cli *tcpClient) Echo(action, msg string, field, latency, payload, isStream, isTCPCostTest int64) error {
	if isStream == 1 {
		if action != runner.EchoAction {
			return nil
		}
		conn, _ := net.Dial("tcp", cli.address)

		// 把tcp的前8个字节设置为每次streaming传输的大小
		sendBytes := []byte(msg)
		binary.PutVarint(sendBytes, payload)
		conn.Write(sendBytes)

		buf := make([]byte, payload)
		num := 0
		for num != 1024 * 1024 * 1024 {
			delta, _ := conn.Read(buf)
			num += delta
		}
		conn.Close()
		return nil
	}
	return nil
}
