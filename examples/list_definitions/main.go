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

	list, err := client.ListEventsDefinition(context.Background(), &pb.ListEventsDefinitionRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range list.Definitions {
		log.Printf("definition: %+v\n", event)
	}
}
