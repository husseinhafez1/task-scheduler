// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/task.proto

package task

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Payload       string                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobRequest) Reset() {
	*x = JobRequest{}
	mi := &file_proto_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequest) ProtoMessage() {}

func (x *JobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequest.ProtoReflect.Descriptor instead.
func (*JobRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{0}
}

func (x *JobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type JobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobResponse) Reset() {
	*x = JobResponse{}
	mi := &file_proto_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResponse) ProtoMessage() {}

func (x *JobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResponse.ProtoReflect.Descriptor instead.
func (*JobResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{1}
}

func (x *JobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type JobStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobStatusRequest) Reset() {
	*x = JobStatusRequest{}
	mi := &file_proto_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusRequest) ProtoMessage() {}

func (x *JobStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusRequest.ProtoReflect.Descriptor instead.
func (*JobStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{2}
}

func (x *JobStatusRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type JobStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobStatusResponse) Reset() {
	*x = JobStatusResponse{}
	mi := &file_proto_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusResponse) ProtoMessage() {}

func (x *JobStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusResponse.ProtoReflect.Descriptor instead.
func (*JobStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{3}
}

func (x *JobStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type JobLogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Logs          []string               `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobLogsResponse) Reset() {
	*x = JobLogsResponse{}
	mi := &file_proto_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLogsResponse) ProtoMessage() {}

func (x *JobLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLogsResponse.ProtoReflect.Descriptor instead.
func (*JobLogsResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{4}
}

func (x *JobLogsResponse) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

var File_proto_task_proto protoreflect.FileDescriptor

const file_proto_task_proto_rawDesc = "" +
	"\n" +
	"\x10proto/task.proto\x12\x04task\"=\n" +
	"\n" +
	"JobRequest\x12\x15\n" +
	"\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x18\n" +
	"\apayload\x18\x02 \x01(\tR\apayload\"'\n" +
	"\vJobResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\")\n" +
	"\x10JobStatusRequest\x12\x15\n" +
	"\x06job_id\x18\x01 \x01(\tR\x05jobId\"+\n" +
	"\x11JobStatusResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"%\n" +
	"\x0fJobLogsResponse\x12\x12\n" +
	"\x04logs\x18\x01 \x03(\tR\x04logs2\xbd\x01\n" +
	"\vTaskService\x120\n" +
	"\tSubmitJob\x12\x10.task.JobRequest\x1a\x11.task.JobResponse\x12?\n" +
	"\fGetJobStatus\x12\x16.task.JobStatusRequest\x1a\x17.task.JobStatusResponse\x12;\n" +
	"\n" +
	"GetJobLogs\x12\x16.task.JobStatusRequest\x1a\x15.task.JobLogsResponseB\x11Z\x0ftask/proto;taskb\x06proto3"

var (
	file_proto_task_proto_rawDescOnce sync.Once
	file_proto_task_proto_rawDescData []byte
)

func file_proto_task_proto_rawDescGZIP() []byte {
	file_proto_task_proto_rawDescOnce.Do(func() {
		file_proto_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_task_proto_rawDesc), len(file_proto_task_proto_rawDesc)))
	})
	return file_proto_task_proto_rawDescData
}

var file_proto_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_task_proto_goTypes = []any{
	(*JobRequest)(nil),        // 0: task.JobRequest
	(*JobResponse)(nil),       // 1: task.JobResponse
	(*JobStatusRequest)(nil),  // 2: task.JobStatusRequest
	(*JobStatusResponse)(nil), // 3: task.JobStatusResponse
	(*JobLogsResponse)(nil),   // 4: task.JobLogsResponse
}
var file_proto_task_proto_depIdxs = []int32{
	0, // 0: task.TaskService.SubmitJob:input_type -> task.JobRequest
	2, // 1: task.TaskService.GetJobStatus:input_type -> task.JobStatusRequest
	2, // 2: task.TaskService.GetJobLogs:input_type -> task.JobStatusRequest
	1, // 3: task.TaskService.SubmitJob:output_type -> task.JobResponse
	3, // 4: task.TaskService.GetJobStatus:output_type -> task.JobStatusResponse
	4, // 5: task.TaskService.GetJobLogs:output_type -> task.JobLogsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_task_proto_init() }
func file_proto_task_proto_init() {
	if File_proto_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_task_proto_rawDesc), len(file_proto_task_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_task_proto_goTypes,
		DependencyIndexes: file_proto_task_proto_depIdxs,
		MessageInfos:      file_proto_task_proto_msgTypes,
	}.Build()
	File_proto_task_proto = out.File
	file_proto_task_proto_goTypes = nil
	file_proto_task_proto_depIdxs = nil
}
