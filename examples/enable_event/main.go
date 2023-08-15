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

	in := &pb.EnablePolicyRuleRequest{
		PolicyName: "forensics",
		RuleId:     "openat",
	}

	_, err := client.EnablePolicyRule(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
}
