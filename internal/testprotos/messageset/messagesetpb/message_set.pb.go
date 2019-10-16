// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messageset/messagesetpb/message_set.proto

package messagesetpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type MessageSet struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields
}

func (x *MessageSet) Reset() {
	*x = MessageSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messageset_messagesetpb_message_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSet) ProtoMessage() {}

func (x *MessageSet) ProtoReflect() protoreflect.Message {
	mi := &file_messageset_messagesetpb_message_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSet.ProtoReflect.Descriptor instead.
func (*MessageSet) Descriptor() ([]byte, []int) {
	return file_messageset_messagesetpb_message_set_proto_rawDescGZIP(), []int{0}
}

var extRange_MessageSet = []protoiface.ExtensionRangeV1{
	{Start: 4, End: 2147483646},
}

// Deprecated: Use MessageSet.ProtoReflect.Descriptor.ExtensionRanges instead.
func (*MessageSet) ExtensionRangeArray() []protoiface.ExtensionRangeV1 {
	return extRange_MessageSet
}

var File_messageset_messagesetpb_message_set_proto protoreflect.FileDescriptor

var file_messageset_messagesetpb_message_set_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x65, 0x74, 0x22, 0x1a, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x2a, 0x08, 0x08, 0x04, 0x10, 0xff, 0xff, 0xff, 0xff, 0x07, 0x3a, 0x02, 0x08,
	0x01, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x70, 0x62,
}

var (
	file_messageset_messagesetpb_message_set_proto_rawDescOnce sync.Once
	file_messageset_messagesetpb_message_set_proto_rawDescData = file_messageset_messagesetpb_message_set_proto_rawDesc
)

func file_messageset_messagesetpb_message_set_proto_rawDescGZIP() []byte {
	file_messageset_messagesetpb_message_set_proto_rawDescOnce.Do(func() {
		file_messageset_messagesetpb_message_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_messageset_messagesetpb_message_set_proto_rawDescData)
	})
	return file_messageset_messagesetpb_message_set_proto_rawDescData
}

var file_messageset_messagesetpb_message_set_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messageset_messagesetpb_message_set_proto_goTypes = []interface{}{
	(*MessageSet)(nil), // 0: goproto.proto.messageset.MessageSet
}
var file_messageset_messagesetpb_message_set_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messageset_messagesetpb_message_set_proto_init() }
func file_messageset_messagesetpb_message_set_proto_init() {
	if File_messageset_messagesetpb_message_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messageset_messagesetpb_message_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messageset_messagesetpb_message_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messageset_messagesetpb_message_set_proto_goTypes,
		DependencyIndexes: file_messageset_messagesetpb_message_set_proto_depIdxs,
		MessageInfos:      file_messageset_messagesetpb_message_set_proto_msgTypes,
	}.Build()
	File_messageset_messagesetpb_message_set_proto = out.File
	file_messageset_messagesetpb_message_set_proto_rawDesc = nil
	file_messageset_messagesetpb_message_set_proto_goTypes = nil
	file_messageset_messagesetpb_message_set_proto_depIdxs = nil
}
