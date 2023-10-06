package main

import (
	"context"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":4466", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := pb.NewTraceeServiceClient(conn)
	in := &pb.EnableEventRequest{
		Name: "net_packet_udp",
	}

	_, err = client.EnableEvent(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
}
