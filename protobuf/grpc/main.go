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
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	grpcg "github.com/bbbearxyz/kitex-benchmark/codec/protobuf/grpc_gen"
	"github.com/bbbearxyz/kitex-benchmark/perf"
	"github.com/bbbearxyz/kitex-benchmark/runner"
)

const (
	port = 8000
)

var data string
var recorder = perf.NewRecorder("GRPC@Server")

type server struct {
	grpcg.UnimplementedEchoServer
}

func (s *server) Send(ctx context.Context, req *grpcg.Request) (*grpcg.Response, error) {
	time.Sleep(time.Duration(req.Time) * time.Millisecond)
	// 正常只需要返回一个空的msg
	resp := runner.ProcessRequest(recorder, req.Action, "")

	return &grpcg.Response{
		Msg:    resp.Msg,
		Action: resp.Action,
	}, nil
}

func (s *server) StreamTest(stream grpcg.Echo_StreamTestServer) error {
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
			stream.Send(&grpcg.Response{Msg: data[0: lastDataLength], IsEnd: true})
			break
		}
		stream.Send(&grpcg.Response{Msg: data[0: sendDataLength], IsEnd: false})
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 1024),
		grpc.MaxSendMsgSize(1024 * 1024 * 1024),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time: 10,
			Timeout: 3,
		}),
		grpc.InitialWindowSize(1024 * 1024 * 1024),
		grpc.InitialConnWindowSize(1024 * 1024 * 1024),
		grpc.WriteBufferSize(32 * 1024 * 1024),
		grpc.ReadBufferSize(32 * 1024 * 1024),
	}

	s := grpc.NewServer(opts...)
	grpcg.RegisterEchoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
