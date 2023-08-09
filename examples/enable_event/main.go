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

	in := &pb.UpdatePolicyRequest{
		Names:  []string{"test1"},
		Action: pb.UpdaPolicyAction_EnableEvent,
		Event:  &pb.Event{Name: "ptrace"},
	}

	_, err := client.UpdatePolicy(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
}
