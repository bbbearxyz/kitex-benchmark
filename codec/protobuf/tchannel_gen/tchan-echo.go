// @generated Code generated by thrift-gen. Do not modify.

// Package hello_world is generated code used to make or handle TChannel calls using Thrift.
package tchannel_gen

import (
	"fmt"
	tchannel_go "github.com/bbbearxyz/another-tchannel-go"
	"github.com/bbbearxyz/another-tchannel-go/pb"
)


type EchoClient interface {
	Send(ctx pb.Context, arg *Request) (*Response, error)
	StreamTest(ctx pb.Context) (Echo_StreamTest_Client, error)
}

type echoClient struct {
	pbService string
	client    pb.TChanClient
}

func newEchoClient(pbService string, client pb.TChanClient) EchoClient {
	return &echoClient{
		pbService,
		client,
	}
}

func NewEchoClient(client pb.TChanClient) EchoClient {
	return newEchoClient("Echo", client)
}

func (c *echoClient) Send(ctx pb.Context, arg *Request) (*Response, error) {
	resp := new(Response)
	success, err := c.client.Call(ctx, c.pbService, "Send", arg, resp)
	if err != nil || !success {
		return nil, err
	}
	return resp, nil
}

func (c *echoClient) StreamTest(ctx pb.Context) (Echo_StreamTest_Client, error) {
	server := new(echoStreamTest_Client)
	success, err, reader, writer := c.client.CallStreaming(ctx, c.pbService, "StreamTest")
	if err != nil || !success {
		return nil, err
	}
	server.Context = ctx
	server.ArgReader = reader
	server.ArgWriter = writer
	return server, nil
}

type Echo_StreamTest_Client interface {
	tchannel_go.ArgReader
	tchannel_go.ArgWriter
	pb.Context
	Send(*Request) error
	Recv() (*Response, error)
	Close() error
}

type echoStreamTest_Client struct {
	tchannel_go.ArgReader
	tchannel_go.ArgWriter
	pb.Context
}

func (c *echoStreamTest_Client) Send(arg *Request) error {
	err := pb.WriteStruct(c.ArgWriter, arg)
	c.ArgWriter.Flush()
	return err
}

func (c *echoStreamTest_Client) Recv() (*Response, error) {
	resp := new(Response)
	err := pb.ReadStruct(c.ArgReader, resp)
	return resp, err
}

func (c *echoStreamTest_Client) Close() error {
	c.ArgReader.Close()
	c.ArgWriter.Close()
	return nil
}

type EchoServer interface {
	Send(ctx pb.Context, arg *Request) (*Response, error)
	StreamTest(server Echo_StreamTest_Server) error
}

type echoServer struct {
	handler EchoServer
}

func NewEchoServer(handler EchoServer) pb.TChanServer {
	return &echoServer{
		handler,
	}
}

func (s *echoServer) Service() string {
	return "Echo"
}

func (s *echoServer) Methods() []string {
	return []string{
		"Send",
		"StreamTest",
	}
}

func (s *echoServer) Types() []bool {
	return []bool{
		true,
		false,
	}
}

func (s *echoServer) Handle(ctx pb.Context, methodName string, reader tchannel_go.ArgReader, writer tchannel_go.ArgWriter) (bool, pb.PbStruct, error) {
	switch methodName {
	case "Send":
		return s.handle_Echo_Send(ctx, reader)
	case "StreamTest":
		return s.handle_Echo_StreamTest(reader, writer)
	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *echoServer) handle_Echo_Send(ctx pb.Context, reader tchannel_go.ArgReader) (bool, pb.PbStruct, error) {
	req := new(Request)
	if err := pb.ReadStruct(reader, req); err != nil {
		return false, nil, err
	}
	res, err := s.handler.Send(ctx, req)
	if err != nil {
		return false, nil, err
	}
	return err == nil, res, nil
}
func (s *echoServer) handle_Echo_StreamTest(reader tchannel_go.ArgReader, writer tchannel_go.ArgWriter) (bool, pb.PbStruct, error) {
	server := new(echoStreamTest_Server)
	server.ArgReader = reader
	server.ArgWriter = writer
	err := s.handler.StreamTest(server)
	if err != nil {
		return false, nil, err
	}
	return true, nil, nil
}

type Echo_StreamTest_Server interface {
	tchannel_go.ArgReader
	tchannel_go.ArgWriter
	pb.Context
	Send(*Response) error
	Recv() (*Request, error)
	Close() error
}

type echoStreamTest_Server struct {
	tchannel_go.ArgReader
	tchannel_go.ArgWriter
	pb.Context
}

func (c *echoStreamTest_Server) Send(arg *Response) error {
	err := pb.WriteStruct(c.ArgWriter, arg)
	c.ArgWriter.Flush()
	return err
}

func (c *echoStreamTest_Server) Recv() (*Request, error) {
	resp := new(Request)
	err := pb.ReadStruct(c.ArgReader, resp)
	return resp, err
}

func (c *echoStreamTest_Server) Close() error {
	c.ArgReader.Close()
	c.ArgWriter.Close()
	return nil
}