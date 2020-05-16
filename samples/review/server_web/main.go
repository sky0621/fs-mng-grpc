package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sky0621/fs-mng-grpc/pb"
	"github.com/sky0621/fs-mng-grpc/samples/common"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterReviewHandlerFromEndpoint(ctx, mux, "localhost"+common.ReviewGrpcServerPort, opts); err != nil {
		log.Fatalf("failed to gRPC serve: %v", err)
	}

	if err := http.ListenAndServe(common.ReviewWebServerPort, mux); err != nil {
		log.Fatalf("failed to Web serve: %v", err)
	}
}
