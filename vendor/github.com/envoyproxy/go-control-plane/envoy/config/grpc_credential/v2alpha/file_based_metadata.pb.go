// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/grpc_credential/v2alpha/file_based_metadata.proto

package v2alpha

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FileBasedMetadataConfig struct {
	// Location or inline data of secret to use for authentication of the Google gRPC connection
	// this secret will be attached to a header of the gRPC connection
	SecretData *core.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// Metadata header key to use for sending the secret data
	// if no header key is set, "authorization" header will be used
	HeaderKey string `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	// Prefix to prepend to the secret in the metadata header
	// if no prefix is set, the default is to use no prefix
	HeaderPrefix         string   `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileBasedMetadataConfig) Reset()         { *m = FileBasedMetadataConfig{} }
func (m *FileBasedMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*FileBasedMetadataConfig) ProtoMessage()    {}
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_based_metadata_d58d604cef8b60b9, []int{0}
}
func (m *FileBasedMetadataConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileBasedMetadataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileBasedMetadataConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FileBasedMetadataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBasedMetadataConfig.Merge(dst, src)
}
func (m *FileBasedMetadataConfig) XXX_Size() int {
	return m.Size()
}
func (m *FileBasedMetadataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBasedMetadataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FileBasedMetadataConfig proto.InternalMessageInfo

func (m *FileBasedMetadataConfig) GetSecretData() *core.DataSource {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *FileBasedMetadataConfig) GetHeaderKey() string {
	if m != nil {
		return m.HeaderKey
	}
	return ""
}

func (m *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if m != nil {
		return m.HeaderPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*FileBasedMetadataConfig)(nil), "envoy.config.grpc_credential.v2alpha.FileBasedMetadataConfig")
}
func (m *FileBasedMetadataConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileBasedMetadataConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SecretData != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(m.SecretData.Size()))
		n1, err := m.SecretData.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.HeaderKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(len(m.HeaderKey)))
		i += copy(dAtA[i:], m.HeaderKey)
	}
	if len(m.HeaderPrefix) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFileBasedMetadata(dAtA, i, uint64(len(m.HeaderPrefix)))
		i += copy(dAtA[i:], m.HeaderPrefix)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFileBasedMetadata(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FileBasedMetadataConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SecretData != nil {
		l = m.SecretData.Size()
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	l = len(m.HeaderKey)
	if l > 0 {
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	l = len(m.HeaderPrefix)
	if l > 0 {
		n += 1 + l + sovFileBasedMetadata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileBasedMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFileBasedMetadata(x uint64) (n int) {
	return sovFileBasedMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileBasedMetadataConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileBasedMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileBasedMetadataConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileBasedMetadataConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecretData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SecretData == nil {
				m.SecretData = &core.DataSource{}
			}
			if err := m.SecretData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileBasedMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFileBasedMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFileBasedMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFileBasedMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileBasedMetadata
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFileBasedMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFileBasedMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFileBasedMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFileBasedMetadata
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFileBasedMetadata(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFileBasedMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileBasedMetadata   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v2alpha/file_based_metadata.proto", fileDescriptor_file_based_metadata_d58d604cef8b60b9)
}

var fileDescriptor_file_based_metadata_d58d604cef8b60b9 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0xc7, 0x89, 0x82, 0x72, 0x39, 0x5d, 0xba, 0x58, 0xc4, 0x2b, 0x87, 0x3a, 0xdc, 0x94, 0x40,
	0xdc, 0x6f, 0x38, 0x45, 0x10, 0x11, 0xe4, 0xdc, 0x5c, 0xc2, 0x77, 0xe9, 0xd7, 0xbb, 0x60, 0x6d,
	0x42, 0x2e, 0x96, 0xf6, 0x61, 0x7c, 0x1f, 0x47, 0x1f, 0x41, 0xfa, 0x24, 0x92, 0xa4, 0x93, 0x93,
	0x5b, 0xf9, 0x35, 0xbf, 0xdf, 0x9f, 0x8f, 0x2e, 0xb1, 0x69, 0x4d, 0xcf, 0x95, 0x69, 0x2a, 0xbd,
	0xe5, 0x5b, 0x67, 0x95, 0x54, 0x0e, 0x4b, 0x6c, 0xbc, 0x86, 0x9a, 0xb7, 0x02, 0x6a, 0xbb, 0x03,
	0x5e, 0xe9, 0x1a, 0xe5, 0x06, 0xf6, 0x58, 0xca, 0x77, 0xf4, 0x50, 0x82, 0x07, 0x66, 0x9d, 0xf1,
	0x26, 0xbb, 0x8e, 0x3e, 0x4b, 0x3e, 0xfb, 0xe3, 0xb3, 0xd1, 0x3f, 0xbf, 0x48, 0x2b, 0x60, 0x35,
	0x6f, 0x05, 0x57, 0xc6, 0x21, 0x0f, 0xb5, 0xd4, 0xb8, 0xfc, 0x24, 0xf4, 0xec, 0x5e, 0xd7, 0xb8,
	0x0a, 0x03, 0x4f, 0x63, 0xff, 0x36, 0x16, 0xb3, 0x25, 0x9d, 0xee, 0x51, 0x39, 0xf4, 0x32, 0xc0,
	0x9c, 0xcc, 0xc9, 0x62, 0x2a, 0x66, 0x2c, 0xad, 0x82, 0xd5, 0xac, 0x15, 0x2c, 0xf4, 0xd8, 0x1d,
	0x78, 0x78, 0x31, 0x1f, 0x4e, 0xe1, 0x9a, 0x26, 0x23, 0x90, 0x6c, 0x46, 0xe9, 0x0e, 0xa1, 0x44,
	0x27, 0xdf, 0xb0, 0xcf, 0x0f, 0xe6, 0x64, 0x31, 0x59, 0x4f, 0x12, 0x79, 0xc4, 0x3e, 0xbb, 0xa2,
	0xa7, 0xe3, 0x6f, 0xeb, 0xb0, 0xd2, 0x5d, 0x7e, 0x18, 0x5f, 0x9c, 0x24, 0xf8, 0x1c, 0xd9, 0xea,
	0xe1, 0x6b, 0x28, 0xc8, 0xf7, 0x50, 0x90, 0x9f, 0xa1, 0x20, 0x54, 0x68, 0x93, 0xe6, 0xad, 0x33,
	0x5d, 0xcf, 0xfe, 0x73, 0xff, 0xeb, 0xf1, 0xf8, 0xb1, 0x39, 0x8a, 0x17, 0xdf, 0xfc, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xb3, 0x12, 0x47, 0xb3, 0x77, 0x01, 0x00, 0x00,
}