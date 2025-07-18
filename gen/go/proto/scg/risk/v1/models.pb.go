// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/risk/v1/models.proto

package riskv1

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

// Risk represents a risk in the system
type Risk struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type           RiskType               `protobuf:"varint,4,opt,name=type,proto3,enum=proto.scg.notification.risk.enums.RiskType" json:"type,omitempty"`
	Level          RiskLevel              `protobuf:"varint,5,opt,name=level,proto3,enum=proto.scg.notification.risk.enums.RiskLevel" json:"level,omitempty"`
	Status         RiskStatus             `protobuf:"varint,6,opt,name=status,proto3,enum=proto.scg.notification.risk.enums.RiskStatus" json:"status,omitempty"`
	RiskScore      float64                `protobuf:"fixed64,7,opt,name=risk_score,json=riskScore,proto3" json:"risk_score,omitempty"`
	Threshold      float64                `protobuf:"fixed64,8,opt,name=threshold,proto3" json:"threshold,omitempty"`
	EntityId       string                 `protobuf:"bytes,9,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	EntityType     string                 `protobuf:"bytes,10,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	IdentifiedAt   *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=identified_at,json=identifiedAt,proto3" json:"identified_at,omitempty"`
	LastAssessedAt *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=last_assessed_at,json=lastAssessedAt,proto3" json:"last_assessed_at,omitempty"`
	MitigationPlan string                 `protobuf:"bytes,13,opt,name=mitigation_plan,json=mitigationPlan,proto3" json:"mitigation_plan,omitempty"`
	Metadata       map[string]string      `protobuf:"bytes,14,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Risk) Reset() {
	*x = Risk{}
	mi := &file_proto_scg_risk_v1_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Risk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Risk) ProtoMessage() {}

func (x *Risk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_risk_v1_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Risk.ProtoReflect.Descriptor instead.
func (*Risk) Descriptor() ([]byte, []int) {
	return file_proto_scg_risk_v1_models_proto_rawDescGZIP(), []int{0}
}

func (x *Risk) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Risk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Risk) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Risk) GetType() RiskType {
	if x != nil {
		return x.Type
	}
	return RiskType_RISK_TYPE_UNSPECIFIED
}

func (x *Risk) GetLevel() RiskLevel {
	if x != nil {
		return x.Level
	}
	return RiskLevel_RISK_LEVEL_UNSPECIFIED
}

func (x *Risk) GetStatus() RiskStatus {
	if x != nil {
		return x.Status
	}
	return RiskStatus_RISK_STATUS_UNSPECIFIED
}

func (x *Risk) GetRiskScore() float64 {
	if x != nil {
		return x.RiskScore
	}
	return 0
}

func (x *Risk) GetThreshold() float64 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *Risk) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *Risk) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *Risk) GetIdentifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IdentifiedAt
	}
	return nil
}

func (x *Risk) GetLastAssessedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastAssessedAt
	}
	return nil
}

func (x *Risk) GetMitigationPlan() string {
	if x != nil {
		return x.MitigationPlan
	}
	return ""
}

func (x *Risk) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// RiskAssessment represents an assessment of a risk
type RiskAssessment struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RiskId            string                 `protobuf:"bytes,2,opt,name=risk_id,json=riskId,proto3" json:"risk_id,omitempty"`
	AssessorId        string                 `protobuf:"bytes,3,opt,name=assessor_id,json=assessorId,proto3" json:"assessor_id,omitempty"`
	AssessorName      string                 `protobuf:"bytes,4,opt,name=assessor_name,json=assessorName,proto3" json:"assessor_name,omitempty"`
	PreviousRiskScore float64                `protobuf:"fixed64,5,opt,name=previous_risk_score,json=previousRiskScore,proto3" json:"previous_risk_score,omitempty"`
	NewRiskScore      float64                `protobuf:"fixed64,6,opt,name=new_risk_score,json=newRiskScore,proto3" json:"new_risk_score,omitempty"`
	AssessmentNotes   string                 `protobuf:"bytes,7,opt,name=assessment_notes,json=assessmentNotes,proto3" json:"assessment_notes,omitempty"`
	AssessedAt        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=assessed_at,json=assessedAt,proto3" json:"assessed_at,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *RiskAssessment) Reset() {
	*x = RiskAssessment{}
	mi := &file_proto_scg_risk_v1_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RiskAssessment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiskAssessment) ProtoMessage() {}

func (x *RiskAssessment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_risk_v1_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiskAssessment.ProtoReflect.Descriptor instead.
func (*RiskAssessment) Descriptor() ([]byte, []int) {
	return file_proto_scg_risk_v1_models_proto_rawDescGZIP(), []int{1}
}

func (x *RiskAssessment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RiskAssessment) GetRiskId() string {
	if x != nil {
		return x.RiskId
	}
	return ""
}

func (x *RiskAssessment) GetAssessorId() string {
	if x != nil {
		return x.AssessorId
	}
	return ""
}

func (x *RiskAssessment) GetAssessorName() string {
	if x != nil {
		return x.AssessorName
	}
	return ""
}

func (x *RiskAssessment) GetPreviousRiskScore() float64 {
	if x != nil {
		return x.PreviousRiskScore
	}
	return 0
}

func (x *RiskAssessment) GetNewRiskScore() float64 {
	if x != nil {
		return x.NewRiskScore
	}
	return 0
}

func (x *RiskAssessment) GetAssessmentNotes() string {
	if x != nil {
		return x.AssessmentNotes
	}
	return ""
}

func (x *RiskAssessment) GetAssessedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AssessedAt
	}
	return nil
}

var File_proto_scg_risk_v1_models_proto protoreflect.FileDescriptor

const file_proto_scg_risk_v1_models_proto_rawDesc = "" +
	"\n" +
	"\x1eproto/scg/risk/v1/models.proto\x12\"proto.scg.notification.risk.models\x1a\x1dproto/scg/risk/v1/enums.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd4\x05\n" +
	"\x04Risk\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12?\n" +
	"\x04type\x18\x04 \x01(\x0e2+.proto.scg.notification.risk.enums.RiskTypeR\x04type\x12B\n" +
	"\x05level\x18\x05 \x01(\x0e2,.proto.scg.notification.risk.enums.RiskLevelR\x05level\x12E\n" +
	"\x06status\x18\x06 \x01(\x0e2-.proto.scg.notification.risk.enums.RiskStatusR\x06status\x12\x1d\n" +
	"\n" +
	"risk_score\x18\a \x01(\x01R\triskScore\x12\x1c\n" +
	"\tthreshold\x18\b \x01(\x01R\tthreshold\x12\x1b\n" +
	"\tentity_id\x18\t \x01(\tR\bentityId\x12\x1f\n" +
	"\ventity_type\x18\n" +
	" \x01(\tR\n" +
	"entityType\x12?\n" +
	"\ridentified_at\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\fidentifiedAt\x12D\n" +
	"\x10last_assessed_at\x18\f \x01(\v2\x1a.google.protobuf.TimestampR\x0elastAssessedAt\x12'\n" +
	"\x0fmitigation_plan\x18\r \x01(\tR\x0emitigationPlan\x12R\n" +
	"\bmetadata\x18\x0e \x03(\v26.proto.scg.notification.risk.models.Risk.MetadataEntryR\bmetadata\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xbd\x02\n" +
	"\x0eRiskAssessment\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\arisk_id\x18\x02 \x01(\tR\x06riskId\x12\x1f\n" +
	"\vassessor_id\x18\x03 \x01(\tR\n" +
	"assessorId\x12#\n" +
	"\rassessor_name\x18\x04 \x01(\tR\fassessorName\x12.\n" +
	"\x13previous_risk_score\x18\x05 \x01(\x01R\x11previousRiskScore\x12$\n" +
	"\x0enew_risk_score\x18\x06 \x01(\x01R\fnewRiskScore\x12)\n" +
	"\x10assessment_notes\x18\a \x01(\tR\x0fassessmentNotes\x12;\n" +
	"\vassessed_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"assessedAtBKZIgithub.com/hbttundar/scg-notification-contracts/gen/go/scg/risk/v1;riskv1b\x06proto3"

var (
	file_proto_scg_risk_v1_models_proto_rawDescOnce sync.Once
	file_proto_scg_risk_v1_models_proto_rawDescData []byte
)

func file_proto_scg_risk_v1_models_proto_rawDescGZIP() []byte {
	file_proto_scg_risk_v1_models_proto_rawDescOnce.Do(func() {
		file_proto_scg_risk_v1_models_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_risk_v1_models_proto_rawDesc), len(file_proto_scg_risk_v1_models_proto_rawDesc)))
	})
	return file_proto_scg_risk_v1_models_proto_rawDescData
}

var file_proto_scg_risk_v1_models_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_scg_risk_v1_models_proto_goTypes = []any{
	(*Risk)(nil),                  // 0: proto.scg.notification.risk.models.Risk
	(*RiskAssessment)(nil),        // 1: proto.scg.notification.risk.models.RiskAssessment
	nil,                           // 2: proto.scg.notification.risk.models.Risk.MetadataEntry
	(RiskType)(0),                 // 3: proto.scg.notification.risk.enums.RiskType
	(RiskLevel)(0),                // 4: proto.scg.notification.risk.enums.RiskLevel
	(RiskStatus)(0),               // 5: proto.scg.notification.risk.enums.RiskStatus
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_proto_scg_risk_v1_models_proto_depIdxs = []int32{
	3, // 0: proto.scg.notification.risk.models.Risk.type:type_name -> proto.scg.notification.risk.enums.RiskType
	4, // 1: proto.scg.notification.risk.models.Risk.level:type_name -> proto.scg.notification.risk.enums.RiskLevel
	5, // 2: proto.scg.notification.risk.models.Risk.status:type_name -> proto.scg.notification.risk.enums.RiskStatus
	6, // 3: proto.scg.notification.risk.models.Risk.identified_at:type_name -> google.protobuf.Timestamp
	6, // 4: proto.scg.notification.risk.models.Risk.last_assessed_at:type_name -> google.protobuf.Timestamp
	2, // 5: proto.scg.notification.risk.models.Risk.metadata:type_name -> proto.scg.notification.risk.models.Risk.MetadataEntry
	6, // 6: proto.scg.notification.risk.models.RiskAssessment.assessed_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_scg_risk_v1_models_proto_init() }
func file_proto_scg_risk_v1_models_proto_init() {
	if File_proto_scg_risk_v1_models_proto != nil {
		return
	}
	file_proto_scg_risk_v1_enums_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_risk_v1_models_proto_rawDesc), len(file_proto_scg_risk_v1_models_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_risk_v1_models_proto_goTypes,
		DependencyIndexes: file_proto_scg_risk_v1_models_proto_depIdxs,
		MessageInfos:      file_proto_scg_risk_v1_models_proto_msgTypes,
	}.Build()
	File_proto_scg_risk_v1_models_proto = out.File
	file_proto_scg_risk_v1_models_proto_goTypes = nil
	file_proto_scg_risk_v1_models_proto_depIdxs = nil
}
