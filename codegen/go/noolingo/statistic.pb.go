// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: noolingo/statistic.proto

package noolingo

import (
	common "github.com/MelnikovNA/noolingoproto/codegen/go/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatisticObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckId       string               `protobuf:"bytes,1,opt,name=deck_id,json=deckId,proto3" json:"deck_id,omitempty"`
	WordsCount   string               `protobuf:"bytes,2,opt,name=words_count,json=wordsCount,proto3" json:"words_count,omitempty"`
	WordsLearned string               `protobuf:"bytes,3,opt,name=words_learned,json=wordsLearned,proto3" json:"words_learned,omitempty"`
	Duration     *durationpb.Duration `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *StatisticObject) Reset() {
	*x = StatisticObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noolingo_statistic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticObject) ProtoMessage() {}

func (x *StatisticObject) ProtoReflect() protoreflect.Message {
	mi := &file_noolingo_statistic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticObject.ProtoReflect.Descriptor instead.
func (*StatisticObject) Descriptor() ([]byte, []int) {
	return file_noolingo_statistic_proto_rawDescGZIP(), []int{0}
}

func (x *StatisticObject) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

func (x *StatisticObject) GetWordsCount() string {
	if x != nil {
		return x.WordsCount
	}
	return ""
}

func (x *StatisticObject) GetWordsLearned() string {
	if x != nil {
		return x.WordsLearned
	}
	return ""
}

func (x *StatisticObject) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type StatisticUpdataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statistic *StatisticObject `protobuf:"bytes,1,opt,name=statistic,proto3" json:"statistic,omitempty"`
}

func (x *StatisticUpdataRequest) Reset() {
	*x = StatisticUpdataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noolingo_statistic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticUpdataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticUpdataRequest) ProtoMessage() {}

func (x *StatisticUpdataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_noolingo_statistic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticUpdataRequest.ProtoReflect.Descriptor instead.
func (*StatisticUpdataRequest) Descriptor() ([]byte, []int) {
	return file_noolingo_statistic_proto_rawDescGZIP(), []int{1}
}

func (x *StatisticUpdataRequest) GetStatistic() *StatisticObject {
	if x != nil {
		return x.Statistic
	}
	return nil
}

type StatisticbyDeckIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckId string `protobuf:"bytes,1,opt,name=deck_id,json=deckId,proto3" json:"deck_id,omitempty"`
}

func (x *StatisticbyDeckIDRequest) Reset() {
	*x = StatisticbyDeckIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noolingo_statistic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticbyDeckIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticbyDeckIDRequest) ProtoMessage() {}

func (x *StatisticbyDeckIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_noolingo_statistic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticbyDeckIDRequest.ProtoReflect.Descriptor instead.
func (*StatisticbyDeckIDRequest) Descriptor() ([]byte, []int) {
	return file_noolingo_statistic_proto_rawDescGZIP(), []int{2}
}

func (x *StatisticbyDeckIDRequest) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

type StatisticResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statistic *StatisticObject `protobuf:"bytes,1,opt,name=statistic,proto3" json:"statistic,omitempty"`
	Response  *common.Error    `protobuf:"bytes,2,opt,name=response,proto3,oneof" json:"response,omitempty"`
}

func (x *StatisticResponse) Reset() {
	*x = StatisticResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noolingo_statistic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticResponse) ProtoMessage() {}

func (x *StatisticResponse) ProtoReflect() protoreflect.Message {
	mi := &file_noolingo_statistic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticResponse.ProtoReflect.Descriptor instead.
func (*StatisticResponse) Descriptor() ([]byte, []int) {
	return file_noolingo_statistic_proto_rawDescGZIP(), []int{3}
}

func (x *StatisticResponse) GetStatistic() *StatisticObject {
	if x != nil {
		return x.Statistic
	}
	return nil
}

func (x *StatisticResponse) GetResponse() *common.Error {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_noolingo_statistic_proto protoreflect.FileDescriptor

var file_noolingo_statistic_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x6f, 0x6f, 0x6c,
	0x69, 0x6e, 0x67, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51,
	0x0a, 0x16, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f,
	0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x22, 0x33, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x62, 0x79,
	0x44, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6e, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x2e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xb7, 0x02, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x12, 0x5c, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6e, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a,
	0x01, 0x2a, 0x1a, 0x0a, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x58,
	0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x6f, 0x6c, 0x69, 0x6e,
	0x67, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x72, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x62, 0x79, 0x44, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x22, 0x2e,
	0x6e, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x62, 0x79, 0x44, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x2f, 0x7b, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6c, 0x6e, 0x69,
	0x6b, 0x6f, 0x76, 0x4e, 0x41, 0x2f, 0x6e, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6e,
	0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_noolingo_statistic_proto_rawDescOnce sync.Once
	file_noolingo_statistic_proto_rawDescData = file_noolingo_statistic_proto_rawDesc
)

func file_noolingo_statistic_proto_rawDescGZIP() []byte {
	file_noolingo_statistic_proto_rawDescOnce.Do(func() {
		file_noolingo_statistic_proto_rawDescData = protoimpl.X.CompressGZIP(file_noolingo_statistic_proto_rawDescData)
	})
	return file_noolingo_statistic_proto_rawDescData
}

var file_noolingo_statistic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_noolingo_statistic_proto_goTypes = []interface{}{
	(*StatisticObject)(nil),          // 0: noolingo.StatisticObject
	(*StatisticUpdataRequest)(nil),   // 1: noolingo.StatisticUpdataRequest
	(*StatisticbyDeckIDRequest)(nil), // 2: noolingo.StatisticbyDeckIDRequest
	(*StatisticResponse)(nil),        // 3: noolingo.StatisticResponse
	(*durationpb.Duration)(nil),      // 4: google.protobuf.Duration
	(*common.Error)(nil),             // 5: common.Error
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
	(*common.Response)(nil),          // 7: common.Response
}
var file_noolingo_statistic_proto_depIdxs = []int32{
	4, // 0: noolingo.StatisticObject.duration:type_name -> google.protobuf.Duration
	0, // 1: noolingo.StatisticUpdataRequest.statistic:type_name -> noolingo.StatisticObject
	0, // 2: noolingo.StatisticResponse.statistic:type_name -> noolingo.StatisticObject
	5, // 3: noolingo.StatisticResponse.response:type_name -> common.Error
	1, // 4: noolingo.Statistic.StatisticUpdate:input_type -> noolingo.StatisticUpdataRequest
	6, // 5: noolingo.Statistic.StatisticList:input_type -> google.protobuf.Empty
	2, // 6: noolingo.Statistic.StatisticbyDeckID:input_type -> noolingo.StatisticbyDeckIDRequest
	7, // 7: noolingo.Statistic.StatisticUpdate:output_type -> common.Response
	3, // 8: noolingo.Statistic.StatisticList:output_type -> noolingo.StatisticResponse
	3, // 9: noolingo.Statistic.StatisticbyDeckID:output_type -> noolingo.StatisticResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_noolingo_statistic_proto_init() }
func file_noolingo_statistic_proto_init() {
	if File_noolingo_statistic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_noolingo_statistic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticObject); i {
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
		file_noolingo_statistic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticUpdataRequest); i {
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
		file_noolingo_statistic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticbyDeckIDRequest); i {
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
		file_noolingo_statistic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticResponse); i {
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
	file_noolingo_statistic_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_noolingo_statistic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_noolingo_statistic_proto_goTypes,
		DependencyIndexes: file_noolingo_statistic_proto_depIdxs,
		MessageInfos:      file_noolingo_statistic_proto_msgTypes,
	}.Build()
	File_noolingo_statistic_proto = out.File
	file_noolingo_statistic_proto_rawDesc = nil
	file_noolingo_statistic_proto_goTypes = nil
	file_noolingo_statistic_proto_depIdxs = nil
}
