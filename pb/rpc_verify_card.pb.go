// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: rpc_verify_card.proto

package pb

import (
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

type VerifyCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardNumber      int64 `protobuf:"varint,1,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	ExpirationMonth int64 `protobuf:"varint,2,opt,name=expiration_month,json=expirationMonth,proto3" json:"expiration_month,omitempty"`
	ExpirationYear  int64 `protobuf:"varint,3,opt,name=expiration_year,json=expirationYear,proto3" json:"expiration_year,omitempty"`
}

func (x *VerifyCardRequest) Reset() {
	*x = VerifyCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_verify_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCardRequest) ProtoMessage() {}

func (x *VerifyCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_verify_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCardRequest.ProtoReflect.Descriptor instead.
func (*VerifyCardRequest) Descriptor() ([]byte, []int) {
	return file_rpc_verify_card_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyCardRequest) GetCardNumber() int64 {
	if x != nil {
		return x.CardNumber
	}
	return 0
}

func (x *VerifyCardRequest) GetExpirationMonth() int64 {
	if x != nil {
		return x.ExpirationMonth
	}
	return 0
}

func (x *VerifyCardRequest) GetExpirationYear() int64 {
	if x != nil {
		return x.ExpirationYear
	}
	return 0
}

type VerifyCardErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VerifyCardErrorResponse) Reset() {
	*x = VerifyCardErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_verify_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCardErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCardErrorResponse) ProtoMessage() {}

func (x *VerifyCardErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_verify_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCardErrorResponse.ProtoReflect.Descriptor instead.
func (*VerifyCardErrorResponse) Descriptor() ([]byte, []int) {
	return file_rpc_verify_card_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyCardErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type VerifyCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool                     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Error *VerifyCardErrorResponse `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *VerifyCardResponse) Reset() {
	*x = VerifyCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_verify_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCardResponse) ProtoMessage() {}

func (x *VerifyCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_verify_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCardResponse.ProtoReflect.Descriptor instead.
func (*VerifyCardResponse) Descriptor() ([]byte, []int) {
	return file_rpc_verify_card_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyCardResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *VerifyCardResponse) GetError() *VerifyCardErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_rpc_verify_card_proto protoreflect.FileDescriptor

var file_rpc_verify_card_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x88, 0x01, 0x0a, 0x11,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x22, 0x33, 0x0a, 0x17, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x61, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a, 0x12, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x61, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x76, 0x61, 0x62, 0x6e, 0x64, 0x72,
	0x2f, 0x63, 0x61, 0x72, 0x64, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_verify_card_proto_rawDescOnce sync.Once
	file_rpc_verify_card_proto_rawDescData = file_rpc_verify_card_proto_rawDesc
)

func file_rpc_verify_card_proto_rawDescGZIP() []byte {
	file_rpc_verify_card_proto_rawDescOnce.Do(func() {
		file_rpc_verify_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_verify_card_proto_rawDescData)
	})
	return file_rpc_verify_card_proto_rawDescData
}

var file_rpc_verify_card_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_verify_card_proto_goTypes = []interface{}{
	(*VerifyCardRequest)(nil),       // 0: pb.VerifyCardRequest
	(*VerifyCardErrorResponse)(nil), // 1: pb.VerifyCardErrorResponse
	(*VerifyCardResponse)(nil),      // 2: pb.VerifyCardResponse
}
var file_rpc_verify_card_proto_depIdxs = []int32{
	1, // 0: pb.VerifyCardResponse.error:type_name -> pb.VerifyCardErrorResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_verify_card_proto_init() }
func file_rpc_verify_card_proto_init() {
	if File_rpc_verify_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_verify_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCardRequest); i {
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
		file_rpc_verify_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCardErrorResponse); i {
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
		file_rpc_verify_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCardResponse); i {
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
	file_rpc_verify_card_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_verify_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_verify_card_proto_goTypes,
		DependencyIndexes: file_rpc_verify_card_proto_depIdxs,
		MessageInfos:      file_rpc_verify_card_proto_msgTypes,
	}.Build()
	File_rpc_verify_card_proto = out.File
	file_rpc_verify_card_proto_rawDesc = nil
	file_rpc_verify_card_proto_goTypes = nil
	file_rpc_verify_card_proto_depIdxs = nil
}
