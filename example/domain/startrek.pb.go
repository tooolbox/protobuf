// Code generated by protoc-gen-go. DO NOT EDIT.
// source: startrek.proto

/*
Package startrek is a generated protocol buffer package.

It is generated from these files:
	startrek.proto

It has these top-level messages:
	StarfleetShip
*/
package startrek

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dkfbasel_protobuf "github.com/dkfbasel/protobuf/types/timestamp"
import dkfbasel_protobuf1 "github.com/dkfbasel/protobuf/types/nullstring"
import dkfbasel_protobuf2 "github.com/dkfbasel/protobuf/types/nullint"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StarfleetShip struct {
	Name             string                         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NoOfPassengers   *dkfbasel_protobuf2.NullInt    `protobuf:"bytes,2,opt,name=no_of_passengers,json=noOfPassengers" json:"no_of_passengers,omitempty"`
	MissionStatement *dkfbasel_protobuf1.NullString `protobuf:"bytes,3,opt,name=mission_statement,json=missionStatement" json:"mission_statement,omitempty"`
	// use a different db column name for the departure time
	// `db:"we_are_leaving_at"`
	DepartureTime *dkfbasel_protobuf.Timestamp `protobuf:"bytes,4,opt,name=departure_time,json=departureTime" json:"departure_time,omitempty" db:"we_are_leaving_at"`
}

func (m *StarfleetShip) Reset()                    { *m = StarfleetShip{} }
func (m *StarfleetShip) String() string            { return proto.CompactTextString(m) }
func (*StarfleetShip) ProtoMessage()               {}
func (*StarfleetShip) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StarfleetShip) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StarfleetShip) GetNoOfPassengers() *dkfbasel_protobuf2.NullInt {
	if m != nil {
		return m.NoOfPassengers
	}
	return nil
}

func (m *StarfleetShip) GetMissionStatement() *dkfbasel_protobuf1.NullString {
	if m != nil {
		return m.MissionStatement
	}
	return nil
}

func (m *StarfleetShip) GetDepartureTime() *dkfbasel_protobuf.Timestamp {
	if m != nil {
		return m.DepartureTime
	}
	return nil
}

func init() {
	proto.RegisterType((*StarfleetShip)(nil), "startrek.StarfleetShip")
}

func init() { proto.RegisterFile("startrek.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0xa8, 0x10, 0x18, 0x35, 0x2a, 0x9e, 0xa2, 0x08, 0x44, 0xe9, 0xd4, 0x29, 0x95,
	0xe0, 0x27, 0xc0, 0x02, 0x03, 0xa0, 0x84, 0x3d, 0x72, 0xd4, 0x73, 0xb1, 0x6a, 0x9f, 0x2d, 0xdf,
	0xe5, 0xcf, 0x33, 0xa1, 0xba, 0x75, 0x17, 0xca, 0xe6, 0x8f, 0xe7, 0x7d, 0x4e, 0xf7, 0x8a, 0x92,
	0x58, 0x45, 0x8e, 0xb0, 0x6d, 0x42, 0xf4, 0xec, 0xe5, 0x65, 0xbe, 0xd7, 0x0f, 0xeb, 0xad, 0x1e,
	0x14, 0x81, 0x5d, 0xa5, 0x9f, 0x61, 0xd4, 0x2b, 0x36, 0x0e, 0x88, 0x95, 0x0b, 0x7b, 0xb8, 0x5e,
	0xfc, 0x45, 0x70, 0xb4, 0x96, 0x38, 0x1a, 0xdc, 0x1c, 0x98, 0xfb, 0xd3, 0x8c, 0x41, 0xde, 0x03,
	0x8b, 0x9f, 0x42, 0x4c, 0x3b, 0x56, 0x51, 0x5b, 0x00, 0xee, 0xbe, 0x4d, 0x90, 0x52, 0x4c, 0x50,
	0x39, 0xa8, 0x8a, 0x79, 0xb1, 0xbc, 0x6a, 0xd3, 0x59, 0xbe, 0x88, 0x19, 0xfa, 0xde, 0xeb, 0x3e,
	0x28, 0x22, 0xc0, 0x0d, 0x44, 0xaa, 0xce, 0xe6, 0xc5, 0xf2, 0xfa, 0xb1, 0x6e, 0xf2, 0x84, 0x26,
	0x4f, 0x68, 0xde, 0x47, 0x6b, 0x5f, 0x91, 0xdb, 0x12, 0xfd, 0x87, 0xfe, 0x3c, 0x26, 0xe4, 0x9b,
	0xb8, 0x71, 0x86, 0xc8, 0x78, 0xec, 0x89, 0x15, 0x83, 0x03, 0xe4, 0xea, 0x3c, 0x69, 0xee, 0xfe,
	0xd1, 0x74, 0x69, 0x99, 0x76, 0x76, 0xc8, 0x75, 0x39, 0x26, 0x9f, 0x45, 0xb9, 0x86, 0xa0, 0x22,
	0x8f, 0x11, 0xfa, 0x5d, 0x33, 0xd5, 0x24, 0x89, 0x6e, 0x4f, 0x88, 0xbe, 0x72, 0x71, 0xed, 0xf4,
	0x98, 0xd9, 0xbd, 0x0d, 0x17, 0x09, 0x79, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x11, 0xcd, 0xac,
	0x81, 0x87, 0x01, 0x00, 0x00,
}
