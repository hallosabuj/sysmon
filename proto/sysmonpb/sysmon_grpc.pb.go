// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sysmonpb

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

// SysmonServiceClient is the client API for SysmonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysmonServiceClient interface {
	Dong(ctx context.Context, in *DongReq, opts ...grpc.CallOption) (*DongRes, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Pong(ctx context.Context, in *PongRequest, opts ...grpc.CallOption) (*PongResponse, error)
	AddRule(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*IPResponse, error)
	DelRule(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*IPResponse, error)
	AddRoute(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*IPResponse, error)
	DelRoute(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*IPResponse, error)
	Rules(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*RuleRespone, error)
	Routes(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*RouteRespone, error)
	RoutesByTableName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RouteRespone, error)
	InterfaceAddresses(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*InterfaceAddressesResponse, error)
	Interfaces(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*InterfacesResponse, error)
	InterfaceDetailsByName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*InterfaceDetailsResponse, error)
	IpTables(ctx context.Context, in *Request, opts ...grpc.CallOption) (*IPTablesResponse, error)
	AddTable(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*Response, error)
	ListAllocatedIp(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*AllocatedIPsResponse, error)
	ListAllocatedIpGTP(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*AllocatedIPsResponse, error)
}

type sysmonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSysmonServiceClient(cc grpc.ClientConnInterface) SysmonServiceClient {
	return &sysmonServiceClient{cc}
}

func (c *sysmonServiceClient) Dong(ctx context.Context, in *DongReq, opts ...grpc.CallOption) (*DongRes, error) {
	out := new(DongRes)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/Dong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) Pong(ctx context.Context, in *PongRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) AddRule(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*IPResponse, error) {
	out := new(IPResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/AddRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) DelRule(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*IPResponse, error) {
	out := new(IPResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/DelRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) AddRoute(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*IPResponse, error) {
	out := new(IPResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/AddRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) DelRoute(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*IPResponse, error) {
	out := new(IPResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/DelRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) Rules(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*RuleRespone, error) {
	out := new(RuleRespone)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/Rules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) Routes(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*RouteRespone, error) {
	out := new(RouteRespone)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/Routes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) RoutesByTableName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RouteRespone, error) {
	out := new(RouteRespone)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/RoutesByTableName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) InterfaceAddresses(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*InterfaceAddressesResponse, error) {
	out := new(InterfaceAddressesResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/InterfaceAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) Interfaces(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*InterfacesResponse, error) {
	out := new(InterfacesResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/Interfaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) InterfaceDetailsByName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*InterfaceDetailsResponse, error) {
	out := new(InterfaceDetailsResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/InterfaceDetailsByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) IpTables(ctx context.Context, in *Request, opts ...grpc.CallOption) (*IPTablesResponse, error) {
	out := new(IPTablesResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/IpTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) AddTable(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/AddTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) ListAllocatedIp(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*AllocatedIPsResponse, error) {
	out := new(AllocatedIPsResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/ListAllocatedIp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysmonServiceClient) ListAllocatedIpGTP(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*AllocatedIPsResponse, error) {
	out := new(AllocatedIPsResponse)
	err := c.cc.Invoke(ctx, "/sysmonpb.SysmonService/ListAllocatedIpGTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysmonServiceServer is the server API for SysmonService service.
// All implementations must embed UnimplementedSysmonServiceServer
// for forward compatibility
type SysmonServiceServer interface {
	Dong(context.Context, *DongReq) (*DongRes, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Pong(context.Context, *PongRequest) (*PongResponse, error)
	AddRule(context.Context, *IPRequest) (*IPResponse, error)
	DelRule(context.Context, *IPRequest) (*IPResponse, error)
	AddRoute(context.Context, *IPRequest) (*IPResponse, error)
	DelRoute(context.Context, *IPRequest) (*IPResponse, error)
	Rules(context.Context, *IPRequest) (*RuleRespone, error)
	Routes(context.Context, *IPRequest) (*RouteRespone, error)
	RoutesByTableName(context.Context, *Request) (*RouteRespone, error)
	InterfaceAddresses(context.Context, *IPRequest) (*InterfaceAddressesResponse, error)
	Interfaces(context.Context, *IPRequest) (*InterfacesResponse, error)
	InterfaceDetailsByName(context.Context, *Request) (*InterfaceDetailsResponse, error)
	IpTables(context.Context, *Request) (*IPTablesResponse, error)
	AddTable(context.Context, *IPRequest) (*Response, error)
	ListAllocatedIp(context.Context, *IPRequest) (*AllocatedIPsResponse, error)
	ListAllocatedIpGTP(context.Context, *IPRequest) (*AllocatedIPsResponse, error)
	mustEmbedUnimplementedSysmonServiceServer()
}

// UnimplementedSysmonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSysmonServiceServer struct {
}

func (UnimplementedSysmonServiceServer) Dong(context.Context, *DongReq) (*DongRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dong not implemented")
}
func (UnimplementedSysmonServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSysmonServiceServer) Pong(context.Context, *PongRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (UnimplementedSysmonServiceServer) AddRule(context.Context, *IPRequest) (*IPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRule not implemented")
}
func (UnimplementedSysmonServiceServer) DelRule(context.Context, *IPRequest) (*IPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelRule not implemented")
}
func (UnimplementedSysmonServiceServer) AddRoute(context.Context, *IPRequest) (*IPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoute not implemented")
}
func (UnimplementedSysmonServiceServer) DelRoute(context.Context, *IPRequest) (*IPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelRoute not implemented")
}
func (UnimplementedSysmonServiceServer) Rules(context.Context, *IPRequest) (*RuleRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rules not implemented")
}
func (UnimplementedSysmonServiceServer) Routes(context.Context, *IPRequest) (*RouteRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Routes not implemented")
}
func (UnimplementedSysmonServiceServer) RoutesByTableName(context.Context, *Request) (*RouteRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoutesByTableName not implemented")
}
func (UnimplementedSysmonServiceServer) InterfaceAddresses(context.Context, *IPRequest) (*InterfaceAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InterfaceAddresses not implemented")
}
func (UnimplementedSysmonServiceServer) Interfaces(context.Context, *IPRequest) (*InterfacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Interfaces not implemented")
}
func (UnimplementedSysmonServiceServer) InterfaceDetailsByName(context.Context, *Request) (*InterfaceDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InterfaceDetailsByName not implemented")
}
func (UnimplementedSysmonServiceServer) IpTables(context.Context, *Request) (*IPTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IpTables not implemented")
}
func (UnimplementedSysmonServiceServer) AddTable(context.Context, *IPRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTable not implemented")
}
func (UnimplementedSysmonServiceServer) ListAllocatedIp(context.Context, *IPRequest) (*AllocatedIPsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllocatedIp not implemented")
}
func (UnimplementedSysmonServiceServer) ListAllocatedIpGTP(context.Context, *IPRequest) (*AllocatedIPsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllocatedIpGTP not implemented")
}
func (UnimplementedSysmonServiceServer) mustEmbedUnimplementedSysmonServiceServer() {}

// UnsafeSysmonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysmonServiceServer will
// result in compilation errors.
type UnsafeSysmonServiceServer interface {
	mustEmbedUnimplementedSysmonServiceServer()
}

func RegisterSysmonServiceServer(s grpc.ServiceRegistrar, srv SysmonServiceServer) {
	s.RegisterService(&SysmonService_ServiceDesc, srv)
}

func _SysmonService_Dong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DongReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).Dong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/Dong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).Dong(ctx, req.(*DongReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).Pong(ctx, req.(*PongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_AddRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).AddRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/AddRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).AddRule(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_DelRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).DelRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/DelRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).DelRule(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_AddRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).AddRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/AddRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).AddRoute(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_DelRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).DelRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/DelRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).DelRoute(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_Rules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).Rules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/Rules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).Rules(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_Routes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).Routes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/Routes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).Routes(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_RoutesByTableName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).RoutesByTableName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/RoutesByTableName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).RoutesByTableName(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_InterfaceAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).InterfaceAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/InterfaceAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).InterfaceAddresses(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_Interfaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).Interfaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/Interfaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).Interfaces(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_InterfaceDetailsByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).InterfaceDetailsByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/InterfaceDetailsByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).InterfaceDetailsByName(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_IpTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).IpTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/IpTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).IpTables(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_AddTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).AddTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/AddTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).AddTable(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_ListAllocatedIp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).ListAllocatedIp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/ListAllocatedIp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).ListAllocatedIp(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysmonService_ListAllocatedIpGTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysmonServiceServer).ListAllocatedIpGTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysmonpb.SysmonService/ListAllocatedIpGTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysmonServiceServer).ListAllocatedIpGTP(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SysmonService_ServiceDesc is the grpc.ServiceDesc for SysmonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysmonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sysmonpb.SysmonService",
	HandlerType: (*SysmonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dong",
			Handler:    _SysmonService_Dong_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _SysmonService_Ping_Handler,
		},
		{
			MethodName: "Pong",
			Handler:    _SysmonService_Pong_Handler,
		},
		{
			MethodName: "AddRule",
			Handler:    _SysmonService_AddRule_Handler,
		},
		{
			MethodName: "DelRule",
			Handler:    _SysmonService_DelRule_Handler,
		},
		{
			MethodName: "AddRoute",
			Handler:    _SysmonService_AddRoute_Handler,
		},
		{
			MethodName: "DelRoute",
			Handler:    _SysmonService_DelRoute_Handler,
		},
		{
			MethodName: "Rules",
			Handler:    _SysmonService_Rules_Handler,
		},
		{
			MethodName: "Routes",
			Handler:    _SysmonService_Routes_Handler,
		},
		{
			MethodName: "RoutesByTableName",
			Handler:    _SysmonService_RoutesByTableName_Handler,
		},
		{
			MethodName: "InterfaceAddresses",
			Handler:    _SysmonService_InterfaceAddresses_Handler,
		},
		{
			MethodName: "Interfaces",
			Handler:    _SysmonService_Interfaces_Handler,
		},
		{
			MethodName: "InterfaceDetailsByName",
			Handler:    _SysmonService_InterfaceDetailsByName_Handler,
		},
		{
			MethodName: "IpTables",
			Handler:    _SysmonService_IpTables_Handler,
		},
		{
			MethodName: "AddTable",
			Handler:    _SysmonService_AddTable_Handler,
		},
		{
			MethodName: "ListAllocatedIp",
			Handler:    _SysmonService_ListAllocatedIp_Handler,
		},
		{
			MethodName: "ListAllocatedIpGTP",
			Handler:    _SysmonService_ListAllocatedIpGTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sysmon.proto",
}
