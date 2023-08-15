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

	in := &pb.DisablePolicyRuleRequest{
		PolicyName: "test1",
		RuleId:     "ptrace",
	}

	_, err := client.DisablePolicyRule(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
}
