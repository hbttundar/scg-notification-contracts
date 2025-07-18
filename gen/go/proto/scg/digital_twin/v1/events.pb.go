// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/digital_twin/v1/events.proto

package digitaltwinv1

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

// DigitalTwinAlertNotification is sent when a digital twin generates an alert
type DigitalTwinAlertNotification struct {
	state             protoimpl.MessageState  `protogen:"open.v1"`
	Base              *v1.NotificationEvent   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DigitalTwinId     string                  `protobuf:"bytes,2,opt,name=digital_twin_id,json=digitalTwinId,proto3" json:"digital_twin_id,omitempty"`
	PhysicalAssetId   string                  `protobuf:"bytes,3,opt,name=physical_asset_id,json=physicalAssetId,proto3" json:"physical_asset_id,omitempty"`
	PhysicalAssetType string                  `protobuf:"bytes,4,opt,name=physical_asset_type,json=physicalAssetType,proto3" json:"physical_asset_type,omitempty"`
	AlertType         string                  `protobuf:"bytes,5,opt,name=alert_type,json=alertType,proto3" json:"alert_type,omitempty"`
	AlertDescription  string                  `protobuf:"bytes,6,opt,name=alert_description,json=alertDescription,proto3" json:"alert_description,omitempty"`
	SensorData        map[string]string       `protobuf:"bytes,7,rep,name=sensor_data,json=sensorData,proto3" json:"sensor_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Severity          v1.NotificationSeverity `protobuf:"varint,8,opt,name=severity,proto3,enum=proto.scg.notification.shared.enums.NotificationSeverity" json:"severity,omitempty"`
	DetectedAt        *timestamppb.Timestamp  `protobuf:"bytes,9,opt,name=detected_at,json=detectedAt,proto3" json:"detected_at,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DigitalTwinAlertNotification) Reset() {
	*x = DigitalTwinAlertNotification{}
	mi := &file_proto_scg_digital_twin_v1_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DigitalTwinAlertNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigitalTwinAlertNotification) ProtoMessage() {}

func (x *DigitalTwinAlertNotification) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_digital_twin_v1_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigitalTwinAlertNotification.ProtoReflect.Descriptor instead.
func (*DigitalTwinAlertNotification) Descriptor() ([]byte, []int) {
	return file_proto_scg_digital_twin_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *DigitalTwinAlertNotification) GetBase() *v1.NotificationEvent {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *DigitalTwinAlertNotification) GetDigitalTwinId() string {
	if x != nil {
		return x.DigitalTwinId
	}
	return ""
}

func (x *DigitalTwinAlertNotification) GetPhysicalAssetId() string {
	if x != nil {
		return x.PhysicalAssetId
	}
	return ""
}

func (x *DigitalTwinAlertNotification) GetPhysicalAssetType() string {
	if x != nil {
		return x.PhysicalAssetType
	}
	return ""
}

func (x *DigitalTwinAlertNotification) GetAlertType() string {
	if x != nil {
		return x.AlertType
	}
	return ""
}

func (x *DigitalTwinAlertNotification) GetAlertDescription() string {
	if x != nil {
		return x.AlertDescription
	}
	return ""
}

func (x *DigitalTwinAlertNotification) GetSensorData() map[string]string {
	if x != nil {
		return x.SensorData
	}
	return nil
}

func (x *DigitalTwinAlertNotification) GetSeverity() v1.NotificationSeverity {
	if x != nil {
		return x.Severity
	}
	return v1.NotificationSeverity(0)
}

func (x *DigitalTwinAlertNotification) GetDetectedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DetectedAt
	}
	return nil
}

var File_proto_scg_digital_twin_v1_events_proto protoreflect.FileDescriptor

const file_proto_scg_digital_twin_v1_events_proto_rawDesc = "" +
	"\n" +
	"&proto/scg/digital_twin/v1/events.proto\x12*proto.scg.notification.digital_twin.events\x1a proto/scg/shared/v1/events.proto\x1a\x1fproto/scg/shared/v1/enums.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x89\x05\n" +
	"\x1cDigitalTwinAlertNotification\x12K\n" +
	"\x04base\x18\x01 \x01(\v27.proto.scg.notification.shared.events.NotificationEventR\x04base\x12&\n" +
	"\x0fdigital_twin_id\x18\x02 \x01(\tR\rdigitalTwinId\x12*\n" +
	"\x11physical_asset_id\x18\x03 \x01(\tR\x0fphysicalAssetId\x12.\n" +
	"\x13physical_asset_type\x18\x04 \x01(\tR\x11physicalAssetType\x12\x1d\n" +
	"\n" +
	"alert_type\x18\x05 \x01(\tR\talertType\x12+\n" +
	"\x11alert_description\x18\x06 \x01(\tR\x10alertDescription\x12y\n" +
	"\vsensor_data\x18\a \x03(\v2X.proto.scg.notification.digital_twin.events.DigitalTwinAlertNotification.SensorDataEntryR\n" +
	"sensorData\x12U\n" +
	"\bseverity\x18\b \x01(\x0e29.proto.scg.notification.shared.enums.NotificationSeverityR\bseverity\x12;\n" +
	"\vdetected_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"detectedAt\x1a=\n" +
	"\x0fSensorDataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01BZZXgithub.com/hbttundar/scg-notification-contracts/gen/go/scg/digital_twin/v1;digitaltwinv1b\x06proto3"

var (
	file_proto_scg_digital_twin_v1_events_proto_rawDescOnce sync.Once
	file_proto_scg_digital_twin_v1_events_proto_rawDescData []byte
)

func file_proto_scg_digital_twin_v1_events_proto_rawDescGZIP() []byte {
	file_proto_scg_digital_twin_v1_events_proto_rawDescOnce.Do(func() {
		file_proto_scg_digital_twin_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_digital_twin_v1_events_proto_rawDesc), len(file_proto_scg_digital_twin_v1_events_proto_rawDesc)))
	})
	return file_proto_scg_digital_twin_v1_events_proto_rawDescData
}

var file_proto_scg_digital_twin_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_scg_digital_twin_v1_events_proto_goTypes = []any{
	(*DigitalTwinAlertNotification)(nil), // 0: proto.scg.notification.digital_twin.events.DigitalTwinAlertNotification
	nil,                                  // 1: proto.scg.notification.digital_twin.events.DigitalTwinAlertNotification.SensorDataEntry
	(*v1.NotificationEvent)(nil),         // 2: proto.scg.notification.shared.events.NotificationEvent
	(v1.NotificationSeverity)(0),         // 3: proto.scg.notification.shared.enums.NotificationSeverity
	(*timestamppb.Timestamp)(nil),        // 4: google.protobuf.Timestamp
}
var file_proto_scg_digital_twin_v1_events_proto_depIdxs = []int32{
	2, // 0: proto.scg.notification.digital_twin.events.DigitalTwinAlertNotification.base:type_name -> proto.scg.notification.shared.events.NotificationEvent
	1, // 1: proto.scg.notification.digital_twin.events.DigitalTwinAlertNotification.sensor_data:type_name -> proto.scg.notification.digital_twin.events.DigitalTwinAlertNotification.SensorDataEntry
	3, // 2: proto.scg.notification.digital_twin.events.DigitalTwinAlertNotification.severity:type_name -> proto.scg.notification.shared.enums.NotificationSeverity
	4, // 3: proto.scg.notification.digital_twin.events.DigitalTwinAlertNotification.detected_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_scg_digital_twin_v1_events_proto_init() }
func file_proto_scg_digital_twin_v1_events_proto_init() {
	if File_proto_scg_digital_twin_v1_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_digital_twin_v1_events_proto_rawDesc), len(file_proto_scg_digital_twin_v1_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_digital_twin_v1_events_proto_goTypes,
		DependencyIndexes: file_proto_scg_digital_twin_v1_events_proto_depIdxs,
		MessageInfos:      file_proto_scg_digital_twin_v1_events_proto_msgTypes,
	}.Build()
	File_proto_scg_digital_twin_v1_events_proto = out.File
	file_proto_scg_digital_twin_v1_events_proto_goTypes = nil
	file_proto_scg_digital_twin_v1_events_proto_depIdxs = nil
}
