// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: supervisor/service.proto

package supervisorpb

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

type DeployRequest_Region int32

const (
	DeployRequest_euw DeployRequest_Region = 0
)

// Enum value maps for DeployRequest_Region.
var (
	DeployRequest_Region_name = map[int32]string{
		0: "euw",
	}
	DeployRequest_Region_value = map[string]int32{
		"euw": 0,
	}
)

func (x DeployRequest_Region) Enum() *DeployRequest_Region {
	p := new(DeployRequest_Region)
	*p = x
	return p
}

func (x DeployRequest_Region) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeployRequest_Region) Descriptor() protoreflect.EnumDescriptor {
	return file_supervisor_service_proto_enumTypes[0].Descriptor()
}

func (DeployRequest_Region) Type() protoreflect.EnumType {
	return &file_supervisor_service_proto_enumTypes[0]
}

func (x DeployRequest_Region) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeployRequest_Region.Descriptor instead.
func (DeployRequest_Region) EnumDescriptor() ([]byte, []int) {
	return file_supervisor_service_proto_rawDescGZIP(), []int{0, 0}
}

type DeployRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64                 `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Region *DeployRequest_Region `protobuf:"varint,2,opt,name=region,proto3,enum=service.DeployRequest_Region,oneof" json:"region,omitempty"`
}

func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployRequest) ProtoMessage() {}

func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployRequest.ProtoReflect.Descriptor instead.
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_supervisor_service_proto_rawDescGZIP(), []int{0}
}

func (x *DeployRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *DeployRequest) GetRegion() DeployRequest_Region {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return DeployRequest_euw
}

type DeployResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeployResponse) Reset() {
	*x = DeployResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployResponse) ProtoMessage() {}

func (x *DeployResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployResponse.ProtoReflect.Descriptor instead.
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return file_supervisor_service_proto_rawDescGZIP(), []int{1}
}

var File_supervisor_service_proto protoreflect.FileDescriptor

var file_supervisor_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x3a, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x22, 0x11, 0x0a, 0x06, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x65, 0x75, 0x77, 0x10, 0x00, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x50, 0x0a, 0x11, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6b, 0x75, 0x72, 0x61,
	0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_supervisor_service_proto_rawDescOnce sync.Once
	file_supervisor_service_proto_rawDescData = file_supervisor_service_proto_rawDesc
)

func file_supervisor_service_proto_rawDescGZIP() []byte {
	file_supervisor_service_proto_rawDescOnce.Do(func() {
		file_supervisor_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_supervisor_service_proto_rawDescData)
	})
	return file_supervisor_service_proto_rawDescData
}

var file_supervisor_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_supervisor_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_supervisor_service_proto_goTypes = []interface{}{
	(DeployRequest_Region)(0), // 0: service.DeployRequest.Region
	(*DeployRequest)(nil),     // 1: service.DeployRequest
	(*DeployResponse)(nil),    // 2: service.DeployResponse
}
var file_supervisor_service_proto_depIdxs = []int32{
	0, // 0: service.DeployRequest.region:type_name -> service.DeployRequest.Region
	1, // 1: service.SupervisorService.Deploy:input_type -> service.DeployRequest
	2, // 2: service.SupervisorService.Deploy:output_type -> service.DeployResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_supervisor_service_proto_init() }
func file_supervisor_service_proto_init() {
	if File_supervisor_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supervisor_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployRequest); i {
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
		file_supervisor_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployResponse); i {
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
	file_supervisor_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_supervisor_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supervisor_service_proto_goTypes,
		DependencyIndexes: file_supervisor_service_proto_depIdxs,
		EnumInfos:         file_supervisor_service_proto_enumTypes,
		MessageInfos:      file_supervisor_service_proto_msgTypes,
	}.Build()
	File_supervisor_service_proto = out.File
	file_supervisor_service_proto_rawDesc = nil
	file_supervisor_service_proto_goTypes = nil
	file_supervisor_service_proto_depIdxs = nil
}
