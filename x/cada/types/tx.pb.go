// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cada/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BlobStatus defines the statuses for a blob submission
type BlobStatus int32

const (
	// Indicates that the blob status is unspecified or not set.
	BLOB_STATUS_UNSPECIFIED BlobStatus = 0
	// Indicates that the blob submission failed.
	BLOB_STATUS_FAILURE BlobStatus = 1
	// Indicates that the blob submission was successful.
	BLOB_STATUS_SUCCESS BlobStatus = 2
	// Indicates that the blob submission is still pending and has not yet been processed.
	BLOB_STATUS_PENDING BlobStatus = 3
)

var BlobStatus_name = map[int32]string{
	0: "BLOB_STATUS_UNSPECIFIED",
	1: "BLOB_STATUS_FAILURE",
	2: "BLOB_STATUS_SUCCESS",
	3: "BLOB_STATUS_PENDING",
}

var BlobStatus_value = map[string]int32{
	"BLOB_STATUS_UNSPECIFIED": 0,
	"BLOB_STATUS_FAILURE":     1,
	"BLOB_STATUS_SUCCESS":     2,
	"BLOB_STATUS_PENDING":     3,
}

func (x BlobStatus) String() string {
	return proto.EnumName(BlobStatus_name, int32(x))
}

func (BlobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c917a06c648c4256, []int{0}
}

// Range defines the range of blocks for which the blob is being submitted.
type Range struct {
	// The starting block height in the range. Indicates the beginning of the block range.
	From uint64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	// The ending block height in the range. Indicates the end of the block range.
	To uint64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_c917a06c648c4256, []int{0}
}
func (m *Range) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Range.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return m.Size()
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetFrom() uint64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Range) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

// MsgUpdateBlobStatusRequest define a message to update the status of a previously submitted blob.
type MsgUpdateBlobStatusRequest struct {
	// Address of the validator updating the blob status.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// range of blocks for which the blob status is being updated.
	BlocksRange *Range `protobuf:"bytes,2,opt,name=blocks_range,json=blocksRange,proto3" json:"blocks_range,omitempty"`
	// The height at which the blob is stored in the Avail system. This indicates where the blob data is available.
	AvailHeight uint64 `protobuf:"varint,3,opt,name=avail_height,json=availHeight,proto3" json:"avail_height,omitempty"`
	// The status of the blob submission.
	IsSuccess bool `protobuf:"varint,4,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (m *MsgUpdateBlobStatusRequest) Reset()         { *m = MsgUpdateBlobStatusRequest{} }
func (m *MsgUpdateBlobStatusRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateBlobStatusRequest) ProtoMessage()    {}
func (*MsgUpdateBlobStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c917a06c648c4256, []int{1}
}
func (m *MsgUpdateBlobStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateBlobStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateBlobStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateBlobStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateBlobStatusRequest.Merge(m, src)
}
func (m *MsgUpdateBlobStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateBlobStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateBlobStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateBlobStatusRequest proto.InternalMessageInfo

func (m *MsgUpdateBlobStatusRequest) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *MsgUpdateBlobStatusRequest) GetBlocksRange() *Range {
	if m != nil {
		return m.BlocksRange
	}
	return nil
}

func (m *MsgUpdateBlobStatusRequest) GetAvailHeight() uint64 {
	if m != nil {
		return m.AvailHeight
	}
	return 0
}

func (m *MsgUpdateBlobStatusRequest) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

// MsgUpdateBlobStatusResponse is the response type for the Msg/UpdateBlobStatus RPC method.
type MsgUpdateBlobStatusResponse struct {
}

func (m *MsgUpdateBlobStatusResponse) Reset()         { *m = MsgUpdateBlobStatusResponse{} }
func (m *MsgUpdateBlobStatusResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateBlobStatusResponse) ProtoMessage()    {}
func (*MsgUpdateBlobStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c917a06c648c4256, []int{2}
}
func (m *MsgUpdateBlobStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateBlobStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateBlobStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateBlobStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateBlobStatusResponse.Merge(m, src)
}
func (m *MsgUpdateBlobStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateBlobStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateBlobStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateBlobStatusResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cada.v1beta1.BlobStatus", BlobStatus_name, BlobStatus_value)
	proto.RegisterType((*Range)(nil), "cada.v1beta1.Range")
	proto.RegisterType((*MsgUpdateBlobStatusRequest)(nil), "cada.v1beta1.MsgUpdateBlobStatusRequest")
	proto.RegisterType((*MsgUpdateBlobStatusResponse)(nil), "cada.v1beta1.MsgUpdateBlobStatusResponse")
}

func init() { proto.RegisterFile("cada/v1beta1/tx.proto", fileDescriptor_c917a06c648c4256) }

var fileDescriptor_c917a06c648c4256 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x49, 0x8a, 0xe8, 0x25, 0x42, 0xe6, 0x0a, 0x24, 0x72, 0x55, 0xab, 0x64, 0x0a,
	0xa9, 0xea, 0x53, 0x8a, 0xc4, 0xc0, 0x96, 0xa4, 0x2e, 0x8d, 0xd4, 0x86, 0xca, 0xae, 0x17, 0x16,
	0xeb, 0x6c, 0x1f, 0x17, 0x0b, 0xbb, 0x17, 0x7c, 0x67, 0x53, 0x98, 0x10, 0x13, 0x23, 0xdf, 0x81,
	0x2f, 0xd0, 0x8f, 0xc1, 0xd8, 0x11, 0x89, 0x05, 0x25, 0x43, 0xbf, 0x06, 0xf2, 0x19, 0x41, 0x09,
	0x20, 0x75, 0xf2, 0xd3, 0xef, 0xff, 0xe4, 0xf7, 0xff, 0xdf, 0x7b, 0xf0, 0x7e, 0x48, 0x22, 0x82,
	0x8b, 0x41, 0x40, 0x25, 0x19, 0x60, 0x79, 0x6e, 0xcd, 0x33, 0x2e, 0x39, 0x6a, 0x95, 0xd8, 0xfa,
	0x89, 0x8d, 0x76, 0xc8, 0x45, 0xca, 0x05, 0x4e, 0x05, 0xc3, 0xc5, 0xa0, 0xfc, 0x54, 0x6d, 0xc6,
	0x3d, 0xc6, 0x19, 0x57, 0x25, 0x2e, 0xab, 0x8a, 0x76, 0x77, 0xe0, 0x9a, 0x43, 0xce, 0x18, 0x45,
	0x08, 0x36, 0x5e, 0x66, 0x3c, 0xed, 0x80, 0x6d, 0xd0, 0x6b, 0x38, 0xaa, 0x46, 0x77, 0x60, 0x4d,
	0xf2, 0x4e, 0x4d, 0x91, 0x9a, 0xe4, 0xdd, 0x6f, 0x00, 0x1a, 0xc7, 0x82, 0x79, 0xf3, 0x88, 0x48,
	0x3a, 0x4a, 0x78, 0xe0, 0x4a, 0x22, 0x73, 0xe1, 0xd0, 0xd7, 0x39, 0x15, 0x12, 0xed, 0xc0, 0xbb,
	0x05, 0x49, 0xe2, 0x88, 0x48, 0x9e, 0xf9, 0x24, 0x8a, 0x32, 0x2a, 0x84, 0xfa, 0xdf, 0xba, 0xa3,
	0xff, 0x12, 0x86, 0x15, 0x47, 0x4f, 0x60, 0x2b, 0x48, 0x78, 0xf8, 0x4a, 0xf8, 0x59, 0x39, 0x5f,
	0x4d, 0x69, 0xee, 0x6d, 0x58, 0xd7, 0xc3, 0x58, 0xca, 0x9a, 0xd3, 0xac, 0x1a, 0x2b, 0x9f, 0x0f,
	0x61, 0x8b, 0x14, 0x24, 0x4e, 0xfc, 0x19, 0x8d, 0xd9, 0x4c, 0x76, 0xea, 0xca, 0x5d, 0x53, 0xb1,
	0x43, 0x85, 0xd0, 0x16, 0x84, 0xb1, 0xf0, 0x45, 0x1e, 0x86, 0xa5, 0x81, 0xc6, 0x36, 0xe8, 0xdd,
	0x76, 0xd6, 0x63, 0xe1, 0x56, 0xe0, 0xe9, 0x83, 0x0f, 0x57, 0x17, 0xfd, 0xbf, 0x9d, 0x76, 0xb7,
	0xe0, 0xe6, 0x3f, 0xc3, 0x89, 0x39, 0x3f, 0x13, 0xb4, 0xff, 0x0e, 0xc2, 0xdf, 0x14, 0x6d, 0xc2,
	0xf6, 0xe8, 0xe8, 0xf9, 0xc8, 0x77, 0x4f, 0x87, 0xa7, 0x9e, 0xeb, 0x7b, 0x53, 0xf7, 0xc4, 0x1e,
	0x4f, 0x0e, 0x26, 0xf6, 0xbe, 0xae, 0xa1, 0x36, 0xdc, 0xb8, 0x2e, 0x1e, 0x0c, 0x27, 0x47, 0x9e,
	0x63, 0xeb, 0x60, 0x55, 0x70, 0xbd, 0xf1, 0xd8, 0x76, 0x5d, 0xbd, 0xb6, 0x2a, 0x9c, 0xd8, 0xd3,
	0xfd, 0xc9, 0xf4, 0x99, 0x5e, 0x37, 0x1a, 0x1f, 0x3f, 0x9b, 0xda, 0x5e, 0x0e, 0xeb, 0xc7, 0x82,
	0x21, 0x06, 0xf5, 0x55, 0x7b, 0xa8, 0xf7, 0xe7, 0x8b, 0xfd, 0x7f, 0x3d, 0xc6, 0xa3, 0x1b, 0x74,
	0x56, 0x59, 0x8d, 0xb5, 0xf7, 0x57, 0x17, 0x7d, 0x30, 0x3a, 0xfc, 0xb2, 0x30, 0xc1, 0xe5, 0xc2,
	0x04, 0xdf, 0x17, 0x26, 0xf8, 0xb4, 0x34, 0xb5, 0xcb, 0xa5, 0xa9, 0x7d, 0x5d, 0x9a, 0xda, 0x0b,
	0x8b, 0xc5, 0x72, 0x96, 0x07, 0x56, 0xc8, 0x53, 0x5c, 0xc4, 0xf2, 0x4d, 0x2c, 0xb1, 0xda, 0xc0,
	0x6e, 0x44, 0x76, 0x53, 0x1e, 0xe5, 0x09, 0xc5, 0xe7, 0x58, 0x9d, 0xab, 0x7c, 0x3b, 0xa7, 0x22,
	0xb8, 0xa5, 0xae, 0xed, 0xf1, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x43, 0x13, 0x63, 0xc3,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateBlobStatus updates the status of a blob submission.
	UpdateBlobStatus(ctx context.Context, in *MsgUpdateBlobStatusRequest, opts ...grpc.CallOption) (*MsgUpdateBlobStatusResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateBlobStatus(ctx context.Context, in *MsgUpdateBlobStatusRequest, opts ...grpc.CallOption) (*MsgUpdateBlobStatusResponse, error) {
	out := new(MsgUpdateBlobStatusResponse)
	err := c.cc.Invoke(ctx, "/cada.v1beta1.Msg/UpdateBlobStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateBlobStatus updates the status of a blob submission.
	UpdateBlobStatus(context.Context, *MsgUpdateBlobStatusRequest) (*MsgUpdateBlobStatusResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateBlobStatus(ctx context.Context, req *MsgUpdateBlobStatusRequest) (*MsgUpdateBlobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlobStatus not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateBlobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateBlobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateBlobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cada.v1beta1.Msg/UpdateBlobStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateBlobStatus(ctx, req.(*MsgUpdateBlobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cada.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateBlobStatus",
			Handler:    _Msg_UpdateBlobStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cada/v1beta1/tx.proto",
}

func (m *Range) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Range) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Range) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.To != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.To))
		i--
		dAtA[i] = 0x10
	}
	if m.From != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.From))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateBlobStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateBlobStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateBlobStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsSuccess {
		i--
		if m.IsSuccess {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.AvailHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AvailHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.BlocksRange != nil {
		{
			size, err := m.BlocksRange.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateBlobStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateBlobStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateBlobStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Range) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != 0 {
		n += 1 + sovTx(uint64(m.From))
	}
	if m.To != 0 {
		n += 1 + sovTx(uint64(m.To))
	}
	return n
}

func (m *MsgUpdateBlobStatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BlocksRange != nil {
		l = m.BlocksRange.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AvailHeight != 0 {
		n += 1 + sovTx(uint64(m.AvailHeight))
	}
	if m.IsSuccess {
		n += 2
	}
	return n
}

func (m *MsgUpdateBlobStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Range) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Range: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Range: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			m.From = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.From |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			m.To = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.To |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateBlobStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateBlobStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateBlobStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlocksRange == nil {
				m.BlocksRange = &Range{}
			}
			if err := m.BlocksRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvailHeight", wireType)
			}
			m.AvailHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AvailHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSuccess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsSuccess = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateBlobStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateBlobStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateBlobStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
