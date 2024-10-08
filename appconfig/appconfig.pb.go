// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.1
// source: appconfig/appconfig.proto

// Define the package name (optional but recommended)

package appconfig

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

// AppConfig used to configure the garage pi application
type AppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The telegram bot token
	TelegramBotToken string `protobuf:"bytes,1,opt,name=telegram_bot_token,json=telegramBotToken,proto3" json:"telegram_bot_token,omitempty"`
	// group_channel_id if non zero the events and errors will
	// additionally be sent to this channel
	GroupChannelId int64 `protobuf:"varint,2,opt,name=group_channel_id,json=groupChannelId,proto3" json:"group_channel_id,omitempty"`
	// allow_list_ids as a repeated list of integers
	AllowListIds []int64 `protobuf:"varint,3,rep,packed,name=allow_list_ids,json=allowListIds,proto3" json:"allow_list_ids,omitempty"`
	// door_relay_number the relay on the PiFACE used to open the door
	DoorRelayNumber int32 `protobuf:"varint,4,opt,name=door_relay_number,json=doorRelayNumber,proto3" json:"door_relay_number,omitempty"`
	// door_input_number the input on the PiFACE used to detect the door state
	DoorInputNumber int32 `protobuf:"varint,5,opt,name=door_input_number,json=doorInputNumber,proto3" json:"door_input_number,omitempty"`
	// auth token for webhook
	WebhookAuthToken string `protobuf:"bytes,6,opt,name=webhook_auth_token,json=webhookAuthToken,proto3" json:"webhook_auth_token,omitempty"`
}

func (x *AppConfig) Reset() {
	*x = AppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appconfig_appconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfig) ProtoMessage() {}

func (x *AppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_appconfig_appconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfig.ProtoReflect.Descriptor instead.
func (*AppConfig) Descriptor() ([]byte, []int) {
	return file_appconfig_appconfig_proto_rawDescGZIP(), []int{0}
}

func (x *AppConfig) GetTelegramBotToken() string {
	if x != nil {
		return x.TelegramBotToken
	}
	return ""
}

func (x *AppConfig) GetGroupChannelId() int64 {
	if x != nil {
		return x.GroupChannelId
	}
	return 0
}

func (x *AppConfig) GetAllowListIds() []int64 {
	if x != nil {
		return x.AllowListIds
	}
	return nil
}

func (x *AppConfig) GetDoorRelayNumber() int32 {
	if x != nil {
		return x.DoorRelayNumber
	}
	return 0
}

func (x *AppConfig) GetDoorInputNumber() int32 {
	if x != nil {
		return x.DoorInputNumber
	}
	return 0
}

func (x *AppConfig) GetWebhookAuthToken() string {
	if x != nil {
		return x.WebhookAuthToken
	}
	return ""
}

var File_appconfig_appconfig_proto protoreflect.FileDescriptor

var file_appconfig_appconfig_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x70,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x8f, 0x02, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d,
	0x5f, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x6f, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x64, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x6f, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64,
	0x6f, 0x6f, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x11, 0x64, 0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x6f, 0x6f, 0x72, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6a, 0x2d, 0x72, 0x69, 0x63, 0x68, 0x61,
	0x72, 0x64, 0x73, 0x2f, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appconfig_appconfig_proto_rawDescOnce sync.Once
	file_appconfig_appconfig_proto_rawDescData = file_appconfig_appconfig_proto_rawDesc
)

func file_appconfig_appconfig_proto_rawDescGZIP() []byte {
	file_appconfig_appconfig_proto_rawDescOnce.Do(func() {
		file_appconfig_appconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_appconfig_appconfig_proto_rawDescData)
	})
	return file_appconfig_appconfig_proto_rawDescData
}

var file_appconfig_appconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_appconfig_appconfig_proto_goTypes = []interface{}{
	(*AppConfig)(nil), // 0: appconfig.AppConfig
}
var file_appconfig_appconfig_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_appconfig_appconfig_proto_init() }
func file_appconfig_appconfig_proto_init() {
	if File_appconfig_appconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appconfig_appconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfig); i {
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
			RawDescriptor: file_appconfig_appconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_appconfig_appconfig_proto_goTypes,
		DependencyIndexes: file_appconfig_appconfig_proto_depIdxs,
		MessageInfos:      file_appconfig_appconfig_proto_msgTypes,
	}.Build()
	File_appconfig_appconfig_proto = out.File
	file_appconfig_appconfig_proto_rawDesc = nil
	file_appconfig_appconfig_proto_goTypes = nil
	file_appconfig_appconfig_proto_depIdxs = nil
}
