package main

import (
	"context"
	"log"
	"time"

	"github.com/sky0621/fs-mng-grpc/pb"
	"github.com/sky0621/fs-mng-grpc/samples/common"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "localhost"+common.ReviewGrpcServerPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if conn == nil {
			return
		}
		if err := conn.Close(); err != nil {
			log.Fatalf("did not close: %v", err)
		}
	}()

	cli := pb.NewReviewClient(conn)

	{
		res, err := cli.ListFacility(ctx, &pb.ListFacilityRequest{
			// MEMO: no parameter
		})
		if err != nil {
			log.Fatalf("could not ListFacility: %v", err)
		}
		if res != nil {
			for _, facility := range res.Facilities {
				log.Printf("Facility: %#v\n", facility)
			}
		}
	}

	{
		res, err := cli.ListRecord(ctx, &pb.ListRecordRequest{
			// MEMO: no parameter
		})
		if err != nil {
			log.Fatalf("could not ListRecord: %v", err)
		}
		if res != nil {
			for _, record := range res.Records {
				log.Printf("Record: %#v\n", record)
			}
		}
	}
}
