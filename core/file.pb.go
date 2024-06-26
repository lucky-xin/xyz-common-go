// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: com/pistonint/grpc/core/file.proto

package core

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

// 单个文件
type FileEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// 存储卷名称
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// *
	// 存储key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// *
	// 文件名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// *
	// 存储数据
	Bytes []byte `protobuf:"bytes,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *FileEntry) Reset() {
	*x = FileEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileEntry) ProtoMessage() {}

func (x *FileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileEntry.ProtoReflect.Descriptor instead.
func (*FileEntry) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileEntry) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *FileEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FileEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileEntry) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

// 单个文件返回实体
type FileEntryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回状态码 0: 成功标记， 1: 失败标记， 2: 没有数据
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 查询结果集
	Data *FileEntry `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	// 请求unique id
	ReqId string `protobuf:"bytes,4,opt,name=reqId,proto3" json:"reqId,omitempty"`
	// 服务执行用时，单位毫秒
	Took int64 `protobuf:"varint,5,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *FileEntryResp) Reset() {
	*x = FileEntryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileEntryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileEntryResp) ProtoMessage() {}

func (x *FileEntryResp) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileEntryResp.ProtoReflect.Descriptor instead.
func (*FileEntryResp) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileEntryResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FileEntryResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FileEntryResp) GetData() *FileEntry {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FileEntryResp) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *FileEntryResp) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// 存储卷名称
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// *
	// 存储key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{2}
}

func (x *Profile) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Profile) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// 分段文件信息
type ChunkFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// 上传请求id
	UploadId string `protobuf:"bytes,1,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	// *
	// 存储卷bucket
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// *
	// 存储 Object key
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// *
	// 文件分片字节数组
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// *
	// 过期时间
	ExpirationTime int64 `protobuf:"varint,5,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// *
	// 当前分段下标 从0开始。如1000字节大小文件整个文件分段下标为0-999
	Index int32 `protobuf:"varint,6,opt,name=index,proto3" json:"index,omitempty"`
	// *
	// 总分段数量
	Parts int32 `protobuf:"varint,7,opt,name=parts,proto3" json:"parts,omitempty"`
	// *
	// The predefined Access Control List (ACL) default: This is only for object, means the permission inherits the
	// bucket's permission. private: The owner has the Permission, other {AllUsers} does not have access. public-read:
	// The owner has the Permission, other {allUsers} have read-only access. public-read-write: Both the owner and
	// allUsers have fullControl. It's not safe and thus not recommended.
	Acl string `protobuf:"bytes,8,opt,name=acl,proto3" json:"acl,omitempty"`
	// *
	// 文件名称
	Name string `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	// *
	// 文件类型
	ContentType string `protobuf:"bytes,10,opt,name=contentType,proto3" json:"contentType,omitempty"`
}

func (x *ChunkFile) Reset() {
	*x = ChunkFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkFile) ProtoMessage() {}

func (x *ChunkFile) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkFile.ProtoReflect.Descriptor instead.
func (*ChunkFile) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{3}
}

func (x *ChunkFile) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *ChunkFile) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *ChunkFile) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ChunkFile) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChunkFile) GetExpirationTime() int64 {
	if x != nil {
		return x.ExpirationTime
	}
	return 0
}

func (x *ChunkFile) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ChunkFile) GetParts() int32 {
	if x != nil {
		return x.Parts
	}
	return 0
}

func (x *ChunkFile) GetAcl() string {
	if x != nil {
		return x.Acl
	}
	return ""
}

func (x *ChunkFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChunkFile) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type BatchFileInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// 存储卷bucket
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// *
	// 存储 Object key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// *
	// Optional parameter that causes keys that contain the same string between
	// the prefix and the first occurrence of the delimiter to be rolled up into
	// a single result element in the {@link ListObjectsV2Result#getCommonPrefixes()} list.
	// The most commonly used delimiter is "/", which simulates a hierarchical organization similar to
	// a file system directory structure.
	Delimiter string `protobuf:"bytes,3,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	// *
	// Optional parameter indicating the encoding method to be applied on the
	// response. An object key can contain any Unicode character; however, XML
	// 1.0 parser cannot parse some characters, such as characters with an ASCII
	// value from 0 to 10. you can add this parameter to request that OSS encode the keys in the
	// response.
	EncodingType string `protobuf:"bytes,4,opt,name=encodingType,proto3" json:"encodingType,omitempty"`
	// *
	// Optional parameter indicating the maximum number of keys to include in
	// the response.
	MaxKeys int32 `protobuf:"varint,5,opt,name=maxKeys,proto3" json:"maxKeys,omitempty"`
	// *
	// Optional parameter restricting the response to keys which begin with the
	// specified prefix.
	Prefix string `protobuf:"bytes,6,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// *
	// Optional parameter which allows list to be continued from a specific point.
	// ContinuationToken is provided in truncated list results.
	ContinuationToken string `protobuf:"bytes,7,opt,name=continuationToken,proto3" json:"continuationToken,omitempty"`
	// *
	// The owner field is not present in ListObjectsV2 results by default. If this flag is
	// set to true the owner field in {@link OSSObjectSummary} will be included.
	FetchOwner bool `protobuf:"varint,8,opt,name=fetchOwner,proto3" json:"fetchOwner,omitempty"`
	// *
	// Optional parameter indicating where you want OSS to start the object listing
	// from.
	StartAfter string `protobuf:"bytes,9,opt,name=startAfter,proto3" json:"startAfter,omitempty"`
}

func (x *BatchFileInfoReq) Reset() {
	*x = BatchFileInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchFileInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchFileInfoReq) ProtoMessage() {}

func (x *BatchFileInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchFileInfoReq.ProtoReflect.Descriptor instead.
func (*BatchFileInfoReq) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{4}
}

func (x *BatchFileInfoReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *BatchFileInfoReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *BatchFileInfoReq) GetDelimiter() string {
	if x != nil {
		return x.Delimiter
	}
	return ""
}

func (x *BatchFileInfoReq) GetEncodingType() string {
	if x != nil {
		return x.EncodingType
	}
	return ""
}

func (x *BatchFileInfoReq) GetMaxKeys() int32 {
	if x != nil {
		return x.MaxKeys
	}
	return 0
}

func (x *BatchFileInfoReq) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *BatchFileInfoReq) GetContinuationToken() string {
	if x != nil {
		return x.ContinuationToken
	}
	return ""
}

func (x *BatchFileInfoReq) GetFetchOwner() bool {
	if x != nil {
		return x.FetchOwner
	}
	return false
}

func (x *BatchFileInfoReq) GetStartAfter() string {
	if x != nil {
		return x.StartAfter
	}
	return ""
}

// 分段文件查询请求实体
type ChunkFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// 存储卷bucket
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// *
	// 存储 Object key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// *
	// 对于大小为1000 Bytes的文件，正常的字节范围为0~999，开始下标为0
	Start int64 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	// *
	// 对于大小为1000 Bytes的文件，正常的字节范围为0~999，结束下标为999
	End int64 `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ChunkFileReq) Reset() {
	*x = ChunkFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkFileReq) ProtoMessage() {}

func (x *ChunkFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkFileReq.ProtoReflect.Descriptor instead.
func (*ChunkFileReq) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{5}
}

func (x *ChunkFileReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *ChunkFileReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ChunkFileReq) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ChunkFileReq) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type BatchFileInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回状态码 0: 成功标记， 1: 失败标记， 2: 没有数据
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 查询结果集
	Data []*OSSObjSummary `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	// 请求unique id
	ReqId string `protobuf:"bytes,4,opt,name=reqId,proto3" json:"reqId,omitempty"`
	// 服务执行用时，单位毫秒
	Took int64 `protobuf:"varint,5,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *BatchFileInfoResp) Reset() {
	*x = BatchFileInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchFileInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchFileInfoResp) ProtoMessage() {}

func (x *BatchFileInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchFileInfoResp.ProtoReflect.Descriptor instead.
func (*BatchFileInfoResp) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{6}
}

func (x *BatchFileInfoResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BatchFileInfoResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BatchFileInfoResp) GetData() []*OSSObjSummary {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BatchFileInfoResp) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *BatchFileInfoResp) GetTook() int64 {
	if x != nil {
		return x.Took
	}
	return 0
}

// 文件信息
type OSSObjSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// 存储卷bucket
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// *
	// 存储 Object key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// *
	// E tag
	ETag string `protobuf:"bytes,3,opt,name=eTag,proto3" json:"eTag,omitempty"`
	// *
	// 文件大小
	Size int64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// *
	// 最后修改时间（毫秒）
	LastModified int64 `protobuf:"varint,5,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	// *
	// 存储 class
	StorageClass string `protobuf:"bytes,6,opt,name=storageClass,proto3" json:"storageClass,omitempty"`
	// *
	// 文件类型
	Type string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *OSSObjSummary) Reset() {
	*x = OSSObjSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSSObjSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSSObjSummary) ProtoMessage() {}

func (x *OSSObjSummary) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSSObjSummary.ProtoReflect.Descriptor instead.
func (*OSSObjSummary) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{7}
}

func (x *OSSObjSummary) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *OSSObjSummary) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *OSSObjSummary) GetETag() string {
	if x != nil {
		return x.ETag
	}
	return ""
}

func (x *OSSObjSummary) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *OSSObjSummary) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

func (x *OSSObjSummary) GetStorageClass() string {
	if x != nil {
		return x.StorageClass
	}
	return ""
}

func (x *OSSObjSummary) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AbortMultUploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// 分段上传请求id
	UploadId string `protobuf:"bytes,1,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	// *
	// 存储卷bucket
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// *
	// 存储 Object key
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *AbortMultUploadReq) Reset() {
	*x = AbortMultUploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbortMultUploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbortMultUploadReq) ProtoMessage() {}

func (x *AbortMultUploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_pistonint_grpc_core_file_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbortMultUploadReq.ProtoReflect.Descriptor instead.
func (*AbortMultUploadReq) Descriptor() ([]byte, []int) {
	return file_com_pistonint_grpc_core_file_proto_rawDescGZIP(), []int{8}
}

func (x *AbortMultUploadReq) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *AbortMultUploadReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *AbortMultUploadReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_com_pistonint_grpc_core_file_proto protoreflect.FileDescriptor

var file_com_pistonint_grpc_core_file_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e,
	0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x5f, 0x0a,
	0x09, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0xa5,
	0x01, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f,
	0x6e, 0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x81, 0x02, 0x0a, 0x09,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x9e, 0x02, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x22, 0x60, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x53, 0x53, 0x4f, 0x62, 0x6a, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x71, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x6f, 0x6f, 0x6b, 0x22, 0xbd, 0x01, 0x0a, 0x0d, 0x4f, 0x53, 0x53, 0x4f, 0x62, 0x6a, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x65, 0x54, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x65, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x42, 0x3c, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x69, 0x6e,
	0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x09, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x70, 0x69, 0x73, 0x74, 0x6f, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_pistonint_grpc_core_file_proto_rawDescOnce sync.Once
	file_com_pistonint_grpc_core_file_proto_rawDescData = file_com_pistonint_grpc_core_file_proto_rawDesc
)

func file_com_pistonint_grpc_core_file_proto_rawDescGZIP() []byte {
	file_com_pistonint_grpc_core_file_proto_rawDescOnce.Do(func() {
		file_com_pistonint_grpc_core_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_pistonint_grpc_core_file_proto_rawDescData)
	})
	return file_com_pistonint_grpc_core_file_proto_rawDescData
}

var file_com_pistonint_grpc_core_file_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_com_pistonint_grpc_core_file_proto_goTypes = []interface{}{
	(*FileEntry)(nil),          // 0: com.pistonint.grpc.core.FileEntry
	(*FileEntryResp)(nil),      // 1: com.pistonint.grpc.core.FileEntryResp
	(*Profile)(nil),            // 2: com.pistonint.grpc.core.Profile
	(*ChunkFile)(nil),          // 3: com.pistonint.grpc.core.ChunkFile
	(*BatchFileInfoReq)(nil),   // 4: com.pistonint.grpc.core.BatchFileInfoReq
	(*ChunkFileReq)(nil),       // 5: com.pistonint.grpc.core.ChunkFileReq
	(*BatchFileInfoResp)(nil),  // 6: com.pistonint.grpc.core.BatchFileInfoResp
	(*OSSObjSummary)(nil),      // 7: com.pistonint.grpc.core.OSSObjSummary
	(*AbortMultUploadReq)(nil), // 8: com.pistonint.grpc.core.AbortMultUploadReq
}
var file_com_pistonint_grpc_core_file_proto_depIdxs = []int32{
	0, // 0: com.pistonint.grpc.core.FileEntryResp.data:type_name -> com.pistonint.grpc.core.FileEntry
	7, // 1: com.pistonint.grpc.core.BatchFileInfoResp.data:type_name -> com.pistonint.grpc.core.OSSObjSummary
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_pistonint_grpc_core_file_proto_init() }
func file_com_pistonint_grpc_core_file_proto_init() {
	if File_com_pistonint_grpc_core_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_pistonint_grpc_core_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileEntry); i {
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
		file_com_pistonint_grpc_core_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileEntryResp); i {
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
		file_com_pistonint_grpc_core_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_com_pistonint_grpc_core_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkFile); i {
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
		file_com_pistonint_grpc_core_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchFileInfoReq); i {
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
		file_com_pistonint_grpc_core_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkFileReq); i {
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
		file_com_pistonint_grpc_core_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchFileInfoResp); i {
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
		file_com_pistonint_grpc_core_file_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSSObjSummary); i {
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
		file_com_pistonint_grpc_core_file_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbortMultUploadReq); i {
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
	file_com_pistonint_grpc_core_file_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_pistonint_grpc_core_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_pistonint_grpc_core_file_proto_goTypes,
		DependencyIndexes: file_com_pistonint_grpc_core_file_proto_depIdxs,
		MessageInfos:      file_com_pistonint_grpc_core_file_proto_msgTypes,
	}.Build()
	File_com_pistonint_grpc_core_file_proto = out.File
	file_com_pistonint_grpc_core_file_proto_rawDesc = nil
	file_com_pistonint_grpc_core_file_proto_goTypes = nil
	file_com_pistonint_grpc_core_file_proto_depIdxs = nil
}
