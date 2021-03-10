// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: app/impl/grpc/snippet.proto

package generated

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

type CreateSnippetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateSnippetRequest) Reset() {
	*x = CreateSnippetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_impl_grpc_snippet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnippetRequest) ProtoMessage() {}

func (x *CreateSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_impl_grpc_snippet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnippetRequest.ProtoReflect.Descriptor instead.
func (*CreateSnippetRequest) Descriptor() ([]byte, []int) {
	return file_app_impl_grpc_snippet_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSnippetRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateSnippetRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateSnippetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateSnippetResponse) Reset() {
	*x = CreateSnippetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_impl_grpc_snippet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnippetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnippetResponse) ProtoMessage() {}

func (x *CreateSnippetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_impl_grpc_snippet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnippetResponse.ProtoReflect.Descriptor instead.
func (*CreateSnippetResponse) Descriptor() ([]byte, []int) {
	return file_app_impl_grpc_snippet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSnippetResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_app_impl_grpc_snippet_proto protoreflect.FileDescriptor

var file_app_impl_grpc_snippet_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x66, 0x0a, 0x0e, 0x53, 0x6e, 0x69,
	0x70, 0x70, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_impl_grpc_snippet_proto_rawDescOnce sync.Once
	file_app_impl_grpc_snippet_proto_rawDescData = file_app_impl_grpc_snippet_proto_rawDesc
)

func file_app_impl_grpc_snippet_proto_rawDescGZIP() []byte {
	file_app_impl_grpc_snippet_proto_rawDescOnce.Do(func() {
		file_app_impl_grpc_snippet_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_impl_grpc_snippet_proto_rawDescData)
	})
	return file_app_impl_grpc_snippet_proto_rawDescData
}

var file_app_impl_grpc_snippet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_impl_grpc_snippet_proto_goTypes = []interface{}{
	(*CreateSnippetRequest)(nil),  // 0: generated.CreateSnippetRequest
	(*CreateSnippetResponse)(nil), // 1: generated.CreateSnippetResponse
}
var file_app_impl_grpc_snippet_proto_depIdxs = []int32{
	0, // 0: generated.SnippetService.CreateSnippet:input_type -> generated.CreateSnippetRequest
	1, // 1: generated.SnippetService.CreateSnippet:output_type -> generated.CreateSnippetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_impl_grpc_snippet_proto_init() }
func file_app_impl_grpc_snippet_proto_init() {
	if File_app_impl_grpc_snippet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_impl_grpc_snippet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnippetRequest); i {
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
		file_app_impl_grpc_snippet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnippetResponse); i {
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
			RawDescriptor: file_app_impl_grpc_snippet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_impl_grpc_snippet_proto_goTypes,
		DependencyIndexes: file_app_impl_grpc_snippet_proto_depIdxs,
		MessageInfos:      file_app_impl_grpc_snippet_proto_msgTypes,
	}.Build()
	File_app_impl_grpc_snippet_proto = out.File
	file_app_impl_grpc_snippet_proto_rawDesc = nil
	file_app_impl_grpc_snippet_proto_goTypes = nil
	file_app_impl_grpc_snippet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SnippetServiceClient is the client API for SnippetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SnippetServiceClient interface {
	CreateSnippet(ctx context.Context, in *CreateSnippetRequest, opts ...grpc.CallOption) (*CreateSnippetResponse, error)
}

type snippetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnippetServiceClient(cc grpc.ClientConnInterface) SnippetServiceClient {
	return &snippetServiceClient{cc}
}

func (c *snippetServiceClient) CreateSnippet(ctx context.Context, in *CreateSnippetRequest, opts ...grpc.CallOption) (*CreateSnippetResponse, error) {
	out := new(CreateSnippetResponse)
	err := c.cc.Invoke(ctx, "/generated.SnippetService/CreateSnippet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnippetServiceServer is the server API for SnippetService service.
type SnippetServiceServer interface {
	CreateSnippet(context.Context, *CreateSnippetRequest) (*CreateSnippetResponse, error)
}

// UnimplementedSnippetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSnippetServiceServer struct {
}

func (*UnimplementedSnippetServiceServer) CreateSnippet(context.Context, *CreateSnippetRequest) (*CreateSnippetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnippet not implemented")
}

func RegisterSnippetServiceServer(s *grpc.Server, srv SnippetServiceServer) {
	s.RegisterService(&_SnippetService_serviceDesc, srv)
}

func _SnippetService_CreateSnippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnippetServiceServer).CreateSnippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.SnippetService/CreateSnippet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnippetServiceServer).CreateSnippet(ctx, req.(*CreateSnippetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SnippetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.SnippetService",
	HandlerType: (*SnippetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSnippet",
			Handler:    _SnippetService_CreateSnippet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/impl/grpc/snippet.proto",
}
