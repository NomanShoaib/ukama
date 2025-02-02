// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: package.proto

package gen

import (
	_ "github.com/mwitkow/go-proto-validators"
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

type GetPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetPackageRequest) Reset() {
	*x = GetPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageRequest) ProtoMessage() {}

func (x *GetPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageRequest.ProtoReflect.Descriptor instead.
func (*GetPackageRequest) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{0}
}

func (x *GetPackageRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetByOrgPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId string `protobuf:"bytes,1,opt,name=orgId,json=org_id,proto3" json:"orgId,omitempty"`
}

func (x *GetByOrgPackageRequest) Reset() {
	*x = GetByOrgPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByOrgPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByOrgPackageRequest) ProtoMessage() {}

func (x *GetByOrgPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByOrgPackageRequest.ProtoReflect.Descriptor instead.
func (*GetByOrgPackageRequest) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{1}
}

func (x *GetByOrgPackageRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

type GetByOrgPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packages []*Package `protobuf:"bytes,1,rep,name=packages,proto3" json:"packages,omitempty"`
}

func (x *GetByOrgPackageResponse) Reset() {
	*x = GetByOrgPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByOrgPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByOrgPackageResponse) ProtoMessage() {}

func (x *GetByOrgPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByOrgPackageResponse.ProtoReflect.Descriptor instead.
func (*GetByOrgPackageResponse) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{2}
}

func (x *GetByOrgPackageResponse) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

type GetPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *GetPackageResponse) Reset() {
	*x = GetPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageResponse) ProtoMessage() {}

func (x *GetPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageResponse.ProtoReflect.Descriptor instead.
func (*GetPackageResponse) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{3}
}

func (x *GetPackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type DeletePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DeletePackageRequest) Reset() {
	*x = DeletePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackageRequest) ProtoMessage() {}

func (x *DeletePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackageRequest.ProtoReflect.Descriptor instead.
func (*DeletePackageRequest) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePackageRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type DeletePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DeletePackageResponse) Reset() {
	*x = DeletePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackageResponse) ProtoMessage() {}

func (x *DeletePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackageResponse.ProtoReflect.Descriptor instead.
func (*DeletePackageResponse) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePackageResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type UpdatePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Active      bool   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Duration    uint64 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	SimType     string `protobuf:"bytes,5,opt,name=simType,json=sim_type,proto3" json:"simType,omitempty"`
	SmsVolume   int64  `protobuf:"varint,6,opt,name=smsVolume,json=sms_volume,proto3" json:"smsVolume,omitempty"`
	DataVolume  int64  `protobuf:"varint,7,opt,name=dataVolume,json=data_volume,proto3" json:"dataVolume,omitempty"`
	VoiceVolume int64  `protobuf:"varint,8,opt,name=voiceVolume,json=voice_volume,proto3" json:"voiceVolume,omitempty"`
	OrgRatesId  uint64 `protobuf:"varint,9,opt,name=orgRatesId,json=org_rates_id,proto3" json:"orgRatesId,omitempty"`
}

func (x *UpdatePackageRequest) Reset() {
	*x = UpdatePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePackageRequest) ProtoMessage() {}

func (x *UpdatePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePackageRequest.ProtoReflect.Descriptor instead.
func (*UpdatePackageRequest) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePackageRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UpdatePackageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePackageRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *UpdatePackageRequest) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *UpdatePackageRequest) GetSimType() string {
	if x != nil {
		return x.SimType
	}
	return ""
}

func (x *UpdatePackageRequest) GetSmsVolume() int64 {
	if x != nil {
		return x.SmsVolume
	}
	return 0
}

func (x *UpdatePackageRequest) GetDataVolume() int64 {
	if x != nil {
		return x.DataVolume
	}
	return 0
}

func (x *UpdatePackageRequest) GetVoiceVolume() int64 {
	if x != nil {
		return x.VoiceVolume
	}
	return 0
}

func (x *UpdatePackageRequest) GetOrgRatesId() uint64 {
	if x != nil {
		return x.OrgRatesId
	}
	return 0
}

type UpdatePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *UpdatePackageResponse) Reset() {
	*x = UpdatePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePackageResponse) ProtoMessage() {}

func (x *UpdatePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePackageResponse.ProtoReflect.Descriptor instead.
func (*UpdatePackageResponse) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type AddPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrgId       string `protobuf:"bytes,2,opt,name=orgId,json=org_id,proto3" json:"orgId,omitempty"`
	Active      bool   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Duration    uint64 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	SimType     string `protobuf:"bytes,5,opt,name=simType,json=sim_type,proto3" json:"simType,omitempty"`
	SmsVolume   int64  `protobuf:"varint,6,opt,name=smsVolume,json=sms_volume,proto3" json:"smsVolume,omitempty"`
	DataVolume  int64  `protobuf:"varint,7,opt,name=dataVolume,json=data_volume,proto3" json:"dataVolume,omitempty"`
	VoiceVolume int64  `protobuf:"varint,8,opt,name=voiceVolume,json=voice_volume,proto3" json:"voiceVolume,omitempty"`
	OrgRatesId  uint64 `protobuf:"varint,9,opt,name=orgRatesId,json=org_rates_id,proto3" json:"orgRatesId,omitempty"`
}

func (x *AddPackageRequest) Reset() {
	*x = AddPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPackageRequest) ProtoMessage() {}

func (x *AddPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPackageRequest.ProtoReflect.Descriptor instead.
func (*AddPackageRequest) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{8}
}

func (x *AddPackageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddPackageRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *AddPackageRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *AddPackageRequest) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *AddPackageRequest) GetSimType() string {
	if x != nil {
		return x.SimType
	}
	return ""
}

func (x *AddPackageRequest) GetSmsVolume() int64 {
	if x != nil {
		return x.SmsVolume
	}
	return 0
}

func (x *AddPackageRequest) GetDataVolume() int64 {
	if x != nil {
		return x.DataVolume
	}
	return 0
}

func (x *AddPackageRequest) GetVoiceVolume() int64 {
	if x != nil {
		return x.VoiceVolume
	}
	return 0
}

func (x *AddPackageRequest) GetOrgRatesId() uint64 {
	if x != nil {
		return x.OrgRatesId
	}
	return 0
}

type AddPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *AddPackageResponse) Reset() {
	*x = AddPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPackageResponse) ProtoMessage() {}

func (x *AddPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPackageResponse.ProtoReflect.Descriptor instead.
func (*AddPackageResponse) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{9}
}

func (x *AddPackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OrgId       string `protobuf:"bytes,3,opt,name=orgId,json=org_id,proto3" json:"orgId,omitempty"`
	Active      bool   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Duration    uint64 `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	SimType     string `protobuf:"bytes,6,opt,name=simType,json=sim_type,proto3" json:"simType,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=createdAt,json=created_at,proto3" json:"createdAt,omitempty"`
	DeletedAt   string `protobuf:"bytes,8,opt,name=deletedAt,json=deleted_at,proto3" json:"deletedAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,9,opt,name=updatedAt,json=updated_at,proto3" json:"updatedAt,omitempty"`
	SmsVolume   int64  `protobuf:"varint,10,opt,name=smsVolume,json=sms_volume,proto3" json:"smsVolume,omitempty"`
	DataVolume  int64  `protobuf:"varint,11,opt,name=dataVolume,json=data_volume,proto3" json:"dataVolume,omitempty"`
	VoiceVolume int64  `protobuf:"varint,12,opt,name=voiceVolume,json=voice_volume,proto3" json:"voiceVolume,omitempty"`
	OrgRatesId  uint64 `protobuf:"varint,13,opt,name=orgRatesId,json=org_rates_id,proto3" json:"orgRatesId,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{10}
}

func (x *Package) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Package) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Package) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Package) GetSimType() string {
	if x != nil {
		return x.SimType
	}
	return ""
}

func (x *Package) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Package) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Package) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Package) GetSmsVolume() int64 {
	if x != nil {
		return x.SmsVolume
	}
	return 0
}

func (x *Package) GetDataVolume() int64 {
	if x != nil {
		return x.DataVolume
	}
	return 0
}

func (x *Package) GetVoiceVolume() int64 {
	if x != nil {
		return x.VoiceVolume
	}
	return 0
}

func (x *Package) GetOrgRatesId() uint64 {
	if x != nil {
		return x.OrgRatesId
	}
	return 0
}

var File_package_proto protoreflect.FileDescriptor

var file_package_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x0f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x3a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58,
	0x01, 0x90, 0x01, 0x04, 0x52, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x6b, 0x61, 0x6d,
	0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f,
	0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x9b, 0x02, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x90, 0x01, 0x04, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x07, 0x73, 0x69,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x6d,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x6d, 0x73, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x6d, 0x73, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x6f, 0x72, 0x67, 0x52,
	0x61, 0x74, 0x65, 0x73, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6f, 0x72,
	0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x22, 0x9d, 0x02, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f,
	0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x07, 0x73, 0x69, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x09, 0x73, 0x6d, 0x73, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x73, 0x6d, 0x73, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0a, 0x6f, 0x72, 0x67, 0x52, 0x61, 0x74, 0x65, 0x73, 0x49, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x5f,
	0x69, 0x64, 0x22, 0x53, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x6b, 0x61, 0x6d,
	0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x8d, 0x03, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x90, 0x01, 0x04, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x07, 0x73, 0x69,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x6d,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x6d, 0x73, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x6d, 0x73, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x6f, 0x72, 0x67, 0x52, 0x61, 0x74, 0x65,
	0x73, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x32, 0xba, 0x04, 0x0a, 0x0f, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x2d, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x2d, 0x2e, 0x75, 0x6b, 0x61,
	0x6d, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x75, 0x6b, 0x61, 0x6d,
	0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x75, 0x6b, 0x61, 0x6d, 0x61,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x12, 0x32, 0x2e, 0x75, 0x6b, 0x61, 0x6d,
	0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e,
	0x75, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x4f, 0x72, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_proto_rawDescOnce sync.Once
	file_package_proto_rawDescData = file_package_proto_rawDesc
)

func file_package_proto_rawDescGZIP() []byte {
	file_package_proto_rawDescOnce.Do(func() {
		file_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_proto_rawDescData)
	})
	return file_package_proto_rawDescData
}

var file_package_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_package_proto_goTypes = []interface{}{
	(*GetPackageRequest)(nil),       // 0: ukama.data_plan.package.v1.GetPackageRequest
	(*GetByOrgPackageRequest)(nil),  // 1: ukama.data_plan.package.v1.GetByOrgPackageRequest
	(*GetByOrgPackageResponse)(nil), // 2: ukama.data_plan.package.v1.GetByOrgPackageResponse
	(*GetPackageResponse)(nil),      // 3: ukama.data_plan.package.v1.GetPackageResponse
	(*DeletePackageRequest)(nil),    // 4: ukama.data_plan.package.v1.DeletePackageRequest
	(*DeletePackageResponse)(nil),   // 5: ukama.data_plan.package.v1.DeletePackageResponse
	(*UpdatePackageRequest)(nil),    // 6: ukama.data_plan.package.v1.UpdatePackageRequest
	(*UpdatePackageResponse)(nil),   // 7: ukama.data_plan.package.v1.UpdatePackageResponse
	(*AddPackageRequest)(nil),       // 8: ukama.data_plan.package.v1.AddPackageRequest
	(*AddPackageResponse)(nil),      // 9: ukama.data_plan.package.v1.AddPackageResponse
	(*Package)(nil),                 // 10: ukama.data_plan.package.v1.Package
}
var file_package_proto_depIdxs = []int32{
	10, // 0: ukama.data_plan.package.v1.GetByOrgPackageResponse.packages:type_name -> ukama.data_plan.package.v1.Package
	10, // 1: ukama.data_plan.package.v1.GetPackageResponse.package:type_name -> ukama.data_plan.package.v1.Package
	10, // 2: ukama.data_plan.package.v1.UpdatePackageResponse.package:type_name -> ukama.data_plan.package.v1.Package
	10, // 3: ukama.data_plan.package.v1.AddPackageResponse.package:type_name -> ukama.data_plan.package.v1.Package
	0,  // 4: ukama.data_plan.package.v1.PackagesService.Get:input_type -> ukama.data_plan.package.v1.GetPackageRequest
	8,  // 5: ukama.data_plan.package.v1.PackagesService.Add:input_type -> ukama.data_plan.package.v1.AddPackageRequest
	4,  // 6: ukama.data_plan.package.v1.PackagesService.Delete:input_type -> ukama.data_plan.package.v1.DeletePackageRequest
	6,  // 7: ukama.data_plan.package.v1.PackagesService.Update:input_type -> ukama.data_plan.package.v1.UpdatePackageRequest
	1,  // 8: ukama.data_plan.package.v1.PackagesService.GetByOrg:input_type -> ukama.data_plan.package.v1.GetByOrgPackageRequest
	3,  // 9: ukama.data_plan.package.v1.PackagesService.Get:output_type -> ukama.data_plan.package.v1.GetPackageResponse
	9,  // 10: ukama.data_plan.package.v1.PackagesService.Add:output_type -> ukama.data_plan.package.v1.AddPackageResponse
	5,  // 11: ukama.data_plan.package.v1.PackagesService.Delete:output_type -> ukama.data_plan.package.v1.DeletePackageResponse
	7,  // 12: ukama.data_plan.package.v1.PackagesService.Update:output_type -> ukama.data_plan.package.v1.UpdatePackageResponse
	2,  // 13: ukama.data_plan.package.v1.PackagesService.GetByOrg:output_type -> ukama.data_plan.package.v1.GetByOrgPackageResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_package_proto_init() }
func file_package_proto_init() {
	if File_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageRequest); i {
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
		file_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByOrgPackageRequest); i {
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
		file_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByOrgPackageResponse); i {
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
		file_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageResponse); i {
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
		file_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackageRequest); i {
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
		file_package_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackageResponse); i {
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
		file_package_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePackageRequest); i {
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
		file_package_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePackageResponse); i {
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
		file_package_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPackageRequest); i {
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
		file_package_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPackageResponse); i {
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
		file_package_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
			RawDescriptor: file_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_package_proto_goTypes,
		DependencyIndexes: file_package_proto_depIdxs,
		MessageInfos:      file_package_proto_msgTypes,
	}.Build()
	File_package_proto = out.File
	file_package_proto_rawDesc = nil
	file_package_proto_goTypes = nil
	file_package_proto_depIdxs = nil
}
