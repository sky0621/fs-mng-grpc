// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: review.proto

package review

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListFacilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFacilityRequest) Reset() {
	*x = ListFacilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFacilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFacilityRequest) ProtoMessage() {}

func (x *ListFacilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFacilityRequest.ProtoReflect.Descriptor instead.
func (*ListFacilityRequest) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{0}
}

type ListFacilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Facilities []*Facility `protobuf:"bytes,1,rep,name=facilities,proto3" json:"facilities,omitempty"`
}

func (x *ListFacilityResponse) Reset() {
	*x = ListFacilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFacilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFacilityResponse) ProtoMessage() {}

func (x *ListFacilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFacilityResponse.ProtoReflect.Descriptor instead.
func (*ListFacilityResponse) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{1}
}

func (x *ListFacilityResponse) GetFacilities() []*Facility {
	if x != nil {
		return x.Facilities
	}
	return nil
}

// 審査機関
type Facility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Facility) Reset() {
	*x = Facility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Facility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Facility) ProtoMessage() {}

func (x *Facility) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Facility.ProtoReflect.Descriptor instead.
func (*Facility) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{2}
}

func (x *Facility) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Facility) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *ListRecordRequest) Reset() {
	*x = ListRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecordRequest) ProtoMessage() {}

func (x *ListRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecordRequest.ProtoReflect.Descriptor instead.
func (*ListRecordRequest) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{3}
}

func (x *ListRecordRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type ListRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListRecordResponse) Reset() {
	*x = ListRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecordResponse) ProtoMessage() {}

func (x *ListRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecordResponse.ProtoReflect.Descriptor instead.
func (*ListRecordResponse) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{4}
}

func (x *ListRecordResponse) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

// 審査記録
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Note           string `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	RecordDatetime int64  `protobuf:"varint,3,opt,name=recordDatetime,proto3" json:"recordDatetime,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{5}
}

func (x *Record) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Record) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Record) GetRecordDatetime() int64 {
	if x != nil {
		return x.RecordDatetime
	}
	return 0
}

var File_review_proto protoreflect.FileDescriptor

var file_review_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x63,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x0a, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x66, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x37, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x22, 0x54, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x80, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x3d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_proto_rawDescOnce sync.Once
	file_review_proto_rawDescData = file_review_proto_rawDesc
)

func file_review_proto_rawDescGZIP() []byte {
	file_review_proto_rawDescOnce.Do(func() {
		file_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_proto_rawDescData)
	})
	return file_review_proto_rawDescData
}

var file_review_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_review_proto_goTypes = []interface{}{
	(*ListFacilityRequest)(nil),  // 0: ListFacilityRequest
	(*ListFacilityResponse)(nil), // 1: ListFacilityResponse
	(*Facility)(nil),             // 2: Facility
	(*ListRecordRequest)(nil),    // 3: ListRecordRequest
	(*ListRecordResponse)(nil),   // 4: ListRecordResponse
	(*Record)(nil),               // 5: Record
}
var file_review_proto_depIdxs = []int32{
	2, // 0: ListFacilityResponse.facilities:type_name -> Facility
	5, // 1: ListRecordResponse.records:type_name -> Record
	0, // 2: Review.ListFacility:input_type -> ListFacilityRequest
	3, // 3: Review.ListRecord:input_type -> ListRecordRequest
	1, // 4: Review.ListFacility:output_type -> ListFacilityResponse
	4, // 5: Review.ListRecord:output_type -> ListRecordResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_review_proto_init() }
func file_review_proto_init() {
	if File_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFacilityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFacilityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Facility); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRecordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_review_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRecordResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_review_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_proto_goTypes,
		DependencyIndexes: file_review_proto_depIdxs,
		MessageInfos:      file_review_proto_msgTypes,
	}.Build()
	File_review_proto = out.File
	file_review_proto_rawDesc = nil
	file_review_proto_goTypes = nil
	file_review_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReviewClient is the client API for Review service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewClient interface {
	// 審査機関の一覧を取得
	ListFacility(ctx context.Context, in *ListFacilityRequest, opts ...grpc.CallOption) (*ListFacilityResponse, error)
	// 特定(プレ)オーダーの審査記録一覧を取得
	ListRecord(ctx context.Context, in *ListRecordRequest, opts ...grpc.CallOption) (*ListRecordResponse, error)
}

type reviewClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewClient(cc grpc.ClientConnInterface) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) ListFacility(ctx context.Context, in *ListFacilityRequest, opts ...grpc.CallOption) (*ListFacilityResponse, error) {
	out := new(ListFacilityResponse)
	err := c.cc.Invoke(ctx, "/Review/ListFacility", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ListRecord(ctx context.Context, in *ListRecordRequest, opts ...grpc.CallOption) (*ListRecordResponse, error) {
	out := new(ListRecordResponse)
	err := c.cc.Invoke(ctx, "/Review/ListRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServer is the server API for Review service.
type ReviewServer interface {
	// 審査機関の一覧を取得
	ListFacility(context.Context, *ListFacilityRequest) (*ListFacilityResponse, error)
	// 特定(プレ)オーダーの審査記録一覧を取得
	ListRecord(context.Context, *ListRecordRequest) (*ListRecordResponse, error)
}

// UnimplementedReviewServer can be embedded to have forward compatible implementations.
type UnimplementedReviewServer struct {
}

func (*UnimplementedReviewServer) ListFacility(context.Context, *ListFacilityRequest) (*ListFacilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFacility not implemented")
}
func (*UnimplementedReviewServer) ListRecord(context.Context, *ListRecordRequest) (*ListRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecord not implemented")
}

func RegisterReviewServer(s *grpc.Server, srv ReviewServer) {
	s.RegisterService(&_Review_serviceDesc, srv)
}

func _Review_ListFacility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFacilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ListFacility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Review/ListFacility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ListFacility(ctx, req.(*ListFacilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ListRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ListRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Review/ListRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ListRecord(ctx, req.(*ListRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Review_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFacility",
			Handler:    _Review_ListFacility_Handler,
		},
		{
			MethodName: "ListRecord",
			Handler:    _Review_ListRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}
