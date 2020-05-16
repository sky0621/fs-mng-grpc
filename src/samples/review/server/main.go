package main

import (
	"context"
	"github.com/sky0621/fs-mng-grpc/samples/common"
	"github.com/sky0621/fs-mng-grpc/review"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	review.ReviewServer
}

// 審査機関の一覧を取得
func (s *server) ListFacility(ctx context.Context, req *review.ListFacilityRequest) (*review.ListFacilityResponse, error) {
	// FIXME:
	log.Printf("req:%#v\n", req)
	return &review.ListFacilityResponse{
		Facilities: []*review.Facility{
			{ID: "34aab2a6-d28d-4ef3-9f2d-1d84780ea994", Name: "審査機関１"},
			{ID: "1579f658-0821-4834-ad82-d3eaf12ac3f4", Name: "審査機関２"},
			{ID: "a205aa40-3f3f-4dc6-8278-f1653030c658", Name: "審査機関３"},
		},
	},nil
}

// 特定(プレ)オーダーの審査記録一覧を取得
func (s *server) ListRecord(ctx context.Context, req *review.ListRecordRequest) (*review.ListRecordResponse, error) {
	// FIXME:
	log.Printf("req:%#v\n", req)
	return &review.ListRecordResponse{
		Records: []*review.Record{
			{ID: "0b139be0-ba4a-45d7-bd05-37df81a5f9b3", Note: "審査メモ１", RecordDatetime: 1589415917},
			{ID: "d5ea2a4b-377b-4e04-8002-9deddcd75ac9", Note: "審査メモ２", RecordDatetime: 1589416160},
			{ID: "887b225b-5945-445c-babd-b0c9f12bd623", Note: "審査メモ３", RecordDatetime: 1589416188},
		},
	},nil
}

func main() {
	lis, err := net.Listen("tcp", common.ReviewServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	review.RegisterReviewServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
