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
	"dubbo.apache.org/dubbo-go/v3/config"
	"fmt"
	dubbo "github.com/bbbearxyz/kitex-benchmark/codec/protobuf/dubbo_gen"
	"github.com/bbbearxyz/kitex-benchmark/perf"
	"github.com/bbbearxyz/kitex-benchmark/runner"
	"time"

	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

const (
	port = 8004
)

var data string
var recorder = perf.NewRecorder("DUBBO@Server")

type Server struct {
	dubbo.UnimplementedEchoServer
}

func (s *Server) Send(ctx context.Context, req *dubbo.Request) (*dubbo.Response, error) {
	time.Sleep(time.Duration(req.Time) * time.Millisecond)
	// 正常只需要返回一个空的msg
	resp := runner.ProcessRequest(recorder, req.Action, "")

	return &dubbo.Response{
		Msg:    resp.Msg,
		Action: resp.Action,
	}, nil
}

func (s *Server) StreamTest(stream dubbo.Echo_StreamTestServer) error {
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
			stream.Send(&dubbo.Response{Msg: data[0: lastDataLength], IsEnd: true})
			break
		}
		stream.Send(&dubbo.Response{Msg: data[0: sendDataLength], IsEnd: false})
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

	config.SetProviderService(&Server{})
	// start dubbo-go framework with configuration
	if err := config.Load(); err != nil{
		panic(err)
	}

	select {}
}
