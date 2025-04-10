// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: service/frontend/client/v1/client.proto

package client

import (
	v1 "github.com/shynggys9219/ap2-apis-gen-user-service/base/bo/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// email
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// password
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_service_frontend_client_v1_client_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_frontend_client_v1_client_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_service_frontend_client_v1_client_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id of a new entity
	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_service_frontend_client_v1_client_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_frontend_client_v1_client_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_service_frontend_client_v1_client_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_service_frontend_client_v1_client_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_frontend_client_v1_client_proto_msgTypes[2]
	if x != nil {
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
	return file_service_frontend_client_v1_client_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// client struct
	Client        *v1.Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_service_frontend_client_v1_client_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_frontend_client_v1_client_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_service_frontend_client_v1_client_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetClient() *v1.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

var File_service_frontend_client_v1_client_proto protoreflect.FileDescriptor

const file_service_frontend_client_v1_client_proto_rawDesc = "" +
	"\n" +
	"'service/frontend/client/v1/client.proto\x12\x1fuser.service.frontend.client.v1\x1a\x1dbase/frontend/v1/client.proto\"A\n" +
	"\rCreateRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\" \n" +
	"\x0eCreateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"\x1c\n" +
	"\n" +
	"GetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"D\n" +
	"\vGetResponse\x125\n" +
	"\x06client\x18\x01 \x01(\v2\x1d.user.base.frontend.v1.ClientR\x06client2\xdc\x01\n" +
	"\rClientService\x12i\n" +
	"\x06Create\x12..user.service.frontend.client.v1.CreateRequest\x1a/.user.service.frontend.client.v1.CreateResponse\x12`\n" +
	"\x03Get\x12+.user.service.frontend.client.v1.GetRequest\x1a,.user.service.frontend.client.v1.GetResponseBUZSgithub.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1;clientb\x06proto3"

var (
	file_service_frontend_client_v1_client_proto_rawDescOnce sync.Once
	file_service_frontend_client_v1_client_proto_rawDescData []byte
)

func file_service_frontend_client_v1_client_proto_rawDescGZIP() []byte {
	file_service_frontend_client_v1_client_proto_rawDescOnce.Do(func() {
		file_service_frontend_client_v1_client_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_service_frontend_client_v1_client_proto_rawDesc), len(file_service_frontend_client_v1_client_proto_rawDesc)))
	})
	return file_service_frontend_client_v1_client_proto_rawDescData
}

var file_service_frontend_client_v1_client_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_frontend_client_v1_client_proto_goTypes = []any{
	(*CreateRequest)(nil),  // 0: user.service.frontend.client.v1.CreateRequest
	(*CreateResponse)(nil), // 1: user.service.frontend.client.v1.CreateResponse
	(*GetRequest)(nil),     // 2: user.service.frontend.client.v1.GetRequest
	(*GetResponse)(nil),    // 3: user.service.frontend.client.v1.GetResponse
	(*v1.Client)(nil),      // 4: user.base.frontend.v1.Client
}
var file_service_frontend_client_v1_client_proto_depIdxs = []int32{
	4, // 0: user.service.frontend.client.v1.GetResponse.client:type_name -> user.base.frontend.v1.Client
	0, // 1: user.service.frontend.client.v1.ClientService.Create:input_type -> user.service.frontend.client.v1.CreateRequest
	2, // 2: user.service.frontend.client.v1.ClientService.Get:input_type -> user.service.frontend.client.v1.GetRequest
	1, // 3: user.service.frontend.client.v1.ClientService.Create:output_type -> user.service.frontend.client.v1.CreateResponse
	3, // 4: user.service.frontend.client.v1.ClientService.Get:output_type -> user.service.frontend.client.v1.GetResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_frontend_client_v1_client_proto_init() }
func file_service_frontend_client_v1_client_proto_init() {
	if File_service_frontend_client_v1_client_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_frontend_client_v1_client_proto_rawDesc), len(file_service_frontend_client_v1_client_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_frontend_client_v1_client_proto_goTypes,
		DependencyIndexes: file_service_frontend_client_v1_client_proto_depIdxs,
		MessageInfos:      file_service_frontend_client_v1_client_proto_msgTypes,
	}.Build()
	File_service_frontend_client_v1_client_proto = out.File
	file_service_frontend_client_v1_client_proto_goTypes = nil
	file_service_frontend_client_v1_client_proto_depIdxs = nil
}
