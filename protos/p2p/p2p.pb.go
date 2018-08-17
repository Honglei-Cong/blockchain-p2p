// Code generated by protoc-gen-go.
// source: p2p/p2p.proto
// DO NOT EDIT!

/*
Package p2p is a generated protocol buffer package.

It is generated from these files:
	p2p/p2p.proto

It has these top-level messages:
	ID
	PacketPing
	PacketPong
	PacketMsg
	NodeInfo
	PexRequestMessage
	NetAddress
	PexAddrsMessage
*/
package p2p

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

type ID struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PacketPing struct {
}

func (m *PacketPing) Reset()                    { *m = PacketPing{} }
func (m *PacketPing) String() string            { return proto.CompactTextString(m) }
func (*PacketPing) ProtoMessage()               {}
func (*PacketPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PacketPong struct {
}

func (m *PacketPong) Reset()                    { *m = PacketPong{} }
func (m *PacketPong) String() string            { return proto.CompactTextString(m) }
func (*PacketPong) ProtoMessage()               {}
func (*PacketPong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PacketMsg struct {
	ChannelID uint32 `protobuf:"varint,1,opt,name=channelID" json:"channelID,omitempty"`
	EOF       uint32 `protobuf:"varint,2,opt,name=EOF" json:"EOF,omitempty"`
	Bytes     []byte `protobuf:"bytes,3,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
}

func (m *PacketMsg) Reset()                    { *m = PacketMsg{} }
func (m *PacketMsg) String() string            { return proto.CompactTextString(m) }
func (*PacketMsg) ProtoMessage()               {}
func (*PacketMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type NodeInfo struct {
	ID         *ID      `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	ListenAddr string   `protobuf:"bytes,2,opt,name=listenAddr" json:"listenAddr,omitempty"`
	Network    string   `protobuf:"bytes,3,opt,name=network" json:"network,omitempty"`
	Version    string   `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	Channels   []byte   `protobuf:"bytes,5,opt,name=channels,proto3" json:"channels,omitempty"`
	Moniker    string   `protobuf:"bytes,6,opt,name=moniker" json:"moniker,omitempty"`
	Other      []string `protobuf:"bytes,7,rep,name=other" json:"other,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NodeInfo) GetID() *ID {
	if m != nil {
		return m.ID
	}
	return nil
}

type PexRequestMessage struct {
}

func (m *PexRequestMessage) Reset()                    { *m = PexRequestMessage{} }
func (m *PexRequestMessage) String() string            { return proto.CompactTextString(m) }
func (*PexRequestMessage) ProtoMessage()               {}
func (*PexRequestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type NetAddress struct {
	ID   *ID    `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	IP   []byte `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Port uint32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Str  string `protobuf:"bytes,4,opt,name=str" json:"str,omitempty"`
}

func (m *NetAddress) Reset()                    { *m = NetAddress{} }
func (m *NetAddress) String() string            { return proto.CompactTextString(m) }
func (*NetAddress) ProtoMessage()               {}
func (*NetAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NetAddress) GetID() *ID {
	if m != nil {
		return m.ID
	}
	return nil
}

type PexAddrsMessage struct {
	Addrs []*NetAddress `protobuf:"bytes,1,rep,name=Addrs" json:"Addrs,omitempty"`
}

func (m *PexAddrsMessage) Reset()                    { *m = PexAddrsMessage{} }
func (m *PexAddrsMessage) String() string            { return proto.CompactTextString(m) }
func (*PexAddrsMessage) ProtoMessage()               {}
func (*PexAddrsMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PexAddrsMessage) GetAddrs() []*NetAddress {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func init() {
	proto.RegisterType((*ID)(nil), "p2p.ID")
	proto.RegisterType((*PacketPing)(nil), "p2p.PacketPing")
	proto.RegisterType((*PacketPong)(nil), "p2p.PacketPong")
	proto.RegisterType((*PacketMsg)(nil), "p2p.PacketMsg")
	proto.RegisterType((*NodeInfo)(nil), "p2p.NodeInfo")
	proto.RegisterType((*PexRequestMessage)(nil), "p2p.PexRequestMessage")
	proto.RegisterType((*NetAddress)(nil), "p2p.NetAddress")
	proto.RegisterType((*PexAddrsMessage)(nil), "p2p.PexAddrsMessage")
}

func init() { proto.RegisterFile("p2p/p2p.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0x64, 0xbb, 0xdd, 0xcc, 0xb6, 0x2c, 0x98, 0x95, 0xb0, 0x10, 0x42, 0x51, 0x24,
	0x50, 0x2e, 0x9b, 0x48, 0xe1, 0x02, 0x47, 0x56, 0x05, 0x29, 0x87, 0x2d, 0xc1, 0x17, 0x24, 0x38,
	0xa5, 0xe9, 0x90, 0x44, 0x69, 0x6d, 0x63, 0xbb, 0x50, 0x7e, 0x1f, 0x7f, 0x0c, 0xd9, 0x49, 0x3f,
	0x38, 0x70, 0x89, 0xe6, 0x99, 0x2f, 0xbf, 0x6f, 0x6c, 0x98, 0xcb, 0x5c, 0x66, 0x32, 0x97, 0xa9,
	0x54, 0xc2, 0x08, 0x12, 0xc8, 0x5c, 0xc6, 0xb7, 0xe0, 0x17, 0x0b, 0xf2, 0xc8, 0x7e, 0xa9, 0x17,
	0x79, 0x49, 0xc8, 0xfc, 0x62, 0x11, 0xcf, 0x00, 0xca, 0xaa, 0xee, 0xd1, 0x94, 0x1d, 0x6f, 0xce,
	0x48, 0xf0, 0x26, 0xfe, 0x0c, 0xe1, 0x40, 0x0f, 0xba, 0x21, 0x2f, 0x20, 0xac, 0xdb, 0x8a, 0x73,
	0xdc, 0x8c, 0xf3, 0x73, 0x76, 0x4a, 0x90, 0xc7, 0x10, 0x7c, 0xf8, 0xf4, 0x91, 0xfa, 0x2e, 0x6f,
	0x43, 0x72, 0x0b, 0x93, 0xfb, 0xdf, 0x06, 0x35, 0x0d, 0x22, 0x2f, 0x99, 0xb1, 0x01, 0xe2, 0x3f,
	0x1e, 0x5c, 0x2d, 0xc5, 0x1a, 0x0b, 0xfe, 0x5d, 0x90, 0x67, 0x47, 0x2d, 0xd7, 0xf9, 0x34, 0xb5,
	0x72, 0x8b, 0x85, 0x15, 0x45, 0x5e, 0x02, 0x6c, 0x3a, 0x6d, 0x90, 0xbf, 0x5f, 0xaf, 0x95, 0x5b,
	0x1a, 0xb2, 0xb3, 0x0c, 0xa1, 0x30, 0xe5, 0x68, 0x7e, 0x09, 0xd5, 0xbb, 0xed, 0x21, 0x3b, 0xa0,
	0xad, 0xfc, 0x44, 0xa5, 0x3b, 0xc1, 0xe9, 0xc5, 0x50, 0x19, 0x91, 0x3c, 0x87, 0xab, 0x51, 0xae,
	0xa6, 0x13, 0x27, 0xe9, 0xc8, 0x76, 0x6a, 0x2b, 0x78, 0xd7, 0xa3, 0xa2, 0x97, 0xc3, 0xd4, 0x88,
	0xd6, 0x85, 0x30, 0x2d, 0x2a, 0x3a, 0x8d, 0x82, 0x24, 0x64, 0x03, 0xc4, 0x4f, 0xe1, 0x49, 0x89,
	0x7b, 0x86, 0x3f, 0x76, 0xa8, 0xcd, 0x03, 0x6a, 0x5d, 0x35, 0x18, 0x7f, 0x03, 0x58, 0xa2, 0xb1,
	0xfa, 0x50, 0xeb, 0xff, 0x7b, 0xb3, 0x17, 0x50, 0x3a, 0x4f, 0x33, 0xe6, 0x17, 0x25, 0x21, 0x70,
	0x21, 0x85, 0x32, 0xce, 0xc8, 0x9c, 0xb9, 0xd8, 0xfe, 0x4d, 0x6d, 0xd4, 0xe8, 0xc0, 0x86, 0xf1,
	0x5b, 0xb8, 0x29, 0x71, 0x6f, 0x97, 0xeb, 0xf1, 0x3c, 0xf2, 0x0a, 0x26, 0x8e, 0xa9, 0x17, 0x05,
	0xc9, 0x75, 0x7e, 0xe3, 0x0e, 0x39, 0x29, 0x60, 0x43, 0xf5, 0xfe, 0x0b, 0xbc, 0x16, 0xaa, 0x49,
	0xdf, 0x99, 0xb6, 0x6e, 0xab, 0x8e, 0xa7, 0xab, 0x8d, 0xa8, 0x7b, 0x17, 0xde, 0x1d, 0x1f, 0x88,
	0xb6, 0xb3, 0x5f, 0xef, 0x9a, 0xce, 0xb4, 0xbb, 0x55, 0x5a, 0x8b, 0x6d, 0x76, 0x68, 0xcf, 0xfe,
	0x6d, 0xcf, 0x86, 0x76, 0xfb, 0xb4, 0x56, 0x97, 0x2e, 0x7e, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x8d, 0x64, 0x1e, 0x6c, 0x02, 0x00, 0x00,
}