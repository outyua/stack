// Code generated by protoc-gen-stack. DO NOT EDIT.
// source: test.proto

package test

import (
	fmt "fmt"
	math "math"

	context "context"

	proto "github.com/golang/protobuf/proto"

	api "github.com/stack-labs/stack/api"

	client "github.com/stack-labs/stack/client"

	server "github.com/stack-labs/stack/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Test service

func NewTestEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Test service

type TestService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type testService struct {
	c    client.Client
	name string
}

func NewTestService(name string, c client.Client) TestService {
	return &testService{
		c:    c,
		name: name,
	}
}

func (c *testService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Test.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Test service

type TestHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterTestHandler(s server.Server, hdlr TestHandler, opts ...server.HandlerOption) error {
	type test interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type Test struct {
		test
	}
	h := &testHandler{hdlr}
	return s.Handle(s.NewHandler(&Test{h}, opts...))
}

type testHandler struct {
	TestHandler
}

func (h *testHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.TestHandler.Call(ctx, in, out)
}
