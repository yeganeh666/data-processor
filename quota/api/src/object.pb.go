// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: quota/api/object.proto

package api

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

type PreUploadObjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"key"
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"key"`
	// @inject_tag: json:"user_id"
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"user_id"`
	// @inject_tag: json:"size"
	Size int64 `protobuf:"varint,3,opt,name=Size,proto3" json:"size"`
}

func (x *PreUploadObjectReq) Reset() {
	*x = PreUploadObjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quota_api_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreUploadObjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreUploadObjectReq) ProtoMessage() {}

func (x *PreUploadObjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_quota_api_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreUploadObjectReq.ProtoReflect.Descriptor instead.
func (*PreUploadObjectReq) Descriptor() ([]byte, []int) {
	return file_quota_api_object_proto_rawDescGZIP(), []int{0}
}

func (x *PreUploadObjectReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PreUploadObjectReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PreUploadObjectReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type PreUploadObjectRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"key"
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"key"`
	// @inject_tag: json:"user_id"
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"user_id"`
	// @inject_tag: json:"upload_id"
	UploadID string `protobuf:"bytes,3,opt,name=UploadID,proto3" json:"upload_id"`
}

func (x *PreUploadObjectRes) Reset() {
	*x = PreUploadObjectRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quota_api_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreUploadObjectRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreUploadObjectRes) ProtoMessage() {}

func (x *PreUploadObjectRes) ProtoReflect() protoreflect.Message {
	mi := &file_quota_api_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreUploadObjectRes.ProtoReflect.Descriptor instead.
func (*PreUploadObjectRes) Descriptor() ([]byte, []int) {
	return file_quota_api_object_proto_rawDescGZIP(), []int{1}
}

func (x *PreUploadObjectRes) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PreUploadObjectRes) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PreUploadObjectRes) GetUploadID() string {
	if x != nil {
		return x.UploadID
	}
	return ""
}

var File_quota_api_object_proto protoreflect.FileDescriptor

var file_quota_api_object_proto_rawDesc = []byte{
	0x0a, 0x16, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f,
	0x61, 0x70, 0x69, 0x22, 0x52, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x44, 0x32, 0x5a, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1d, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x1d, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quota_api_object_proto_rawDescOnce sync.Once
	file_quota_api_object_proto_rawDescData = file_quota_api_object_proto_rawDesc
)

func file_quota_api_object_proto_rawDescGZIP() []byte {
	file_quota_api_object_proto_rawDescOnce.Do(func() {
		file_quota_api_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_quota_api_object_proto_rawDescData)
	})
	return file_quota_api_object_proto_rawDescData
}

var file_quota_api_object_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_quota_api_object_proto_goTypes = []interface{}{
	(*PreUploadObjectReq)(nil), // 0: quota_api.PreUploadObjectReq
	(*PreUploadObjectRes)(nil), // 1: quota_api.PreUploadObjectRes
}
var file_quota_api_object_proto_depIdxs = []int32{
	0, // 0: quota_api.ObjectService.PreUpload:input_type -> quota_api.PreUploadObjectReq
	1, // 1: quota_api.ObjectService.PreUpload:output_type -> quota_api.PreUploadObjectRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_quota_api_object_proto_init() }
func file_quota_api_object_proto_init() {
	if File_quota_api_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quota_api_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreUploadObjectReq); i {
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
		file_quota_api_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreUploadObjectRes); i {
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
			RawDescriptor: file_quota_api_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quota_api_object_proto_goTypes,
		DependencyIndexes: file_quota_api_object_proto_depIdxs,
		MessageInfos:      file_quota_api_object_proto_msgTypes,
	}.Build()
	File_quota_api_object_proto = out.File
	file_quota_api_object_proto_rawDesc = nil
	file_quota_api_object_proto_goTypes = nil
	file_quota_api_object_proto_depIdxs = nil
}