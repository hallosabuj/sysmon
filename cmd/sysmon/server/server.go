package server

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	///////////////////////////////////////////
	// This portion is for UI
	// "io/fs"
	///////////////////////////////////////////
	"net"
	"net/http"
	"strings"
	"sysmon/proto/sysmonpb"
	///////////////////////////////////////////
	// This portion is for UI
	// build "sysmon/web"
	///////////////////////////////////////////
)

func GRPCStart() (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	port := listener.Addr().(*net.TCPAddr).Port
	service := grpc.NewServer()
	sysmonpb.RegisterSysmonServiceServer(service, NewServer())
	go service.Serve(listener)
	return port, nil
}

func GRPCClient(port int) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		fmt.Sprintf("localhost:%d", port),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}

func RESTStart(conn *grpc.ClientConn, port int) (int, error) {
	mux := runtime.NewServeMux()
	err := sysmonpb.RegisterSysmonServiceHandler(context.Background(), mux, conn)
	if nil != err {
		fmt.Println("Failed initializing grpc gateway handler, error: ", err)
		return 0, err
	}

	///////////////////////////////////////////
	// This portion is for UI
	// filesystem := fs.FS(build.WWW)
	// child, err := fs.Sub(filesystem, "build")
	///////////////////////////////////////////
	if nil != err {
		return 0, err
	}
	var (
		///////////////////////////////////////////
		// This portion is for UI
		// web      = http.FileServer(http.FS(child))
		///////////////////////////////////////////
		handlers = http.NewServeMux()
	)
	handlers.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if strings.HasPrefix(request.URL.Path, "/api") {
			mux.ServeHTTP(writer, request)
			return
		}
		///////////////////////////////////////////
		// This portion is for UI
		// if strings.HasPrefix(request.URL.Path, "/web") {
		// 	http.Redirect(writer, request, "/", http.StatusFound)
		// 	return
		// }
		// web.ServeHTTP(writer, request)
		///////////////////////////////////////////
	})
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if nil != err {
		return 0, err
	}
	port = listener.Addr().(*net.TCPAddr).Port
	go http.Serve(listener, handlers)
	return port, nil
}

func Start(port int) error {
	gport, err := GRPCStart()
	if nil != err {
		fmt.Println("Failed starting grpc service, error: ", err)
		return err
	}
	fmt.Println("Starting controller grpc service on port:", gport)
	conn, err := GRPCClient(gport)
	if nil != err {
		fmt.Println("Failed initializing grpc connector, error: ", err)
		return err
	}
	rport, err := RESTStart(conn, port)
	if nil != err {
		fmt.Println("Failed starting rest service, error: ", err)
		return err
	}
	fmt.Println("Starting controller rest service on port:", rport)
	return nil
}
