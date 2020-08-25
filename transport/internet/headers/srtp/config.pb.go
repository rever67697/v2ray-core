package srtp

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Padding     bool   `protobuf:"varint,2,opt,name=padding,proto3" json:"padding,omitempty"`
	Extension   bool   `protobuf:"varint,3,opt,name=extension,proto3" json:"extension,omitempty"`
	CsrcCount   uint32 `protobuf:"varint,4,opt,name=csrc_count,json=csrcCount,proto3" json:"csrc_count,omitempty"`
	Marker      bool   `protobuf:"varint,5,opt,name=marker,proto3" json:"marker,omitempty"`
	PayloadType uint32 `protobuf:"varint,6,opt,name=payload_type,json=payloadType,proto3" json:"payload_type,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_transport_internet_headers_srtp_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_transport_internet_headers_srtp_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Config) GetPadding() bool {
	if x != nil {
		return x.Padding
	}
	return false
}

func (x *Config) GetExtension() bool {
	if x != nil {
		return x.Extension
	}
	return false
}

func (x *Config) GetCsrcCount() uint32 {
	if x != nil {
		return x.CsrcCount
	}
	return 0
}

func (x *Config) GetMarker() bool {
	if x != nil {
		return x.Marker
	}
	return false
}

func (x *Config) GetPayloadType() uint32 {
	if x != nil {
		return x.PayloadType
	}
	return 0
}

var File_v2ray_com_core_transport_internet_headers_srtp_config_proto protoreflect.FileDescriptor

var file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x72, 0x74, 0x70,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x72, 0x74, 0x70, 0x22, 0xb4, 0x01, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x73, 0x72, 0x63, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x73, 0x72, 0x63,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x65, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x72,
	0x74, 0x70, 0x50, 0x01, 0x5a, 0x04, 0x73, 0x72, 0x74, 0x70, 0xaa, 0x02, 0x2a, 0x56, 0x32, 0x52,
	0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x53, 0x72, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDescOnce sync.Once
	file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDescData = file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDesc
)

func file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDescGZIP() []byte {
	file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDescOnce.Do(func() {
		file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDescData)
	})
	return file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDescData
}

var file_v2ray_com_core_transport_internet_headers_srtp_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v2ray_com_core_transport_internet_headers_srtp_config_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: v2ray.core.transport.internet.headers.srtp.Config
}
var file_v2ray_com_core_transport_internet_headers_srtp_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v2ray_com_core_transport_internet_headers_srtp_config_proto_init() }
func file_v2ray_com_core_transport_internet_headers_srtp_config_proto_init() {
	if File_v2ray_com_core_transport_internet_headers_srtp_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2ray_com_core_transport_internet_headers_srtp_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2ray_com_core_transport_internet_headers_srtp_config_proto_goTypes,
		DependencyIndexes: file_v2ray_com_core_transport_internet_headers_srtp_config_proto_depIdxs,
		MessageInfos:      file_v2ray_com_core_transport_internet_headers_srtp_config_proto_msgTypes,
	}.Build()
	File_v2ray_com_core_transport_internet_headers_srtp_config_proto = out.File
	file_v2ray_com_core_transport_internet_headers_srtp_config_proto_rawDesc = nil
	file_v2ray_com_core_transport_internet_headers_srtp_config_proto_goTypes = nil
	file_v2ray_com_core_transport_internet_headers_srtp_config_proto_depIdxs = nil
}
