// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.7.1
// source: api/calendar.proto

package calendar

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       int32                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Datetimefrom *timestamp.Timestamp `protobuf:"bytes,3,opt,name=datetimefrom,proto3" json:"datetimefrom,omitempty"`
	Datetimeto   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=datetimeto,proto3" json:"datetimeto,omitempty"`
	Description  string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Event) GetDatetimefrom() *timestamp.Timestamp {
	if x != nil {
		return x.Datetimefrom
	}
	return nil
}

func (x *Event) GetDatetimeto() *timestamp.Timestamp {
	if x != nil {
		return x.Datetimeto
	}
	return nil
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *AddEventRequest) Reset() {
	*x = AddEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRequest) ProtoMessage() {}

func (x *AddEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRequest.ProtoReflect.Descriptor instead.
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *AddEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *SuccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DropEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	EventId int32 `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *DropEventRequest) Reset() {
	*x = DropEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropEventRequest) ProtoMessage() {}

func (x *DropEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropEventRequest.ProtoReflect.Descriptor instead.
func (*DropEventRequest) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{3}
}

func (x *DropEventRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DropEventRequest) GetEventId() int32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

type EditEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewEvent *Event `protobuf:"bytes,1,opt,name=new_event,json=newEvent,proto3" json:"new_event,omitempty"`
}

func (x *EditEventRequest) Reset() {
	*x = EditEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditEventRequest) ProtoMessage() {}

func (x *EditEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditEventRequest.ProtoReflect.Descriptor instead.
func (*EditEventRequest) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{4}
}

func (x *EditEventRequest) GetNewEvent() *Event {
	if x != nil {
		return x.NewEvent
	}
	return nil
}

type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	EventId int32 `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{5}
}

func (x *GetEventRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetEventRequest) GetEventId() int32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Event   *Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{6}
}

func (x *EventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type AllEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AllEventsRequest) Reset() {
	*x = AllEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllEventsRequest) ProtoMessage() {}

func (x *AllEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllEventsRequest.ProtoReflect.Descriptor instead.
func (*AllEventsRequest) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{7}
}

func (x *AllEventsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AllEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *AllEventsResponse) Reset() {
	*x = AllEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calendar_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllEventsResponse) ProtoMessage() {}

func (x *AllEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calendar_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllEventsResponse.ProtoReflect.Descriptor instead.
func (*AllEventsResponse) Descriptor() ([]byte, []int) {
	return file_api_calendar_proto_rawDescGZIP(), []int{8}
}

func (x *AllEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_api_calendar_proto protoreflect.FileDescriptor

var file_api_calendar_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xce, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x38, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x46, 0x0a, 0x10, 0x44, 0x72, 0x6f, 0x70, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x40, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x10, 0x41, 0x6c,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xe0, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x44, 0x72, 0x6f, 0x70, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x44, 0x72, 0x6f,
	0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x45, 0x64,
	0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x2e,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x3b,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_calendar_proto_rawDescOnce sync.Once
	file_api_calendar_proto_rawDescData = file_api_calendar_proto_rawDesc
)

func file_api_calendar_proto_rawDescGZIP() []byte {
	file_api_calendar_proto_rawDescOnce.Do(func() {
		file_api_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_calendar_proto_rawDescData)
	})
	return file_api_calendar_proto_rawDescData
}

var file_api_calendar_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_calendar_proto_goTypes = []interface{}{
	(*Event)(nil),               // 0: calendar.Event
	(*AddEventRequest)(nil),     // 1: calendar.AddEventRequest
	(*SuccessResponse)(nil),     // 2: calendar.SuccessResponse
	(*DropEventRequest)(nil),    // 3: calendar.DropEventRequest
	(*EditEventRequest)(nil),    // 4: calendar.EditEventRequest
	(*GetEventRequest)(nil),     // 5: calendar.GetEventRequest
	(*EventResponse)(nil),       // 6: calendar.EventResponse
	(*AllEventsRequest)(nil),    // 7: calendar.AllEventsRequest
	(*AllEventsResponse)(nil),   // 8: calendar.AllEventsResponse
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_api_calendar_proto_depIdxs = []int32{
	9,  // 0: calendar.Event.datetimefrom:type_name -> google.protobuf.Timestamp
	9,  // 1: calendar.Event.datetimeto:type_name -> google.protobuf.Timestamp
	0,  // 2: calendar.AddEventRequest.event:type_name -> calendar.Event
	0,  // 3: calendar.EditEventRequest.new_event:type_name -> calendar.Event
	0,  // 4: calendar.EventResponse.event:type_name -> calendar.Event
	0,  // 5: calendar.AllEventsResponse.events:type_name -> calendar.Event
	1,  // 6: calendar.Calendar.AddEvent:input_type -> calendar.AddEventRequest
	3,  // 7: calendar.Calendar.DropEvent:input_type -> calendar.DropEventRequest
	4,  // 8: calendar.Calendar.EditEvent:input_type -> calendar.EditEventRequest
	5,  // 9: calendar.Calendar.GetEvent:input_type -> calendar.GetEventRequest
	7,  // 10: calendar.Calendar.AllEvents:input_type -> calendar.AllEventsRequest
	6,  // 11: calendar.Calendar.AddEvent:output_type -> calendar.EventResponse
	2,  // 12: calendar.Calendar.DropEvent:output_type -> calendar.SuccessResponse
	6,  // 13: calendar.Calendar.EditEvent:output_type -> calendar.EventResponse
	6,  // 14: calendar.Calendar.GetEvent:output_type -> calendar.EventResponse
	8,  // 15: calendar.Calendar.AllEvents:output_type -> calendar.AllEventsResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_calendar_proto_init() }
func file_api_calendar_proto_init() {
	if File_api_calendar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_api_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventRequest); i {
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
		file_api_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_api_calendar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropEventRequest); i {
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
		file_api_calendar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditEventRequest); i {
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
		file_api_calendar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_api_calendar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_api_calendar_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllEventsRequest); i {
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
		file_api_calendar_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllEventsResponse); i {
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
			RawDescriptor: file_api_calendar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_calendar_proto_goTypes,
		DependencyIndexes: file_api_calendar_proto_depIdxs,
		MessageInfos:      file_api_calendar_proto_msgTypes,
	}.Build()
	File_api_calendar_proto = out.File
	file_api_calendar_proto_rawDesc = nil
	file_api_calendar_proto_goTypes = nil
	file_api_calendar_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	DropEvent(ctx context.Context, in *DropEventRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	EditEvent(ctx context.Context, in *EditEventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	AllEvents(ctx context.Context, in *AllEventsRequest, opts ...grpc.CallOption) (*AllEventsResponse, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) DropEvent(ctx context.Context, in *DropEventRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/DropEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) EditEvent(ctx context.Context, in *EditEventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/EditEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) AllEvents(ctx context.Context, in *AllEventsRequest, opts ...grpc.CallOption) (*AllEventsResponse, error) {
	out := new(AllEventsResponse)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/AllEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	AddEvent(context.Context, *AddEventRequest) (*EventResponse, error)
	DropEvent(context.Context, *DropEventRequest) (*SuccessResponse, error)
	EditEvent(context.Context, *EditEventRequest) (*EventResponse, error)
	GetEvent(context.Context, *GetEventRequest) (*EventResponse, error)
	AllEvents(context.Context, *AllEventsRequest) (*AllEventsResponse, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) AddEvent(context.Context, *AddEventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (*UnimplementedCalendarServer) DropEvent(context.Context, *DropEventRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropEvent not implemented")
}
func (*UnimplementedCalendarServer) EditEvent(context.Context, *EditEventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEvent not implemented")
}
func (*UnimplementedCalendarServer) GetEvent(context.Context, *GetEventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServer) AllEvents(context.Context, *AllEventsRequest) (*AllEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllEvents not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_DropEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).DropEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/DropEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).DropEvent(ctx, req.(*DropEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_EditEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).EditEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/EditEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).EditEvent(ctx, req.(*EditEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_AllEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).AllEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/AllEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).AllEvents(ctx, req.(*AllEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _Calendar_AddEvent_Handler,
		},
		{
			MethodName: "DropEvent",
			Handler:    _Calendar_DropEvent_Handler,
		},
		{
			MethodName: "EditEvent",
			Handler:    _Calendar_EditEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Calendar_GetEvent_Handler,
		},
		{
			MethodName: "AllEvents",
			Handler:    _Calendar_AllEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/calendar.proto",
}
