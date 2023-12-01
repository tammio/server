// Copyright 2020-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: buf/alpha/registry/v1alpha1/studio_request.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StudioRequestService_CreateStudioRequest_FullMethodName = "/buf.alpha.registry.v1alpha1.StudioRequestService/CreateStudioRequest"
	StudioRequestService_RenameStudioRequest_FullMethodName = "/buf.alpha.registry.v1alpha1.StudioRequestService/RenameStudioRequest"
	StudioRequestService_DeleteStudioRequest_FullMethodName = "/buf.alpha.registry.v1alpha1.StudioRequestService/DeleteStudioRequest"
	StudioRequestService_ListStudioRequests_FullMethodName  = "/buf.alpha.registry.v1alpha1.StudioRequestService/ListStudioRequests"
)

// StudioRequestServiceClient is the client API for StudioRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudioRequestServiceClient interface {
	// CreateStudioRequest registers a favorite Studio Requests to the caller's
	// BSR profile.
	CreateStudioRequest(ctx context.Context, in *CreateStudioRequestRequest, opts ...grpc.CallOption) (*CreateStudioRequestResponse, error)
	// RenameStudioRequest renames an existing Studio Request.
	RenameStudioRequest(ctx context.Context, in *RenameStudioRequestRequest, opts ...grpc.CallOption) (*RenameStudioRequestResponse, error)
	// DeleteStudioRequest removes a favorite Studio Request from the caller's BSR
	// profile.
	DeleteStudioRequest(ctx context.Context, in *DeleteStudioRequestRequest, opts ...grpc.CallOption) (*DeleteStudioRequestResponse, error)
	// ListStudioRequests shows the caller's favorited Studio Requests.
	ListStudioRequests(ctx context.Context, in *ListStudioRequestsRequest, opts ...grpc.CallOption) (*ListStudioRequestsResponse, error)
}

type studioRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudioRequestServiceClient(cc grpc.ClientConnInterface) StudioRequestServiceClient {
	return &studioRequestServiceClient{cc}
}

func (c *studioRequestServiceClient) CreateStudioRequest(ctx context.Context, in *CreateStudioRequestRequest, opts ...grpc.CallOption) (*CreateStudioRequestResponse, error) {
	out := new(CreateStudioRequestResponse)
	err := c.cc.Invoke(ctx, StudioRequestService_CreateStudioRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studioRequestServiceClient) RenameStudioRequest(ctx context.Context, in *RenameStudioRequestRequest, opts ...grpc.CallOption) (*RenameStudioRequestResponse, error) {
	out := new(RenameStudioRequestResponse)
	err := c.cc.Invoke(ctx, StudioRequestService_RenameStudioRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studioRequestServiceClient) DeleteStudioRequest(ctx context.Context, in *DeleteStudioRequestRequest, opts ...grpc.CallOption) (*DeleteStudioRequestResponse, error) {
	out := new(DeleteStudioRequestResponse)
	err := c.cc.Invoke(ctx, StudioRequestService_DeleteStudioRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studioRequestServiceClient) ListStudioRequests(ctx context.Context, in *ListStudioRequestsRequest, opts ...grpc.CallOption) (*ListStudioRequestsResponse, error) {
	out := new(ListStudioRequestsResponse)
	err := c.cc.Invoke(ctx, StudioRequestService_ListStudioRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudioRequestServiceServer is the server API for StudioRequestService service.
// All implementations should embed UnimplementedStudioRequestServiceServer
// for forward compatibility
type StudioRequestServiceServer interface {
	// CreateStudioRequest registers a favorite Studio Requests to the caller's
	// BSR profile.
	CreateStudioRequest(context.Context, *CreateStudioRequestRequest) (*CreateStudioRequestResponse, error)
	// RenameStudioRequest renames an existing Studio Request.
	RenameStudioRequest(context.Context, *RenameStudioRequestRequest) (*RenameStudioRequestResponse, error)
	// DeleteStudioRequest removes a favorite Studio Request from the caller's BSR
	// profile.
	DeleteStudioRequest(context.Context, *DeleteStudioRequestRequest) (*DeleteStudioRequestResponse, error)
	// ListStudioRequests shows the caller's favorited Studio Requests.
	ListStudioRequests(context.Context, *ListStudioRequestsRequest) (*ListStudioRequestsResponse, error)
}

// UnimplementedStudioRequestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStudioRequestServiceServer struct {
}

func (UnimplementedStudioRequestServiceServer) CreateStudioRequest(context.Context, *CreateStudioRequestRequest) (*CreateStudioRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudioRequest not implemented")
}
func (UnimplementedStudioRequestServiceServer) RenameStudioRequest(context.Context, *RenameStudioRequestRequest) (*RenameStudioRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameStudioRequest not implemented")
}
func (UnimplementedStudioRequestServiceServer) DeleteStudioRequest(context.Context, *DeleteStudioRequestRequest) (*DeleteStudioRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudioRequest not implemented")
}
func (UnimplementedStudioRequestServiceServer) ListStudioRequests(context.Context, *ListStudioRequestsRequest) (*ListStudioRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudioRequests not implemented")
}

// UnsafeStudioRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudioRequestServiceServer will
// result in compilation errors.
type UnsafeStudioRequestServiceServer interface {
	mustEmbedUnimplementedStudioRequestServiceServer()
}

func RegisterStudioRequestServiceServer(s grpc.ServiceRegistrar, srv StudioRequestServiceServer) {
	s.RegisterService(&StudioRequestService_ServiceDesc, srv)
}

func _StudioRequestService_CreateStudioRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudioRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudioRequestServiceServer).CreateStudioRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudioRequestService_CreateStudioRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudioRequestServiceServer).CreateStudioRequest(ctx, req.(*CreateStudioRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudioRequestService_RenameStudioRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameStudioRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudioRequestServiceServer).RenameStudioRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudioRequestService_RenameStudioRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudioRequestServiceServer).RenameStudioRequest(ctx, req.(*RenameStudioRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudioRequestService_DeleteStudioRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudioRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudioRequestServiceServer).DeleteStudioRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudioRequestService_DeleteStudioRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudioRequestServiceServer).DeleteStudioRequest(ctx, req.(*DeleteStudioRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudioRequestService_ListStudioRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStudioRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudioRequestServiceServer).ListStudioRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudioRequestService_ListStudioRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudioRequestServiceServer).ListStudioRequests(ctx, req.(*ListStudioRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudioRequestService_ServiceDesc is the grpc.ServiceDesc for StudioRequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudioRequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.StudioRequestService",
	HandlerType: (*StudioRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudioRequest",
			Handler:    _StudioRequestService_CreateStudioRequest_Handler,
		},
		{
			MethodName: "RenameStudioRequest",
			Handler:    _StudioRequestService_RenameStudioRequest_Handler,
		},
		{
			MethodName: "DeleteStudioRequest",
			Handler:    _StudioRequestService_DeleteStudioRequest_Handler,
		},
		{
			MethodName: "ListStudioRequests",
			Handler:    _StudioRequestService_ListStudioRequests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/studio_request.proto",
}