// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/daoying/pb/appearence.proto

package pb

import (
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

type GetUAppearenceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUAppearenceReq) Reset() {
	*x = GetUAppearenceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_daoying_pb_appearence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUAppearenceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUAppearenceReq) ProtoMessage() {}

func (x *GetUAppearenceReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_daoying_pb_appearence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUAppearenceReq.ProtoReflect.Descriptor instead.
func (*GetUAppearenceReq) Descriptor() ([]byte, []int) {
	return file_api_daoying_pb_appearence_proto_rawDescGZIP(), []int{0}
}

type Appearence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc        string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Cost        int32  `protobuf:"varint,4,opt,name=cost,proto3" json:"cost,omitempty"`
	Url         string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	VersionName string `protobuf:"bytes,6,opt,name=versionName,proto3" json:"versionName,omitempty"`
	Version     int32  `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Appearence) Reset() {
	*x = Appearence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_daoying_pb_appearence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Appearence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Appearence) ProtoMessage() {}

func (x *Appearence) ProtoReflect() protoreflect.Message {
	mi := &file_api_daoying_pb_appearence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Appearence.ProtoReflect.Descriptor instead.
func (*Appearence) Descriptor() ([]byte, []int) {
	return file_api_daoying_pb_appearence_proto_rawDescGZIP(), []int{1}
}

func (x *Appearence) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Appearence) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Appearence) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Appearence) GetCost() int32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Appearence) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Appearence) GetVersionName() string {
	if x != nil {
		return x.VersionName
	}
	return ""
}

func (x *Appearence) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type GetUAppearenceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locked   []*Appearence `protobuf:"bytes,1,rep,name=locked,proto3" json:"locked,omitempty"`
	Unlocked []*Appearence `protobuf:"bytes,2,rep,name=unlocked,proto3" json:"unlocked,omitempty"`
}

func (x *GetUAppearenceResp) Reset() {
	*x = GetUAppearenceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_daoying_pb_appearence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUAppearenceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUAppearenceResp) ProtoMessage() {}

func (x *GetUAppearenceResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_daoying_pb_appearence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUAppearenceResp.ProtoReflect.Descriptor instead.
func (*GetUAppearenceResp) Descriptor() ([]byte, []int) {
	return file_api_daoying_pb_appearence_proto_rawDescGZIP(), []int{2}
}

func (x *GetUAppearenceResp) GetLocked() []*Appearence {
	if x != nil {
		return x.Locked
	}
	return nil
}

func (x *GetUAppearenceResp) GetUnlocked() []*Appearence {
	if x != nil {
		return x.Unlocked
	}
	return nil
}

type UpdateUAppearenceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUAppearenceReq) Reset() {
	*x = UpdateUAppearenceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_daoying_pb_appearence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUAppearenceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUAppearenceReq) ProtoMessage() {}

func (x *UpdateUAppearenceReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_daoying_pb_appearence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUAppearenceReq.ProtoReflect.Descriptor instead.
func (*UpdateUAppearenceReq) Descriptor() ([]byte, []int) {
	return file_api_daoying_pb_appearence_proto_rawDescGZIP(), []int{3}
}

type UpdateUAppearenceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUAppearenceResp) Reset() {
	*x = UpdateUAppearenceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_daoying_pb_appearence_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUAppearenceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUAppearenceResp) ProtoMessage() {}

func (x *UpdateUAppearenceResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_daoying_pb_appearence_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUAppearenceResp.ProtoReflect.Descriptor instead.
func (*UpdateUAppearenceResp) Descriptor() ([]byte, []int) {
	return file_api_daoying_pb_appearence_proto_rawDescGZIP(), []int{4}
}

var File_api_daoying_pb_appearence_proto protoreflect.FileDescriptor

var file_api_daoying_pb_appearence_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x6f, 0x79, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x62,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x64, 0x61, 0x6f, 0x79, 0x69, 0x6e, 0x67, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x55, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x22,
	0xaa, 0x01, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x55, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x61, 0x6f, 0x79, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x70, 0x70,
	0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12,
	0x2f, 0x0a, 0x08, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x64, 0x61, 0x6f, 0x79, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x65,
	0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x41, 0x70, 0x70, 0x65, 0x61,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x32, 0xb2, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x41,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x6f, 0x79,
	0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x6f, 0x79, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x52, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x41, 0x70, 0x70,
	0x65, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x6f, 0x79, 0x69, 0x6e,
	0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x64, 0x61, 0x6f, 0x79, 0x69, 0x6e, 0x67,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61,
	0x6f, 0x79, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_daoying_pb_appearence_proto_rawDescOnce sync.Once
	file_api_daoying_pb_appearence_proto_rawDescData = file_api_daoying_pb_appearence_proto_rawDesc
)

func file_api_daoying_pb_appearence_proto_rawDescGZIP() []byte {
	file_api_daoying_pb_appearence_proto_rawDescOnce.Do(func() {
		file_api_daoying_pb_appearence_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_daoying_pb_appearence_proto_rawDescData)
	})
	return file_api_daoying_pb_appearence_proto_rawDescData
}

var file_api_daoying_pb_appearence_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_daoying_pb_appearence_proto_goTypes = []interface{}{
	(*GetUAppearenceReq)(nil),     // 0: daoying.GetUAppearenceReq
	(*Appearence)(nil),            // 1: daoying.Appearence
	(*GetUAppearenceResp)(nil),    // 2: daoying.GetUAppearenceResp
	(*UpdateUAppearenceReq)(nil),  // 3: daoying.UpdateUAppearenceReq
	(*UpdateUAppearenceResp)(nil), // 4: daoying.UpdateUAppearenceResp
}
var file_api_daoying_pb_appearence_proto_depIdxs = []int32{
	1, // 0: daoying.GetUAppearenceResp.locked:type_name -> daoying.Appearence
	1, // 1: daoying.GetUAppearenceResp.unlocked:type_name -> daoying.Appearence
	0, // 2: daoying.AppearenceService.GetUAppearence:input_type -> daoying.GetUAppearenceReq
	3, // 3: daoying.AppearenceService.UpdateUAppearence:input_type -> daoying.UpdateUAppearenceReq
	2, // 4: daoying.AppearenceService.GetUAppearence:output_type -> daoying.GetUAppearenceResp
	4, // 5: daoying.AppearenceService.UpdateUAppearence:output_type -> daoying.UpdateUAppearenceResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_daoying_pb_appearence_proto_init() }
func file_api_daoying_pb_appearence_proto_init() {
	if File_api_daoying_pb_appearence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_daoying_pb_appearence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUAppearenceReq); i {
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
		file_api_daoying_pb_appearence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Appearence); i {
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
		file_api_daoying_pb_appearence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUAppearenceResp); i {
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
		file_api_daoying_pb_appearence_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUAppearenceReq); i {
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
		file_api_daoying_pb_appearence_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUAppearenceResp); i {
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
			RawDescriptor: file_api_daoying_pb_appearence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_daoying_pb_appearence_proto_goTypes,
		DependencyIndexes: file_api_daoying_pb_appearence_proto_depIdxs,
		MessageInfos:      file_api_daoying_pb_appearence_proto_msgTypes,
	}.Build()
	File_api_daoying_pb_appearence_proto = out.File
	file_api_daoying_pb_appearence_proto_rawDesc = nil
	file_api_daoying_pb_appearence_proto_goTypes = nil
	file_api_daoying_pb_appearence_proto_depIdxs = nil
}
