// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/user/v1/events.proto

package userv1

import (
	v1 "github.com/hbttundar/scg-notification-contracts/gen/go/scg/shared/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// UserInvitedNotification is sent when a user is invited to the platform
type UserInvitedNotification struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Base                *v1.NotificationEvent  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	InviterUuid         string                 `protobuf:"bytes,2,opt,name=inviter_uuid,json=inviterUuid,proto3" json:"inviter_uuid,omitempty"`
	InviterName         string                 `protobuf:"bytes,3,opt,name=inviter_name,json=inviterName,proto3" json:"inviter_name,omitempty"`
	InviteeEmail        string                 `protobuf:"bytes,4,opt,name=invitee_email,json=inviteeEmail,proto3" json:"invitee_email,omitempty"`
	OrganizationUuid    string                 `protobuf:"bytes,5,opt,name=organization_uuid,json=organizationUuid,proto3" json:"organization_uuid,omitempty"`
	OrganizationName    string                 `protobuf:"bytes,6,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	Role                string                 `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`
	InvitationExpiresAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=invitation_expires_at,json=invitationExpiresAt,proto3" json:"invitation_expires_at,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *UserInvitedNotification) Reset() {
	*x = UserInvitedNotification{}
	mi := &file_proto_scg_user_v1_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInvitedNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInvitedNotification) ProtoMessage() {}

func (x *UserInvitedNotification) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_user_v1_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInvitedNotification.ProtoReflect.Descriptor instead.
func (*UserInvitedNotification) Descriptor() ([]byte, []int) {
	return file_proto_scg_user_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *UserInvitedNotification) GetBase() *v1.NotificationEvent {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserInvitedNotification) GetInviterUuid() string {
	if x != nil {
		return x.InviterUuid
	}
	return ""
}

func (x *UserInvitedNotification) GetInviterName() string {
	if x != nil {
		return x.InviterName
	}
	return ""
}

func (x *UserInvitedNotification) GetInviteeEmail() string {
	if x != nil {
		return x.InviteeEmail
	}
	return ""
}

func (x *UserInvitedNotification) GetOrganizationUuid() string {
	if x != nil {
		return x.OrganizationUuid
	}
	return ""
}

func (x *UserInvitedNotification) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *UserInvitedNotification) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserInvitedNotification) GetInvitationExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.InvitationExpiresAt
	}
	return nil
}

// UserRegisteredNotification is sent when a user completes registration
type UserRegisteredNotification struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Base             *v1.NotificationEvent  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	UserUuid         string                 `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	UserName         string                 `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	OrganizationUuid string                 `protobuf:"bytes,4,opt,name=organization_uuid,json=organizationUuid,proto3" json:"organization_uuid,omitempty"`
	OrganizationName string                 `protobuf:"bytes,5,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	Role             string                 `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UserRegisteredNotification) Reset() {
	*x = UserRegisteredNotification{}
	mi := &file_proto_scg_user_v1_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegisteredNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisteredNotification) ProtoMessage() {}

func (x *UserRegisteredNotification) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_user_v1_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisteredNotification.ProtoReflect.Descriptor instead.
func (*UserRegisteredNotification) Descriptor() ([]byte, []int) {
	return file_proto_scg_user_v1_events_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegisteredNotification) GetBase() *v1.NotificationEvent {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserRegisteredNotification) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *UserRegisteredNotification) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserRegisteredNotification) GetOrganizationUuid() string {
	if x != nil {
		return x.OrganizationUuid
	}
	return ""
}

func (x *UserRegisteredNotification) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *UserRegisteredNotification) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_proto_scg_user_v1_events_proto protoreflect.FileDescriptor

const file_proto_scg_user_v1_events_proto_rawDesc = "" +
	"\n" +
	"\x1eproto/scg/user/v1/events.proto\x12\"proto.scg.notification.user.events\x1a proto/scg/shared/v1/events.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8f\x03\n" +
	"\x17UserInvitedNotification\x12K\n" +
	"\x04base\x18\x01 \x01(\v27.proto.scg.notification.shared.events.NotificationEventR\x04base\x12!\n" +
	"\finviter_uuid\x18\x02 \x01(\tR\vinviterUuid\x12!\n" +
	"\finviter_name\x18\x03 \x01(\tR\vinviterName\x12#\n" +
	"\rinvitee_email\x18\x04 \x01(\tR\finviteeEmail\x12+\n" +
	"\x11organization_uuid\x18\x05 \x01(\tR\x10organizationUuid\x12+\n" +
	"\x11organization_name\x18\x06 \x01(\tR\x10organizationName\x12\x12\n" +
	"\x04role\x18\a \x01(\tR\x04role\x12N\n" +
	"\x15invitation_expires_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\x13invitationExpiresAt\"\x91\x02\n" +
	"\x1aUserRegisteredNotification\x12K\n" +
	"\x04base\x18\x01 \x01(\v27.proto.scg.notification.shared.events.NotificationEventR\x04base\x12\x1b\n" +
	"\tuser_uuid\x18\x02 \x01(\tR\buserUuid\x12\x1b\n" +
	"\tuser_name\x18\x03 \x01(\tR\buserName\x12+\n" +
	"\x11organization_uuid\x18\x04 \x01(\tR\x10organizationUuid\x12+\n" +
	"\x11organization_name\x18\x05 \x01(\tR\x10organizationName\x12\x12\n" +
	"\x04role\x18\x06 \x01(\tR\x04roleBKZIgithub.com/hbttundar/scg-notification-contracts/gen/go/scg/user/v1;userv1b\x06proto3"

var (
	file_proto_scg_user_v1_events_proto_rawDescOnce sync.Once
	file_proto_scg_user_v1_events_proto_rawDescData []byte
)

func file_proto_scg_user_v1_events_proto_rawDescGZIP() []byte {
	file_proto_scg_user_v1_events_proto_rawDescOnce.Do(func() {
		file_proto_scg_user_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_user_v1_events_proto_rawDesc), len(file_proto_scg_user_v1_events_proto_rawDesc)))
	})
	return file_proto_scg_user_v1_events_proto_rawDescData
}

var file_proto_scg_user_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_scg_user_v1_events_proto_goTypes = []any{
	(*UserInvitedNotification)(nil),    // 0: proto.scg.notification.user.events.UserInvitedNotification
	(*UserRegisteredNotification)(nil), // 1: proto.scg.notification.user.events.UserRegisteredNotification
	(*v1.NotificationEvent)(nil),       // 2: proto.scg.notification.shared.events.NotificationEvent
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
}
var file_proto_scg_user_v1_events_proto_depIdxs = []int32{
	2, // 0: proto.scg.notification.user.events.UserInvitedNotification.base:type_name -> proto.scg.notification.shared.events.NotificationEvent
	3, // 1: proto.scg.notification.user.events.UserInvitedNotification.invitation_expires_at:type_name -> google.protobuf.Timestamp
	2, // 2: proto.scg.notification.user.events.UserRegisteredNotification.base:type_name -> proto.scg.notification.shared.events.NotificationEvent
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_scg_user_v1_events_proto_init() }
func file_proto_scg_user_v1_events_proto_init() {
	if File_proto_scg_user_v1_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_user_v1_events_proto_rawDesc), len(file_proto_scg_user_v1_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_user_v1_events_proto_goTypes,
		DependencyIndexes: file_proto_scg_user_v1_events_proto_depIdxs,
		MessageInfos:      file_proto_scg_user_v1_events_proto_msgTypes,
	}.Build()
	File_proto_scg_user_v1_events_proto = out.File
	file_proto_scg_user_v1_events_proto_goTypes = nil
	file_proto_scg_user_v1_events_proto_depIdxs = nil
}
