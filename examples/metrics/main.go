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

	metrics, err := client.GetMetrics(context.Background(), &pb.GetMetricsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("EventCount", metrics.EventCount)
	fmt.Println("EventsFiltered", metrics.EventsFiltered)
	fmt.Println("NetCapCount", metrics.NetCapCount)
	fmt.Println("BPFLogsCount", metrics.BPFLogsCount)
	fmt.Println("ErrorCount", metrics.ErrorCount)
	fmt.Println("LostEvCount", metrics.LostEvCount)
	fmt.Println("LostWrCount", metrics.LostWrCount)
	fmt.Println("LostNtCapCount", metrics.LostNtCapCount)
	fmt.Println("LostBPFLogsCount", metrics.LostBPFLogsCount)
}
