// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: com/pistonint/grpc/core/result.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 返回实体
type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回状态码 0: 成功标记， 1: 失败标记， 2: 没有数据
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 查询结果集
	Data *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	// 请求unique id
	ReqId string `protobuf:"bytes,4,opt,name=reqId,proto3" json:"reqId,omitempty"`
	// 服务执行用时，单位毫秒
	Took int64 `protobuf:"varint,5,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_result_proto_rawDescGZIP(), []int{0}
}

func (x *Resp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Resp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Resp) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Resp) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *Resp) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

// 分页返回结果
type PageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回状态码 0: 成功标记， 1: 失败标记， 2: 没有数据
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 查询结果集
	Data *Page `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	// 请求unique id
	ReqId string `protobuf:"bytes,4,opt,name=reqId,proto3" json:"reqId,omitempty"`
	// 服务执行用时，单位毫秒
	Took int64 `protobuf:"varint,5,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *PageResp) Reset() {
	*x = PageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResp) ProtoMessage() {}

func (x *PageResp) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResp.ProtoReflect.Descriptor instead.
func (*PageResp) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_result_proto_rawDescGZIP(), []int{1}
}

func (x *PageResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PageResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PageResp) GetData() *Page {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PageResp) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *PageResp) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

var File_com_pistonint_grpc_core_result_proto protoreflect.FileDescriptor

var file_com_pistonint_grpc_core_result_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74,
	0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e,
	0x01, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x9b, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x71, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x6f, 0x6f, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x3c, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x09, 0x52, 0x65, 0x73, 0x70, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_pistonint_grpc_core_result_proto_rawDescOnce sync.Once
	file_com_pistonint_grpc_core_result_proto_rawDescData = file_com_pistonint_grpc_core_result_proto_rawDesc
)

func file_com_pistonint_grpc_core_result_proto_rawDescGZIP() []byte {
	file_com_pistonint_grpc_core_result_proto_rawDescOnce.Do(func() {
		file_com_pistonint_grpc_core_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_pistonint_grpc_core_result_proto_rawDescData)
	})
	return file_com_pistonint_grpc_core_result_proto_rawDescData
}

var file_com_pistonint_grpc_core_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_pistonint_grpc_core_result_proto_goTypes = []interface{}{
	(*Resp)(nil),      // 0: com.pistonint.grpc.core.Resp
	(*PageResp)(nil),  // 1: com.pistonint.grpc.core.PageResp
	(*anypb.Any)(nil), // 2: google.protobuf.Any
	(*Page)(nil),      // 3: com.pistonint.grpc.core.Page
}
var file_com_pistonint_grpc_core_result_proto_depIdxs = []int32{
	2, // 0: com.pistonint.grpc.core.Resp.data:type_name -> google.protobuf.Any
	3, // 1: com.pistonint.grpc.core.PageResp.data:type_name -> com.pistonint.grpc.core.Page
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_pistonint_grpc_core_result_proto_init() }
func file_com_pistonint_grpc_core_result_proto_init() {
	if File_com_pistonint_grpc_core_result_proto != nil {
		return
	}
	file_com_pistonint_grpc_core_page_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_pistonint_grpc_core_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
		file_com_pistonint_grpc_core_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResp); i {
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
	file_com_pistonint_grpc_core_result_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_com_pistonint_grpc_core_result_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_pistonint_grpc_core_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_pistonint_grpc_core_result_proto_goTypes,
		DependencyIndexes: file_com_pistonint_grpc_core_result_proto_depIdxs,
		MessageInfos:      file_com_pistonint_grpc_core_result_proto_msgTypes,
	}.Build()
	File_com_pistonint_grpc_core_result_proto = out.File
	file_com_pistonint_grpc_core_result_proto_rawDesc = nil
	file_com_pistonint_grpc_core_result_proto_goTypes = nil
	file_com_pistonint_grpc_core_result_proto_depIdxs = nil
}