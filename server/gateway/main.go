package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	m := runtime.JSONPb{}
	m.UseProtoNames = true
	m.UseEnumNumbers = true
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &m))
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(c, mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal("connot register", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("connot http", err)
	}
}
