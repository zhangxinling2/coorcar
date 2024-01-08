package main

import (
	"context"
	trippb "coorcar/proto/gen/go"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connot connect server:%v", err)
	}
	tsClient := trippb.NewTripServiceClient(conn)
	resp, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "trip456"})
	if err != nil {
		log.Printf("resp error:%v", err)
	}
	fmt.Println(resp)
}
