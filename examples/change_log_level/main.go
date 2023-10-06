package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"

	"github.com/josedonizetti/tracee-grpc-client/pkg/client"
)

func main() {

	// todo pass which log level to change to as an argument

	conn := client.NewGRPCConn()
	defer conn.Close()

	client := client.NewTraceeDiagnosticClient(conn)

	changelog := &pb.ChangeLogLevelRequest{
		Level: pb.LogLevel_Debug,
	}

	_, err := client.ChangeLogLevel(context.Background(), changelog)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Log level changed to DEBUG")
}
