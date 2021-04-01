package server

import (
	"context"
	"strconv"
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

func (s *GRPCServer) Ping(
	ctxt context.Context, request *sysmonpb.PingRequest,
) (*sysmonpb.PingResponse, error) {
	temp := []string{"a", "s", "d", "f"}
	response := &sysmonpb.PingResponse{
		Body: &sysmonpb.PP{Name: temp},
	}
	return response, nil
}

func (s *GRPCServer) Pong(
	ctxt context.Context, request *sysmonpb.PongRequest,
) (*sysmonpb.PongResponse, error) {
	response := &sysmonpb.PongResponse{
		Body: request.Body,
	}
	return response, nil
}

func (s *GRPCServer) AddRule(
	ctxt context.Context, request *sysmonpb.IPRequest,
) (*sysmonpb.IPResponse, error) {
	api.MakeSudo()
	msg := api.AddIPRule(request)
	response := &sysmonpb.IPResponse{
		Response: &sysmonpb.Response{Msg: msg},
	}
	return response, nil
}

func (s *GRPCServer) DelRule(
	ctxt context.Context, request *sysmonpb.IPRequest,
) (*sysmonpb.IPResponse, error) {
	api.MakeSudo()
	msg := api.DelIPRule(request)
	response := &sysmonpb.IPResponse{
		Response: &sysmonpb.Response{Msg: msg},
	}
	return response, nil
}

func (s *GRPCServer) AddRoute(
	ctxt context.Context, request *sysmonpb.IPRequest,
) (*sysmonpb.IPResponse, error) {
	api.MakeSudo()
	msg := api.AddIPRoute(request)
	response := &sysmonpb.IPResponse{
		Response: &sysmonpb.Response{Msg: msg},
	}
	return response, nil
}

func (s *GRPCServer) DelRoute(
	ctxt context.Context, request *sysmonpb.IPRequest,
) (*sysmonpb.IPResponse, error) {
	api.MakeSudo()
	msg := api.DelIPRoute(request)
	response := &sysmonpb.IPResponse{
		Response: &sysmonpb.Response{Msg: msg},
	}
	return response, nil
}

func (s *GRPCServer) ListRule(
	ctxt context.Context, request *sysmonpb.IPRequest,
) (*sysmonpb.RuleRespone, error) {
	api.MakeSudo()
	msg := api.ListIPRules()
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

func (s *GRPCServer) ListRoute(
	ctxt context.Context, request *sysmonpb.IPRequest,
) (*sysmonpb.RouteRespone, error) {
	api.MakeSudo()
	msg := api.ListIPRoutes()
	var temp []*sysmonpb.Route
	length := len(msg)
	for i := 1; i < length+1; i++ {
		temp = append(temp, &sysmonpb.Route{Index: strconv.Itoa(i), Route: msg[strconv.Itoa(i)]})
	}
	//temp := sysmonpb.Rule{Priority: msg[0].Priority, Rule: msg[0].Rule}
	result := sysmonpb.Routes{Routes: temp}
	response := &sysmonpb.RouteRespone{
		Response: &result,
	}
	return response, nil
}
