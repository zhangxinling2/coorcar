package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/server"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	lg, err := server.NewZapLog()
	if err != nil {
		lg.Sugar().Fatal("cannot new log")
	}
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	m := runtime.JSONPb{}
	m.UseProtoNames = true
	m.UseEnumNumbers = true
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &m))
	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth service",
			addr:         "localhost:8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "trip service",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
		{
			name:         "profile service",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterProfileServiceHandlerFromEndpoint,
		},
	}
	for _, s := range serverConfig {
		err := s.registerFunc(c, mux, s.addr, []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			log.Fatal("connot register", err)
		}
	}
	lg.Sugar().Fatal(http.ListenAndServe(":8080", mux))
}
