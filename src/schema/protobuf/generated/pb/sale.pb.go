// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: sale.proto

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

type Sale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId   string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SupplierId  string `protobuf:"bytes,3,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	Quantity    string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TotalAmount int64  `protobuf:"varint,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Date        string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	IsDeleted   int64  `protobuf:"varint,7,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *Sale) Reset() {
	*x = Sale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sale) ProtoMessage() {}

func (x *Sale) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sale.ProtoReflect.Descriptor instead.
func (*Sale) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{0}
}

func (x *Sale) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sale) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Sale) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *Sale) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *Sale) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Sale) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Sale) GetIsDeleted() int64 {
	if x != nil {
		return x.IsDeleted
	}
	return 0
}

type CreateSaleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SupplierId  string `protobuf:"bytes,2,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	TotalAmount int64  `protobuf:"varint,3,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
}

func (x *CreateSaleRequest) Reset() {
	*x = CreateSaleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSaleRequest) ProtoMessage() {}

func (x *CreateSaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSaleRequest.ProtoReflect.Descriptor instead.
func (*CreateSaleRequest) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSaleRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateSaleRequest) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *CreateSaleRequest) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

type UpdateSaleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId   string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SupplierId  string `protobuf:"bytes,3,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	TotalAmount int64  `protobuf:"varint,4,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
}

func (x *UpdateSaleRequest) Reset() {
	*x = UpdateSaleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSaleRequest) ProtoMessage() {}

func (x *UpdateSaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSaleRequest.ProtoReflect.Descriptor instead.
func (*UpdateSaleRequest) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSaleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSaleRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateSaleRequest) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *UpdateSaleRequest) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

type DeleteSaleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSaleRequest) Reset() {
	*x = DeleteSaleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSaleRequest) ProtoMessage() {}

func (x *DeleteSaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSaleRequest.ProtoReflect.Descriptor instead.
func (*DeleteSaleRequest) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSaleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSaleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSaleRequest) Reset() {
	*x = GetSaleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSaleRequest) ProtoMessage() {}

func (x *GetSaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSaleRequest.ProtoReflect.Descriptor instead.
func (*GetSaleRequest) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{4}
}

func (x *GetSaleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SaleResponsePagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Data       []*Sale     `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SaleResponsePagination) Reset() {
	*x = SaleResponsePagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleResponsePagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleResponsePagination) ProtoMessage() {}

func (x *SaleResponsePagination) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleResponsePagination.ProtoReflect.Descriptor instead.
func (*SaleResponsePagination) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{5}
}

func (x *SaleResponsePagination) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *SaleResponsePagination) GetData() []*Sale {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_sale_proto protoreflect.FileDescriptor

var file_sale_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x04, 0x53, 0x61, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x22, 0x76, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x16, 0x53,
	0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0xbe, 0x02, 0x0a, 0x0b, 0x53, 0x61, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x49, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x61, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x61, 0x6c, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sale_proto_rawDescOnce sync.Once
	file_sale_proto_rawDescData = file_sale_proto_rawDesc
)

func file_sale_proto_rawDescGZIP() []byte {
	file_sale_proto_rawDescOnce.Do(func() {
		file_sale_proto_rawDescData = protoimpl.X.CompressGZIP(file_sale_proto_rawDescData)
	})
	return file_sale_proto_rawDescData
}

var file_sale_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sale_proto_goTypes = []interface{}{
	(*Sale)(nil),                   // 0: protobuf.Sale
	(*CreateSaleRequest)(nil),      // 1: protobuf.CreateSaleRequest
	(*UpdateSaleRequest)(nil),      // 2: protobuf.UpdateSaleRequest
	(*DeleteSaleRequest)(nil),      // 3: protobuf.DeleteSaleRequest
	(*GetSaleRequest)(nil),         // 4: protobuf.GetSaleRequest
	(*SaleResponsePagination)(nil), // 5: protobuf.SaleResponsePagination
	(*Pagination)(nil),             // 6: protobuf.Pagination
	(*PaginationRequest)(nil),      // 7: protobuf.PaginationRequest
}
var file_sale_proto_depIdxs = []int32{
	6, // 0: protobuf.SaleResponsePagination.pagination:type_name -> protobuf.Pagination
	0, // 1: protobuf.SaleResponsePagination.data:type_name -> protobuf.Sale
	7, // 2: protobuf.SaleService.GetSales:input_type -> protobuf.PaginationRequest
	4, // 3: protobuf.SaleService.GetSale:input_type -> protobuf.GetSaleRequest
	1, // 4: protobuf.SaleService.CreateSale:input_type -> protobuf.CreateSaleRequest
	2, // 5: protobuf.SaleService.UpdateSale:input_type -> protobuf.UpdateSaleRequest
	3, // 6: protobuf.SaleService.DeleteSale:input_type -> protobuf.DeleteSaleRequest
	5, // 7: protobuf.SaleService.GetSales:output_type -> protobuf.SaleResponsePagination
	0, // 8: protobuf.SaleService.GetSale:output_type -> protobuf.Sale
	0, // 9: protobuf.SaleService.CreateSale:output_type -> protobuf.Sale
	0, // 10: protobuf.SaleService.UpdateSale:output_type -> protobuf.Sale
	0, // 11: protobuf.SaleService.DeleteSale:output_type -> protobuf.Sale
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sale_proto_init() }
func file_sale_proto_init() {
	if File_sale_proto != nil {
		return
	}
	file_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sale_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sale); i {
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
		file_sale_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSaleRequest); i {
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
		file_sale_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSaleRequest); i {
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
		file_sale_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSaleRequest); i {
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
		file_sale_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSaleRequest); i {
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
		file_sale_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleResponsePagination); i {
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
			RawDescriptor: file_sale_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sale_proto_goTypes,
		DependencyIndexes: file_sale_proto_depIdxs,
		MessageInfos:      file_sale_proto_msgTypes,
	}.Build()
	File_sale_proto = out.File
	file_sale_proto_rawDesc = nil
	file_sale_proto_goTypes = nil
	file_sale_proto_depIdxs = nil
}
