// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enums.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Valid schema types
type SchemaType int32

const (
	SchemaType_PEDERSEN                            SchemaType = 0
	SchemaType_PEDERSEN_EC                         SchemaType = 1
	SchemaType_SCHNORR                             SchemaType = 2
	SchemaType_SCHNORR_EC                          SchemaType = 3
	SchemaType_CSPAILLIER                          SchemaType = 4
	SchemaType_PSEUDONYMSYS_CA                     SchemaType = 5
	SchemaType_PSEUDONYMSYS_NYM_GEN                SchemaType = 6
	SchemaType_PSEUDONYMSYS_ISSUE_CREDENTIAL       SchemaType = 7
	SchemaType_PSEUDONYMSYS_TRANSFER_CREDENTIAL    SchemaType = 8
	SchemaType_PSEUDONYMSYS_CA_EC                  SchemaType = 9
	SchemaType_PSEUDONYMSYS_NYM_GEN_EC             SchemaType = 10
	SchemaType_PSEUDONYMSYS_ISSUE_CREDENTIAL_EC    SchemaType = 11
	SchemaType_PSEUDONYMSYS_TRANSFER_CREDENTIAL_EC SchemaType = 12
	SchemaType_QR                                  SchemaType = 13
	SchemaType_QNR                                 SchemaType = 14
)

var SchemaType_name = map[int32]string{
	0:  "PEDERSEN",
	1:  "PEDERSEN_EC",
	2:  "SCHNORR",
	3:  "SCHNORR_EC",
	4:  "CSPAILLIER",
	5:  "PSEUDONYMSYS_CA",
	6:  "PSEUDONYMSYS_NYM_GEN",
	7:  "PSEUDONYMSYS_ISSUE_CREDENTIAL",
	8:  "PSEUDONYMSYS_TRANSFER_CREDENTIAL",
	9:  "PSEUDONYMSYS_CA_EC",
	10: "PSEUDONYMSYS_NYM_GEN_EC",
	11: "PSEUDONYMSYS_ISSUE_CREDENTIAL_EC",
	12: "PSEUDONYMSYS_TRANSFER_CREDENTIAL_EC",
	13: "QR",
	14: "QNR",
}
var SchemaType_value = map[string]int32{
	"PEDERSEN":                            0,
	"PEDERSEN_EC":                         1,
	"SCHNORR":                             2,
	"SCHNORR_EC":                          3,
	"CSPAILLIER":                          4,
	"PSEUDONYMSYS_CA":                     5,
	"PSEUDONYMSYS_NYM_GEN":                6,
	"PSEUDONYMSYS_ISSUE_CREDENTIAL":       7,
	"PSEUDONYMSYS_TRANSFER_CREDENTIAL":    8,
	"PSEUDONYMSYS_CA_EC":                  9,
	"PSEUDONYMSYS_NYM_GEN_EC":             10,
	"PSEUDONYMSYS_ISSUE_CREDENTIAL_EC":    11,
	"PSEUDONYMSYS_TRANSFER_CREDENTIAL_EC": 12,
	"QR":  13,
	"QNR": 14,
}

func (x SchemaType) String() string {
	return proto.EnumName(SchemaType_name, int32(x))
}
func (SchemaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// Valid schema variants
type SchemaVariant int32

const (
	SchemaVariant_SIGMA SchemaVariant = 0
	SchemaVariant_ZKP   SchemaVariant = 1
	SchemaVariant_ZKPOK SchemaVariant = 2
)

var SchemaVariant_name = map[int32]string{
	0: "SIGMA",
	1: "ZKP",
	2: "ZKPOK",
}
var SchemaVariant_value = map[string]int32{
	"SIGMA": 0,
	"ZKP":   1,
	"ZKPOK": 2,
}

func (x SchemaVariant) String() string {
	return proto.EnumName(SchemaVariant_name, int32(x))
}
func (SchemaVariant) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterEnum("protobuf.SchemaType", SchemaType_name, SchemaType_value)
	proto.RegisterEnum("protobuf.SchemaVariant", SchemaVariant_name, SchemaVariant_value)
}

func init() { proto.RegisterFile("enums.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0xdb, 0xd4, 0x36, 0xe9, 0x4d, 0x7f, 0x2e, 0x57, 0x51, 0x41, 0x04, 0x45, 0x41, 0xe8,
	0xa2, 0x1b, 0x9f, 0x20, 0x24, 0x63, 0x0d, 0x49, 0x26, 0xe9, 0xdc, 0x54, 0x48, 0x37, 0x21, 0x95,
	0x88, 0x2e, 0xfa, 0x43, 0x6d, 0x17, 0x3e, 0xb0, 0xef, 0x21, 0x53, 0x15, 0x4c, 0x15, 0x5c, 0x0d,
	0xe7, 0x9c, 0x8f, 0x39, 0x87, 0x19, 0xb0, 0xcb, 0xc5, 0x76, 0xfe, 0x3a, 0x5c, 0xad, 0x97, 0x9b,
	0x25, 0x59, 0xbb, 0x63, 0xb6, 0x7d, 0x1a, 0xbc, 0x1b, 0x00, 0xfc, 0xf8, 0x5c, 0xce, 0x8b, 0xf4,
	0x6d, 0x55, 0x52, 0x07, 0xac, 0x44, 0x78, 0x42, 0xb1, 0x90, 0x58, 0xa3, 0x3e, 0xd8, 0xdf, 0x2a,
	0x17, 0x2e, 0xd6, 0xc9, 0x06, 0x93, 0xdd, 0x7b, 0x19, 0x2b, 0x85, 0x06, 0xf5, 0x00, 0xbe, 0x84,
	0x0e, 0x1b, 0x5a, 0xbb, 0x9c, 0x38, 0x7e, 0x18, 0xfa, 0x42, 0xe1, 0x01, 0x1d, 0x42, 0x3f, 0x61,
	0x31, 0xf1, 0x62, 0x99, 0x45, 0x9c, 0x71, 0xee, 0x3a, 0xd8, 0xa4, 0x53, 0x38, 0xaa, 0x98, 0x32,
	0x8b, 0xf2, 0x91, 0x90, 0xd8, 0xa2, 0x4b, 0x38, 0xaf, 0x24, 0x3e, 0xf3, 0x44, 0xe4, 0xae, 0x12,
	0x9e, 0x90, 0xa9, 0xef, 0x84, 0x68, 0xd2, 0x35, 0x5c, 0x54, 0x90, 0x54, 0x39, 0x92, 0xef, 0x84,
	0xfa, 0x49, 0x59, 0x74, 0x0c, 0xb4, 0xd7, 0xab, 0xf7, 0xb5, 0xe9, 0x0c, 0x4e, 0xfe, 0xaa, 0xd6,
	0x21, 0xfc, 0xba, 0x7a, 0xbf, 0x5d, 0x53, 0x36, 0xdd, 0xc0, 0xd5, 0x7f, 0x03, 0x34, 0xd8, 0xa1,
	0x16, 0x18, 0x63, 0x85, 0x5d, 0x32, 0xa1, 0x31, 0x96, 0x0a, 0x7b, 0x83, 0x21, 0x74, 0x3f, 0x9f,
	0xf9, 0xa1, 0x58, 0xbf, 0x14, 0x8b, 0x0d, 0xb5, 0xa1, 0xc9, 0xfe, 0x28, 0x72, 0xb0, 0xa6, 0xa1,
	0x69, 0x90, 0x60, 0x5d, 0x7b, 0xd3, 0x20, 0x89, 0x03, 0x34, 0x66, 0xad, 0xdd, 0x0f, 0xdd, 0x7e,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x85, 0x92, 0x24, 0xb7, 0x01, 0x00, 0x00,
}