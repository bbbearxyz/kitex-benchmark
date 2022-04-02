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
	"github.com/bbbearxyz/kitex-benchmark/runner"
	"net"
	"sync"
)

func NewPBGrpcClient(opt *runner.Options) runner.Client {
	cli := &tcpClient{}
	//cli.reqPool = &sync.Pool{
	//	New: func() interface{} {
	//		return &grpcg.Request{}
	//	},
	//}
	//cli.connpool = runner.NewPool(func() interface{} {
	//	// Set up a connection to the server.
	//	// 配置参数
	//	conn, err := grpc.Dial(opt.Address, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100 * 1024 * 1024)))
	//	if err != nil {
	//		log.Fatalf("did not connect: %v", err)
	//	}
	//	return grpcg.NewEchoClient(conn)
	//}, opt.PoolSize)
	cli.address = opt.Address
	return cli
}

type tcpClient struct {
	reqPool  *sync.Pool
	connpool *runner.Pool
	address   string
}

func (cli *tcpClient) Echo(action, msg string, field, latency, payload, isStream int64) error {
	if isStream == 1 {
		if action != runner.EchoAction {
			return nil
		}
		conn, _ := net.Dial("tcp", cli.address)
		buf := make([]byte, 10 * 1024)
		conn.Write([]byte(msg))
		num := 0
		for num != 1024 * 1024 * 1024 {
			delta, _ := conn.Read(buf)
			num += delta
		}
		conn.Close()
		return nil
	}
	//ctx := context.Background()
	//req := cli.reqPool.Get().(*grpcg.Request)
	//defer cli.reqPool.Put(req)
	//
	//req.Action = action
	//req.Time = latency
	//
	//if req.Action == runner.EchoAction{
	//	if field == 1 || isStream == 1 {
	//		req.Field1 = msg
	//	} else if field == 5 {
	//		averageLen := (payload) / field
	//		req.Field1 = msg[0: averageLen]
	//		req.Field2 = msg[averageLen: 2 * averageLen]
	//		req.Field3 = msg[averageLen * 2: 3 * averageLen]
	//		req.Field4 = msg[averageLen * 3: 4 * averageLen]
	//		req.Field5 = msg[averageLen * 4:]
	//	} else if field == 10 {
	//		averageLen := (payload) / field
	//		req.Field1 = msg[0: averageLen]
	//		req.Field2 = msg[averageLen: 2 * averageLen]
	//		req.Field3 = msg[averageLen * 2: 3 * averageLen]
	//		req.Field4 = msg[averageLen * 3: 4 * averageLen]
	//		req.Field5 = msg[averageLen * 4: 5 * averageLen]
	//		req.Field6 = msg[averageLen * 5: 6 * averageLen]
	//		req.Field7 = msg[averageLen * 6: 7 * averageLen]
	//		req.Field8 = msg[averageLen * 7: 8 * averageLen]
	//		req.Field9 = msg[averageLen * 8: 9 * averageLen]
	//		req.Field10 = msg[averageLen * 9:]
	//	}
	//}
	//
	//pbcli := cli.connpool.Get().(grpcg.EchoClient)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//var err error
	//var reply *grpcg.Response
	//reply, err = pbcli.Send(ctx, req)
	//
	//if reply != nil {
	//	runner.ProcessResponse(reply.Action, reply.Msg)
	//}
	//return err
	return nil
}
