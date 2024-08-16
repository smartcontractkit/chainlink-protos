// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package node

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	// DisableNode marks a node as disabled, disabling any active operations on
	// it.
	DisableNode(ctx context.Context, in *DisableNodeRequest, opts ...grpc.CallOption) (*DisableNodeResponse, error)
	// EnableNode enabled a disabled node, allowing operations to resume.
	EnableNode(ctx context.Context, in *EnableNodeRequest, opts ...grpc.CallOption) (*EnableNodeResponse, error)
	// GetNode retrieves the details of a node by its unique identifier.
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	// ListNodes returns a list of nodes, optionally filtered by the provided
	// criteria.
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	ListNodeChainConfigs(ctx context.Context, in *ListNodeChainConfigsRequest, opts ...grpc.CallOption) (*ListNodeChainConfigsResponse, error)
	// RegisterNode registers a new node to the system.
	RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
	// UpdateNode updates the details of an existing node.
	UpdateNode(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) DisableNode(ctx context.Context, in *DisableNodeRequest, opts ...grpc.CallOption) (*DisableNodeResponse, error) {
	out := new(DisableNodeResponse)
	err := c.cc.Invoke(ctx, "/tooling.distributor.node.NodeService/DisableNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) EnableNode(ctx context.Context, in *EnableNodeRequest, opts ...grpc.CallOption) (*EnableNodeResponse, error) {
	out := new(EnableNodeResponse)
	err := c.cc.Invoke(ctx, "/tooling.distributor.node.NodeService/EnableNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, "/tooling.distributor.node.NodeService/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := c.cc.Invoke(ctx, "/tooling.distributor.node.NodeService/ListNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ListNodeChainConfigs(ctx context.Context, in *ListNodeChainConfigsRequest, opts ...grpc.CallOption) (*ListNodeChainConfigsResponse, error) {
	out := new(ListNodeChainConfigsResponse)
	err := c.cc.Invoke(ctx, "/tooling.distributor.node.NodeService/ListNodeChainConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) RegisterNode(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, "/tooling.distributor.node.NodeService/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) UpdateNode(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error) {
	out := new(UpdateNodeResponse)
	err := c.cc.Invoke(ctx, "/tooling.distributor.node.NodeService/UpdateNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	// DisableNode marks a node as disabled, disabling any active operations on
	// it.
	DisableNode(context.Context, *DisableNodeRequest) (*DisableNodeResponse, error)
	// EnableNode enabled a disabled node, allowing operations to resume.
	EnableNode(context.Context, *EnableNodeRequest) (*EnableNodeResponse, error)
	// GetNode retrieves the details of a node by its unique identifier.
	GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	// ListNodes returns a list of nodes, optionally filtered by the provided
	// criteria.
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	ListNodeChainConfigs(context.Context, *ListNodeChainConfigsRequest) (*ListNodeChainConfigsResponse, error)
	// RegisterNode registers a new node to the system.
	RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
	// UpdateNode updates the details of an existing node.
	UpdateNode(context.Context, *UpdateNodeRequest) (*UpdateNodeResponse, error)
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) DisableNode(context.Context, *DisableNodeRequest) (*DisableNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableNode not implemented")
}
func (UnimplementedNodeServiceServer) EnableNode(context.Context, *EnableNodeRequest) (*EnableNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableNode not implemented")
}
func (UnimplementedNodeServiceServer) GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (UnimplementedNodeServiceServer) ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (UnimplementedNodeServiceServer) ListNodeChainConfigs(context.Context, *ListNodeChainConfigsRequest) (*ListNodeChainConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodeChainConfigs not implemented")
}
func (UnimplementedNodeServiceServer) RegisterNode(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (UnimplementedNodeServiceServer) UpdateNode(context.Context, *UpdateNodeRequest) (*UpdateNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNode not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_DisableNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).DisableNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tooling.distributor.node.NodeService/DisableNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).DisableNode(ctx, req.(*DisableNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_EnableNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).EnableNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tooling.distributor.node.NodeService/EnableNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).EnableNode(ctx, req.(*EnableNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tooling.distributor.node.NodeService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tooling.distributor.node.NodeService/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ListNodeChainConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeChainConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ListNodeChainConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tooling.distributor.node.NodeService/ListNodeChainConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ListNodeChainConfigs(ctx, req.(*ListNodeChainConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tooling.distributor.node.NodeService/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).RegisterNode(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).UpdateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tooling.distributor.node.NodeService/UpdateNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).UpdateNode(ctx, req.(*UpdateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tooling.distributor.node.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DisableNode",
			Handler:    _NodeService_DisableNode_Handler,
		},
		{
			MethodName: "EnableNode",
			Handler:    _NodeService_EnableNode_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _NodeService_GetNode_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _NodeService_ListNodes_Handler,
		},
		{
			MethodName: "ListNodeChainConfigs",
			Handler:    _NodeService_ListNodeChainConfigs_Handler,
		},
		{
			MethodName: "RegisterNode",
			Handler:    _NodeService_RegisterNode_Handler,
		},
		{
			MethodName: "UpdateNode",
			Handler:    _NodeService_UpdateNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job-distributor/node/node.proto",
}
