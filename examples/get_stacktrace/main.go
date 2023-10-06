package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"

	"github.com/josedonizetti/tracee-grpc-client/pkg/client"
)

func main() {
	conn := client.NewGRPCConn()
	defer conn.Close()

	client := client.NewTraceeDiagnosticClient(conn)

	stacktrace, err := client.GetStacktrace(context.Background(), &pb.GetStacktraceRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(stacktrace.Stacktrace))
}
