// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: claude/text/v1/chat.proto

package textv1

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

type UserMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserMeta) Reset() {
	*x = UserMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claude_text_v1_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMeta) ProtoMessage() {}

func (x *UserMeta) ProtoReflect() protoreflect.Message {
	mi := &file_claude_text_v1_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMeta.ProtoReflect.Descriptor instead.
func (*UserMeta) Descriptor() ([]byte, []int) {
	return file_claude_text_v1_chat_proto_rawDescGZIP(), []int{0}
}

func (x *UserMeta) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前支持以下：
	// Human: 表示用户
	// Assistant: 表示对话助手
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// 对话内容，不能为空。
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claude_text_v1_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_claude_text_v1_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_claude_text_v1_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ChatCompletionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 目前只能取下面的值
	// claude-1
	// claude-1-100k
	// claude-instant-1
	// claude-instant-1-100k
	// or sub model
	// claude-1.3
	// claude-1.3-100k
	// claude-1.2
	// claude-1.0
	// claude-instant-1.1
	// claude-instant-1.1-100k
	// claude-instant-1.0
	Model string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	// 停用词
	StopSequences []string `protobuf:"bytes,2,rep,name=stop_sequences,json=stopSequences,proto3" json:"stop_sequences,omitempty"`
	// 否 	是否以流式接口的形式返回数据，默认false。
	Stream bool `protobuf:"varint,3,opt,name=stream,proto3" json:"stream,omitempty"`
	// 最大token
	MaxTokensToSample int32 `protobuf:"varint,4,opt,name=max_tokens_to_sample,json=maxTokensToSample,proto3" json:"max_tokens_to_sample,omitempty"`
	// 对话问题、人物或功能设定 ，不可为空
	Messages []*Message `protobuf:"bytes,5,rep,name=messages,proto3" json:"messages,omitempty"`
	// 对话上下文信息。说明："user_name": "我", "bot_name": "专家"
	Metadata    *UserMeta `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Temperature float32   `protobuf:"fixed32,7,opt,name=temperature,proto3" json:"temperature,omitempty"`
	TopP        float32   `protobuf:"fixed32,8,opt,name=top_p,json=topP,proto3" json:"top_p,omitempty"`
	TopK        int32     `protobuf:"varint,9,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
	Prompt      string    `protobuf:"bytes,10,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *ChatCompletionsRequest) Reset() {
	*x = ChatCompletionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claude_text_v1_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionsRequest) ProtoMessage() {}

func (x *ChatCompletionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_claude_text_v1_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionsRequest.ProtoReflect.Descriptor instead.
func (*ChatCompletionsRequest) Descriptor() ([]byte, []int) {
	return file_claude_text_v1_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatCompletionsRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatCompletionsRequest) GetStopSequences() []string {
	if x != nil {
		return x.StopSequences
	}
	return nil
}

func (x *ChatCompletionsRequest) GetStream() bool {
	if x != nil {
		return x.Stream
	}
	return false
}

func (x *ChatCompletionsRequest) GetMaxTokensToSample() int32 {
	if x != nil {
		return x.MaxTokensToSample
	}
	return 0
}

func (x *ChatCompletionsRequest) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ChatCompletionsRequest) GetMetadata() *UserMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ChatCompletionsRequest) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *ChatCompletionsRequest) GetTopP() float32 {
	if x != nil {
		return x.TopP
	}
	return 0
}

func (x *ChatCompletionsRequest) GetTopK() int32 {
	if x != nil {
		return x.TopK
	}
	return 0
}

func (x *ChatCompletionsRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误详情
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// 错误详情
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claude_text_v1_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_claude_text_v1_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_claude_text_v1_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChatCompletionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文本结果
	Completion string `protobuf:"bytes,1,opt,name=completion,proto3" json:"completion,omitempty"`
	// 模型名称
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	// 结束原因，枚举值
	StopReason string `protobuf:"bytes,3,opt,name=stop_reason,json=stopReason,proto3" json:"stop_reason,omitempty"`
	// 是否截断
	Truncated bool `protobuf:"varint,4,opt,name=truncated,proto3" json:"truncated,omitempty"`
	// 停止词
	Stop string `protobuf:"bytes,5,opt,name=stop,proto3" json:"stop,omitempty"`
	// 日志id
	LogId string `protobuf:"bytes,6,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	// 异常信息
	Exception string `protobuf:"bytes,7,opt,name=exception,proto3" json:"exception,omitempty"`
	// 错误信息
	Error *ErrorResponse `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ChatCompletionsResponse) Reset() {
	*x = ChatCompletionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claude_text_v1_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionsResponse) ProtoMessage() {}

func (x *ChatCompletionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_claude_text_v1_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionsResponse.ProtoReflect.Descriptor instead.
func (*ChatCompletionsResponse) Descriptor() ([]byte, []int) {
	return file_claude_text_v1_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatCompletionsResponse) GetCompletion() string {
	if x != nil {
		return x.Completion
	}
	return ""
}

func (x *ChatCompletionsResponse) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatCompletionsResponse) GetStopReason() string {
	if x != nil {
		return x.StopReason
	}
	return ""
}

func (x *ChatCompletionsResponse) GetTruncated() bool {
	if x != nil {
		return x.Truncated
	}
	return false
}

func (x *ChatCompletionsResponse) GetStop() string {
	if x != nil {
		return x.Stop
	}
	return ""
}

func (x *ChatCompletionsResponse) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *ChatCompletionsResponse) GetException() string {
	if x != nil {
		return x.Exception
	}
	return ""
}

func (x *ChatCompletionsResponse) GetError() *ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_claude_text_v1_chat_proto protoreflect.FileDescriptor

var file_claude_text_v1_chat_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6c, 0x61,
	0x75, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x23, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x37, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xed, 0x02, 0x0a, 0x16, 0x43, 0x68,
	0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74,
	0x6f, 0x70, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2f, 0x0a, 0x14, 0x6d, 0x61, 0x78,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x54, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x6c, 0x61, 0x75, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x12, 0x13, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x6f, 0x70,
	0x4b, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x3d, 0x0a, 0x0d, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8c, 0x02, 0x0a, 0x17, 0x43, 0x68, 0x61,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x6f, 0x70, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x6f,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x15, 0x0a,
	0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6f, 0x67, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0xb8, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x43, 0x68, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41,
	0x49, 0x2d, 0x45, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x78, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x78, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x58,
	0xaa, 0x02, 0x0e, 0x43, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0e, 0x43, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x5c, 0x54, 0x65, 0x78, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1a, 0x43, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x5c, 0x54, 0x65, 0x78, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x10, 0x43, 0x6c, 0x61, 0x75, 0x64, 0x65, 0x3a, 0x3a, 0x54, 0x65, 0x78, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_claude_text_v1_chat_proto_rawDescOnce sync.Once
	file_claude_text_v1_chat_proto_rawDescData = file_claude_text_v1_chat_proto_rawDesc
)

func file_claude_text_v1_chat_proto_rawDescGZIP() []byte {
	file_claude_text_v1_chat_proto_rawDescOnce.Do(func() {
		file_claude_text_v1_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_claude_text_v1_chat_proto_rawDescData)
	})
	return file_claude_text_v1_chat_proto_rawDescData
}

var file_claude_text_v1_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_claude_text_v1_chat_proto_goTypes = []interface{}{
	(*UserMeta)(nil),                // 0: claude.text.v1.UserMeta
	(*Message)(nil),                 // 1: claude.text.v1.Message
	(*ChatCompletionsRequest)(nil),  // 2: claude.text.v1.ChatCompletionsRequest
	(*ErrorResponse)(nil),           // 3: claude.text.v1.ErrorResponse
	(*ChatCompletionsResponse)(nil), // 4: claude.text.v1.ChatCompletionsResponse
}
var file_claude_text_v1_chat_proto_depIdxs = []int32{
	1, // 0: claude.text.v1.ChatCompletionsRequest.messages:type_name -> claude.text.v1.Message
	0, // 1: claude.text.v1.ChatCompletionsRequest.metadata:type_name -> claude.text.v1.UserMeta
	3, // 2: claude.text.v1.ChatCompletionsResponse.error:type_name -> claude.text.v1.ErrorResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_claude_text_v1_chat_proto_init() }
func file_claude_text_v1_chat_proto_init() {
	if File_claude_text_v1_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_claude_text_v1_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMeta); i {
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
		file_claude_text_v1_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_claude_text_v1_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionsRequest); i {
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
		file_claude_text_v1_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
		file_claude_text_v1_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionsResponse); i {
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
			RawDescriptor: file_claude_text_v1_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_claude_text_v1_chat_proto_goTypes,
		DependencyIndexes: file_claude_text_v1_chat_proto_depIdxs,
		MessageInfos:      file_claude_text_v1_chat_proto_msgTypes,
	}.Build()
	File_claude_text_v1_chat_proto = out.File
	file_claude_text_v1_chat_proto_rawDesc = nil
	file_claude_text_v1_chat_proto_goTypes = nil
	file_claude_text_v1_chat_proto_depIdxs = nil
}
