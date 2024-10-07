// Code generated by protoc-gen-go-wsrpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-wsrpc v0.0.1
// - protoc             v5.28.1

package node

import (
	context "context"
	wsrpc "github.com/smartcontractkit/wsrpc"
)

// NodeServiceClient is the client API for NodeService service.
type NodeServiceClient interface {
	// ProposeJob is called by the JD to propose a job to the node.
	ProposeJob(ctx context.Context, in *ProposeJobRequest) (*ProposeJobResponse, error)
	// DeleteJob is called by the JD to delete a job from the node.
	DeleteJob(ctx context.Context, in *DeleteJobRequest) (*DeleteJobResponse, error)
	// RevokeJob is called by the JD to revoke a job from the node.
	RevokeJob(ctx context.Context, in *RevokeJobRequest) (*RevokeJobResponse, error)
}

type nodeServiceClient struct {
	cc wsrpc.ClientInterface
}

func NewNodeServiceClient(cc wsrpc.ClientInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) ProposeJob(ctx context.Context, in *ProposeJobRequest) (*ProposeJobResponse, error) {
	out := new(ProposeJobResponse)
	err := c.cc.Invoke(ctx, "ProposeJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) DeleteJob(ctx context.Context, in *DeleteJobRequest) (*DeleteJobResponse, error) {
	out := new(DeleteJobResponse)
	err := c.cc.Invoke(ctx, "DeleteJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) RevokeJob(ctx context.Context, in *RevokeJobRequest) (*RevokeJobResponse, error) {
	out := new(RevokeJobResponse)
	err := c.cc.Invoke(ctx, "RevokeJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	// ProposeJob is called by the JD to propose a job to the node.
	ProposeJob(context.Context, *ProposeJobRequest) (*ProposeJobResponse, error)
	// DeleteJob is called by the JD to delete a job from the node.
	DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobResponse, error)
	// RevokeJob is called by the JD to revoke a job from the node.
	RevokeJob(context.Context, *RevokeJobRequest) (*RevokeJobResponse, error)
}

func RegisterNodeServiceServer(s wsrpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_ProposeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProposeJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(NodeServiceServer).ProposeJob(ctx, in)
}

func _NodeService_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(NodeServiceServer).DeleteJob(ctx, in)
}

func _NodeService_RevokeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RevokeJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(NodeServiceServer).RevokeJob(ctx, in)
}

// NodeService_ServiceDesc is the wsrpc.ServiceDesc for NodeService service.
// It's only intended for direct use with wsrpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = wsrpc.ServiceDesc{
	ServiceName: "node.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []wsrpc.MethodDesc{
		{
			MethodName: "ProposeJob",
			Handler:    _NodeService_ProposeJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _NodeService_DeleteJob_Handler,
		},
		{
			MethodName: "RevokeJob",
			Handler:    _NodeService_RevokeJob_Handler,
		},
	},
}
