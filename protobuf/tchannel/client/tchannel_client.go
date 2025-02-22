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
	tchannel "github.com/bbbearxyz/another-tchannel-go"
	"github.com/bbbearxyz/another-tchannel-go/pb"
	"github.com/bbbearxyz/kitex-benchmark/codec/protobuf/tchannel_gen"
	"sync"
	"time"

	"github.com/bbbearxyz/kitex-benchmark/runner"
)

func NewPBTchannelClient(opt *runner.Options) runner.Client {
	cli := &pbGrpcClient{}
	cli.reqPool = &sync.Pool{
		New: func() interface{} {
			return &tchannel_gen.Request{}
		},
	}
	tchan, _ := tchannel.NewChannel("client1", nil)
	tchan.Peers().Add(opt.Address)
	cli.connpool = runner.NewPool(func() interface{} {
		// Set up a connection to the server.
		// 配置参数
		tclient := pb.NewClient(tchan, "server", nil)

		return tchannel_gen.NewEchoClient(tclient)
	}, opt.PoolSize)
	return cli
}

type pbGrpcClient struct {
	reqPool  *sync.Pool
	connpool *runner.Pool
}

func (cli *pbGrpcClient) Echo(action, msg string, field, latency, payload, isStream, isTCPCostTest int64) error {
	req := cli.reqPool.Get().(*tchannel_gen.Request)
	defer cli.reqPool.Put(req)

	req.Action = action
	req.Time = latency

	if req.Action == runner.EchoAction {
		if field == 1 || isStream == 1 {
			req.Field1 = msg
		} else if field == 5 {
			averageLen := (payload) / field
			req.Field1 = msg[0: averageLen]
			req.Field2 = msg[averageLen: 2 * averageLen]
			req.Field3 = msg[averageLen * 2: 3 * averageLen]
			req.Field4 = msg[averageLen * 3: 4 * averageLen]
			req.Field5 = msg[averageLen * 4:]
		} else if field == 10 {
			averageLen := (payload) / field
			req.Field1 = msg[0: averageLen]
			req.Field2 = msg[averageLen: 2 * averageLen]
			req.Field3 = msg[averageLen * 2: 3 * averageLen]
			req.Field4 = msg[averageLen * 3: 4 * averageLen]
			req.Field5 = msg[averageLen * 4: 5 * averageLen]
			req.Field6 = msg[averageLen * 5: 6 * averageLen]
			req.Field7 = msg[averageLen * 6: 7 * averageLen]
			req.Field8 = msg[averageLen * 7: 8 * averageLen]
			req.Field9 = msg[averageLen * 8: 9 * averageLen]
			req.Field10 = msg[averageLen * 9:]
		}
	}

	pbcli := cli.connpool.Get().(tchannel_gen.EchoClient)
	ctx, cancel := pb.NewContext(time.Hour)
	defer cancel()
	var err error
	var reply *tchannel_gen.Response
	if isStream == 1 {
		stream, _ := pbcli.StreamTest(ctx)
		req.Length = payload
		// 目前实现的bug
		stream.Send(req)
		stream.Recv()
		for true {
			res, err := stream.Recv()
			if err != nil {
				println(err.Error())
			}
			if res.IsEnd {
				break
			}
		}
		stream.Close()
	} else {
		reply, err = pbcli.Send(ctx, req)
		if reply != nil {
			runner.ProcessResponse(reply.Action, reply.Msg)
		}
	}
	return err
}
