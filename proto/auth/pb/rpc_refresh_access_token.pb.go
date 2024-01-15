// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: rpc_refresh_access_token.proto

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

type RefreshAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAddress string `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
}

func (x *RefreshAccessTokenRequest) Reset() {
	*x = RefreshAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_refresh_access_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAccessTokenRequest) ProtoMessage() {}

func (x *RefreshAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_refresh_access_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_rpc_refresh_access_token_proto_rawDescGZIP(), []int{0}
}

func (x *RefreshAccessTokenRequest) GetWalletAddress() string {
	if x != nil {
		return x.WalletAddress
	}
	return ""
}

type RefreshAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *RefreshAccessTokenResponse) Reset() {
	*x = RefreshAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_refresh_access_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAccessTokenResponse) ProtoMessage() {}

func (x *RefreshAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_refresh_access_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_rpc_refresh_access_token_proto_rawDescGZIP(), []int{1}
}

func (x *RefreshAccessTokenResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_rpc_refresh_access_token_proto protoreflect.FileDescriptor

var file_rpc_refresh_access_token_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x19, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x43, 0x0a, 0x1a, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x1f, 0x5a, 0x1d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x79, 0x61, 0x6d, 0x61,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_refresh_access_token_proto_rawDescOnce sync.Once
	file_rpc_refresh_access_token_proto_rawDescData = file_rpc_refresh_access_token_proto_rawDesc
)

func file_rpc_refresh_access_token_proto_rawDescGZIP() []byte {
	file_rpc_refresh_access_token_proto_rawDescOnce.Do(func() {
		file_rpc_refresh_access_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_refresh_access_token_proto_rawDescData)
	})
	return file_rpc_refresh_access_token_proto_rawDescData
}

var file_rpc_refresh_access_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_refresh_access_token_proto_goTypes = []interface{}{
	(*RefreshAccessTokenRequest)(nil),  // 0: pb.RefreshAccessTokenRequest
	(*RefreshAccessTokenResponse)(nil), // 1: pb.RefreshAccessTokenResponse
	(*Session)(nil),                    // 2: pb.Session
}
var file_rpc_refresh_access_token_proto_depIdxs = []int32{
	2, // 0: pb.RefreshAccessTokenResponse.session:type_name -> pb.Session
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_refresh_access_token_proto_init() }
func file_rpc_refresh_access_token_proto_init() {
	if File_rpc_refresh_access_token_proto != nil {
		return
	}
	file_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_refresh_access_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshAccessTokenRequest); i {
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
		file_rpc_refresh_access_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshAccessTokenResponse); i {
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
			RawDescriptor: file_rpc_refresh_access_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_refresh_access_token_proto_goTypes,
		DependencyIndexes: file_rpc_refresh_access_token_proto_depIdxs,
		MessageInfos:      file_rpc_refresh_access_token_proto_msgTypes,
	}.Build()
	File_rpc_refresh_access_token_proto = out.File
	file_rpc_refresh_access_token_proto_rawDesc = nil
	file_rpc_refresh_access_token_proto_goTypes = nil
	file_rpc_refresh_access_token_proto_depIdxs = nil
}
