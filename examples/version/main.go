package main

import (
	"context"
	"log"

	pb "github.com/aquasecurity/tracee/types/api/v1beta1"

	"github.com/josedonizetti/tracee-grpc-client/pkg/client"
)

func main() {

	conn := client.NewGRPCConn()
	defer conn.Close()

	client := client.NewTraceeServciceClient(conn)

	version, err := client.GetVersion(context.Background(), &pb.GetVersionRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("version: %+v\n", version.Version)
}