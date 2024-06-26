// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: com/pistonint/grpc/query/mongo/export_info.proto

package mongo

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

// 搜索条件实体
type ExportRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 导出记录主键
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 记录md5
	Md5 string `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5,omitempty"`
	// 租户id
	TenantId int32 `protobuf:"varint,3,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// 总数据量
	Total int64 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	// 查询条件
	Searcher string `protobuf:"bytes,5,opt,name=searcher,proto3" json:"searcher,omitempty"`
	// oss bucket
	OssBucket string `protobuf:"bytes,6,opt,name=ossBucket,proto3" json:"ossBucket,omitempty"`
	// oss key
	OssKey string `protobuf:"bytes,7,opt,name=ossKey,proto3" json:"ossKey,omitempty"`
	// db name
	DbName string `protobuf:"bytes,8,opt,name=dbName,proto3" json:"dbName,omitempty"`
	// collection name
	CollectionName string `protobuf:"bytes,9,opt,name=collectionName,proto3" json:"collectionName,omitempty"`
	// zip文件名称 如 brand.zip
	FileName string `protobuf:"bytes,10,opt,name=fileName,proto3" json:"fileName,omitempty"`
	// 创建时间
	CreateTime string `protobuf:"bytes,11,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// 状态
	Status int32 `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ExportRecord) Reset() {
	*x = ExportRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_query_mongo_export_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRecord) ProtoMessage() {}

func (x *ExportRecord) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_query_mongo_export_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportRecord.ProtoReflect.Descriptor instead.
func (*ExportRecord) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_query_mongo_export_info_proto_rawDescGZIP(), []int{0}
}

func (x *ExportRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExportRecord) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *ExportRecord) GetTenantId() int32 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *ExportRecord) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ExportRecord) GetSearcher() string {
	if x != nil {
		return x.Searcher
	}
	return ""
}

func (x *ExportRecord) GetOssBucket() string {
	if x != nil {
		return x.OssBucket
	}
	return ""
}

func (x *ExportRecord) GetOssKey() string {
	if x != nil {
		return x.OssKey
	}
	return ""
}

func (x *ExportRecord) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *ExportRecord) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *ExportRecord) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ExportRecord) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ExportRecord) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_com_pistonint_grpc_query_mongo_export_info_proto protoreflect.FileDescriptor

var file_com_pistonint_grpc_query_mongo_export_info_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e,
	0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x73, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x50, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x42,
	0x0f, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x1b, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_pistonint_grpc_query_mongo_export_info_proto_rawDescOnce sync.Once
	file_com_pistonint_grpc_query_mongo_export_info_proto_rawDescData = file_com_pistonint_grpc_query_mongo_export_info_proto_rawDesc
)

func file_com_pistonint_grpc_query_mongo_export_info_proto_rawDescGZIP() []byte {
	file_com_pistonint_grpc_query_mongo_export_info_proto_rawDescOnce.Do(func() {
		file_com_pistonint_grpc_query_mongo_export_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_pistonint_grpc_query_mongo_export_info_proto_rawDescData)
	})
	return file_com_pistonint_grpc_query_mongo_export_info_proto_rawDescData
}

var file_com_pistonint_grpc_query_mongo_export_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_pistonint_grpc_query_mongo_export_info_proto_goTypes = []interface{}{
	(*ExportRecord)(nil), // 0: com.pistonint.grpc.query.mongo.ExportRecord
}
var file_com_pistonint_grpc_query_mongo_export_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_pistonint_grpc_query_mongo_export_info_proto_init() }
func file_com_pistonint_grpc_query_mongo_export_info_proto_init() {
	if File_com_pistonint_grpc_query_mongo_export_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_pistonint_grpc_query_mongo_export_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportRecord); i {
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
			RawDescriptor: file_com_pistonint_grpc_query_mongo_export_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_pistonint_grpc_query_mongo_export_info_proto_goTypes,
		DependencyIndexes: file_com_pistonint_grpc_query_mongo_export_info_proto_depIdxs,
		MessageInfos:      file_com_pistonint_grpc_query_mongo_export_info_proto_msgTypes,
	}.Build()
	File_com_pistonint_grpc_query_mongo_export_info_proto = out.File
	file_com_pistonint_grpc_query_mongo_export_info_proto_rawDesc = nil
	file_com_pistonint_grpc_query_mongo_export_info_proto_goTypes = nil
	file_com_pistonint_grpc_query_mongo_export_info_proto_depIdxs = nil
}
