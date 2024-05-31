// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: listing/service.proto

package listing

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetRequest is the request object for the Get method
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListingId string `protobuf:"bytes,1,opt,name=listing_id,json=listingId,proto3" json:"listing_id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listing_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_listing_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetListingId() string {
	if x != nil {
		return x.ListingId
	}
	return ""
}

// ListRequest is the request object for the List method
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastUpdateTs *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=last_update_ts,json=lastUpdateTs,proto3" json:"last_update_ts,omitempty"`
	NextToken    string                 `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
	PageLength   int32                  `protobuf:"varint,3,opt,name=page_length,json=pageLength,proto3" json:"page_length,omitempty"`
	// protolint:disable:next REPEATED_FIELD_NAMES_PLURALIZED
	Skus []string `protobuf:"bytes,4,rep,name=skus,proto3" json:"skus,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listing_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_listing_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetLastUpdateTs() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTs
	}
	return nil
}

func (x *ListRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

func (x *ListRequest) GetPageLength() int32 {
	if x != nil {
		return x.PageLength
	}
	return 0
}

func (x *ListRequest) GetSkus() []string {
	if x != nil {
		return x.Skus
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listings []*Listing `protobuf:"bytes,1,rep,name=listings,proto3" json:"listings,omitempty"`
	// next_token is a token that can be provided to List request
	// to get the next page of results.
	// Next tokens are only valid for 5 minutes after being returned.
	// If no `nextToken` is provided, there are no more results to return.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listing_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listing_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_listing_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetListings() []*Listing {
	if x != nil {
		return x.Listings
	}
	return nil
}

func (x *ListResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

var File_listing_service_proto protoreflect.FileDescriptor

var file_listing_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x22, 0x5b, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xec, 0x02, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x13, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12,
	0x23, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x53, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0xa9, 0x01, 0x92, 0x41, 0xa5, 0x01,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x90, 0x01, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x20, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x61,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x20, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x0a, 0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x73, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x69, 0x6e, 0x67,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x20, 0x61, 0x73, 0x20, 0x77, 0x65, 0x6c,
	0x6c, 0x20, 0x61, 0x73, 0x20, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x62, 0x61, 0x63,
	0x6b, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x5a, 0x65,
	0x6e, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x48, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x92, 0x41, 0x14, 0x12, 0x12, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x20, 0x41, 0x50, 0x49, 0x32, 0x02, 0x76, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_listing_service_proto_rawDescOnce sync.Once
	file_listing_service_proto_rawDescData = file_listing_service_proto_rawDesc
)

func file_listing_service_proto_rawDescGZIP() []byte {
	file_listing_service_proto_rawDescOnce.Do(func() {
		file_listing_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_listing_service_proto_rawDescData)
	})
	return file_listing_service_proto_rawDescData
}

var file_listing_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_listing_service_proto_goTypes = []interface{}{
	(*GetRequest)(nil),            // 0: listing.GetRequest
	(*ListRequest)(nil),           // 1: listing.ListRequest
	(*ListResponse)(nil),          // 2: listing.ListResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Listing)(nil),               // 4: listing.Listing
}
var file_listing_service_proto_depIdxs = []int32{
	3, // 0: listing.ListRequest.last_update_ts:type_name -> google.protobuf.Timestamp
	4, // 1: listing.ListResponse.listings:type_name -> listing.Listing
	0, // 2: listing.ListingService.Get:input_type -> listing.GetRequest
	1, // 3: listing.ListingService.List:input_type -> listing.ListRequest
	4, // 4: listing.ListingService.Get:output_type -> listing.Listing
	2, // 5: listing.ListingService.List:output_type -> listing.ListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_listing_service_proto_init() }
func file_listing_service_proto_init() {
	if File_listing_service_proto != nil {
		return
	}
	file_listing_listing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_listing_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_listing_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_listing_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_listing_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listing_service_proto_goTypes,
		DependencyIndexes: file_listing_service_proto_depIdxs,
		MessageInfos:      file_listing_service_proto_msgTypes,
	}.Build()
	File_listing_service_proto = out.File
	file_listing_service_proto_rawDesc = nil
	file_listing_service_proto_goTypes = nil
	file_listing_service_proto_depIdxs = nil
}