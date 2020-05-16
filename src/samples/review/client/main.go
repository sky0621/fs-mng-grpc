package main

import (
	"context"
	"github.com/sky0621/fs-mng-grpc/samples/common"
	"github.com/sky0621/fs-mng-grpc/review"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost"+common.ReviewServerPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(){
		if conn == nil {
			return
		}
		if err := conn.Close(); err != nil {
			log.Fatalf("did not close: %v", err)
		}
	}()

	cli := review.NewReviewClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	{
		res, err := cli.ListFacility(ctx, &review.ListFacilityRequest{
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
		res, err := cli.ListRecord(ctx, &review.ListRecordRequest{
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
