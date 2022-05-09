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
	"fmt"
	tchannel "github.com/bbbearxyz/another-tchannel-go"
	"github.com/bbbearxyz/another-tchannel-go/pb"
	"github.com/bbbearxyz/kitex-benchmark/codec/protobuf/tchannel_gen"
	"net"
	"time"

	"github.com/bbbearxyz/kitex-benchmark/perf"
	"github.com/bbbearxyz/kitex-benchmark/runner"
)

const (
	port = 8004
)

var data string
var recorder = perf.NewRecorder("GRPC@Server")

type server struct {

}

func (s *server) Send(ctx pb.Context, req *tchannel_gen.Request) (*tchannel_gen.Response, error) {
	time.Sleep(time.Duration(req.Time) * time.Millisecond)
	// 正常只需要返回一个空的msg
	resp := runner.ProcessRequest(recorder, req.Action, "")
	return &tchannel_gen.Response{
		Msg:    resp.Msg,
		Action: resp.Action,
	}, nil
}

func (s *server) StreamTest(stream tchannel_gen.Echo_StreamTest_Server) error {
	// 计算1GB / length的次数
	// 空的msg
	stream.Send(&tchannel_gen.Response{Msg: ""})
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
			stream.Send(&tchannel_gen.Response{Msg: data[0: lastDataLength], IsEnd: true})
			break
		}
		stream.Send(&tchannel_gen.Response{Msg: data[0: sendDataLength], IsEnd: false})
	}
	stream.Close()
	return nil
}

func main() {
	// 产生100mb的数据为了测试流的性能
	data = runner.GetRandomString(100 * 1024 * 1024)

	// start pprof server
	go func() {
		perf.ServeMonitor(fmt.Sprintf(":%d", port+10000))
	}()


	tchan, _ := tchannel.NewChannel("server", nil)

	listener, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))


	ser := pb.NewServer(tchan)
	ser.Register(tchannel_gen.NewEchoServer(&server{}))

	// Serve will set the local peer info, and start accepting sockets in a separate goroutine.
	tchan.Serve(listener)

	for {

	}
}
