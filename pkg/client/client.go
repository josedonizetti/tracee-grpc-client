package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/aquasecurity/tracee/types/api/v1beta1"
)

func NewGRPCConn() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":4466", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return conn
}

func NewTraceeServciceClient(conn *grpc.ClientConn) pb.TraceeServiceClient {
	return pb.NewTraceeServiceClient(conn)
}

func NewTraceeDiagnosticClient(conn *grpc.ClientConn) pb.DiagnosticServiceClient {
	return pb.NewDiagnosticServiceClient(conn)
}
