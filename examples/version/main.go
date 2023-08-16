package main

import (
	"context"
	"log"

	pb "github.com/aquasecurity/tracee/types/api/v1beta1"
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

	version, err := client.GetVersion(context.Background(), &pb.GetVersionRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("version: %+v\n", version.Version)
}
