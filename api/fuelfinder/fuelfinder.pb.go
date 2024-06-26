// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: api/fuelfinder/fuelfinder.proto

package fuelfinder

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_fuelfinder_fuelfinder_proto_rawDescGZIP(), []int{0}
}

type UploadedItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UploadedItems) Reset() {
	*x = UploadedItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadedItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadedItems) ProtoMessage() {}

func (x *UploadedItems) ProtoReflect() protoreflect.Message {
	mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadedItems.ProtoReflect.Descriptor instead.
func (*UploadedItems) Descriptor() ([]byte, []int) {
	return file_api_fuelfinder_fuelfinder_proto_rawDescGZIP(), []int{1}
}

func (x *UploadedItems) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Geofence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Longitude float32  `protobuf:"fixed32,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float32  `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Radius    float32  `protobuf:"fixed32,3,opt,name=radius,proto3" json:"radius,omitempty"`
	Brands    []string `protobuf:"bytes,4,rep,name=brands,proto3" json:"brands,omitempty"`
	Fueltype  []string `protobuf:"bytes,5,rep,name=fueltype,proto3" json:"fueltype,omitempty"`
}

func (x *Geofence) Reset() {
	*x = Geofence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Geofence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geofence) ProtoMessage() {}

func (x *Geofence) ProtoReflect() protoreflect.Message {
	mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geofence.ProtoReflect.Descriptor instead.
func (*Geofence) Descriptor() ([]byte, []int) {
	return file_api_fuelfinder_fuelfinder_proto_rawDescGZIP(), []int{2}
}

func (x *Geofence) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Geofence) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Geofence) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *Geofence) GetBrands() []string {
	if x != nil {
		return x.Brands
	}
	return nil
}

func (x *Geofence) GetFueltype() []string {
	if x != nil {
		return x.Fueltype
	}
	return nil
}

type StationItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt int64     `protobuf:"varint,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	SiteId    string    `protobuf:"bytes,2,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	Brand     string    `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	Address   string    `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Postcode  string    `protobuf:"bytes,5,opt,name=postcode,proto3" json:"postcode,omitempty"`
	Location  *Location `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	E5        float32   `protobuf:"fixed32,7,opt,name=e5,proto3" json:"e5,omitempty"`
	E10       float32   `protobuf:"fixed32,8,opt,name=e10,proto3" json:"e10,omitempty"`
	B7        float32   `protobuf:"fixed32,9,opt,name=b7,proto3" json:"b7,omitempty"`
	Sdv       float32   `protobuf:"fixed32,10,opt,name=sdv,proto3" json:"sdv,omitempty"`
	Distance  float64   `protobuf:"fixed64,11,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *StationItem) Reset() {
	*x = StationItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationItem) ProtoMessage() {}

func (x *StationItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationItem.ProtoReflect.Descriptor instead.
func (*StationItem) Descriptor() ([]byte, []int) {
	return file_api_fuelfinder_fuelfinder_proto_rawDescGZIP(), []int{3}
}

func (x *StationItem) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *StationItem) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

func (x *StationItem) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *StationItem) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StationItem) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

func (x *StationItem) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *StationItem) GetE5() float32 {
	if x != nil {
		return x.E5
	}
	return 0
}

func (x *StationItem) GetE10() float32 {
	if x != nil {
		return x.E10
	}
	return 0
}

func (x *StationItem) GetB7() float32 {
	if x != nil {
		return x.B7
	}
	return 0
}

func (x *StationItem) GetSdv() float32 {
	if x != nil {
		return x.Sdv
	}
	return 0
}

func (x *StationItem) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,2,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_api_fuelfinder_fuelfinder_proto_rawDescGZIP(), []int{4}
}

func (x *Location) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Location) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type StationItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*StationItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *StationItems) Reset() {
	*x = StationItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationItems) ProtoMessage() {}

func (x *StationItems) ProtoReflect() protoreflect.Message {
	mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationItems.ProtoReflect.Descriptor instead.
func (*StationItems) Descriptor() ([]byte, []int) {
	return file_api_fuelfinder_fuelfinder_proto_rawDescGZIP(), []int{5}
}

func (x *StationItems) GetItems() []*StationItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type Brands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brands []string `protobuf:"bytes,1,rep,name=brands,proto3" json:"brands,omitempty"`
}

func (x *Brands) Reset() {
	*x = Brands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brands) ProtoMessage() {}

func (x *Brands) ProtoReflect() protoreflect.Message {
	mi := &file_api_fuelfinder_fuelfinder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brands.ProtoReflect.Descriptor instead.
func (*Brands) Descriptor() ([]byte, []int) {
	return file_api_fuelfinder_fuelfinder_proto_rawDescGZIP(), []int{6}
}

func (x *Brands) GetBrands() []string {
	if x != nil {
		return x.Brands
	}
	return nil
}

var File_api_fuelfinder_fuelfinder_proto protoreflect.FileDescriptor

var file_api_fuelfinder_fuelfinder_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x2f, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x25, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x90, 0x01,
	0x0a, 0x08, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x74, 0x79, 0x70, 0x65,
	0x22, 0xa3, 0x02, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x35, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x02, 0x65, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x31, 0x30, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x65, 0x31, 0x30, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x37, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x62, 0x37, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x64, 0x76, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x73, 0x64, 0x76, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x40, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x32, 0xc1, 0x01, 0x0a, 0x0a, 0x46, 0x75,
	0x65, 0x6c, 0x46, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x72, 0x65, 0x61, 0x12, 0x14, 0x2e, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x18, 0x2e, 0x66, 0x75,
	0x65, 0x6c, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x18, 0x2e, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x19, 0x2e, 0x66, 0x75, 0x65, 0x6c,
	0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x11, 0x2e, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x66, 0x75, 0x65, 0x6c,
	0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x42, 0x15, 0x5a,
	0x13, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x65, 0x6c, 0x66, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_fuelfinder_fuelfinder_proto_rawDescOnce sync.Once
	file_api_fuelfinder_fuelfinder_proto_rawDescData = file_api_fuelfinder_fuelfinder_proto_rawDesc
)

func file_api_fuelfinder_fuelfinder_proto_rawDescGZIP() []byte {
	file_api_fuelfinder_fuelfinder_proto_rawDescOnce.Do(func() {
		file_api_fuelfinder_fuelfinder_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_fuelfinder_fuelfinder_proto_rawDescData)
	})
	return file_api_fuelfinder_fuelfinder_proto_rawDescData
}

var file_api_fuelfinder_fuelfinder_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_fuelfinder_fuelfinder_proto_goTypes = []interface{}{
	(*Empty)(nil),         // 0: fuelfinder.Empty
	(*UploadedItems)(nil), // 1: fuelfinder.UploadedItems
	(*Geofence)(nil),      // 2: fuelfinder.Geofence
	(*StationItem)(nil),   // 3: fuelfinder.StationItem
	(*Location)(nil),      // 4: fuelfinder.Location
	(*StationItems)(nil),  // 5: fuelfinder.StationItems
	(*Brands)(nil),        // 6: fuelfinder.Brands
}
var file_api_fuelfinder_fuelfinder_proto_depIdxs = []int32{
	4, // 0: fuelfinder.StationItem.location:type_name -> fuelfinder.Location
	3, // 1: fuelfinder.StationItems.items:type_name -> fuelfinder.StationItem
	2, // 2: fuelfinder.FuelFinder.QueryArea:input_type -> fuelfinder.Geofence
	5, // 3: fuelfinder.FuelFinder.Upload:input_type -> fuelfinder.StationItems
	0, // 4: fuelfinder.FuelFinder.DistinctBrands:input_type -> fuelfinder.Empty
	5, // 5: fuelfinder.FuelFinder.QueryArea:output_type -> fuelfinder.StationItems
	1, // 6: fuelfinder.FuelFinder.Upload:output_type -> fuelfinder.UploadedItems
	6, // 7: fuelfinder.FuelFinder.DistinctBrands:output_type -> fuelfinder.Brands
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_fuelfinder_fuelfinder_proto_init() }
func file_api_fuelfinder_fuelfinder_proto_init() {
	if File_api_fuelfinder_fuelfinder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_fuelfinder_fuelfinder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_fuelfinder_fuelfinder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadedItems); i {
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
		file_api_fuelfinder_fuelfinder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Geofence); i {
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
		file_api_fuelfinder_fuelfinder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationItem); i {
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
		file_api_fuelfinder_fuelfinder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_api_fuelfinder_fuelfinder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationItems); i {
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
		file_api_fuelfinder_fuelfinder_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brands); i {
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
			RawDescriptor: file_api_fuelfinder_fuelfinder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_fuelfinder_fuelfinder_proto_goTypes,
		DependencyIndexes: file_api_fuelfinder_fuelfinder_proto_depIdxs,
		MessageInfos:      file_api_fuelfinder_fuelfinder_proto_msgTypes,
	}.Build()
	File_api_fuelfinder_fuelfinder_proto = out.File
	file_api_fuelfinder_fuelfinder_proto_rawDesc = nil
	file_api_fuelfinder_fuelfinder_proto_goTypes = nil
	file_api_fuelfinder_fuelfinder_proto_depIdxs = nil
}
