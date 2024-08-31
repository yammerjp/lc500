// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: worker.proto

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

type InitVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitVMRequest) Reset() {
	*x = InitVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitVMRequest) ProtoMessage() {}

func (x *InitVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitVMRequest.ProtoReflect.Descriptor instead.
func (*InitVMRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

type InitVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vmid string `protobuf:"bytes,1,opt,name=vmid,proto3" json:"vmid,omitempty"`
}

func (x *InitVMResponse) Reset() {
	*x = InitVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitVMResponse) ProtoMessage() {}

func (x *InitVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitVMResponse.ProtoReflect.Descriptor instead.
func (*InitVMResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *InitVMResponse) GetVmid() string {
	if x != nil {
		return x.Vmid
	}
	return ""
}

type DisposeVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vmid string `protobuf:"bytes,1,opt,name=vmid,proto3" json:"vmid,omitempty"`
}

func (x *DisposeVMRequest) Reset() {
	*x = DisposeVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisposeVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisposeVMRequest) ProtoMessage() {}

func (x *DisposeVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisposeVMRequest.ProtoReflect.Descriptor instead.
func (*DisposeVMRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

func (x *DisposeVMRequest) GetVmid() string {
	if x != nil {
		return x.Vmid
	}
	return ""
}

type DisposeVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisposeVMResponse) Reset() {
	*x = DisposeVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisposeVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisposeVMResponse) ProtoMessage() {}

func (x *DisposeVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisposeVMResponse.ProtoReflect.Descriptor instead.
func (*DisposeVMResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{3}
}

type CompileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vmid   string `protobuf:"bytes,1,opt,name=vmid,proto3" json:"vmid,omitempty"`
	Script string `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *CompileRequest) Reset() {
	*x = CompileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileRequest) ProtoMessage() {}

func (x *CompileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileRequest.ProtoReflect.Descriptor instead.
func (*CompileRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{4}
}

func (x *CompileRequest) GetVmid() string {
	if x != nil {
		return x.Vmid
	}
	return ""
}

func (x *CompileRequest) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

type CompileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompileResponse) Reset() {
	*x = CompileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileResponse) ProtoMessage() {}

func (x *CompileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileResponse.ProtoReflect.Descriptor instead.
func (*CompileResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{5}
}

type SetContextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vmid               string            `protobuf:"bytes,1,opt,name=vmid,proto3" json:"vmid,omitempty"`
	HttpRequestMethod  string            `protobuf:"bytes,2,opt,name=httpRequestMethod,proto3" json:"httpRequestMethod,omitempty"`
	HttpRequestUrl     string            `protobuf:"bytes,3,opt,name=httpRequestUrl,proto3" json:"httpRequestUrl,omitempty"`
	HttpRequestBody    string            `protobuf:"bytes,4,opt,name=httpRequestBody,proto3" json:"httpRequestBody,omitempty"`
	HttpRequestHeaders map[string]string `protobuf:"bytes,5,rep,name=httpRequestHeaders,proto3" json:"httpRequestHeaders,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AdditionalContext  string            `protobuf:"bytes,6,opt,name=additionalContext,proto3" json:"additionalContext,omitempty"`
}

func (x *SetContextRequest) Reset() {
	*x = SetContextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetContextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetContextRequest) ProtoMessage() {}

func (x *SetContextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetContextRequest.ProtoReflect.Descriptor instead.
func (*SetContextRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{6}
}

func (x *SetContextRequest) GetVmid() string {
	if x != nil {
		return x.Vmid
	}
	return ""
}

func (x *SetContextRequest) GetHttpRequestMethod() string {
	if x != nil {
		return x.HttpRequestMethod
	}
	return ""
}

func (x *SetContextRequest) GetHttpRequestUrl() string {
	if x != nil {
		return x.HttpRequestUrl
	}
	return ""
}

func (x *SetContextRequest) GetHttpRequestBody() string {
	if x != nil {
		return x.HttpRequestBody
	}
	return ""
}

func (x *SetContextRequest) GetHttpRequestHeaders() map[string]string {
	if x != nil {
		return x.HttpRequestHeaders
	}
	return nil
}

func (x *SetContextRequest) GetAdditionalContext() string {
	if x != nil {
		return x.AdditionalContext
	}
	return ""
}

type SetContextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetContextResponse) Reset() {
	*x = SetContextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetContextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetContextResponse) ProtoMessage() {}

func (x *SetContextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetContextResponse.ProtoReflect.Descriptor instead.
func (*SetContextResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{7}
}

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vmid    string `protobuf:"bytes,1,opt,name=vmid,proto3" json:"vmid,omitempty"`
	Dispose bool   `protobuf:"varint,2,opt,name=dispose,proto3" json:"dispose,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{8}
}

func (x *RunRequest) GetVmid() string {
	if x != nil {
		return x.Vmid
	}
	return ""
}

func (x *RunRequest) GetDispose() bool {
	if x != nil {
		return x.Dispose
	}
	return false
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpResponseStatusCode int32             `protobuf:"varint,1,opt,name=httpResponseStatusCode,proto3" json:"httpResponseStatusCode,omitempty"`
	HttpResponseHeaders    map[string]string `protobuf:"bytes,2,rep,name=httpResponseHeaders,proto3" json:"httpResponseHeaders,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HttpResponseBody       string            `protobuf:"bytes,3,opt,name=httpResponseBody,proto3" json:"httpResponseBody,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{9}
}

func (x *RunResponse) GetHttpResponseStatusCode() int32 {
	if x != nil {
		return x.HttpResponseStatusCode
	}
	return 0
}

func (x *RunResponse) GetHttpResponseHeaders() map[string]string {
	if x != nil {
		return x.HttpResponseHeaders
	}
	return nil
}

func (x *RunResponse) GetHttpResponseBody() string {
	if x != nil {
		return x.HttpResponseBody
	}
	return ""
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x22, 0x0f, 0x0a, 0x0d, 0x49, 0x6e, 0x69, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6d, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x44, 0x69, 0x73,
	0x70, 0x6f, 0x73, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6d, 0x69,
	0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x56, 0x4d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6d, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6d, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xfb, 0x02, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6d, 0x69,
	0x64, 0x12, 0x2c, 0x0a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x68, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x5d, 0x0a, 0x12, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x68, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x45,
	0x0a, 0x17, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x0a, 0x52,
	0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6d, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6d, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x64, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x95, 0x02, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x68, 0x74, 0x74, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x5a, 0x0a, 0x13, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x68,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x46, 0x0a, 0x18, 0x48, 0x74, 0x74, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0x96, 0x02, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x49, 0x6e,
	0x69, 0x74, 0x56, 0x4d, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x56, 0x4d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x09, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x56, 0x4d, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x56, 0x4d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x15, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6d, 0x6d, 0x65, 0x72, 0x6a, 0x70, 0x2f,
	0x6c, 0x63, 0x35, 0x30, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_worker_proto_goTypes = []any{
	(*InitVMRequest)(nil),      // 0: v1.InitVMRequest
	(*InitVMResponse)(nil),     // 1: v1.InitVMResponse
	(*DisposeVMRequest)(nil),   // 2: v1.DisposeVMRequest
	(*DisposeVMResponse)(nil),  // 3: v1.DisposeVMResponse
	(*CompileRequest)(nil),     // 4: v1.CompileRequest
	(*CompileResponse)(nil),    // 5: v1.CompileResponse
	(*SetContextRequest)(nil),  // 6: v1.SetContextRequest
	(*SetContextResponse)(nil), // 7: v1.SetContextResponse
	(*RunRequest)(nil),         // 8: v1.RunRequest
	(*RunResponse)(nil),        // 9: v1.RunResponse
	nil,                        // 10: v1.SetContextRequest.HttpRequestHeadersEntry
	nil,                        // 11: v1.RunResponse.HttpResponseHeadersEntry
}
var file_worker_proto_depIdxs = []int32{
	10, // 0: v1.SetContextRequest.httpRequestHeaders:type_name -> v1.SetContextRequest.HttpRequestHeadersEntry
	11, // 1: v1.RunResponse.httpResponseHeaders:type_name -> v1.RunResponse.HttpResponseHeadersEntry
	0,  // 2: v1.Worker.InitVM:input_type -> v1.InitVMRequest
	2,  // 3: v1.Worker.DisposeVM:input_type -> v1.DisposeVMRequest
	4,  // 4: v1.Worker.Compile:input_type -> v1.CompileRequest
	6,  // 5: v1.Worker.SetContext:input_type -> v1.SetContextRequest
	8,  // 6: v1.Worker.Run:input_type -> v1.RunRequest
	1,  // 7: v1.Worker.InitVM:output_type -> v1.InitVMResponse
	3,  // 8: v1.Worker.DisposeVM:output_type -> v1.DisposeVMResponse
	5,  // 9: v1.Worker.Compile:output_type -> v1.CompileResponse
	7,  // 10: v1.Worker.SetContext:output_type -> v1.SetContextResponse
	9,  // 11: v1.Worker.Run:output_type -> v1.RunResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InitVMRequest); i {
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
		file_worker_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InitVMResponse); i {
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
		file_worker_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DisposeVMRequest); i {
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
		file_worker_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DisposeVMResponse); i {
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
		file_worker_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CompileRequest); i {
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
		file_worker_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CompileResponse); i {
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
		file_worker_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SetContextRequest); i {
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
		file_worker_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SetContextResponse); i {
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
		file_worker_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RunRequest); i {
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
		file_worker_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RunResponse); i {
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
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}
