// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messageset/msetextpb/msetextpb.proto

package msetextpb

import (
	messagesetpb "google.golang.org/protobuf/internal/testprotos/messageset/messagesetpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Ext1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ext1Field1 *int32 `protobuf:"varint,1,opt,name=ext1_field1,json=ext1Field1" json:"ext1_field1,omitempty"`
	Ext1Field2 *int32 `protobuf:"varint,2,opt,name=ext1_field2,json=ext1Field2" json:"ext1_field2,omitempty"`
}

func (x *Ext1) Reset() {
	*x = Ext1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messageset_msetextpb_msetextpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ext1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ext1) ProtoMessage() {}

func (x *Ext1) ProtoReflect() protoreflect.Message {
	mi := &file_messageset_msetextpb_msetextpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ext1.ProtoReflect.Descriptor instead.
func (*Ext1) Descriptor() ([]byte, []int) {
	return file_messageset_msetextpb_msetextpb_proto_rawDescGZIP(), []int{0}
}

func (x *Ext1) GetExt1Field1() int32 {
	if x != nil && x.Ext1Field1 != nil {
		return *x.Ext1Field1
	}
	return 0
}

func (x *Ext1) GetExt1Field2() int32 {
	if x != nil && x.Ext1Field2 != nil {
		return *x.Ext1Field2
	}
	return 0
}

type Ext2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ext2Field1 *int32 `protobuf:"varint,1,opt,name=ext2_field1,json=ext2Field1" json:"ext2_field1,omitempty"`
}

func (x *Ext2) Reset() {
	*x = Ext2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messageset_msetextpb_msetextpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ext2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ext2) ProtoMessage() {}

func (x *Ext2) ProtoReflect() protoreflect.Message {
	mi := &file_messageset_msetextpb_msetextpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ext2.ProtoReflect.Descriptor instead.
func (*Ext2) Descriptor() ([]byte, []int) {
	return file_messageset_msetextpb_msetextpb_proto_rawDescGZIP(), []int{1}
}

func (x *Ext2) GetExt2Field1() int32 {
	if x != nil && x.Ext2Field1 != nil {
		return *x.Ext2Field1
	}
	return 0
}

var file_messageset_msetextpb_msetextpb_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*messagesetpb.MessageSet)(nil),
		ExtensionType: (*Ext1)(nil),
		Field:         1000,
		Name:          "goproto.proto.messageset.Ext1",
		Tag:           "bytes,1000,opt,name=message_set_extension",
		Filename:      "messageset/msetextpb/msetextpb.proto",
	},
	{
		ExtendedType:  (*messagesetpb.MessageSet)(nil),
		ExtensionType: (*Ext2)(nil),
		Field:         1001,
		Name:          "goproto.proto.messageset.Ext2",
		Tag:           "bytes,1001,opt,name=message_set_extension",
		Filename:      "messageset/msetextpb/msetextpb.proto",
	},
}

// Extension fields to messagesetpb.MessageSet.
var (
	// optional goproto.proto.messageset.Ext1 message_set_extension = 1000;
	E_Ext1_MessageSetExtension = &file_messageset_msetextpb_msetextpb_proto_extTypes[0]
	// optional goproto.proto.messageset.Ext2 message_set_extension = 1001;
	E_Ext2_MessageSetExtension = &file_messageset_msetextpb_msetextpb_proto_extTypes[1]
)

var File_messageset_msetextpb_msetextpb_proto protoreflect.FileDescriptor

var file_messageset_msetextpb_msetextpb_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x6d, 0x73, 0x65,
	0x74, 0x65, 0x78, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x65, 0x74, 0x65, 0x78, 0x74, 0x70, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74,
	0x1a, 0x29, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x04,
	0x45, 0x78, 0x74, 0x31, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x31, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x31, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x31, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x31,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x79, 0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x31, 0x52, 0x13, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x04, 0x45, 0x78, 0x74, 0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x74, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x79, 0x0a, 0x15, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x74,
	0x32, 0x52, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x65, 0x74, 0x2f, 0x6d, 0x73, 0x65, 0x74, 0x65, 0x78, 0x74, 0x70, 0x62,
}

var (
	file_messageset_msetextpb_msetextpb_proto_rawDescOnce sync.Once
	file_messageset_msetextpb_msetextpb_proto_rawDescData = file_messageset_msetextpb_msetextpb_proto_rawDesc
)

func file_messageset_msetextpb_msetextpb_proto_rawDescGZIP() []byte {
	file_messageset_msetextpb_msetextpb_proto_rawDescOnce.Do(func() {
		file_messageset_msetextpb_msetextpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_messageset_msetextpb_msetextpb_proto_rawDescData)
	})
	return file_messageset_msetextpb_msetextpb_proto_rawDescData
}

var file_messageset_msetextpb_msetextpb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messageset_msetextpb_msetextpb_proto_goTypes = []interface{}{
	(*Ext1)(nil),                    // 0: goproto.proto.messageset.Ext1
	(*Ext2)(nil),                    // 1: goproto.proto.messageset.Ext2
	(*messagesetpb.MessageSet)(nil), // 2: goproto.proto.messageset.MessageSet
}
var file_messageset_msetextpb_msetextpb_proto_depIdxs = []int32{
	2, // 0: goproto.proto.messageset.Ext1.message_set_extension:extendee -> goproto.proto.messageset.MessageSet
	2, // 1: goproto.proto.messageset.Ext2.message_set_extension:extendee -> goproto.proto.messageset.MessageSet
	0, // 2: goproto.proto.messageset.Ext1.message_set_extension:type_name -> goproto.proto.messageset.Ext1
	1, // 3: goproto.proto.messageset.Ext2.message_set_extension:type_name -> goproto.proto.messageset.Ext2
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messageset_msetextpb_msetextpb_proto_init() }
func file_messageset_msetextpb_msetextpb_proto_init() {
	if File_messageset_msetextpb_msetextpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messageset_msetextpb_msetextpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ext1); i {
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
		file_messageset_msetextpb_msetextpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ext2); i {
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
			RawDescriptor: file_messageset_msetextpb_msetextpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_messageset_msetextpb_msetextpb_proto_goTypes,
		DependencyIndexes: file_messageset_msetextpb_msetextpb_proto_depIdxs,
		MessageInfos:      file_messageset_msetextpb_msetextpb_proto_msgTypes,
		ExtensionInfos:    file_messageset_msetextpb_msetextpb_proto_extTypes,
	}.Build()
	File_messageset_msetextpb_msetextpb_proto = out.File
	file_messageset_msetextpb_msetextpb_proto_rawDesc = nil
	file_messageset_msetextpb_msetextpb_proto_goTypes = nil
	file_messageset_msetextpb_msetextpb_proto_depIdxs = nil
}
