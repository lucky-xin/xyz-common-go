// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: com/pistonint/grpc/query/mongo/searcher.proto

package mongo

import (
	"github.com/lucky-xin/xyz-common-go/query"
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

// 搜索条件实体
type Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 页码（获取第几页数据）,分页从1开始，current必须大于等于1
	Current int32 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	// 分页大小最大限制为500
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// 查询返回字段
	Includes []string `protobuf:"bytes,3,rep,name=includes,proto3" json:"includes,omitempty"`
	// 查询结果集不返回字段
	Excludes []string `protobuf:"bytes,4,rep,name=excludes,proto3" json:"excludes,omitempty"`
	// 排序字段
	Sort *query.Sort `protobuf:"bytes,5,opt,name=sort,proto3,oneof" json:"sort,omitempty"`
	// 查询条件内容
	Condition *Criteria `protobuf:"bytes,6,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *Search) Reset() {
	*x = Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_query_mongo_searcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Search) ProtoMessage() {}

func (x *Search) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_query_mongo_searcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Search.ProtoReflect.Descriptor instead.
func (*Search) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_query_mongo_searcher_proto_rawDescGZIP(), []int{0}
}

func (x *Search) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *Search) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Search) GetIncludes() []string {
	if x != nil {
		return x.Includes
	}
	return nil
}

func (x *Search) GetExcludes() []string {
	if x != nil {
		return x.Excludes
	}
	return nil
}

func (x *Search) GetSort() *query.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *Search) GetCondition() *Criteria {
	if x != nil {
		return x.Condition
	}
	return nil
}

var File_com_pistonint_grpc_query_mongo_searcher_proto protoreflect.FileDescriptor

var file_com_pistonint_grpc_query_mongo_searcher_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x1a,
	0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e,
	0x69, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf8, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x73, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x4e, 0x0a, 0x1e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x42, 0x0d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1b,
	0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_pistonint_grpc_query_mongo_searcher_proto_rawDescOnce sync.Once
	file_com_pistonint_grpc_query_mongo_searcher_proto_rawDescData = file_com_pistonint_grpc_query_mongo_searcher_proto_rawDesc
)

func file_com_pistonint_grpc_query_mongo_searcher_proto_rawDescGZIP() []byte {
	file_com_pistonint_grpc_query_mongo_searcher_proto_rawDescOnce.Do(func() {
		file_com_pistonint_grpc_query_mongo_searcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_pistonint_grpc_query_mongo_searcher_proto_rawDescData)
	})
	return file_com_pistonint_grpc_query_mongo_searcher_proto_rawDescData
}

var file_com_pistonint_grpc_query_mongo_searcher_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_pistonint_grpc_query_mongo_searcher_proto_goTypes = []interface{}{
	(*Search)(nil),     // 0: com.pistonint.grpc.query.mongo.Search
	(*query.Sort)(nil), // 1: com.pistonint.grpc.query.Sort
	(*Criteria)(nil),   // 2: com.pistonint.grpc.query.mongo.Criteria
}
var file_com_pistonint_grpc_query_mongo_searcher_proto_depIdxs = []int32{
	1, // 0: com.pistonint.grpc.query.mongo.Search.sort:type_name -> com.pistonint.grpc.query.Sort
	2, // 1: com.pistonint.grpc.query.mongo.Search.condition:type_name -> com.pistonint.grpc.query.mongo.Criteria
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_pistonint_grpc_query_mongo_searcher_proto_init() }
func file_com_pistonint_grpc_query_mongo_searcher_proto_init() {
	if File_com_pistonint_grpc_query_mongo_searcher_proto != nil {
		return
	}
	file_com_pistonint_grpc_query_mongo_query_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_pistonint_grpc_query_mongo_searcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Search); i {
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
	file_com_pistonint_grpc_query_mongo_searcher_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_pistonint_grpc_query_mongo_searcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_pistonint_grpc_query_mongo_searcher_proto_goTypes,
		DependencyIndexes: file_com_pistonint_grpc_query_mongo_searcher_proto_depIdxs,
		MessageInfos:      file_com_pistonint_grpc_query_mongo_searcher_proto_msgTypes,
	}.Build()
	File_com_pistonint_grpc_query_mongo_searcher_proto = out.File
	file_com_pistonint_grpc_query_mongo_searcher_proto_rawDesc = nil
	file_com_pistonint_grpc_query_mongo_searcher_proto_goTypes = nil
	file_com_pistonint_grpc_query_mongo_searcher_proto_depIdxs = nil
}
