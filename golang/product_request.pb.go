// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: product_request.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credential *Credential           `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	Query      *ProductRequest_Query `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_product_request_proto_rawDescGZIP(), []int{0}
}

func (x *ProductRequest) GetCredential() *Credential {
	if x != nil {
		return x.Credential
	}
	return nil
}

func (x *ProductRequest) GetQuery() *ProductRequest_Query {
	if x != nil {
		return x.Query
	}
	return nil
}

type ProductRequest_Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductRequest_Query) Reset() {
	*x = ProductRequest_Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest_Query) ProtoMessage() {}

func (x *ProductRequest_Query) ProtoReflect() protoreflect.Message {
	mi := &file_product_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest_Query.ProtoReflect.Descriptor instead.
func (*ProductRequest_Query) Descriptor() ([]byte, []int) {
	return file_product_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProductRequest_Query) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_product_request_proto protoreflect.FileDescriptor

var file_product_request_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x12, 0x2e, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x17, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_request_proto_rawDescOnce sync.Once
	file_product_request_proto_rawDescData = file_product_request_proto_rawDesc
)

func file_product_request_proto_rawDescGZIP() []byte {
	file_product_request_proto_rawDescOnce.Do(func() {
		file_product_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_request_proto_rawDescData)
	})
	return file_product_request_proto_rawDescData
}

var file_product_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_request_proto_goTypes = []interface{}{
	(*ProductRequest)(nil),       // 0: pb.ProductRequest
	(*ProductRequest_Query)(nil), // 1: pb.ProductRequest.Query
	(*Credential)(nil),           // 2: pb.Credential
}
var file_product_request_proto_depIdxs = []int32{
	2, // 0: pb.ProductRequest.credential:type_name -> pb.Credential
	1, // 1: pb.ProductRequest.query:type_name -> pb.ProductRequest.Query
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_product_request_proto_init() }
func file_product_request_proto_init() {
	if File_product_request_proto != nil {
		return
	}
	file_credential_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
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
		file_product_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest_Query); i {
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
			RawDescriptor: file_product_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_request_proto_goTypes,
		DependencyIndexes: file_product_request_proto_depIdxs,
		MessageInfos:      file_product_request_proto_msgTypes,
	}.Build()
	File_product_request_proto = out.File
	file_product_request_proto_rawDesc = nil
	file_product_request_proto_goTypes = nil
	file_product_request_proto_depIdxs = nil
}
