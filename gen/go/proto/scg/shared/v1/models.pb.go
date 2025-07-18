// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/shared/v1/models.proto

package sharedv1

import (
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

// NotificationTemplate represents a template for a notification
type NotificationTemplate struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Uuid            string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type            NotificationType       `protobuf:"varint,4,opt,name=type,proto3,enum=proto.scg.notification.shared.enums.NotificationType" json:"type,omitempty"`
	Category        NotificationCategory   `protobuf:"varint,5,opt,name=category,proto3,enum=proto.scg.notification.shared.enums.NotificationCategory" json:"category,omitempty"`
	DeliveryChannel DeliveryChannel        `protobuf:"varint,6,opt,name=delivery_channel,json=deliveryChannel,proto3,enum=proto.scg.notification.shared.enums.DeliveryChannel" json:"delivery_channel,omitempty"`
	SubjectTemplate string                 `protobuf:"bytes,7,opt,name=subject_template,json=subjectTemplate,proto3" json:"subject_template,omitempty"`
	BodyTemplate    string                 `protobuf:"bytes,8,opt,name=body_template,json=bodyTemplate,proto3" json:"body_template,omitempty"`
	Metadata        map[string]string      `protobuf:"bytes,9,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Version         string                 `protobuf:"bytes,12,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *NotificationTemplate) Reset() {
	*x = NotificationTemplate{}
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTemplate) ProtoMessage() {}

func (x *NotificationTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTemplate.ProtoReflect.Descriptor instead.
func (*NotificationTemplate) Descriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_models_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationTemplate) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *NotificationTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotificationTemplate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NotificationTemplate) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_NOTIFICATION_TYPE_UNSPECIFIED
}

func (x *NotificationTemplate) GetCategory() NotificationCategory {
	if x != nil {
		return x.Category
	}
	return NotificationCategory_NOTIFICATION_CATEGORY_UNSPECIFIED
}

func (x *NotificationTemplate) GetDeliveryChannel() DeliveryChannel {
	if x != nil {
		return x.DeliveryChannel
	}
	return DeliveryChannel_DELIVERY_CHANNEL_UNSPECIFIED
}

func (x *NotificationTemplate) GetSubjectTemplate() string {
	if x != nil {
		return x.SubjectTemplate
	}
	return ""
}

func (x *NotificationTemplate) GetBodyTemplate() string {
	if x != nil {
		return x.BodyTemplate
	}
	return ""
}

func (x *NotificationTemplate) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *NotificationTemplate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NotificationTemplate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *NotificationTemplate) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// NotificationPreference represents a user's preferences for receiving notifications
type NotificationPreference struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserUuid        string                 `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Category        NotificationCategory   `protobuf:"varint,2,opt,name=category,proto3,enum=proto.scg.notification.shared.enums.NotificationCategory" json:"category,omitempty"`
	Type            NotificationType       `protobuf:"varint,3,opt,name=type,proto3,enum=proto.scg.notification.shared.enums.NotificationType" json:"type,omitempty"`
	EnabledChannels []DeliveryChannel      `protobuf:"varint,4,rep,packed,name=enabled_channels,json=enabledChannels,proto3,enum=proto.scg.notification.shared.enums.DeliveryChannel" json:"enabled_channels,omitempty"`
	IsOptedOut      bool                   `protobuf:"varint,5,opt,name=is_opted_out,json=isOptedOut,proto3" json:"is_opted_out,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *NotificationPreference) Reset() {
	*x = NotificationPreference{}
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationPreference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationPreference) ProtoMessage() {}

func (x *NotificationPreference) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationPreference.ProtoReflect.Descriptor instead.
func (*NotificationPreference) Descriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_models_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationPreference) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *NotificationPreference) GetCategory() NotificationCategory {
	if x != nil {
		return x.Category
	}
	return NotificationCategory_NOTIFICATION_CATEGORY_UNSPECIFIED
}

func (x *NotificationPreference) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_NOTIFICATION_TYPE_UNSPECIFIED
}

func (x *NotificationPreference) GetEnabledChannels() []DeliveryChannel {
	if x != nil {
		return x.EnabledChannels
	}
	return nil
}

func (x *NotificationPreference) GetIsOptedOut() bool {
	if x != nil {
		return x.IsOptedOut
	}
	return false
}

func (x *NotificationPreference) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// NotificationDelivery represents the delivery status of a notification
type NotificationDelivery struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Uuid             string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	NotificationUuid string                 `protobuf:"bytes,2,opt,name=notification_uuid,json=notificationUuid,proto3" json:"notification_uuid,omitempty"`
	RecipientUuid    string                 `protobuf:"bytes,3,opt,name=recipient_uuid,json=recipientUuid,proto3" json:"recipient_uuid,omitempty"`
	Channel          DeliveryChannel        `protobuf:"varint,4,opt,name=channel,proto3,enum=proto.scg.notification.shared.enums.DeliveryChannel" json:"channel,omitempty"`
	Status           NotificationStatus     `protobuf:"varint,5,opt,name=status,proto3,enum=proto.scg.notification.shared.enums.NotificationStatus" json:"status,omitempty"`
	ErrorMessage     string                 `protobuf:"bytes,6,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	SentAt           *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	DeliveredAt      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=delivered_at,json=deliveredAt,proto3" json:"delivered_at,omitempty"`
	ReadAt           *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=read_at,json=readAt,proto3" json:"read_at,omitempty"`
	DeliveryMetadata map[string]string      `protobuf:"bytes,10,rep,name=delivery_metadata,json=deliveryMetadata,proto3" json:"delivery_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *NotificationDelivery) Reset() {
	*x = NotificationDelivery{}
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationDelivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationDelivery) ProtoMessage() {}

func (x *NotificationDelivery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationDelivery.ProtoReflect.Descriptor instead.
func (*NotificationDelivery) Descriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_models_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationDelivery) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *NotificationDelivery) GetNotificationUuid() string {
	if x != nil {
		return x.NotificationUuid
	}
	return ""
}

func (x *NotificationDelivery) GetRecipientUuid() string {
	if x != nil {
		return x.RecipientUuid
	}
	return ""
}

func (x *NotificationDelivery) GetChannel() DeliveryChannel {
	if x != nil {
		return x.Channel
	}
	return DeliveryChannel_DELIVERY_CHANNEL_UNSPECIFIED
}

func (x *NotificationDelivery) GetStatus() NotificationStatus {
	if x != nil {
		return x.Status
	}
	return NotificationStatus_NOTIFICATION_STATUS_UNSPECIFIED
}

func (x *NotificationDelivery) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *NotificationDelivery) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

func (x *NotificationDelivery) GetDeliveredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveredAt
	}
	return nil
}

func (x *NotificationDelivery) GetReadAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadAt
	}
	return nil
}

func (x *NotificationDelivery) GetDeliveryMetadata() map[string]string {
	if x != nil {
		return x.DeliveryMetadata
	}
	return nil
}

// Notification represents a notification to be sent to a recipient
type Notification struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Uuid           string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	EventUuid      string                 `protobuf:"bytes,2,opt,name=event_uuid,json=eventUuid,proto3" json:"event_uuid,omitempty"`
	EventType      string                 `protobuf:"bytes,3,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	EntityUuid     string                 `protobuf:"bytes,4,opt,name=entity_uuid,json=entityUuid,proto3" json:"entity_uuid,omitempty"`
	EntityType     string                 `protobuf:"bytes,5,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	WorkflowStep   string                 `protobuf:"bytes,6,opt,name=workflow_step,json=workflowStep,proto3" json:"workflow_step,omitempty"`
	Reason         string                 `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	Type           NotificationType       `protobuf:"varint,8,opt,name=type,proto3,enum=proto.scg.notification.shared.enums.NotificationType" json:"type,omitempty"`
	Category       NotificationCategory   `protobuf:"varint,9,opt,name=category,proto3,enum=proto.scg.notification.shared.enums.NotificationCategory" json:"category,omitempty"`
	Severity       NotificationSeverity   `protobuf:"varint,10,opt,name=severity,proto3,enum=proto.scg.notification.shared.enums.NotificationSeverity" json:"severity,omitempty"`
	Audience       NotificationAudience   `protobuf:"varint,11,opt,name=audience,proto3,enum=proto.scg.notification.shared.enums.NotificationAudience" json:"audience,omitempty"`
	Subject        string                 `protobuf:"bytes,12,opt,name=subject,proto3" json:"subject,omitempty"`
	Body           string                 `protobuf:"bytes,13,opt,name=body,proto3" json:"body,omitempty"`
	Data           map[string]string      `protobuf:"bytes,14,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	RecipientUuids []string               `protobuf:"bytes,15,rep,name=recipient_uuids,json=recipientUuids,proto3" json:"recipient_uuids,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt      *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Status         NotificationStatus     `protobuf:"varint,18,opt,name=status,proto3,enum=proto.scg.notification.shared.enums.NotificationStatus" json:"status,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_models_proto_rawDescGZIP(), []int{3}
}

func (x *Notification) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Notification) GetEventUuid() string {
	if x != nil {
		return x.EventUuid
	}
	return ""
}

func (x *Notification) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Notification) GetEntityUuid() string {
	if x != nil {
		return x.EntityUuid
	}
	return ""
}

func (x *Notification) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *Notification) GetWorkflowStep() string {
	if x != nil {
		return x.WorkflowStep
	}
	return ""
}

func (x *Notification) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Notification) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_NOTIFICATION_TYPE_UNSPECIFIED
}

func (x *Notification) GetCategory() NotificationCategory {
	if x != nil {
		return x.Category
	}
	return NotificationCategory_NOTIFICATION_CATEGORY_UNSPECIFIED
}

func (x *Notification) GetSeverity() NotificationSeverity {
	if x != nil {
		return x.Severity
	}
	return NotificationSeverity_NOTIFICATION_SEVERITY_UNSPECIFIED
}

func (x *Notification) GetAudience() NotificationAudience {
	if x != nil {
		return x.Audience
	}
	return NotificationAudience_NOTIFICATION_AUDIENCE_UNSPECIFIED
}

func (x *Notification) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Notification) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Notification) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Notification) GetRecipientUuids() []string {
	if x != nil {
		return x.RecipientUuids
	}
	return nil
}

func (x *Notification) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Notification) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *Notification) GetStatus() NotificationStatus {
	if x != nil {
		return x.Status
	}
	return NotificationStatus_NOTIFICATION_STATUS_UNSPECIFIED
}

// UserNotification represents a notification for a specific user
type UserNotification struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Uuid             string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	NotificationUuid string                 `protobuf:"bytes,2,opt,name=notification_uuid,json=notificationUuid,proto3" json:"notification_uuid,omitempty"`
	UserUuid         string                 `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	IsRead           bool                   `protobuf:"varint,4,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	ReadAt           *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=read_at,json=readAt,proto3" json:"read_at,omitempty"`
	IsArchived       bool                   `protobuf:"varint,6,opt,name=is_archived,json=isArchived,proto3" json:"is_archived,omitempty"`
	ArchivedAt       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=archived_at,json=archivedAt,proto3" json:"archived_at,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UserNotification) Reset() {
	*x = UserNotification{}
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotification) ProtoMessage() {}

func (x *UserNotification) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotification.ProtoReflect.Descriptor instead.
func (*UserNotification) Descriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_models_proto_rawDescGZIP(), []int{4}
}

func (x *UserNotification) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserNotification) GetNotificationUuid() string {
	if x != nil {
		return x.NotificationUuid
	}
	return ""
}

func (x *UserNotification) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *UserNotification) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *UserNotification) GetReadAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadAt
	}
	return nil
}

func (x *UserNotification) GetIsArchived() bool {
	if x != nil {
		return x.IsArchived
	}
	return false
}

func (x *UserNotification) GetArchivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ArchivedAt
	}
	return nil
}

// NotificationDigest represents a digest of multiple notifications
type NotificationDigest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Uuid              string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid          string                 `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Category          NotificationCategory   `protobuf:"varint,3,opt,name=category,proto3,enum=proto.scg.notification.shared.enums.NotificationCategory" json:"category,omitempty"`
	DigestPeriodStart *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=digest_period_start,json=digestPeriodStart,proto3" json:"digest_period_start,omitempty"`
	DigestPeriodEnd   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=digest_period_end,json=digestPeriodEnd,proto3" json:"digest_period_end,omitempty"`
	NotificationUuids []string               `protobuf:"bytes,6,rep,name=notification_uuids,json=notificationUuids,proto3" json:"notification_uuids,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Status            NotificationStatus     `protobuf:"varint,8,opt,name=status,proto3,enum=proto.scg.notification.shared.enums.NotificationStatus" json:"status,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *NotificationDigest) Reset() {
	*x = NotificationDigest{}
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationDigest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationDigest) ProtoMessage() {}

func (x *NotificationDigest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_shared_v1_models_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationDigest.ProtoReflect.Descriptor instead.
func (*NotificationDigest) Descriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_models_proto_rawDescGZIP(), []int{5}
}

func (x *NotificationDigest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *NotificationDigest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *NotificationDigest) GetCategory() NotificationCategory {
	if x != nil {
		return x.Category
	}
	return NotificationCategory_NOTIFICATION_CATEGORY_UNSPECIFIED
}

func (x *NotificationDigest) GetDigestPeriodStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DigestPeriodStart
	}
	return nil
}

func (x *NotificationDigest) GetDigestPeriodEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.DigestPeriodEnd
	}
	return nil
}

func (x *NotificationDigest) GetNotificationUuids() []string {
	if x != nil {
		return x.NotificationUuids
	}
	return nil
}

func (x *NotificationDigest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NotificationDigest) GetStatus() NotificationStatus {
	if x != nil {
		return x.Status
	}
	return NotificationStatus_NOTIFICATION_STATUS_UNSPECIFIED
}

var File_proto_scg_shared_v1_models_proto protoreflect.FileDescriptor

const file_proto_scg_shared_v1_models_proto_rawDesc = "" +
	"\n" +
	" proto/scg/shared/v1/models.proto\x12$proto.scg.notification.shared.models\x1a\x1fproto/scg/shared/v1/enums.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe6\x05\n" +
	"\x14NotificationTemplate\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12I\n" +
	"\x04type\x18\x04 \x01(\x0e25.proto.scg.notification.shared.enums.NotificationTypeR\x04type\x12U\n" +
	"\bcategory\x18\x05 \x01(\x0e29.proto.scg.notification.shared.enums.NotificationCategoryR\bcategory\x12_\n" +
	"\x10delivery_channel\x18\x06 \x01(\x0e24.proto.scg.notification.shared.enums.DeliveryChannelR\x0fdeliveryChannel\x12)\n" +
	"\x10subject_template\x18\a \x01(\tR\x0fsubjectTemplate\x12#\n" +
	"\rbody_template\x18\b \x01(\tR\fbodyTemplate\x12d\n" +
	"\bmetadata\x18\t \x03(\v2H.proto.scg.notification.shared.models.NotificationTemplate.MetadataEntryR\bmetadata\x129\n" +
	"\n" +
	"created_at\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x18\n" +
	"\aversion\x18\f \x01(\tR\aversion\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x95\x03\n" +
	"\x16NotificationPreference\x12\x1b\n" +
	"\tuser_uuid\x18\x01 \x01(\tR\buserUuid\x12U\n" +
	"\bcategory\x18\x02 \x01(\x0e29.proto.scg.notification.shared.enums.NotificationCategoryR\bcategory\x12I\n" +
	"\x04type\x18\x03 \x01(\x0e25.proto.scg.notification.shared.enums.NotificationTypeR\x04type\x12_\n" +
	"\x10enabled_channels\x18\x04 \x03(\x0e24.proto.scg.notification.shared.enums.DeliveryChannelR\x0fenabledChannels\x12 \n" +
	"\fis_opted_out\x18\x05 \x01(\bR\n" +
	"isOptedOut\x129\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\xb1\x05\n" +
	"\x14NotificationDelivery\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12+\n" +
	"\x11notification_uuid\x18\x02 \x01(\tR\x10notificationUuid\x12%\n" +
	"\x0erecipient_uuid\x18\x03 \x01(\tR\rrecipientUuid\x12N\n" +
	"\achannel\x18\x04 \x01(\x0e24.proto.scg.notification.shared.enums.DeliveryChannelR\achannel\x12O\n" +
	"\x06status\x18\x05 \x01(\x0e27.proto.scg.notification.shared.enums.NotificationStatusR\x06status\x12#\n" +
	"\rerror_message\x18\x06 \x01(\tR\ferrorMessage\x123\n" +
	"\asent_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\x06sentAt\x12=\n" +
	"\fdelivered_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\vdeliveredAt\x123\n" +
	"\aread_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\x06readAt\x12}\n" +
	"\x11delivery_metadata\x18\n" +
	" \x03(\v2P.proto.scg.notification.shared.models.NotificationDelivery.DeliveryMetadataEntryR\x10deliveryMetadata\x1aC\n" +
	"\x15DeliveryMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xd8\a\n" +
	"\fNotification\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x1d\n" +
	"\n" +
	"event_uuid\x18\x02 \x01(\tR\teventUuid\x12\x1d\n" +
	"\n" +
	"event_type\x18\x03 \x01(\tR\teventType\x12\x1f\n" +
	"\ventity_uuid\x18\x04 \x01(\tR\n" +
	"entityUuid\x12\x1f\n" +
	"\ventity_type\x18\x05 \x01(\tR\n" +
	"entityType\x12#\n" +
	"\rworkflow_step\x18\x06 \x01(\tR\fworkflowStep\x12\x16\n" +
	"\x06reason\x18\a \x01(\tR\x06reason\x12I\n" +
	"\x04type\x18\b \x01(\x0e25.proto.scg.notification.shared.enums.NotificationTypeR\x04type\x12U\n" +
	"\bcategory\x18\t \x01(\x0e29.proto.scg.notification.shared.enums.NotificationCategoryR\bcategory\x12U\n" +
	"\bseverity\x18\n" +
	" \x01(\x0e29.proto.scg.notification.shared.enums.NotificationSeverityR\bseverity\x12U\n" +
	"\baudience\x18\v \x01(\x0e29.proto.scg.notification.shared.enums.NotificationAudienceR\baudience\x12\x18\n" +
	"\asubject\x18\f \x01(\tR\asubject\x12\x12\n" +
	"\x04body\x18\r \x01(\tR\x04body\x12P\n" +
	"\x04data\x18\x0e \x03(\v2<.proto.scg.notification.shared.models.Notification.DataEntryR\x04data\x12'\n" +
	"\x0frecipient_uuids\x18\x0f \x03(\tR\x0erecipientUuids\x129\n" +
	"\n" +
	"created_at\x18\x10 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"expires_at\x18\x11 \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt\x12O\n" +
	"\x06status\x18\x12 \x01(\x0e27.proto.scg.notification.shared.enums.NotificationStatusR\x06status\x1a7\n" +
	"\tDataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x9c\x02\n" +
	"\x10UserNotification\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12+\n" +
	"\x11notification_uuid\x18\x02 \x01(\tR\x10notificationUuid\x12\x1b\n" +
	"\tuser_uuid\x18\x03 \x01(\tR\buserUuid\x12\x17\n" +
	"\ais_read\x18\x04 \x01(\bR\x06isRead\x123\n" +
	"\aread_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\x06readAt\x12\x1f\n" +
	"\vis_archived\x18\x06 \x01(\bR\n" +
	"isArchived\x12;\n" +
	"\varchived_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"archivedAt\"\xeb\x03\n" +
	"\x12NotificationDigest\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x1b\n" +
	"\tuser_uuid\x18\x02 \x01(\tR\buserUuid\x12U\n" +
	"\bcategory\x18\x03 \x01(\x0e29.proto.scg.notification.shared.enums.NotificationCategoryR\bcategory\x12J\n" +
	"\x13digest_period_start\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\x11digestPeriodStart\x12F\n" +
	"\x11digest_period_end\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\x0fdigestPeriodEnd\x12-\n" +
	"\x12notification_uuids\x18\x06 \x03(\tR\x11notificationUuids\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12O\n" +
	"\x06status\x18\b \x01(\x0e27.proto.scg.notification.shared.enums.NotificationStatusR\x06statusBOZMgithub.com/hbttundar/scg-notification-contracts/gen/go/scg/shared/v1;sharedv1b\x06proto3"

var (
	file_proto_scg_shared_v1_models_proto_rawDescOnce sync.Once
	file_proto_scg_shared_v1_models_proto_rawDescData []byte
)

func file_proto_scg_shared_v1_models_proto_rawDescGZIP() []byte {
	file_proto_scg_shared_v1_models_proto_rawDescOnce.Do(func() {
		file_proto_scg_shared_v1_models_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_shared_v1_models_proto_rawDesc), len(file_proto_scg_shared_v1_models_proto_rawDesc)))
	})
	return file_proto_scg_shared_v1_models_proto_rawDescData
}

var file_proto_scg_shared_v1_models_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_scg_shared_v1_models_proto_goTypes = []any{
	(*NotificationTemplate)(nil),   // 0: proto.scg.notification.shared.models.NotificationTemplate
	(*NotificationPreference)(nil), // 1: proto.scg.notification.shared.models.NotificationPreference
	(*NotificationDelivery)(nil),   // 2: proto.scg.notification.shared.models.NotificationDelivery
	(*Notification)(nil),           // 3: proto.scg.notification.shared.models.Notification
	(*UserNotification)(nil),       // 4: proto.scg.notification.shared.models.UserNotification
	(*NotificationDigest)(nil),     // 5: proto.scg.notification.shared.models.NotificationDigest
	nil,                            // 6: proto.scg.notification.shared.models.NotificationTemplate.MetadataEntry
	nil,                            // 7: proto.scg.notification.shared.models.NotificationDelivery.DeliveryMetadataEntry
	nil,                            // 8: proto.scg.notification.shared.models.Notification.DataEntry
	(NotificationType)(0),          // 9: proto.scg.notification.shared.enums.NotificationType
	(NotificationCategory)(0),      // 10: proto.scg.notification.shared.enums.NotificationCategory
	(DeliveryChannel)(0),           // 11: proto.scg.notification.shared.enums.DeliveryChannel
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
	(NotificationStatus)(0),        // 13: proto.scg.notification.shared.enums.NotificationStatus
	(NotificationSeverity)(0),      // 14: proto.scg.notification.shared.enums.NotificationSeverity
	(NotificationAudience)(0),      // 15: proto.scg.notification.shared.enums.NotificationAudience
}
var file_proto_scg_shared_v1_models_proto_depIdxs = []int32{
	9,  // 0: proto.scg.notification.shared.models.NotificationTemplate.type:type_name -> proto.scg.notification.shared.enums.NotificationType
	10, // 1: proto.scg.notification.shared.models.NotificationTemplate.category:type_name -> proto.scg.notification.shared.enums.NotificationCategory
	11, // 2: proto.scg.notification.shared.models.NotificationTemplate.delivery_channel:type_name -> proto.scg.notification.shared.enums.DeliveryChannel
	6,  // 3: proto.scg.notification.shared.models.NotificationTemplate.metadata:type_name -> proto.scg.notification.shared.models.NotificationTemplate.MetadataEntry
	12, // 4: proto.scg.notification.shared.models.NotificationTemplate.created_at:type_name -> google.protobuf.Timestamp
	12, // 5: proto.scg.notification.shared.models.NotificationTemplate.updated_at:type_name -> google.protobuf.Timestamp
	10, // 6: proto.scg.notification.shared.models.NotificationPreference.category:type_name -> proto.scg.notification.shared.enums.NotificationCategory
	9,  // 7: proto.scg.notification.shared.models.NotificationPreference.type:type_name -> proto.scg.notification.shared.enums.NotificationType
	11, // 8: proto.scg.notification.shared.models.NotificationPreference.enabled_channels:type_name -> proto.scg.notification.shared.enums.DeliveryChannel
	12, // 9: proto.scg.notification.shared.models.NotificationPreference.updated_at:type_name -> google.protobuf.Timestamp
	11, // 10: proto.scg.notification.shared.models.NotificationDelivery.channel:type_name -> proto.scg.notification.shared.enums.DeliveryChannel
	13, // 11: proto.scg.notification.shared.models.NotificationDelivery.status:type_name -> proto.scg.notification.shared.enums.NotificationStatus
	12, // 12: proto.scg.notification.shared.models.NotificationDelivery.sent_at:type_name -> google.protobuf.Timestamp
	12, // 13: proto.scg.notification.shared.models.NotificationDelivery.delivered_at:type_name -> google.protobuf.Timestamp
	12, // 14: proto.scg.notification.shared.models.NotificationDelivery.read_at:type_name -> google.protobuf.Timestamp
	7,  // 15: proto.scg.notification.shared.models.NotificationDelivery.delivery_metadata:type_name -> proto.scg.notification.shared.models.NotificationDelivery.DeliveryMetadataEntry
	9,  // 16: proto.scg.notification.shared.models.Notification.type:type_name -> proto.scg.notification.shared.enums.NotificationType
	10, // 17: proto.scg.notification.shared.models.Notification.category:type_name -> proto.scg.notification.shared.enums.NotificationCategory
	14, // 18: proto.scg.notification.shared.models.Notification.severity:type_name -> proto.scg.notification.shared.enums.NotificationSeverity
	15, // 19: proto.scg.notification.shared.models.Notification.audience:type_name -> proto.scg.notification.shared.enums.NotificationAudience
	8,  // 20: proto.scg.notification.shared.models.Notification.data:type_name -> proto.scg.notification.shared.models.Notification.DataEntry
	12, // 21: proto.scg.notification.shared.models.Notification.created_at:type_name -> google.protobuf.Timestamp
	12, // 22: proto.scg.notification.shared.models.Notification.expires_at:type_name -> google.protobuf.Timestamp
	13, // 23: proto.scg.notification.shared.models.Notification.status:type_name -> proto.scg.notification.shared.enums.NotificationStatus
	12, // 24: proto.scg.notification.shared.models.UserNotification.read_at:type_name -> google.protobuf.Timestamp
	12, // 25: proto.scg.notification.shared.models.UserNotification.archived_at:type_name -> google.protobuf.Timestamp
	10, // 26: proto.scg.notification.shared.models.NotificationDigest.category:type_name -> proto.scg.notification.shared.enums.NotificationCategory
	12, // 27: proto.scg.notification.shared.models.NotificationDigest.digest_period_start:type_name -> google.protobuf.Timestamp
	12, // 28: proto.scg.notification.shared.models.NotificationDigest.digest_period_end:type_name -> google.protobuf.Timestamp
	12, // 29: proto.scg.notification.shared.models.NotificationDigest.created_at:type_name -> google.protobuf.Timestamp
	13, // 30: proto.scg.notification.shared.models.NotificationDigest.status:type_name -> proto.scg.notification.shared.enums.NotificationStatus
	31, // [31:31] is the sub-list for method output_type
	31, // [31:31] is the sub-list for method input_type
	31, // [31:31] is the sub-list for extension type_name
	31, // [31:31] is the sub-list for extension extendee
	0,  // [0:31] is the sub-list for field type_name
}

func init() { file_proto_scg_shared_v1_models_proto_init() }
func file_proto_scg_shared_v1_models_proto_init() {
	if File_proto_scg_shared_v1_models_proto != nil {
		return
	}
	file_proto_scg_shared_v1_enums_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_shared_v1_models_proto_rawDesc), len(file_proto_scg_shared_v1_models_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_shared_v1_models_proto_goTypes,
		DependencyIndexes: file_proto_scg_shared_v1_models_proto_depIdxs,
		MessageInfos:      file_proto_scg_shared_v1_models_proto_msgTypes,
	}.Build()
	File_proto_scg_shared_v1_models_proto = out.File
	file_proto_scg_shared_v1_models_proto_goTypes = nil
	file_proto_scg_shared_v1_models_proto_depIdxs = nil
}
