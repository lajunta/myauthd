package grpcd

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	// GrpcdAddress will init in main package
	GrpcdAddress string
	// RestPort will init in main package
	RestPort string
	// Rest is a switch to serve http rest request
	Rest bool
)

// restRun serve json api to http client
func restRun() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := RegisterGrpcdHandlerFromEndpoint(ctx, mux, GrpcdAddress, opts)
	if err != nil {
		return
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Printf("Gateway is running on 0.0.0.0:%s \n", RestPort)
	http.ListenAndServe(":"+RestPort, mux)
}
