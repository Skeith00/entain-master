package main

import (
	"context"
	"flag"
	"git.neds.sh/matty/entain/api/proto/racing"
	"git.neds.sh/matty/entain/api/proto/sports"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var (
	apiEndpoint        = flag.String("api-endpoint", "localhost:8000", "API endpoint")
	grpcRacingEndpoint = flag.String("grpc-racing-endpoint", "localhost:9000", "gRPC Racing server endpoint")
	grpcSportsEndpoint = flag.String("grpc-sports-endpoint", "localhost:7000", "gRPC Sports server endpoint")
)

func main() {
	flag.Parse()

	if err := run(RacingServer{}); err != nil {
		log.Printf("failed running api server: %s\n", err)
	}

	if err := run(SportsServer{}); err != nil {
		log.Printf("failed running api server: %s\n", err)
	}
}

type Server interface {
	register(ctx context.Context, mux *runtime.ServeMux) error
}

type RacingServer struct{}
type SportsServer struct{}

func run(server Server) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := server.register(ctx, mux)
	if err != nil {
		return err
	}
	log.Printf("API server listening on: %s\n", *apiEndpoint)

	return http.ListenAndServe(*apiEndpoint, mux)
}

func (s SportsServer) register(ctx context.Context, mux *runtime.ServeMux) error {
	if err := sports.RegisterSportsHandlerFromEndpoint(
		ctx,
		mux,
		*grpcRacingEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}
	return nil
}

func (r RacingServer) register(ctx context.Context, mux *runtime.ServeMux) error {
	if err := racing.RegisterRacingHandlerFromEndpoint(
		ctx,
		mux,
		*grpcSportsEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		return err
	}
	return nil
}
