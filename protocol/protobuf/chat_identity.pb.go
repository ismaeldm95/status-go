// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_identity.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// SourceType are the predefined types of image source allowed
type IdentityImage_SourceType int32

const (
	IdentityImage_UNKNOWN_SOURCE_TYPE IdentityImage_SourceType = 0
	// RAW_PAYLOAD image byte data
	IdentityImage_RAW_PAYLOAD IdentityImage_SourceType = 1
	// ENS_AVATAR uses the ENS record's resolver get-text-data.avatar data
	// The `payload` field will be ignored if ENS_AVATAR is selected
	// The application will read and parse the ENS avatar data as image payload data, URLs will be ignored
	// The parent `ChatMessageIdentity` must have a valid `ens_name` set
	IdentityImage_ENS_AVATAR IdentityImage_SourceType = 2
)

var IdentityImage_SourceType_name = map[int32]string{
	0: "UNKNOWN_SOURCE_TYPE",
	1: "RAW_PAYLOAD",
	2: "ENS_AVATAR",
}

var IdentityImage_SourceType_value = map[string]int32{
	"UNKNOWN_SOURCE_TYPE": 0,
	"RAW_PAYLOAD":         1,
	"ENS_AVATAR":          2,
}

func (x IdentityImage_SourceType) String() string {
	return proto.EnumName(IdentityImage_SourceType_name, int32(x))
}

func (IdentityImage_SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{1, 0}
}

// ChatIdentity represents the user defined identity associated with their public chat key
type ChatIdentity struct {
	// Lamport timestamp of the message
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// ens_name is the valid ENS name associated with the chat key
	EnsName string `protobuf:"bytes,2,opt,name=ens_name,json=ensName,proto3" json:"ens_name,omitempty"`
	// images is a string indexed mapping of images associated with an identity
	Images map[string]*IdentityImage `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// display name is the user set identity, valid only for communitys
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// description is the user set description, valid only for communitys
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Color                string   `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatIdentity) Reset()         { *m = ChatIdentity{} }
func (m *ChatIdentity) String() string { return proto.CompactTextString(m) }
func (*ChatIdentity) ProtoMessage()    {}
func (*ChatIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{0}
}

func (m *ChatIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatIdentity.Unmarshal(m, b)
}
func (m *ChatIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatIdentity.Marshal(b, m, deterministic)
}
func (m *ChatIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatIdentity.Merge(m, src)
}
func (m *ChatIdentity) XXX_Size() int {
	return xxx_messageInfo_ChatIdentity.Size(m)
}
func (m *ChatIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_ChatIdentity proto.InternalMessageInfo

func (m *ChatIdentity) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *ChatIdentity) GetEnsName() string {
	if m != nil {
		return m.EnsName
	}
	return ""
}

func (m *ChatIdentity) GetImages() map[string]*IdentityImage {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *ChatIdentity) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ChatIdentity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ChatIdentity) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

// ProfileImage represents data associated with a user's profile image
type IdentityImage struct {
	// payload is a context based payload for the profile image data,
	// context is determined by the `source_type`
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// source_type signals the image payload source
	SourceType IdentityImage_SourceType `protobuf:"varint,2,opt,name=source_type,json=sourceType,proto3,enum=protobuf.IdentityImage_SourceType" json:"source_type,omitempty"`
	// image_type signals the image type and method of parsing the payload
	ImageType            ImageType `protobuf:"varint,3,opt,name=image_type,json=imageType,proto3,enum=protobuf.ImageType" json:"image_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IdentityImage) Reset()         { *m = IdentityImage{} }
func (m *IdentityImage) String() string { return proto.CompactTextString(m) }
func (*IdentityImage) ProtoMessage()    {}
func (*IdentityImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{1}
}

func (m *IdentityImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityImage.Unmarshal(m, b)
}
func (m *IdentityImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityImage.Marshal(b, m, deterministic)
}
func (m *IdentityImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityImage.Merge(m, src)
}
func (m *IdentityImage) XXX_Size() int {
	return xxx_messageInfo_IdentityImage.Size(m)
}
func (m *IdentityImage) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityImage.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityImage proto.InternalMessageInfo

func (m *IdentityImage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *IdentityImage) GetSourceType() IdentityImage_SourceType {
	if m != nil {
		return m.SourceType
	}
	return IdentityImage_UNKNOWN_SOURCE_TYPE
}

func (m *IdentityImage) GetImageType() ImageType {
	if m != nil {
		return m.ImageType
	}
	return ImageType_UNKNOWN_IMAGE_TYPE
}

func init() {
	proto.RegisterEnum("protobuf.IdentityImage_SourceType", IdentityImage_SourceType_name, IdentityImage_SourceType_value)
	proto.RegisterType((*ChatIdentity)(nil), "protobuf.ChatIdentity")
	proto.RegisterMapType((map[string]*IdentityImage)(nil), "protobuf.ChatIdentity.ImagesEntry")
	proto.RegisterType((*IdentityImage)(nil), "protobuf.IdentityImage")
}

func init() {
	proto.RegisterFile("chat_identity.proto", fileDescriptor_7a652489000a5879)
}

var fileDescriptor_7a652489000a5879 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x8e, 0x9b, 0x30,
	0x10, 0xc6, 0x0b, 0x6c, 0xb2, 0x9b, 0x71, 0xba, 0x45, 0x4e, 0xa5, 0xa5, 0x7b, 0xa2, 0x9c, 0x72,
	0x29, 0x87, 0xf4, 0x52, 0xed, 0x0d, 0xa5, 0x54, 0x5a, 0xb5, 0x22, 0x2b, 0x27, 0x69, 0x94, 0x13,
	0x72, 0xc0, 0x6d, 0x50, 0x00, 0x23, 0x6c, 0x2a, 0xf1, 0x4a, 0x7d, 0xb4, 0x3e, 0x45, 0x15, 0x1b,
	0x36, 0xe4, 0xe4, 0xf9, 0xf3, 0xcd, 0xef, 0x1b, 0x0f, 0xcc, 0x92, 0x23, 0x95, 0x71, 0x96, 0xb2,
	0x52, 0x66, 0xb2, 0xf5, 0xab, 0x9a, 0x4b, 0x8e, 0xef, 0xd4, 0x73, 0x68, 0x7e, 0x3d, 0x22, 0x56,
	0x36, 0x85, 0xd0, 0x65, 0xef, 0xaf, 0x09, 0xd3, 0xe5, 0x91, 0xca, 0xe7, 0x4e, 0x8d, 0xdf, 0xc3,
	0x28, 0xc9, 0x79, 0x72, 0x72, 0x0c, 0xd7, 0x98, 0xdf, 0x10, 0x9d, 0xe0, 0x0f, 0x70, 0xc7, 0x4a,
	0x11, 0x97, 0xb4, 0x60, 0x8e, 0xe9, 0x1a, 0xf3, 0x09, 0xb9, 0x65, 0xa5, 0x88, 0x68, 0xc1, 0xf0,
	0x13, 0x8c, 0xb3, 0x82, 0xfe, 0x66, 0xc2, 0xb1, 0x5c, 0x6b, 0x8e, 0x16, 0x9e, 0xdf, 0x3b, 0xf9,
	0x43, 0xb0, 0xff, 0xac, 0x44, 0x61, 0x29, 0xeb, 0x96, 0x74, 0x13, 0xf8, 0x23, 0x4c, 0xd3, 0x4c,
	0x54, 0x39, 0x6d, 0x35, 0xfa, 0x46, 0xa1, 0x51, 0x57, 0x53, 0x78, 0x17, 0x50, 0xca, 0x44, 0x52,
	0x67, 0x95, 0xcc, 0x78, 0xe9, 0x8c, 0x3a, 0xc5, 0xa5, 0xa4, 0x36, 0xe6, 0x39, 0xaf, 0x9d, 0xb1,
	0xea, 0xe9, 0xe4, 0x91, 0x00, 0x1a, 0x38, 0x62, 0x1b, 0xac, 0x13, 0x6b, 0xd5, 0xa7, 0x26, 0xe4,
	0x1c, 0xe2, 0x4f, 0x30, 0xfa, 0x43, 0xf3, 0x46, 0xff, 0x07, 0x2d, 0x1e, 0x2e, 0x6b, 0xf7, 0x2b,
	0xab, 0x79, 0xa2, 0x55, 0x4f, 0xe6, 0x17, 0xc3, 0xfb, 0x67, 0xc0, 0xdb, 0xab, 0x26, 0x76, 0xe0,
	0xb6, 0xa2, 0x6d, 0xce, 0x69, 0xaa, 0xd0, 0x53, 0xd2, 0xa7, 0x78, 0x09, 0x48, 0xf0, 0xa6, 0x4e,
	0x58, 0x2c, 0xdb, 0x4a, 0x9b, 0xdc, 0x0f, 0x6f, 0x73, 0xc5, 0xf1, 0xd7, 0x4a, 0xba, 0x69, 0x2b,
	0x46, 0x40, 0xbc, 0xc6, 0x78, 0x01, 0xa0, 0x2e, 0xa5, 0x19, 0x96, 0x62, 0xcc, 0x06, 0x8c, 0x73,
	0x4f, 0x0d, 0x4d, 0xb2, 0x3e, 0xf4, 0xbe, 0x01, 0x5c, 0x68, 0xf8, 0x01, 0x66, 0xdb, 0xe8, 0x7b,
	0xb4, 0xda, 0x45, 0xf1, 0x7a, 0xb5, 0x25, 0xcb, 0x30, 0xde, 0xec, 0x5f, 0x42, 0xfb, 0x0d, 0x7e,
	0x07, 0x88, 0x04, 0xbb, 0xf8, 0x25, 0xd8, 0xff, 0x58, 0x05, 0x5f, 0x6d, 0x03, 0xdf, 0x03, 0x84,
	0xd1, 0x3a, 0x0e, 0x7e, 0x06, 0x9b, 0x80, 0xd8, 0xe6, 0x61, 0xac, 0x6c, 0x3e, 0xff, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0xf5, 0x91, 0x10, 0x4e, 0x02, 0x00, 0x00,
}