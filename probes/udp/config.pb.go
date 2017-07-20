// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/udp/config.proto

/*
Package udp is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/probes/udp/config.proto

It has these top-level messages:
	ProbeConf
*/
package udp

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

type ProbeConf struct {
	// Export stats after these many milliseconds
	StatsExportIntervalMsec *int32 `protobuf:"varint,2,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec,def=10000" json:"stats_export_interval_msec,omitempty"`
	// Port to send UDP Ping to (UDP Echo).  Should be same as
	// ProberConfig.udp_echo_server_port.
	// TODO: Can we just read this from ProberConfig?
	Port             *int32 `protobuf:"varint,3,opt,name=port,def=31122" json:"port,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProbeConf) Reset()                    { *m = ProbeConf{} }
func (m *ProbeConf) String() string            { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()               {}
func (*ProbeConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ProbeConf_StatsExportIntervalMsec int32 = 10000
const Default_ProbeConf_Port int32 = 31122

func (m *ProbeConf) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return Default_ProbeConf_StatsExportIntervalMsec
}

func (m *ProbeConf) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return Default_ProbeConf_Port
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.probes.udp.ProbeConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/udp/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0x29, 0x28, 0xca, 0x4f, 0x4a, 0x2d, 0xd2, 0x07, 0x53, 0xc5, 0xfa, 0xa5, 0x29,
	0x05, 0xfa, 0xc9, 0xf9, 0x79, 0x69, 0x99, 0xe9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x62,
	0x48, 0x8a, 0xf4, 0x20, 0x8a, 0xf4, 0x4a, 0x53, 0x0a, 0x94, 0xb2, 0xb8, 0x38, 0x03, 0x40, 0x3c,
	0xe7, 0xfc, 0xbc, 0x34, 0x21, 0x27, 0x2e, 0xa9, 0xe2, 0x92, 0xc4, 0x92, 0xe2, 0xf8, 0xd4, 0x8a,
	0x82, 0xfc, 0xa2, 0x92, 0xf8, 0xcc, 0xbc, 0x92, 0xd4, 0xa2, 0xb2, 0xc4, 0x9c, 0xf8, 0xdc, 0xe2,
	0xd4, 0x64, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x56, 0x2b, 0x56, 0x43, 0x03, 0x03, 0x03, 0x83, 0x20,
	0x71, 0xb0, 0x42, 0x57, 0xb0, 0x3a, 0x4f, 0xa8, 0x32, 0xdf, 0xe2, 0xd4, 0x64, 0x21, 0x49, 0x2e,
	0x16, 0x90, 0x98, 0x04, 0x33, 0x44, 0xb5, 0xb1, 0xa1, 0xa1, 0x91, 0x51, 0x10, 0x58, 0x08, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0xbc, 0x10, 0x7b, 0xbb, 0x00, 0x00, 0x00,
}
