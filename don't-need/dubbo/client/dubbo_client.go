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
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	dubbo "github.com/bbbearxyz/kitex-benchmark/codec/protobuf/dubbo_gen"
	"github.com/bbbearxyz/kitex-benchmark/runner"
	"sync"
	"time"
)


func NewPBDubboClient(opt *runner.Options) runner.Client {
	cli := &pbDubboClient{}
	cli.client = new(dubbo.EchoClientImpl)

	config.SetConsumerService(cli.client)

	// start dubbo-go framework with configuration
	if err := config.Load(); err != nil{
		panic(err)
	}

	cli.reqPool = &sync.Pool{
		New: func() interface{} {
			return &dubbo.Request{}
		},
	}
	return cli
}

type pbDubboClient struct {
	reqPool  *sync.Pool
	client   *dubbo.EchoClientImpl
}

func (cli *pbDubboClient) Echo(action, msg string, field, latency, payload, isStream int64) error {
	ctx := context.Background()
	req := cli.reqPool.Get().(*dubbo.Request)
	defer cli.reqPool.Put(req)

	req.Action = action
	req.Time = latency

	if req.Action == runner.EchoAction{
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

	pbcli := cli.client
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	var err error
	var reply *dubbo.Response
	if isStream == 1 {
		stream, _ := pbcli.StreamTest(ctx)
		req.Length = payload
		stream.Send(req)
		for true {
			res, err := stream.Recv()
			if err != nil {
				println(err.Error())
			}
			if res.IsEnd {
				break
			}
		}
		stream.CloseSend()
	} else {
		reply, err = pbcli.Send(ctx, req)

		if reply != nil {
			runner.ProcessResponse(reply.Action, reply.Msg)
		}
	}
	return err
}
