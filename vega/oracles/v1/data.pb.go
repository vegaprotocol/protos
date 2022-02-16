// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vega/oracles/v1/data.proto

package v1

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

// OracleData describes an oracle data that has been broadcast.
type OracleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pubKeys is the list of authorized public keys that signed the data for this
	// oracle. All the public keys in the oracle data should be contained in these
	// public keys.
	PubKeys []string `protobuf:"bytes,1,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys,omitempty"`
	// data holds all the properties of the oracle data
	Data []*Property `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	// matched_specs_ids lists all the oracle specs that matched this oracle data.
	MatchedSpecIds []string `protobuf:"bytes,3,rep,name=matched_spec_ids,json=matchedSpecIds,proto3" json:"matched_spec_ids,omitempty"`
	// broadcast_at is the time at which the data was broadcast for the first
	// time.
	BroadcastAt int64 `protobuf:"varint,4,opt,name=broadcast_at,json=broadcastAt,proto3" json:"broadcast_at,omitempty"`
}

func (x *OracleData) Reset() {
	*x = OracleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_oracles_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleData) ProtoMessage() {}

func (x *OracleData) ProtoReflect() protoreflect.Message {
	mi := &file_vega_oracles_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleData.ProtoReflect.Descriptor instead.
func (*OracleData) Descriptor() ([]byte, []int) {
	return file_vega_oracles_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *OracleData) GetPubKeys() []string {
	if x != nil {
		return x.PubKeys
	}
	return nil
}

func (x *OracleData) GetData() []*Property {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OracleData) GetMatchedSpecIds() []string {
	if x != nil {
		return x.MatchedSpecIds
	}
	return nil
}

func (x *OracleData) GetBroadcastAt() int64 {
	if x != nil {
		return x.BroadcastAt
	}
	return 0
}

// Property describes one property of an oracle spec with a key with its value.
type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the property.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// value is the value of the property.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_oracles_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_vega_oracles_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_vega_oracles_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *Property) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Property) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_vega_oracles_v1_data_proto protoreflect.FileDescriptor

var file_vega_oracles_v1_data_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x4f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x75, 0x62, 0x5f, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x10,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x53,
	0x70, 0x65, 0x63, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x41, 0x74, 0x22, 0x34, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x2d, 0x5a, 0x2b, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76,
	0x65, 0x67, 0x61, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vega_oracles_v1_data_proto_rawDescOnce sync.Once
	file_vega_oracles_v1_data_proto_rawDescData = file_vega_oracles_v1_data_proto_rawDesc
)

func file_vega_oracles_v1_data_proto_rawDescGZIP() []byte {
	file_vega_oracles_v1_data_proto_rawDescOnce.Do(func() {
		file_vega_oracles_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_vega_oracles_v1_data_proto_rawDescData)
	})
	return file_vega_oracles_v1_data_proto_rawDescData
}

var file_vega_oracles_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vega_oracles_v1_data_proto_goTypes = []interface{}{
	(*OracleData)(nil), // 0: oracles.v1.OracleData
	(*Property)(nil),   // 1: oracles.v1.Property
}
var file_vega_oracles_v1_data_proto_depIdxs = []int32{
	1, // 0: oracles.v1.OracleData.data:type_name -> oracles.v1.Property
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vega_oracles_v1_data_proto_init() }
func file_vega_oracles_v1_data_proto_init() {
	if File_vega_oracles_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vega_oracles_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleData); i {
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
		file_vega_oracles_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
			RawDescriptor: file_vega_oracles_v1_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vega_oracles_v1_data_proto_goTypes,
		DependencyIndexes: file_vega_oracles_v1_data_proto_depIdxs,
		MessageInfos:      file_vega_oracles_v1_data_proto_msgTypes,
	}.Build()
	File_vega_oracles_v1_data_proto = out.File
	file_vega_oracles_v1_data_proto_rawDesc = nil
	file_vega_oracles_v1_data_proto_goTypes = nil
	file_vega_oracles_v1_data_proto_depIdxs = nil
}
