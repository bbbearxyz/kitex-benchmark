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
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/cloudwego/kitex/server"

	"github.com/bbbearxyz/kitex-benchmark/codec/protobuf/kitex_gen/echo"
	echosvr "github.com/bbbearxyz/kitex-benchmark/codec/protobuf/kitex_gen/echo/echo"
	"github.com/bbbearxyz/kitex-benchmark/perf"
	"github.com/bbbearxyz/kitex-benchmark/runner"
)

const (
	port = 8001
)

var data string
var recorder = perf.NewRecorder("KITEX@Server")

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Send(ctx context.Context, req *echo.Request) (*echo.Response, error) {
	time.Sleep(time.Duration(req.Time) * time.Millisecond)
	resp := runner.ProcessRequest(recorder, req.Action, "")

	return &echo.Response{
		Action: resp.Action,
		Msg:    resp.Msg,
	}, nil
}

func (s *EchoImpl) StreamTest(stream echo.Echo_StreamTestServer) (err error) {
	// 计算1GB / length的次数
	req, _ := stream.Recv()
	length := req.Length
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
			stream.Send(&echo.Response{Msg: data[0: lastDataLength], IsEnd: true})
			break
		}
		stream.Send(&echo.Response{Msg: data[0: sendDataLength], IsEnd: false})
	}
	return nil
}

func (s *EchoImpl) TCPCostTest(stream echo.Echo_TCPCostTestServer) error {
	// 计算1GB / length的次数
	req, _ := stream.Recv()
	length := req.Length
	round := int64(100)
	sendDataLength := int64(length)

	for i := int64(0); i < round; i ++ {
		if i == round - 1 {
			stream.Send(&echo.Response{Msg: data[0: sendDataLength], IsEnd: true})
			break
		}
		stream.Send(&echo.Response{Msg: data[0: sendDataLength], IsEnd: false})
	}
	return nil
}


func main() {
	// 产生100mb的数据为了测试流的性能
	data = runner.GetRandomString(100 * 1024 * 1024)
	// start pprof server
	go func() {
		perf.ServeMonitor(fmt.Sprintf(":%d", port+10000))
	}()

	address := &net.UnixAddr{Net: "tcp", Name: fmt.Sprintf(":%d", port)}
	svr := echosvr.NewServer(new(EchoImpl), server.WithServiceAddr(address))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
