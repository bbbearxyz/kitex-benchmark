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
	"fmt"
	"github.com/bbbearxyz/kitex-benchmark/perf"
	"github.com/bbbearxyz/kitex-benchmark/runner"
	"log"
	"net"
)

const (
	port = 8003
)

var data []byte
var recorder = perf.NewRecorder("TCP-STREAMING@Server")

func StreamTest(c net.Conn) error {
	buf := make([]byte, 256)
	c.Read(buf)
	length, _ := binary.Varint(buf)

	// 计算1GB / length的次数
	round := int64(0)
	sendDataLength := int64(length)
	lastDataLength := int64(0)

	if 1024 * 1024 * 1024 % length == 0 {
		round = 1024 * 1024 * 1024 / length
		lastDataLength = sendDataLength
	} else {
		round = 1024 * 1024 * 1024 / length + 1
		lastDataLength = 1024 * 1024 * 1024 - (sendDataLength * (round - 1))
	}

	for i := int64(0); i < round; i ++ {
		if i == round - 1 {
			c.Write(data[0: lastDataLength])
			break
		}
		c.Write(data[0: sendDataLength])
	}
	c.Close()
	return nil
}

func main() {
	// 产生100mb的数据为了测试流的性能
	data = []byte(runner.GetRandomString(100 * 1024 * 1024))

	// start pprof server
	go func() {
		perf.ServeMonitor(fmt.Sprintf(":%d", port+10000))
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %v", err)
		}
		go StreamTest(conn)
	}
}
