package main

import (
	"context"
	trippb "coorcar/proto/gen/go"
	trip "coorcar/tripservice"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	go startGRPCGateway()
	if err != nil {
		log.Panicf("failed to listen :%v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(lis))
}
func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	m := &runtime.JSONPb{}
	m.UseEnumNumbers = true
	m.UseProtoNames = true
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard,
		m))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("can not start grpc gateway:%v", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("can not listen and server:%v", err)
	}
}
