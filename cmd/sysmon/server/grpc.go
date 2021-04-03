package server

import (
	"context"
	"strconv"
	"strings"
	"sysmon/cmd/sysmon/api"
	"sysmon/proto/sysmonpb"
)

type GRPCServer struct {
	sysmonpb.UnimplementedSysmonServiceServer
}

func NewServer() *GRPCServer {
	return new(GRPCServer)
}

type MSG struct {
	n string
}

func (s *GRPCServer) Ping(ctxt context.Context, request *sysmonpb.PingRequest) (*sysmonpb.PingResponse, error) {
	temp := []string{"a", "s", "d", "f"}
	response := &sysmonpb.PingResponse{
		Body: &sysmonpb.PP{Name: temp},
	}
	return response, nil
}

func (s *GRPCServer) Pong(ctxt context.Context, request *sysmonpb.PongRequest) (*sysmonpb.PongResponse, error) {
	response := &sysmonpb.PongResponse{
		Body: request.Body,
	}
	return response, nil
}

func (s *GRPCServer) Dong(ctxt context.Context, request *sysmonpb.DongReq) (*sysmonpb.DongRes, error) {
	response := &sysmonpb.DongRes{
		Response: request.Id,
	}
	return response, nil
}

func (s *GRPCServer) AddRule(ctxt context.Context, request *sysmonpb.IPRequest) (*sysmonpb.IPResponse, error) {
	api.MakeSudo()
	msg := api.AddIPRule(request)
	response := &sysmonpb.IPResponse{
		Response: &sysmonpb.Response{Msg: msg},
	}
	return response, nil
}

func (s *GRPCServer) DelRule(ctxt context.Context, request *sysmonpb.IPRequest) (*sysmonpb.IPResponse, error) {
	api.MakeSudo()
	msg := api.DelIPRule(request)
	response := &sysmonpb.IPResponse{
		Response: &sysmonpb.Response{Msg: msg},
	}
	return response, nil
}

func (s *GRPCServer) AddRoute(ctxt context.Context, request *sysmonpb.IPRequest) (*sysmonpb.IPResponse, error) {
	api.MakeSudo()
	msg := api.AddIPRoute(request)
	response := &sysmonpb.IPResponse{
		Response: &sysmonpb.Response{Msg: msg},
	}
	return response, nil
}

func (s *GRPCServer) DelRoute(ctxt context.Context, request *sysmonpb.IPRequest) (*sysmonpb.IPResponse, error) {
	api.MakeSudo()
	msg := api.DelIPRoute(request)
	response := &sysmonpb.IPResponse{
		Response: &sysmonpb.Response{Msg: msg},
	}
	return response, nil
}

func (s *GRPCServer) Rules(ctxt context.Context, request *sysmonpb.IPRequest) (*sysmonpb.RuleRespone, error) {
	api.MakeSudo()
	msg := api.IPRules()
	var temp []*sysmonpb.Rule
	for i := range msg {
		temp = append(temp, &sysmonpb.Rule{Priority: msg[i].Priority, Rule: msg[i].Rule})
	}
	//temp := sysmonpb.Rule{Priority: msg[0].Priority, Rule: msg[0].Rule}
	result := sysmonpb.Rules{Rules: temp}
	response := &sysmonpb.RuleRespone{
		Response: &result,
	}
	return response, nil
}

func (s *GRPCServer) Routes(ctxt context.Context, request *sysmonpb.IPRequest) (*sysmonpb.RouteRespone, error) {
	api.MakeSudo()
	msg := api.IPRoutes()
	var temp []*sysmonpb.Route
	length := len(msg)
	for i := 1; i < length+1; i++ {
		temp = append(temp, &sysmonpb.Route{Index: strconv.Itoa(i), Route: msg[strconv.Itoa(i)]})
	}
	result := sysmonpb.Routes{Routes: temp}
	response := &sysmonpb.RouteRespone{
		Response: &result,
	}
	return response, nil
}
func (s *GRPCServer) RoutesByTableName(ctxt context.Context, request *sysmonpb.Request) (*sysmonpb.RouteRespone, error) {
	api.MakeSudo()
	msg := api.IPRoutesByTableName(request)
	var temp []*sysmonpb.Route
	length := len(msg)
	for i := 1; i < length+1; i++ {
		temp = append(temp, &sysmonpb.Route{Index: strconv.Itoa(i), Route: msg[strconv.Itoa(i)]})
	}
	result := sysmonpb.Routes{Routes: temp}
	response := &sysmonpb.RouteRespone{
		Response: &result,
	}
	return response, nil
}

func (s *GRPCServer) InterfaceAddresses(ctxt context.Context, request *sysmonpb.IPRequest) (*sysmonpb.InterfaceAddressesResponse, error) {
	api.MakeSudo()
	msg := api.InterfaceAddresses()
	var temp []*sysmonpb.IpAddress
	for i := range msg {
		temp = append(temp, &sysmonpb.IpAddress{IP: msg[i].IP, Type: msg[i].Type})
	}
	//temp := sysmonpb.Rule{Priority: msg[0].Priority, Rule: msg[0].Rule}
	result := sysmonpb.IpAddresses{Addresses: temp}
	response := &sysmonpb.InterfaceAddressesResponse{
		Response: &result,
	}
	return response, nil
}

func (s *GRPCServer) Interfaces(ctxt context.Context, request *sysmonpb.IPRequest) (*sysmonpb.InterfacesResponse, error) {
	api.MakeSudo()
	msg := api.Interfaces()
	var temp []*sysmonpb.Interface
	for i := range msg {
		temp = append(temp, &sysmonpb.Interface{Index: strconv.Itoa(msg[i].Index), Name: msg[i].Name})
	}
	//temp := sysmonpb.Rule{Priority: msg[0].Priority, Rule: msg[0].Rule}
	result := sysmonpb.Interfaces{Interfaces: temp}
	response := &sysmonpb.InterfacesResponse{
		Response: &result,
	}
	return response, nil
}

func (s *GRPCServer) InterfaceDetailsByName(ctxt context.Context, request *sysmonpb.Request) (*sysmonpb.InterfaceDetailsResponse, error) {
	api.MakeSudo()
	msg := api.InterfaceDetailsByName(request)
	var normalAddress []*sysmonpb.IpAddress
	for i := range msg.NormalAddress {
		normalAddress = append(normalAddress, &sysmonpb.IpAddress{IP: msg.NormalAddress[i].IP, Type: msg.NormalAddress[i].Type})
	}

	var multicastAddress []*sysmonpb.IpAddress
	for i := range msg.MulticastAddress {
		multicastAddress = append(multicastAddress, &sysmonpb.IpAddress{IP: msg.MulticastAddress[i].IP, Type: msg.MulticastAddress[i].Type})
	}

	var gateways []*sysmonpb.Gateway
	for i := range msg.Gateway {
		gateways = append(gateways, &sysmonpb.Gateway{Destination: msg.Gateway[i].Destination, Gateway: msg.Gateway[i].Gateway})
	}
	result := sysmonpb.InterfaceDetails{Name: msg.Name, Gateways: gateways, NormalAddress: normalAddress, MulticastAddress: multicastAddress}
	response := &sysmonpb.InterfaceDetailsResponse{
		Response: &result,
	}
	return response, nil
}

func (s *GRPCServer) IpTables(ctxt context.Context, request *sysmonpb.Request) (*sysmonpb.IPTablesResponse, error) {
	api.MakeSudo()
	msg := api.Tables()
	var temp []*sysmonpb.IPTable
	for i := range msg {
		temp = append(temp, &sysmonpb.IPTable{TableNumber: strings.Fields(msg[i])[0], TableName: strings.Fields(msg[i])[1]})
	}
	result := sysmonpb.IPTables{Tables: temp}
	response := &sysmonpb.IPTablesResponse{
		Response: &result,
	}
	return response, nil
}
