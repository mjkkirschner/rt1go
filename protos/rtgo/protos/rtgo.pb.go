// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: protos/rtgo.proto

package rtgo

import (
	proto "github.com/golang/protobuf/proto"
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

type Material_MaterialType int32

const (
	Material_DiffuseMaterial    Material_MaterialType = 0
	Material_MetailMaterial     Material_MaterialType = 1
	Material_RefractiveMaterial Material_MaterialType = 2
	Material_LightMaterial      Material_MaterialType = 3
)

// Enum value maps for Material_MaterialType.
var (
	Material_MaterialType_name = map[int32]string{
		0: "DiffuseMaterial",
		1: "MetailMaterial",
		2: "RefractiveMaterial",
		3: "LightMaterial",
	}
	Material_MaterialType_value = map[string]int32{
		"DiffuseMaterial":    0,
		"MetailMaterial":     1,
		"RefractiveMaterial": 2,
		"LightMaterial":      3,
	}
)

func (x Material_MaterialType) Enum() *Material_MaterialType {
	p := new(Material_MaterialType)
	*p = x
	return p
}

func (x Material_MaterialType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Material_MaterialType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_rtgo_proto_enumTypes[0].Descriptor()
}

func (Material_MaterialType) Type() protoreflect.EnumType {
	return &file_protos_rtgo_proto_enumTypes[0]
}

func (x Material_MaterialType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Material_MaterialType.Descriptor instead.
func (Material_MaterialType) EnumDescriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{3, 0}
}

type RenderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Colors []*Vec3 `protobuf:"bytes,1,rep,name=colors,proto3" json:"colors,omitempty"`
}

func (x *RenderReply) Reset() {
	*x = RenderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rtgo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderReply) ProtoMessage() {}

func (x *RenderReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rtgo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderReply.ProtoReflect.Descriptor instead.
func (*RenderReply) Descriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{0}
}

func (x *RenderReply) GetColors() []*Vec3 {
	if x != nil {
		return x.Colors
	}
	return nil
}

type RenderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mesh     []*Mesh     `protobuf:"bytes,1,rep,name=mesh,proto3" json:"mesh,omitempty"`
	Material []*Material `protobuf:"bytes,2,rep,name=material,proto3" json:"material,omitempty"`
	Camera   *Camera     `protobuf:"bytes,3,opt,name=camera,proto3" json:"camera,omitempty"`
}

func (x *RenderRequest) Reset() {
	*x = RenderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rtgo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderRequest) ProtoMessage() {}

func (x *RenderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rtgo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderRequest.ProtoReflect.Descriptor instead.
func (*RenderRequest) Descriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{1}
}

func (x *RenderRequest) GetMesh() []*Mesh {
	if x != nil {
		return x.Mesh
	}
	return nil
}

func (x *RenderRequest) GetMaterial() []*Material {
	if x != nil {
		return x.Material
	}
	return nil
}

func (x *RenderRequest) GetCamera() *Camera {
	if x != nil {
		return x.Camera
	}
	return nil
}

type Camera struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lookfrom    *Vec3   `protobuf:"bytes,1,opt,name=lookfrom,proto3" json:"lookfrom,omitempty"`
	LookAt      *Vec3   `protobuf:"bytes,2,opt,name=lookAt,proto3" json:"lookAt,omitempty"`
	Vup         *Vec3   `protobuf:"bytes,3,opt,name=vup,proto3" json:"vup,omitempty"`
	FovDegrees  float64 `protobuf:"fixed64,4,opt,name=fovDegrees,proto3" json:"fovDegrees,omitempty"`
	AspectRatio float64 `protobuf:"fixed64,5,opt,name=aspectRatio,proto3" json:"aspectRatio,omitempty"`
	Aperture    float64 `protobuf:"fixed64,6,opt,name=aperture,proto3" json:"aperture,omitempty"`
	Focuslength float64 `protobuf:"fixed64,7,opt,name=focuslength,proto3" json:"focuslength,omitempty"`
}

func (x *Camera) Reset() {
	*x = Camera{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rtgo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Camera) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Camera) ProtoMessage() {}

func (x *Camera) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rtgo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Camera.ProtoReflect.Descriptor instead.
func (*Camera) Descriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{2}
}

func (x *Camera) GetLookfrom() *Vec3 {
	if x != nil {
		return x.Lookfrom
	}
	return nil
}

func (x *Camera) GetLookAt() *Vec3 {
	if x != nil {
		return x.LookAt
	}
	return nil
}

func (x *Camera) GetVup() *Vec3 {
	if x != nil {
		return x.Vup
	}
	return nil
}

func (x *Camera) GetFovDegrees() float64 {
	if x != nil {
		return x.FovDegrees
	}
	return 0
}

func (x *Camera) GetAspectRatio() float64 {
	if x != nil {
		return x.AspectRatio
	}
	return 0
}

func (x *Camera) GetAperture() float64 {
	if x != nil {
		return x.Aperture
	}
	return 0
}

func (x *Camera) GetFocuslength() float64 {
	if x != nil {
		return x.Focuslength
	}
	return 0
}

type Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Texture        *Texture `protobuf:"bytes,1,opt,name=texture,proto3" json:"texture,omitempty"`
	Albedo         *Vec3    `protobuf:"bytes,2,opt,name=albedo,proto3" json:"albedo,omitempty"`
	Lightintensity float64  `protobuf:"fixed64,3,opt,name=lightintensity,proto3" json:"lightintensity,omitempty"`
	Roughness      float64  `protobuf:"fixed64,4,opt,name=roughness,proto3" json:"roughness,omitempty"`
}

func (x *Material) Reset() {
	*x = Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rtgo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Material) ProtoMessage() {}

func (x *Material) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rtgo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Material.ProtoReflect.Descriptor instead.
func (*Material) Descriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{3}
}

func (x *Material) GetTexture() *Texture {
	if x != nil {
		return x.Texture
	}
	return nil
}

func (x *Material) GetAlbedo() *Vec3 {
	if x != nil {
		return x.Albedo
	}
	return nil
}

func (x *Material) GetLightintensity() float64 {
	if x != nil {
		return x.Lightintensity
	}
	return 0
}

func (x *Material) GetRoughness() float64 {
	if x != nil {
		return x.Roughness
	}
	return 0
}

type Texture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Colors []*Vec3 `protobuf:"bytes,1,rep,name=colors,proto3" json:"colors,omitempty"`
}

func (x *Texture) Reset() {
	*x = Texture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rtgo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Texture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Texture) ProtoMessage() {}

func (x *Texture) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rtgo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Texture.ProtoReflect.Descriptor instead.
func (*Texture) Descriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{4}
}

func (x *Texture) GetColors() []*Vec3 {
	if x != nil {
		return x.Colors
	}
	return nil
}

type Mesh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Facedata  []*FaceData `protobuf:"bytes,1,rep,name=facedata,proto3" json:"facedata,omitempty"`
	Verts     []*Vec3     `protobuf:"bytes,2,rep,name=verts,proto3" json:"verts,omitempty"`
	Texcoords []*Vec3     `protobuf:"bytes,3,rep,name=texcoords,proto3" json:"texcoords,omitempty"`
	Normals   []*Vec3     `protobuf:"bytes,4,rep,name=normals,proto3" json:"normals,omitempty"`
}

func (x *Mesh) Reset() {
	*x = Mesh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rtgo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mesh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mesh) ProtoMessage() {}

func (x *Mesh) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rtgo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mesh.ProtoReflect.Descriptor instead.
func (*Mesh) Descriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{5}
}

func (x *Mesh) GetFacedata() []*FaceData {
	if x != nil {
		return x.Facedata
	}
	return nil
}

func (x *Mesh) GetVerts() []*Vec3 {
	if x != nil {
		return x.Verts
	}
	return nil
}

func (x *Mesh) GetTexcoords() []*Vec3 {
	if x != nil {
		return x.Texcoords
	}
	return nil
}

func (x *Mesh) GetNormals() []*Vec3 {
	if x != nil {
		return x.Normals
	}
	return nil
}

type Vec3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vec3) Reset() {
	*x = Vec3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rtgo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vec3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vec3) ProtoMessage() {}

func (x *Vec3) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rtgo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vec3.ProtoReflect.Descriptor instead.
func (*Vec3) Descriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{6}
}

func (x *Vec3) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vec3) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vec3) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type FaceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vertindicies     []int32 `protobuf:"varint,1,rep,packed,name=vertindicies,proto3" json:"vertindicies,omitempty"`
	Normalindicies   []int32 `protobuf:"varint,2,rep,packed,name=normalindicies,proto3" json:"normalindicies,omitempty"`
	Texcoordindicies []int32 `protobuf:"varint,3,rep,packed,name=texcoordindicies,proto3" json:"texcoordindicies,omitempty"`
}

func (x *FaceData) Reset() {
	*x = FaceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_rtgo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceData) ProtoMessage() {}

func (x *FaceData) ProtoReflect() protoreflect.Message {
	mi := &file_protos_rtgo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceData.ProtoReflect.Descriptor instead.
func (*FaceData) Descriptor() ([]byte, []int) {
	return file_protos_rtgo_proto_rawDescGZIP(), []int{7}
}

func (x *FaceData) GetVertindicies() []int32 {
	if x != nil {
		return x.Vertindicies
	}
	return nil
}

func (x *FaceData) GetNormalindicies() []int32 {
	if x != nil {
		return x.Normalindicies
	}
	return nil
}

func (x *FaceData) GetTexcoordindicies() []int32 {
	if x != nil {
		return x.Texcoordindicies
	}
	return nil
}

var File_protos_rtgo_proto protoreflect.FileDescriptor

var file_protos_rtgo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x74, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x74, 0x67, 0x6f, 0x22, 0x31, 0x0a, 0x0b, 0x52, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e,
	0x56, 0x65, 0x63, 0x33, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x22, 0x81, 0x01, 0x0a,
	0x0d, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72,
	0x74, 0x67, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x12, 0x2a,
	0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x74, 0x67,
	0x6f, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x06, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x22, 0xf2, 0x01, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x26, 0x0a, 0x08, 0x6c,
	0x6f, 0x6f, 0x6b, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x72, 0x74, 0x67, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x08, 0x6c, 0x6f, 0x6f, 0x6b, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x06, 0x6c, 0x6f, 0x6f, 0x6b, 0x41, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52,
	0x06, 0x6c, 0x6f, 0x6f, 0x6b, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x03, 0x76, 0x75, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x33,
	0x52, 0x03, 0x76, 0x75, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x76, 0x44, 0x65, 0x67, 0x72,
	0x65, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x66, 0x6f, 0x76, 0x44, 0x65,
	0x67, 0x72, 0x65, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x61, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x81, 0x02, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x07, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x61,
	0x6c, 0x62, 0x65, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x74,
	0x67, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x65, 0x64, 0x6f, 0x12,
	0x26, 0x0a, 0x0e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x67, 0x68,
	0x6e, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x67,
	0x68, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x62, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x69, 0x66, 0x66, 0x75, 0x73, 0x65,
	0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x10, 0x03, 0x22, 0x2d, 0x0a, 0x07, 0x54, 0x65, 0x78,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x33,
	0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x73,
	0x68, 0x12, 0x2a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a,
	0x05, 0x76, 0x65, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72,
	0x74, 0x67, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x05, 0x76, 0x65, 0x72, 0x74, 0x73, 0x12,
	0x28, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x09,
	0x74, 0x65, 0x78, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x6e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x74, 0x67,
	0x6f, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x07, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x73, 0x22,
	0x30, 0x0a, 0x04, 0x56, 0x65, 0x63, 0x33, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x7a, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22,
	0x0a, 0x0c, 0x76, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x65,
	0x78, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x10, 0x74, 0x65, 0x78, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x64, 0x69, 0x63, 0x69, 0x65, 0x73, 0x32, 0x38, 0x0a, 0x04, 0x72, 0x74, 0x67, 0x6f, 0x12, 0x30,
	0x0a, 0x06, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x72, 0x74, 0x67, 0x6f, 0x2e,
	0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x72, 0x74, 0x67, 0x6f, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_rtgo_proto_rawDescOnce sync.Once
	file_protos_rtgo_proto_rawDescData = file_protos_rtgo_proto_rawDesc
)

func file_protos_rtgo_proto_rawDescGZIP() []byte {
	file_protos_rtgo_proto_rawDescOnce.Do(func() {
		file_protos_rtgo_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_rtgo_proto_rawDescData)
	})
	return file_protos_rtgo_proto_rawDescData
}

var file_protos_rtgo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_rtgo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_rtgo_proto_goTypes = []interface{}{
	(Material_MaterialType)(0), // 0: rtgo.Material.MaterialType
	(*RenderReply)(nil),        // 1: rtgo.RenderReply
	(*RenderRequest)(nil),      // 2: rtgo.RenderRequest
	(*Camera)(nil),             // 3: rtgo.Camera
	(*Material)(nil),           // 4: rtgo.Material
	(*Texture)(nil),            // 5: rtgo.Texture
	(*Mesh)(nil),               // 6: rtgo.Mesh
	(*Vec3)(nil),               // 7: rtgo.Vec3
	(*FaceData)(nil),           // 8: rtgo.FaceData
}
var file_protos_rtgo_proto_depIdxs = []int32{
	7,  // 0: rtgo.RenderReply.colors:type_name -> rtgo.Vec3
	6,  // 1: rtgo.RenderRequest.mesh:type_name -> rtgo.Mesh
	4,  // 2: rtgo.RenderRequest.material:type_name -> rtgo.Material
	3,  // 3: rtgo.RenderRequest.camera:type_name -> rtgo.Camera
	7,  // 4: rtgo.Camera.lookfrom:type_name -> rtgo.Vec3
	7,  // 5: rtgo.Camera.lookAt:type_name -> rtgo.Vec3
	7,  // 6: rtgo.Camera.vup:type_name -> rtgo.Vec3
	5,  // 7: rtgo.Material.texture:type_name -> rtgo.Texture
	7,  // 8: rtgo.Material.albedo:type_name -> rtgo.Vec3
	7,  // 9: rtgo.Texture.colors:type_name -> rtgo.Vec3
	8,  // 10: rtgo.Mesh.facedata:type_name -> rtgo.FaceData
	7,  // 11: rtgo.Mesh.verts:type_name -> rtgo.Vec3
	7,  // 12: rtgo.Mesh.texcoords:type_name -> rtgo.Vec3
	7,  // 13: rtgo.Mesh.normals:type_name -> rtgo.Vec3
	2,  // 14: rtgo.rtgo.Render:input_type -> rtgo.RenderRequest
	1,  // 15: rtgo.rtgo.Render:output_type -> rtgo.RenderReply
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_protos_rtgo_proto_init() }
func file_protos_rtgo_proto_init() {
	if File_protos_rtgo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_rtgo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderReply); i {
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
		file_protos_rtgo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderRequest); i {
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
		file_protos_rtgo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Camera); i {
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
		file_protos_rtgo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Material); i {
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
		file_protos_rtgo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Texture); i {
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
		file_protos_rtgo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mesh); i {
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
		file_protos_rtgo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vec3); i {
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
		file_protos_rtgo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceData); i {
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
			RawDescriptor: file_protos_rtgo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_rtgo_proto_goTypes,
		DependencyIndexes: file_protos_rtgo_proto_depIdxs,
		EnumInfos:         file_protos_rtgo_proto_enumTypes,
		MessageInfos:      file_protos_rtgo_proto_msgTypes,
	}.Build()
	File_protos_rtgo_proto = out.File
	file_protos_rtgo_proto_rawDesc = nil
	file_protos_rtgo_proto_goTypes = nil
	file_protos_rtgo_proto_depIdxs = nil
}
