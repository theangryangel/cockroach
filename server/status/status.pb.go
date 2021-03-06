// Code generated by protoc-gen-gogo.
// source: cockroach/server/status/status.proto
// DO NOT EDIT!

/*
	Package status is a generated protocol buffer package.

	It is generated from these files:
		cockroach/server/status/status.proto

	It has these top-level messages:
		StoreStatus
		NodeStatus
*/
package status

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"
import cockroach_util1 "github.com/cockroachdb/cockroach/util"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// StoreStatus records the most recent values of metrics for a store.
type StoreStatus struct {
	Desc    cockroach_roachpb.StoreDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc"`
	Metrics map[string]float64                `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *StoreStatus) Reset()                    { *m = StoreStatus{} }
func (m *StoreStatus) String() string            { return proto.CompactTextString(m) }
func (*StoreStatus) ProtoMessage()               {}
func (*StoreStatus) Descriptor() ([]byte, []int) { return fileDescriptorStatus, []int{0} }

// NodeStatus records the most recent values of metrics for a node.
type NodeStatus struct {
	Desc          cockroach_roachpb.NodeDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc"`
	BuildInfo     cockroach_util1.BuildInfo        `protobuf:"bytes,2,opt,name=build_info,json=buildInfo" json:"build_info"`
	StartedAt     int64                            `protobuf:"varint,3,opt,name=started_at,json=startedAt" json:"started_at"`
	UpdatedAt     int64                            `protobuf:"varint,4,opt,name=updated_at,json=updatedAt" json:"updated_at"`
	Metrics       map[string]float64               `protobuf:"bytes,5,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	StoreStatuses []StoreStatus                    `protobuf:"bytes,6,rep,name=store_statuses,json=storeStatuses" json:"store_statuses"`
}

func (m *NodeStatus) Reset()                    { *m = NodeStatus{} }
func (m *NodeStatus) String() string            { return proto.CompactTextString(m) }
func (*NodeStatus) ProtoMessage()               {}
func (*NodeStatus) Descriptor() ([]byte, []int) { return fileDescriptorStatus, []int{1} }

func init() {
	proto.RegisterType((*StoreStatus)(nil), "cockroach.server.status.StoreStatus")
	proto.RegisterType((*NodeStatus)(nil), "cockroach.server.status.NodeStatus")
}
func (m *StoreStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StoreStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintStatus(data, i, uint64(m.Desc.Size()))
	n1, err := m.Desc.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Metrics) > 0 {
		keysForMetrics := make([]string, 0, len(m.Metrics))
		for k := range m.Metrics {
			keysForMetrics = append(keysForMetrics, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForMetrics)
		for _, k := range keysForMetrics {
			data[i] = 0x12
			i++
			v := m.Metrics[string(k)]
			mapSize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			i = encodeVarintStatus(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintStatus(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x11
			i++
			i = encodeFixed64Status(data, i, uint64(math.Float64bits(float64(v))))
		}
	}
	return i, nil
}

func (m *NodeStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NodeStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintStatus(data, i, uint64(m.Desc.Size()))
	n2, err := m.Desc.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x12
	i++
	i = encodeVarintStatus(data, i, uint64(m.BuildInfo.Size()))
	n3, err := m.BuildInfo.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	data[i] = 0x18
	i++
	i = encodeVarintStatus(data, i, uint64(m.StartedAt))
	data[i] = 0x20
	i++
	i = encodeVarintStatus(data, i, uint64(m.UpdatedAt))
	if len(m.Metrics) > 0 {
		keysForMetrics := make([]string, 0, len(m.Metrics))
		for k := range m.Metrics {
			keysForMetrics = append(keysForMetrics, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForMetrics)
		for _, k := range keysForMetrics {
			data[i] = 0x2a
			i++
			v := m.Metrics[string(k)]
			mapSize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			i = encodeVarintStatus(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintStatus(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x11
			i++
			i = encodeFixed64Status(data, i, uint64(math.Float64bits(float64(v))))
		}
	}
	if len(m.StoreStatuses) > 0 {
		for _, msg := range m.StoreStatuses {
			data[i] = 0x32
			i++
			i = encodeVarintStatus(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Status(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Status(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStatus(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *StoreStatus) Size() (n int) {
	var l int
	_ = l
	l = m.Desc.Size()
	n += 1 + l + sovStatus(uint64(l))
	if len(m.Metrics) > 0 {
		for k, v := range m.Metrics {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			n += mapEntrySize + 1 + sovStatus(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *NodeStatus) Size() (n int) {
	var l int
	_ = l
	l = m.Desc.Size()
	n += 1 + l + sovStatus(uint64(l))
	l = m.BuildInfo.Size()
	n += 1 + l + sovStatus(uint64(l))
	n += 1 + sovStatus(uint64(m.StartedAt))
	n += 1 + sovStatus(uint64(m.UpdatedAt))
	if len(m.Metrics) > 0 {
		for k, v := range m.Metrics {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			n += mapEntrySize + 1 + sovStatus(uint64(mapEntrySize))
		}
	}
	if len(m.StoreStatuses) > 0 {
		for _, e := range m.StoreStatuses {
			l = e.Size()
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	return n
}

func sovStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Desc.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthStatus
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapvaluetemp uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			mapvaluetemp = uint64(data[iNdEx-8])
			mapvaluetemp |= uint64(data[iNdEx-7]) << 8
			mapvaluetemp |= uint64(data[iNdEx-6]) << 16
			mapvaluetemp |= uint64(data[iNdEx-5]) << 24
			mapvaluetemp |= uint64(data[iNdEx-4]) << 32
			mapvaluetemp |= uint64(data[iNdEx-3]) << 40
			mapvaluetemp |= uint64(data[iNdEx-2]) << 48
			mapvaluetemp |= uint64(data[iNdEx-1]) << 56
			mapvalue := math.Float64frombits(mapvaluetemp)
			if m.Metrics == nil {
				m.Metrics = make(map[string]float64)
			}
			m.Metrics[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
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
func (m *NodeStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Desc.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BuildInfo.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedAt", wireType)
			}
			m.StartedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.StartedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UpdatedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthStatus
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapvaluetemp uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			mapvaluetemp = uint64(data[iNdEx-8])
			mapvaluetemp |= uint64(data[iNdEx-7]) << 8
			mapvaluetemp |= uint64(data[iNdEx-6]) << 16
			mapvaluetemp |= uint64(data[iNdEx-5]) << 24
			mapvaluetemp |= uint64(data[iNdEx-4]) << 32
			mapvaluetemp |= uint64(data[iNdEx-3]) << 40
			mapvaluetemp |= uint64(data[iNdEx-2]) << 48
			mapvaluetemp |= uint64(data[iNdEx-1]) << 56
			mapvalue := math.Float64frombits(mapvaluetemp)
			if m.Metrics == nil {
				m.Metrics = make(map[string]float64)
			}
			m.Metrics[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreStatuses = append(m.StoreStatuses, StoreStatus{})
			if err := m.StoreStatuses[len(m.StoreStatuses)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
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
func skipStatus(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStatus
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipStatus(data[start:])
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
	ErrInvalidLengthStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatus   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorStatus = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xce, 0x4f, 0xce,
	0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0x2e, 0x49,
	0x2c, 0x29, 0x2d, 0x86, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xe2, 0x70, 0x55, 0x7a,
	0x10, 0x55, 0x7a, 0x10, 0x69, 0x29, 0x05, 0x84, 0x76, 0x30, 0x59, 0x90, 0xa4, 0x9f, 0x9b, 0x5a,
	0x92, 0x98, 0x92, 0x58, 0x92, 0x08, 0xd1, 0x2a, 0x25, 0x85, 0x50, 0x51, 0x5a, 0x92, 0x99, 0xa3,
	0x9f, 0x54, 0x9a, 0x99, 0x93, 0x02, 0x95, 0x13, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41,
	0x2c, 0x88, 0xa8, 0xd2, 0x15, 0x46, 0x2e, 0xee, 0xe0, 0x92, 0xfc, 0xa2, 0xd4, 0x60, 0xb0, 0x1d,
	0x42, 0x36, 0x5c, 0x2c, 0x29, 0xa9, 0xc5, 0xc9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x4a,
	0x7a, 0x08, 0xb7, 0x40, 0xad, 0xd4, 0x03, 0xab, 0x76, 0x01, 0xaa, 0x29, 0xca, 0x2c, 0x00, 0x32,
	0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x02, 0xeb, 0x12, 0xf2, 0xe6, 0x62, 0x07, 0xba, 0xa8,
	0x28, 0x33, 0xb9, 0x58, 0x82, 0x49, 0x81, 0x19, 0x68, 0x80, 0xa1, 0x1e, 0x0e, 0xcf, 0xe8, 0x21,
	0x59, 0xaa, 0xe7, 0x0b, 0xd1, 0xe3, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0x33, 0x41, 0xca, 0x8a,
	0x8b, 0x07, 0x59, 0x42, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x12, 0xec, 0x32, 0xce, 0x20, 0x10,
	0x53, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0x68, 0x19, 0xa3, 0x06, 0x63, 0x10,
	0x84, 0x63, 0xc5, 0x64, 0xc1, 0xa8, 0xb4, 0x8e, 0x99, 0x8b, 0xcb, 0x2f, 0x3f, 0x05, 0xe6, 0x2b,
	0x6b, 0x14, 0x5f, 0x29, 0x62, 0xf1, 0x15, 0x48, 0x31, 0x0e, 0x4f, 0xd9, 0x71, 0x71, 0x81, 0xc3,
	0x31, 0x3e, 0x33, 0x2f, 0x2d, 0x1f, 0x6c, 0x15, 0xb7, 0x91, 0x24, 0x92, 0x11, 0xa0, 0x90, 0xd6,
	0x73, 0x02, 0xa9, 0xf0, 0x04, 0x2a, 0x80, 0x6a, 0xe5, 0x4c, 0x82, 0x09, 0x08, 0x29, 0x73, 0x71,
	0x01, 0xfd, 0x5c, 0x54, 0x92, 0x9a, 0x12, 0x9f, 0x58, 0x22, 0xc1, 0x0c, 0xd4, 0xcf, 0x0c, 0x53,
	0x04, 0x15, 0x77, 0x2c, 0x01, 0x29, 0x2a, 0x2d, 0x00, 0xc6, 0x24, 0x44, 0x11, 0x0b, 0xb2, 0x22,
	0xa8, 0x38, 0x50, 0x91, 0x17, 0x22, 0x78, 0x59, 0xc1, 0xc1, 0x6b, 0x80, 0x33, 0x78, 0x11, 0x9e,
	0xc7, 0x1e, 0xba, 0x42, 0x81, 0x5c, 0x7c, 0xc5, 0xa0, 0x28, 0x88, 0x87, 0x68, 0x48, 0x2d, 0x96,
	0x60, 0x03, 0x1b, 0xa9, 0x42, 0x4c, 0x8c, 0x41, 0x9d, 0xc6, 0x5b, 0x8c, 0x10, 0x4a, 0xa5, 0x28,
	0xc2, 0x9c, 0x14, 0x4e, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0x02, 0x10, 0xdf, 0x00,
	0xe2, 0x07, 0x40, 0x3c, 0xe1, 0xb1, 0x1c, 0x43, 0x14, 0x1b, 0xc4, 0xfa, 0x08, 0x26, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x74, 0x1b, 0x33, 0x46, 0x3f, 0x03, 0x00, 0x00,
}
