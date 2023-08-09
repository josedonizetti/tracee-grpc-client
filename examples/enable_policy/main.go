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

	in := &pb.EnablePolicyRequest{
		PolicyName: "forensics",
	}

	_, err := client.EnablePolicy(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
}
