// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	model.proto
	state.proto

It has these top-level messages:
	Mesh
	Model
	State
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Mesh struct {
	Indices    []byte `protobuf:"bytes,1,opt,name=indices,proto3" json:"indices,omitempty"`
	Positions  []byte `protobuf:"bytes,2,opt,name=positions,proto3" json:"positions,omitempty"`
	Normals    []byte `protobuf:"bytes,3,opt,name=normals,proto3" json:"normals,omitempty"`
	Tangents   []byte `protobuf:"bytes,4,opt,name=tangents,proto3" json:"tangents,omitempty"`
	Bitangents []byte `protobuf:"bytes,5,opt,name=bitangents,proto3" json:"bitangents,omitempty"`
	Tcoords    []byte `protobuf:"bytes,6,opt,name=tcoords,proto3" json:"tcoords,omitempty"`
	AlbedoMap  []byte `protobuf:"bytes,7,opt,name=albedo_map,json=albedoMap,proto3" json:"albedo_map,omitempty"`
	NormalMap  []byte `protobuf:"bytes,8,opt,name=normal_map,json=normalMap,proto3" json:"normal_map,omitempty"`
	RoughMap   []byte `protobuf:"bytes,9,opt,name=rough_map,json=roughMap,proto3" json:"rough_map,omitempty"`
	MetalMap   []byte `protobuf:"bytes,10,opt,name=metal_map,json=metalMap,proto3" json:"metal_map,omitempty"`
	State      string `protobuf:"bytes,11,opt,name=state" json:"state,omitempty"`
	Name       string `protobuf:"bytes,12,opt,name=name" json:"name,omitempty"`
}

func (m *Mesh) Reset()                    { *m = Mesh{} }
func (m *Mesh) String() string            { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()               {}
func (*Mesh) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Mesh) GetIndices() []byte {
	if m != nil {
		return m.Indices
	}
	return nil
}

func (m *Mesh) GetPositions() []byte {
	if m != nil {
		return m.Positions
	}
	return nil
}

func (m *Mesh) GetNormals() []byte {
	if m != nil {
		return m.Normals
	}
	return nil
}

func (m *Mesh) GetTangents() []byte {
	if m != nil {
		return m.Tangents
	}
	return nil
}

func (m *Mesh) GetBitangents() []byte {
	if m != nil {
		return m.Bitangents
	}
	return nil
}

func (m *Mesh) GetTcoords() []byte {
	if m != nil {
		return m.Tcoords
	}
	return nil
}

func (m *Mesh) GetAlbedoMap() []byte {
	if m != nil {
		return m.AlbedoMap
	}
	return nil
}

func (m *Mesh) GetNormalMap() []byte {
	if m != nil {
		return m.NormalMap
	}
	return nil
}

func (m *Mesh) GetRoughMap() []byte {
	if m != nil {
		return m.RoughMap
	}
	return nil
}

func (m *Mesh) GetMetalMap() []byte {
	if m != nil {
		return m.MetalMap
	}
	return nil
}

func (m *Mesh) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Mesh) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Model struct {
	Meshes []*Mesh `protobuf:"bytes,1,rep,name=meshes" json:"meshes,omitempty"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (m *Model) String() string            { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Model) GetMeshes() []*Mesh {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func init() {
	proto.RegisterType((*Mesh)(nil), "protos.Mesh")
	proto.RegisterType((*Model)(nil), "protos.Model")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4d, 0x6e, 0xb3, 0x30,
	0x10, 0x86, 0x45, 0x02, 0x24, 0x0c, 0xac, 0xac, 0x6f, 0x61, 0x7d, 0xfd, 0x51, 0x14, 0x75, 0xc1,
	0xa6, 0x59, 0xb4, 0xe7, 0x60, 0xc3, 0x05, 0x2a, 0x13, 0xac, 0x80, 0x84, 0x3d, 0x88, 0x71, 0xef,
	0xd9, 0x23, 0x55, 0x9e, 0x71, 0xd3, 0xae, 0xe0, 0x7d, 0x1e, 0xcd, 0x30, 0x7a, 0x81, 0xda, 0xe1,
	0x68, 0x97, 0xcb, 0xba, 0x61, 0x40, 0x55, 0xf2, 0x83, 0xce, 0x5f, 0x3b, 0xc8, 0x3b, 0x4b, 0x93,
	0xd2, 0x70, 0x98, 0xfd, 0x38, 0x5f, 0x2d, 0xe9, 0xec, 0x94, 0xb5, 0x4d, 0xff, 0x13, 0xd5, 0x23,
	0x54, 0x2b, 0xd2, 0x1c, 0x66, 0xf4, 0xa4, 0x77, 0xec, 0x7e, 0x41, 0x9c, 0xf3, 0xb8, 0x39, 0xb3,
	0x90, 0xde, 0xcb, 0x5c, 0x8a, 0xea, 0x3f, 0x1c, 0x83, 0xf1, 0x37, 0xeb, 0x03, 0xe9, 0x9c, 0xd5,
	0x3d, 0xab, 0x67, 0x80, 0x61, 0xbe, 0xdb, 0x82, 0xed, 0x1f, 0x12, 0xb7, 0x86, 0x2b, 0xe2, 0x36,
	0x92, 0x2e, 0x65, 0x6b, 0x8a, 0xea, 0x09, 0xc0, 0x2c, 0x83, 0x1d, 0xf1, 0xc3, 0x99, 0x55, 0x1f,
	0xe4, 0x1c, 0x21, 0x9d, 0x59, 0xa3, 0x96, 0xef, 0xb3, 0x3e, 0x8a, 0x16, 0x12, 0xf5, 0x03, 0x54,
	0x1b, 0x7e, 0xde, 0x26, 0xb6, 0x95, 0x1c, 0xc5, 0x20, 0x49, 0x67, 0x43, 0x1a, 0x05, 0x91, 0x0c,
	0xa2, 0xfc, 0x07, 0x05, 0x05, 0x13, 0xac, 0xae, 0x4f, 0x59, 0x5b, 0xf5, 0x12, 0x94, 0x82, 0xdc,
	0x1b, 0x67, 0x75, 0xc3, 0x90, 0xdf, 0xcf, 0xaf, 0x50, 0x74, 0xb1, 0x69, 0xf5, 0x02, 0xa5, 0xb3,
	0x34, 0x71, 0xa3, 0xfb, 0xb6, 0x7e, 0x6b, 0xa4, 0x7b, 0xba, 0xc4, 0xc2, 0xfb, 0xe4, 0x06, 0xf9,
	0x13, 0xef, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x72, 0x27, 0x9a, 0x9f, 0x01, 0x00, 0x00,
}
