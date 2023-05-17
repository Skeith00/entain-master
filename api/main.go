package main

import (
	"context"
	"flag"
	"git.neds.sh/matty/entain/api/proto/sports"
	"log"
	"net/http"

	"git.neds.sh/matty/entain/api/proto/racing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	apiEndpoint        = flag.String("api-endpoint", "localhost:8000", "API endpoint")
	grpcRacingEndpoint = flag.String("grpc-racing-endpoint", "localhost:9000", "gRPC Racing server endpoint")
	grpcSportsEndpoint = flag.String("grpc-sports-endpoint", "localhost:7000", "gRPC Sports server endpoint")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Printf("failed running api server: %s\n", err)
	}

}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := registerSportsServer(ctx, mux)
	if err != nil {
		return err
	}
	err = registerRacingServer(ctx, mux)
	if err != nil {
		return err
	}

	log.Printf("API server listening on: %s\n", *apiEndpoint)
	return http.ListenAndServe(*apiEndpoint, mux)
}

func registerSportsServer(ctx context.Context, mux *runtime.ServeMux) error {
	if err := sports.RegisterSportsHandlerFromEndpoint(
		ctx,
		mux,
		*grpcSportsEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}
	return nil
}

func registerRacingServer(ctx context.Context, mux *runtime.ServeMux) error {
	if err := racing.RegisterRacingHandlerFromEndpoint(
		ctx,
		mux,
		*grpcRacingEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}
	return nil
}
