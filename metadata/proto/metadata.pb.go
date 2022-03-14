// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: metadata.proto

package proto_metadata

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

type PortalShieldRequestMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    int32  `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	TokenID string `protobuf:"bytes,2,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Proof   []byte `protobuf:"bytes,4,opt,name=Proof,proto3" json:"Proof,omitempty"`
}

func (x *PortalShieldRequestMeta) Reset() {
	*x = PortalShieldRequestMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortalShieldRequestMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortalShieldRequestMeta) ProtoMessage() {}

func (x *PortalShieldRequestMeta) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortalShieldRequestMeta.ProtoReflect.Descriptor instead.
func (*PortalShieldRequestMeta) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *PortalShieldRequestMeta) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PortalShieldRequestMeta) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *PortalShieldRequestMeta) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PortalShieldRequestMeta) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type PortalSubmitConfirmedTxMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    int32  `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	TokenID string `protobuf:"bytes,2,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	BatchID string `protobuf:"bytes,3,opt,name=BatchID,proto3" json:"BatchID,omitempty"`
	Proof   []byte `protobuf:"bytes,4,opt,name=Proof,proto3" json:"Proof,omitempty"`
}

func (x *PortalSubmitConfirmedTxMeta) Reset() {
	*x = PortalSubmitConfirmedTxMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortalSubmitConfirmedTxMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortalSubmitConfirmedTxMeta) ProtoMessage() {}

func (x *PortalSubmitConfirmedTxMeta) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortalSubmitConfirmedTxMeta.ProtoReflect.Descriptor instead.
func (*PortalSubmitConfirmedTxMeta) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *PortalSubmitConfirmedTxMeta) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PortalSubmitConfirmedTxMeta) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *PortalSubmitConfirmedTxMeta) GetBatchID() string {
	if x != nil {
		return x.BatchID
	}
	return ""
}

func (x *PortalSubmitConfirmedTxMeta) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type IssuingEVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      int32    `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	BlockHash []byte   `protobuf:"bytes,2,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	TxIndex   uint64   `protobuf:"varint,3,opt,name=TxIndex,proto3" json:"TxIndex,omitempty"`
	Proofs    [][]byte `protobuf:"bytes,4,rep,name=Proofs,proto3" json:"Proofs,omitempty"`
	TokenID   []byte   `protobuf:"bytes,5,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
}

func (x *IssuingEVMRequest) Reset() {
	*x = IssuingEVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuingEVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuingEVMRequest) ProtoMessage() {}

func (x *IssuingEVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuingEVMRequest.ProtoReflect.Descriptor instead.
func (*IssuingEVMRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *IssuingEVMRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *IssuingEVMRequest) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *IssuingEVMRequest) GetTxIndex() uint64 {
	if x != nil {
		return x.TxIndex
	}
	return 0
}

func (x *IssuingEVMRequest) GetProofs() [][]byte {
	if x != nil {
		return x.Proofs
	}
	return nil
}

func (x *IssuingEVMRequest) GetTokenID() []byte {
	if x != nil {
		return x.TokenID
	}
	return nil
}

type TradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      int32             `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	SellToken []byte            `protobuf:"bytes,2,opt,name=SellToken,proto3" json:"SellToken,omitempty"`
	TradePath []string          `protobuf:"bytes,3,rep,name=TradePath,proto3" json:"TradePath,omitempty"`
	Amounts   []uint64          `protobuf:"varint,4,rep,packed,name=Amounts,proto3" json:"Amounts,omitempty"`
	Receivers map[string][]byte `protobuf:"bytes,5,rep,name=Receivers,proto3" json:"Receivers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TradeRequest) Reset() {
	*x = TradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeRequest) ProtoMessage() {}

func (x *TradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeRequest.ProtoReflect.Descriptor instead.
func (*TradeRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *TradeRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TradeRequest) GetSellToken() []byte {
	if x != nil {
		return x.SellToken
	}
	return nil
}

func (x *TradeRequest) GetTradePath() []string {
	if x != nil {
		return x.TradePath
	}
	return nil
}

func (x *TradeRequest) GetAmounts() []uint64 {
	if x != nil {
		return x.Amounts
	}
	return nil
}

func (x *TradeRequest) GetReceivers() map[string][]byte {
	if x != nil {
		return x.Receivers
	}
	return nil
}

var File_metadata_proto protoreflect.FileDescriptor

var file_metadata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x77, 0x0a, 0x17, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x7b, 0x0a, 0x1b, 0x50, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x54, 0x78, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x91, 0x01, 0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x69,
	0x6e, 0x67, 0x45, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x54, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x54, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x22, 0x81, 0x02, 0x0a, 0x0c, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x49, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x1a, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10,
	0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_proto_rawDescOnce sync.Once
	file_metadata_proto_rawDescData = file_metadata_proto_rawDesc
)

func file_metadata_proto_rawDescGZIP() []byte {
	file_metadata_proto_rawDescOnce.Do(func() {
		file_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_proto_rawDescData)
	})
	return file_metadata_proto_rawDescData
}

var file_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_metadata_proto_goTypes = []interface{}{
	(*PortalShieldRequestMeta)(nil),     // 0: proto_metadata.PortalShieldRequestMeta
	(*PortalSubmitConfirmedTxMeta)(nil), // 1: proto_metadata.PortalSubmitConfirmedTxMeta
	(*IssuingEVMRequest)(nil),           // 2: proto_metadata.IssuingEVMRequest
	(*TradeRequest)(nil),                // 3: proto_metadata.TradeRequest
	nil,                                 // 4: proto_metadata.TradeRequest.ReceiversEntry
}
var file_metadata_proto_depIdxs = []int32{
	4, // 0: proto_metadata.TradeRequest.Receivers:type_name -> proto_metadata.TradeRequest.ReceiversEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metadata_proto_init() }
func file_metadata_proto_init() {
	if File_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortalShieldRequestMeta); i {
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
		file_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortalSubmitConfirmedTxMeta); i {
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
		file_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuingEVMRequest); i {
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
		file_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeRequest); i {
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
			RawDescriptor: file_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metadata_proto_goTypes,
		DependencyIndexes: file_metadata_proto_depIdxs,
		MessageInfos:      file_metadata_proto_msgTypes,
	}.Build()
	File_metadata_proto = out.File
	file_metadata_proto_rawDesc = nil
	file_metadata_proto_goTypes = nil
	file_metadata_proto_depIdxs = nil
}
