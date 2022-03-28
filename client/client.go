package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "simple-grpc/proto/test"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:30001", "The server address in the format of host:port")
)

// printFeature gets the feature for the given point.
func TestGetClient(client pb.TestServiceClient) {
	req := &pb.Request{Id: 2}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.TestGet(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}

func main() {
	var clOpts []grpc.DialOption
	clOpts = append(clOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, clOpts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTestServiceClient(conn)

	// Looking for a valid feature
	TestGetClient(client)
}
