// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/shared/v1/enums.proto

package sharedv1

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

// DeliveryChannel represents the different channels through which notifications can be delivered
type DeliveryChannel int32

const (
	DeliveryChannel_DELIVERY_CHANNEL_UNSPECIFIED DeliveryChannel = 0
	DeliveryChannel_DELIVERY_CHANNEL_EMAIL       DeliveryChannel = 1
	DeliveryChannel_DELIVERY_CHANNEL_SMS         DeliveryChannel = 2
	DeliveryChannel_DELIVERY_CHANNEL_PUSH        DeliveryChannel = 3
	DeliveryChannel_DELIVERY_CHANNEL_WEBHOOK     DeliveryChannel = 4
	DeliveryChannel_DELIVERY_CHANNEL_DASHBOARD   DeliveryChannel = 5
	DeliveryChannel_DELIVERY_CHANNEL_IN_APP      DeliveryChannel = 6
	DeliveryChannel_DELIVERY_CHANNEL_ESCALATION  DeliveryChannel = 7
	DeliveryChannel_DELIVERY_CHANNEL_DIGEST      DeliveryChannel = 8
)

// Enum value maps for DeliveryChannel.
var (
	DeliveryChannel_name = map[int32]string{
		0: "DELIVERY_CHANNEL_UNSPECIFIED",
		1: "DELIVERY_CHANNEL_EMAIL",
		2: "DELIVERY_CHANNEL_SMS",
		3: "DELIVERY_CHANNEL_PUSH",
		4: "DELIVERY_CHANNEL_WEBHOOK",
		5: "DELIVERY_CHANNEL_DASHBOARD",
		6: "DELIVERY_CHANNEL_IN_APP",
		7: "DELIVERY_CHANNEL_ESCALATION",
		8: "DELIVERY_CHANNEL_DIGEST",
	}
	DeliveryChannel_value = map[string]int32{
		"DELIVERY_CHANNEL_UNSPECIFIED": 0,
		"DELIVERY_CHANNEL_EMAIL":       1,
		"DELIVERY_CHANNEL_SMS":         2,
		"DELIVERY_CHANNEL_PUSH":        3,
		"DELIVERY_CHANNEL_WEBHOOK":     4,
		"DELIVERY_CHANNEL_DASHBOARD":   5,
		"DELIVERY_CHANNEL_IN_APP":      6,
		"DELIVERY_CHANNEL_ESCALATION":  7,
		"DELIVERY_CHANNEL_DIGEST":      8,
	}
)

func (x DeliveryChannel) Enum() *DeliveryChannel {
	p := new(DeliveryChannel)
	*p = x
	return p
}

func (x DeliveryChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_shared_v1_enums_proto_enumTypes[0].Descriptor()
}

func (DeliveryChannel) Type() protoreflect.EnumType {
	return &file_proto_scg_shared_v1_enums_proto_enumTypes[0]
}

func (x DeliveryChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryChannel.Descriptor instead.
func (DeliveryChannel) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_enums_proto_rawDescGZIP(), []int{0}
}

// NotificationType represents the different types of notifications
type NotificationType int32

const (
	NotificationType_NOTIFICATION_TYPE_UNSPECIFIED     NotificationType = 0
	NotificationType_NOTIFICATION_TYPE_ALERT           NotificationType = 1
	NotificationType_NOTIFICATION_TYPE_INFORMATION     NotificationType = 2
	NotificationType_NOTIFICATION_TYPE_CONFIRMATION    NotificationType = 3
	NotificationType_NOTIFICATION_TYPE_REMINDER        NotificationType = 4
	NotificationType_NOTIFICATION_TYPE_DIGEST          NotificationType = 5
	NotificationType_NOTIFICATION_TYPE_REPORT          NotificationType = 6
	NotificationType_NOTIFICATION_TYPE_ACTION_REQUIRED NotificationType = 7
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "NOTIFICATION_TYPE_UNSPECIFIED",
		1: "NOTIFICATION_TYPE_ALERT",
		2: "NOTIFICATION_TYPE_INFORMATION",
		3: "NOTIFICATION_TYPE_CONFIRMATION",
		4: "NOTIFICATION_TYPE_REMINDER",
		5: "NOTIFICATION_TYPE_DIGEST",
		6: "NOTIFICATION_TYPE_REPORT",
		7: "NOTIFICATION_TYPE_ACTION_REQUIRED",
	}
	NotificationType_value = map[string]int32{
		"NOTIFICATION_TYPE_UNSPECIFIED":     0,
		"NOTIFICATION_TYPE_ALERT":           1,
		"NOTIFICATION_TYPE_INFORMATION":     2,
		"NOTIFICATION_TYPE_CONFIRMATION":    3,
		"NOTIFICATION_TYPE_REMINDER":        4,
		"NOTIFICATION_TYPE_DIGEST":          5,
		"NOTIFICATION_TYPE_REPORT":          6,
		"NOTIFICATION_TYPE_ACTION_REQUIRED": 7,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_shared_v1_enums_proto_enumTypes[1].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_proto_scg_shared_v1_enums_proto_enumTypes[1]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_enums_proto_rawDescGZIP(), []int{1}
}

// NotificationSeverity represents the severity level of a notification
type NotificationSeverity int32

const (
	NotificationSeverity_NOTIFICATION_SEVERITY_UNSPECIFIED NotificationSeverity = 0
	NotificationSeverity_NOTIFICATION_SEVERITY_LOW         NotificationSeverity = 1
	NotificationSeverity_NOTIFICATION_SEVERITY_MEDIUM      NotificationSeverity = 2
	NotificationSeverity_NOTIFICATION_SEVERITY_HIGH        NotificationSeverity = 3
	NotificationSeverity_NOTIFICATION_SEVERITY_CRITICAL    NotificationSeverity = 4
)

// Enum value maps for NotificationSeverity.
var (
	NotificationSeverity_name = map[int32]string{
		0: "NOTIFICATION_SEVERITY_UNSPECIFIED",
		1: "NOTIFICATION_SEVERITY_LOW",
		2: "NOTIFICATION_SEVERITY_MEDIUM",
		3: "NOTIFICATION_SEVERITY_HIGH",
		4: "NOTIFICATION_SEVERITY_CRITICAL",
	}
	NotificationSeverity_value = map[string]int32{
		"NOTIFICATION_SEVERITY_UNSPECIFIED": 0,
		"NOTIFICATION_SEVERITY_LOW":         1,
		"NOTIFICATION_SEVERITY_MEDIUM":      2,
		"NOTIFICATION_SEVERITY_HIGH":        3,
		"NOTIFICATION_SEVERITY_CRITICAL":    4,
	}
)

func (x NotificationSeverity) Enum() *NotificationSeverity {
	p := new(NotificationSeverity)
	*p = x
	return p
}

func (x NotificationSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_shared_v1_enums_proto_enumTypes[2].Descriptor()
}

func (NotificationSeverity) Type() protoreflect.EnumType {
	return &file_proto_scg_shared_v1_enums_proto_enumTypes[2]
}

func (x NotificationSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationSeverity.Descriptor instead.
func (NotificationSeverity) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_enums_proto_rawDescGZIP(), []int{2}
}

// NotificationAudience represents the target audience for a notification
type NotificationAudience int32

const (
	NotificationAudience_NOTIFICATION_AUDIENCE_UNSPECIFIED NotificationAudience = 0
	NotificationAudience_NOTIFICATION_AUDIENCE_USER        NotificationAudience = 1
	NotificationAudience_NOTIFICATION_AUDIENCE_ADMIN       NotificationAudience = 2
	NotificationAudience_NOTIFICATION_AUDIENCE_PARTNER     NotificationAudience = 3
	NotificationAudience_NOTIFICATION_AUDIENCE_SYSTEM      NotificationAudience = 4
	NotificationAudience_NOTIFICATION_AUDIENCE_INTEGRATION NotificationAudience = 5
)

// Enum value maps for NotificationAudience.
var (
	NotificationAudience_name = map[int32]string{
		0: "NOTIFICATION_AUDIENCE_UNSPECIFIED",
		1: "NOTIFICATION_AUDIENCE_USER",
		2: "NOTIFICATION_AUDIENCE_ADMIN",
		3: "NOTIFICATION_AUDIENCE_PARTNER",
		4: "NOTIFICATION_AUDIENCE_SYSTEM",
		5: "NOTIFICATION_AUDIENCE_INTEGRATION",
	}
	NotificationAudience_value = map[string]int32{
		"NOTIFICATION_AUDIENCE_UNSPECIFIED": 0,
		"NOTIFICATION_AUDIENCE_USER":        1,
		"NOTIFICATION_AUDIENCE_ADMIN":       2,
		"NOTIFICATION_AUDIENCE_PARTNER":     3,
		"NOTIFICATION_AUDIENCE_SYSTEM":      4,
		"NOTIFICATION_AUDIENCE_INTEGRATION": 5,
	}
)

func (x NotificationAudience) Enum() *NotificationAudience {
	p := new(NotificationAudience)
	*p = x
	return p
}

func (x NotificationAudience) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationAudience) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_shared_v1_enums_proto_enumTypes[3].Descriptor()
}

func (NotificationAudience) Type() protoreflect.EnumType {
	return &file_proto_scg_shared_v1_enums_proto_enumTypes[3]
}

func (x NotificationAudience) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationAudience.Descriptor instead.
func (NotificationAudience) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_enums_proto_rawDescGZIP(), []int{3}
}

// NotificationStatus represents the current status of a notification
type NotificationStatus int32

const (
	NotificationStatus_NOTIFICATION_STATUS_UNSPECIFIED NotificationStatus = 0
	NotificationStatus_NOTIFICATION_STATUS_PENDING     NotificationStatus = 1
	NotificationStatus_NOTIFICATION_STATUS_SENT        NotificationStatus = 2
	NotificationStatus_NOTIFICATION_STATUS_DELIVERED   NotificationStatus = 3
	NotificationStatus_NOTIFICATION_STATUS_READ        NotificationStatus = 4
	NotificationStatus_NOTIFICATION_STATUS_FAILED      NotificationStatus = 5
	NotificationStatus_NOTIFICATION_STATUS_CANCELLED   NotificationStatus = 6
)

// Enum value maps for NotificationStatus.
var (
	NotificationStatus_name = map[int32]string{
		0: "NOTIFICATION_STATUS_UNSPECIFIED",
		1: "NOTIFICATION_STATUS_PENDING",
		2: "NOTIFICATION_STATUS_SENT",
		3: "NOTIFICATION_STATUS_DELIVERED",
		4: "NOTIFICATION_STATUS_READ",
		5: "NOTIFICATION_STATUS_FAILED",
		6: "NOTIFICATION_STATUS_CANCELLED",
	}
	NotificationStatus_value = map[string]int32{
		"NOTIFICATION_STATUS_UNSPECIFIED": 0,
		"NOTIFICATION_STATUS_PENDING":     1,
		"NOTIFICATION_STATUS_SENT":        2,
		"NOTIFICATION_STATUS_DELIVERED":   3,
		"NOTIFICATION_STATUS_READ":        4,
		"NOTIFICATION_STATUS_FAILED":      5,
		"NOTIFICATION_STATUS_CANCELLED":   6,
	}
)

func (x NotificationStatus) Enum() *NotificationStatus {
	p := new(NotificationStatus)
	*p = x
	return p
}

func (x NotificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_shared_v1_enums_proto_enumTypes[4].Descriptor()
}

func (NotificationStatus) Type() protoreflect.EnumType {
	return &file_proto_scg_shared_v1_enums_proto_enumTypes[4]
}

func (x NotificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationStatus.Descriptor instead.
func (NotificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_enums_proto_rawDescGZIP(), []int{4}
}

// NotificationCategory represents the business domain category of a notification
type NotificationCategory int32

const (
	NotificationCategory_NOTIFICATION_CATEGORY_UNSPECIFIED  NotificationCategory = 0
	NotificationCategory_NOTIFICATION_CATEGORY_ONBOARDING   NotificationCategory = 1
	NotificationCategory_NOTIFICATION_CATEGORY_COMPLIANCE   NotificationCategory = 2
	NotificationCategory_NOTIFICATION_CATEGORY_RISK         NotificationCategory = 3
	NotificationCategory_NOTIFICATION_CATEGORY_LOGISTICS    NotificationCategory = 4
	NotificationCategory_NOTIFICATION_CATEGORY_DIGITAL_TWIN NotificationCategory = 5
	NotificationCategory_NOTIFICATION_CATEGORY_IDENTITY     NotificationCategory = 6
	NotificationCategory_NOTIFICATION_CATEGORY_SYSTEM       NotificationCategory = 7
	NotificationCategory_NOTIFICATION_CATEGORY_BILLING      NotificationCategory = 8
	NotificationCategory_NOTIFICATION_CATEGORY_SECURITY     NotificationCategory = 9
	NotificationCategory_NOTIFICATION_CATEGORY_CONSIGNMENT  NotificationCategory = 10
	NotificationCategory_NOTIFICATION_CATEGORY_CUSTODY      NotificationCategory = 11
	NotificationCategory_NOTIFICATION_CATEGORY_EVENT_LOG    NotificationCategory = 12
	NotificationCategory_NOTIFICATION_CATEGORY_INVENTORY    NotificationCategory = 13
	NotificationCategory_NOTIFICATION_CATEGORY_PARTNER      NotificationCategory = 14
	NotificationCategory_NOTIFICATION_CATEGORY_PROVENANCE   NotificationCategory = 15
	NotificationCategory_NOTIFICATION_CATEGORY_UNIT         NotificationCategory = 16
	NotificationCategory_NOTIFICATION_CATEGORY_USER         NotificationCategory = 17
	NotificationCategory_NOTIFICATION_CATEGORY_WAREHOUSE    NotificationCategory = 18
)

// Enum value maps for NotificationCategory.
var (
	NotificationCategory_name = map[int32]string{
		0:  "NOTIFICATION_CATEGORY_UNSPECIFIED",
		1:  "NOTIFICATION_CATEGORY_ONBOARDING",
		2:  "NOTIFICATION_CATEGORY_COMPLIANCE",
		3:  "NOTIFICATION_CATEGORY_RISK",
		4:  "NOTIFICATION_CATEGORY_LOGISTICS",
		5:  "NOTIFICATION_CATEGORY_DIGITAL_TWIN",
		6:  "NOTIFICATION_CATEGORY_IDENTITY",
		7:  "NOTIFICATION_CATEGORY_SYSTEM",
		8:  "NOTIFICATION_CATEGORY_BILLING",
		9:  "NOTIFICATION_CATEGORY_SECURITY",
		10: "NOTIFICATION_CATEGORY_CONSIGNMENT",
		11: "NOTIFICATION_CATEGORY_CUSTODY",
		12: "NOTIFICATION_CATEGORY_EVENT_LOG",
		13: "NOTIFICATION_CATEGORY_INVENTORY",
		14: "NOTIFICATION_CATEGORY_PARTNER",
		15: "NOTIFICATION_CATEGORY_PROVENANCE",
		16: "NOTIFICATION_CATEGORY_UNIT",
		17: "NOTIFICATION_CATEGORY_USER",
		18: "NOTIFICATION_CATEGORY_WAREHOUSE",
	}
	NotificationCategory_value = map[string]int32{
		"NOTIFICATION_CATEGORY_UNSPECIFIED":  0,
		"NOTIFICATION_CATEGORY_ONBOARDING":   1,
		"NOTIFICATION_CATEGORY_COMPLIANCE":   2,
		"NOTIFICATION_CATEGORY_RISK":         3,
		"NOTIFICATION_CATEGORY_LOGISTICS":    4,
		"NOTIFICATION_CATEGORY_DIGITAL_TWIN": 5,
		"NOTIFICATION_CATEGORY_IDENTITY":     6,
		"NOTIFICATION_CATEGORY_SYSTEM":       7,
		"NOTIFICATION_CATEGORY_BILLING":      8,
		"NOTIFICATION_CATEGORY_SECURITY":     9,
		"NOTIFICATION_CATEGORY_CONSIGNMENT":  10,
		"NOTIFICATION_CATEGORY_CUSTODY":      11,
		"NOTIFICATION_CATEGORY_EVENT_LOG":    12,
		"NOTIFICATION_CATEGORY_INVENTORY":    13,
		"NOTIFICATION_CATEGORY_PARTNER":      14,
		"NOTIFICATION_CATEGORY_PROVENANCE":   15,
		"NOTIFICATION_CATEGORY_UNIT":         16,
		"NOTIFICATION_CATEGORY_USER":         17,
		"NOTIFICATION_CATEGORY_WAREHOUSE":    18,
	}
)

func (x NotificationCategory) Enum() *NotificationCategory {
	p := new(NotificationCategory)
	*p = x
	return p
}

func (x NotificationCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scg_shared_v1_enums_proto_enumTypes[5].Descriptor()
}

func (NotificationCategory) Type() protoreflect.EnumType {
	return &file_proto_scg_shared_v1_enums_proto_enumTypes[5]
}

func (x NotificationCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationCategory.Descriptor instead.
func (NotificationCategory) EnumDescriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_enums_proto_rawDescGZIP(), []int{5}
}

var File_proto_scg_shared_v1_enums_proto protoreflect.FileDescriptor

const file_proto_scg_shared_v1_enums_proto_rawDesc = "" +
	"\n" +
	"\x1fproto/scg/shared/v1/enums.proto\x12#proto.scg.notification.shared.enums*\x9d\x02\n" +
	"\x0fDeliveryChannel\x12 \n" +
	"\x1cDELIVERY_CHANNEL_UNSPECIFIED\x10\x00\x12\x1a\n" +
	"\x16DELIVERY_CHANNEL_EMAIL\x10\x01\x12\x18\n" +
	"\x14DELIVERY_CHANNEL_SMS\x10\x02\x12\x19\n" +
	"\x15DELIVERY_CHANNEL_PUSH\x10\x03\x12\x1c\n" +
	"\x18DELIVERY_CHANNEL_WEBHOOK\x10\x04\x12\x1e\n" +
	"\x1aDELIVERY_CHANNEL_DASHBOARD\x10\x05\x12\x1b\n" +
	"\x17DELIVERY_CHANNEL_IN_APP\x10\x06\x12\x1f\n" +
	"\x1bDELIVERY_CHANNEL_ESCALATION\x10\a\x12\x1b\n" +
	"\x17DELIVERY_CHANNEL_DIGEST\x10\b*\x9c\x02\n" +
	"\x10NotificationType\x12!\n" +
	"\x1dNOTIFICATION_TYPE_UNSPECIFIED\x10\x00\x12\x1b\n" +
	"\x17NOTIFICATION_TYPE_ALERT\x10\x01\x12!\n" +
	"\x1dNOTIFICATION_TYPE_INFORMATION\x10\x02\x12\"\n" +
	"\x1eNOTIFICATION_TYPE_CONFIRMATION\x10\x03\x12\x1e\n" +
	"\x1aNOTIFICATION_TYPE_REMINDER\x10\x04\x12\x1c\n" +
	"\x18NOTIFICATION_TYPE_DIGEST\x10\x05\x12\x1c\n" +
	"\x18NOTIFICATION_TYPE_REPORT\x10\x06\x12%\n" +
	"!NOTIFICATION_TYPE_ACTION_REQUIRED\x10\a*\xc2\x01\n" +
	"\x14NotificationSeverity\x12%\n" +
	"!NOTIFICATION_SEVERITY_UNSPECIFIED\x10\x00\x12\x1d\n" +
	"\x19NOTIFICATION_SEVERITY_LOW\x10\x01\x12 \n" +
	"\x1cNOTIFICATION_SEVERITY_MEDIUM\x10\x02\x12\x1e\n" +
	"\x1aNOTIFICATION_SEVERITY_HIGH\x10\x03\x12\"\n" +
	"\x1eNOTIFICATION_SEVERITY_CRITICAL\x10\x04*\xea\x01\n" +
	"\x14NotificationAudience\x12%\n" +
	"!NOTIFICATION_AUDIENCE_UNSPECIFIED\x10\x00\x12\x1e\n" +
	"\x1aNOTIFICATION_AUDIENCE_USER\x10\x01\x12\x1f\n" +
	"\x1bNOTIFICATION_AUDIENCE_ADMIN\x10\x02\x12!\n" +
	"\x1dNOTIFICATION_AUDIENCE_PARTNER\x10\x03\x12 \n" +
	"\x1cNOTIFICATION_AUDIENCE_SYSTEM\x10\x04\x12%\n" +
	"!NOTIFICATION_AUDIENCE_INTEGRATION\x10\x05*\xfc\x01\n" +
	"\x12NotificationStatus\x12#\n" +
	"\x1fNOTIFICATION_STATUS_UNSPECIFIED\x10\x00\x12\x1f\n" +
	"\x1bNOTIFICATION_STATUS_PENDING\x10\x01\x12\x1c\n" +
	"\x18NOTIFICATION_STATUS_SENT\x10\x02\x12!\n" +
	"\x1dNOTIFICATION_STATUS_DELIVERED\x10\x03\x12\x1c\n" +
	"\x18NOTIFICATION_STATUS_READ\x10\x04\x12\x1e\n" +
	"\x1aNOTIFICATION_STATUS_FAILED\x10\x05\x12!\n" +
	"\x1dNOTIFICATION_STATUS_CANCELLED\x10\x06*\xc5\x05\n" +
	"\x14NotificationCategory\x12%\n" +
	"!NOTIFICATION_CATEGORY_UNSPECIFIED\x10\x00\x12$\n" +
	" NOTIFICATION_CATEGORY_ONBOARDING\x10\x01\x12$\n" +
	" NOTIFICATION_CATEGORY_COMPLIANCE\x10\x02\x12\x1e\n" +
	"\x1aNOTIFICATION_CATEGORY_RISK\x10\x03\x12#\n" +
	"\x1fNOTIFICATION_CATEGORY_LOGISTICS\x10\x04\x12&\n" +
	"\"NOTIFICATION_CATEGORY_DIGITAL_TWIN\x10\x05\x12\"\n" +
	"\x1eNOTIFICATION_CATEGORY_IDENTITY\x10\x06\x12 \n" +
	"\x1cNOTIFICATION_CATEGORY_SYSTEM\x10\a\x12!\n" +
	"\x1dNOTIFICATION_CATEGORY_BILLING\x10\b\x12\"\n" +
	"\x1eNOTIFICATION_CATEGORY_SECURITY\x10\t\x12%\n" +
	"!NOTIFICATION_CATEGORY_CONSIGNMENT\x10\n" +
	"\x12!\n" +
	"\x1dNOTIFICATION_CATEGORY_CUSTODY\x10\v\x12#\n" +
	"\x1fNOTIFICATION_CATEGORY_EVENT_LOG\x10\f\x12#\n" +
	"\x1fNOTIFICATION_CATEGORY_INVENTORY\x10\r\x12!\n" +
	"\x1dNOTIFICATION_CATEGORY_PARTNER\x10\x0e\x12$\n" +
	" NOTIFICATION_CATEGORY_PROVENANCE\x10\x0f\x12\x1e\n" +
	"\x1aNOTIFICATION_CATEGORY_UNIT\x10\x10\x12\x1e\n" +
	"\x1aNOTIFICATION_CATEGORY_USER\x10\x11\x12#\n" +
	"\x1fNOTIFICATION_CATEGORY_WAREHOUSE\x10\x12BOZMgithub.com/hbttundar/scg-notification-contracts/gen/go/scg/shared/v1;sharedv1b\x06proto3"

var (
	file_proto_scg_shared_v1_enums_proto_rawDescOnce sync.Once
	file_proto_scg_shared_v1_enums_proto_rawDescData []byte
)

func file_proto_scg_shared_v1_enums_proto_rawDescGZIP() []byte {
	file_proto_scg_shared_v1_enums_proto_rawDescOnce.Do(func() {
		file_proto_scg_shared_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_shared_v1_enums_proto_rawDesc), len(file_proto_scg_shared_v1_enums_proto_rawDesc)))
	})
	return file_proto_scg_shared_v1_enums_proto_rawDescData
}

var file_proto_scg_shared_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_proto_scg_shared_v1_enums_proto_goTypes = []any{
	(DeliveryChannel)(0),      // 0: proto.scg.notification.shared.enums.DeliveryChannel
	(NotificationType)(0),     // 1: proto.scg.notification.shared.enums.NotificationType
	(NotificationSeverity)(0), // 2: proto.scg.notification.shared.enums.NotificationSeverity
	(NotificationAudience)(0), // 3: proto.scg.notification.shared.enums.NotificationAudience
	(NotificationStatus)(0),   // 4: proto.scg.notification.shared.enums.NotificationStatus
	(NotificationCategory)(0), // 5: proto.scg.notification.shared.enums.NotificationCategory
}
var file_proto_scg_shared_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_scg_shared_v1_enums_proto_init() }
func file_proto_scg_shared_v1_enums_proto_init() {
	if File_proto_scg_shared_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_shared_v1_enums_proto_rawDesc), len(file_proto_scg_shared_v1_enums_proto_rawDesc)),
			NumEnums:      6,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_shared_v1_enums_proto_goTypes,
		DependencyIndexes: file_proto_scg_shared_v1_enums_proto_depIdxs,
		EnumInfos:         file_proto_scg_shared_v1_enums_proto_enumTypes,
	}.Build()
	File_proto_scg_shared_v1_enums_proto = out.File
	file_proto_scg_shared_v1_enums_proto_goTypes = nil
	file_proto_scg_shared_v1_enums_proto_depIdxs = nil
}
