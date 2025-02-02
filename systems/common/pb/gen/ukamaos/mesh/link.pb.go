//*
// Copyright (c) 2021-present, Ukama Inc.
// All rights reserved.
//
// This source code is licensed under the XXX-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: pb/ukama/link.proto

package mesh

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

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId *string `protobuf:"bytes,1,req,name=nodeId" json:"nodeId,omitempty"` // NodeID of the device.
	Ip     *string `protobuf:"bytes,2,req,name=ip" json:"ip,omitempty"`         // IP address of the device.
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_ukama_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ukama_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_pb_ukama_link_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetNodeId() string {
	if x != nil && x.NodeId != nil {
		return *x.NodeId
	}
	return ""
}

func (x *Link) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

var File_pb_ukama_link_proto protoreflect.FileDescriptor

var file_pb_ukama_link_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x75, 0x6b, 0x61, 0x6d, 0x61, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x68,
}

var (
	file_pb_ukama_link_proto_rawDescOnce sync.Once
	file_pb_ukama_link_proto_rawDescData = file_pb_ukama_link_proto_rawDesc
)

func file_pb_ukama_link_proto_rawDescGZIP() []byte {
	file_pb_ukama_link_proto_rawDescOnce.Do(func() {
		file_pb_ukama_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_ukama_link_proto_rawDescData)
	})
	return file_pb_ukama_link_proto_rawDescData
}

var file_pb_ukama_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_ukama_link_proto_goTypes = []interface{}{
	(*Link)(nil), // 0: Link
}
var file_pb_ukama_link_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_ukama_link_proto_init() }
func file_pb_ukama_link_proto_init() {
	if File_pb_ukama_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_ukama_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
			RawDescriptor: file_pb_ukama_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_ukama_link_proto_goTypes,
		DependencyIndexes: file_pb_ukama_link_proto_depIdxs,
		MessageInfos:      file_pb_ukama_link_proto_msgTypes,
	}.Build()
	File_pb_ukama_link_proto = out.File
	file_pb_ukama_link_proto_rawDesc = nil
	file_pb_ukama_link_proto_goTypes = nil
	file_pb_ukama_link_proto_depIdxs = nil
}
