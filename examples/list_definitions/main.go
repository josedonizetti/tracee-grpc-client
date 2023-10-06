package main

import (
	"context"
	"fmt"
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

	definitions, err := client.GetEventDefinition(
		context.Background(),
		&pb.GetEventDefinitionRequest{Name: ""},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, definition := range definitions.Definitions {
		v := fmt.Sprintf("%d.%d.%d", definition.Version.Major, definition.Version.Minor, definition.Version.Patch)
		fmt.Println("Id:", definition.Id)
		fmt.Println("Name:", definition.Name)
		fmt.Println("Description:", definition.Description)
		fmt.Println("Version:", v)
		fmt.Printf("Tags: %v\n", definition.Tags)
	}
}
