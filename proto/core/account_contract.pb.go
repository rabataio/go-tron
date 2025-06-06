//
// java-tron is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// java-tron is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: core/contract/account_contract.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountCreateContract struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OwnerAddress   []byte                 `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	AccountAddress []byte                 `protobuf:"bytes,2,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	Type           AccountType            `protobuf:"varint,3,opt,name=type,proto3,enum=protocol.AccountType" json:"type,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *AccountCreateContract) Reset() {
	*x = AccountCreateContract{}
	mi := &file_core_contract_account_contract_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountCreateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountCreateContract) ProtoMessage() {}

func (x *AccountCreateContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_account_contract_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountCreateContract.ProtoReflect.Descriptor instead.
func (*AccountCreateContract) Descriptor() ([]byte, []int) {
	return file_core_contract_account_contract_proto_rawDescGZIP(), []int{0}
}

func (x *AccountCreateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *AccountCreateContract) GetAccountAddress() []byte {
	if x != nil {
		return x.AccountAddress
	}
	return nil
}

func (x *AccountCreateContract) GetType() AccountType {
	if x != nil {
		return x.Type
	}
	return AccountType_Normal
}

// Update account name. Account name is not unique now.
type AccountUpdateContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountName   []byte                 `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	OwnerAddress  []byte                 `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountUpdateContract) Reset() {
	*x = AccountUpdateContract{}
	mi := &file_core_contract_account_contract_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountUpdateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUpdateContract) ProtoMessage() {}

func (x *AccountUpdateContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_account_contract_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUpdateContract.ProtoReflect.Descriptor instead.
func (*AccountUpdateContract) Descriptor() ([]byte, []int) {
	return file_core_contract_account_contract_proto_rawDescGZIP(), []int{1}
}

func (x *AccountUpdateContract) GetAccountName() []byte {
	if x != nil {
		return x.AccountName
	}
	return nil
}

func (x *AccountUpdateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

// Set account id if the account has no id. Account id is unique and case insensitive.
type SetAccountIdContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     []byte                 `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	OwnerAddress  []byte                 `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetAccountIdContract) Reset() {
	*x = SetAccountIdContract{}
	mi := &file_core_contract_account_contract_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAccountIdContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAccountIdContract) ProtoMessage() {}

func (x *SetAccountIdContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_account_contract_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAccountIdContract.ProtoReflect.Descriptor instead.
func (*SetAccountIdContract) Descriptor() ([]byte, []int) {
	return file_core_contract_account_contract_proto_rawDescGZIP(), []int{2}
}

func (x *SetAccountIdContract) GetAccountId() []byte {
	if x != nil {
		return x.AccountId
	}
	return nil
}

func (x *SetAccountIdContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

type AccountPermissionUpdateContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerAddress  []byte                 `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Owner         *Permission            `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`     //Empty is invalidate
	Witness       *Permission            `protobuf:"bytes,3,opt,name=witness,proto3" json:"witness,omitempty"` //Can be empty
	Actives       []*Permission          `protobuf:"bytes,4,rep,name=actives,proto3" json:"actives,omitempty"` //Empty is invalidate
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountPermissionUpdateContract) Reset() {
	*x = AccountPermissionUpdateContract{}
	mi := &file_core_contract_account_contract_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountPermissionUpdateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPermissionUpdateContract) ProtoMessage() {}

func (x *AccountPermissionUpdateContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_account_contract_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPermissionUpdateContract.ProtoReflect.Descriptor instead.
func (*AccountPermissionUpdateContract) Descriptor() ([]byte, []int) {
	return file_core_contract_account_contract_proto_rawDescGZIP(), []int{3}
}

func (x *AccountPermissionUpdateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *AccountPermissionUpdateContract) GetOwner() *Permission {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *AccountPermissionUpdateContract) GetWitness() *Permission {
	if x != nil {
		return x.Witness
	}
	return nil
}

func (x *AccountPermissionUpdateContract) GetActives() []*Permission {
	if x != nil {
		return x.Actives
	}
	return nil
}

var File_core_contract_account_contract_proto protoreflect.FileDescriptor

var file_core_contract_account_contract_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x1a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x54, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x90, 0x01, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x5f, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5a, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0xd2, 0x01, 0x0a, 0x1f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x77,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x42, 0x45, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x72, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_core_contract_account_contract_proto_rawDescOnce sync.Once
	file_core_contract_account_contract_proto_rawDescData []byte
)

func file_core_contract_account_contract_proto_rawDescGZIP() []byte {
	file_core_contract_account_contract_proto_rawDescOnce.Do(func() {
		file_core_contract_account_contract_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_core_contract_account_contract_proto_rawDesc), len(file_core_contract_account_contract_proto_rawDesc)))
	})
	return file_core_contract_account_contract_proto_rawDescData
}

var file_core_contract_account_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_core_contract_account_contract_proto_goTypes = []any{
	(*AccountCreateContract)(nil),           // 0: protocol.AccountCreateContract
	(*AccountUpdateContract)(nil),           // 1: protocol.AccountUpdateContract
	(*SetAccountIdContract)(nil),            // 2: protocol.SetAccountIdContract
	(*AccountPermissionUpdateContract)(nil), // 3: protocol.AccountPermissionUpdateContract
	(AccountType)(0),                        // 4: protocol.AccountType
	(*Permission)(nil),                      // 5: protocol.Permission
}
var file_core_contract_account_contract_proto_depIdxs = []int32{
	4, // 0: protocol.AccountCreateContract.type:type_name -> protocol.AccountType
	5, // 1: protocol.AccountPermissionUpdateContract.owner:type_name -> protocol.Permission
	5, // 2: protocol.AccountPermissionUpdateContract.witness:type_name -> protocol.Permission
	5, // 3: protocol.AccountPermissionUpdateContract.actives:type_name -> protocol.Permission
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_core_contract_account_contract_proto_init() }
func file_core_contract_account_contract_proto_init() {
	if File_core_contract_account_contract_proto != nil {
		return
	}
	file_core_Tron_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_core_contract_account_contract_proto_rawDesc), len(file_core_contract_account_contract_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_contract_account_contract_proto_goTypes,
		DependencyIndexes: file_core_contract_account_contract_proto_depIdxs,
		MessageInfos:      file_core_contract_account_contract_proto_msgTypes,
	}.Build()
	File_core_contract_account_contract_proto = out.File
	file_core_contract_account_contract_proto_goTypes = nil
	file_core_contract_account_contract_proto_depIdxs = nil
}
