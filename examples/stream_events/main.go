package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func main() {
	kk := keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithKeepaliveParams(kk))

	conn, err := grpc.Dial(":4466", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewTraceeServiceClient(conn)

	stream, err := client.StreamEvents(context.Background(), &pb.StreamEventsRequest{
		Policies: []string{"udp", "icmp"},
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("pq isso é nil -> %v - %v \n", event, err)
			continue
		}

		fmt.Printf("Event: %s\n", event.Event.Name)

		eventData := event.Event.EventData
		if args, ok := eventData["args"]; ok {
			for i, arg := range args.GetArgs().GetValue() {
				fmt.Printf("%s - %d - %v\n", event.Event.Name, i, arg)
			}
		} else {
			for k, v := range eventData {
				fmt.Printf("%s - %s - %v\n", event.Event.Name, k, v)
			}
		}
	}
}
