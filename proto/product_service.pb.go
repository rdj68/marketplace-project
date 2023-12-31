// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/product_service.proto

package proto

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

type ProductId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *ProductId) Reset() {
	*x = ProductId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductId) ProtoMessage() {}

func (x *ProductId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductId.ProtoReflect.Descriptor instead.
func (*ProductId) Descriptor() ([]byte, []int) {
	return file_proto_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProductId) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type NewProductData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category    string            `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Description string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       int32             `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"` //last two digits represent paisa
	ImageURL    string            `protobuf:"bytes,6,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	ShopId      string            `protobuf:"bytes,7,opt,name=shopId,proto3" json:"shopId,omitempty"`
	Attributes  map[string]string `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NewProductData) Reset() {
	*x = NewProductData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProductData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProductData) ProtoMessage() {}

func (x *NewProductData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProductData.ProtoReflect.Descriptor instead.
func (*NewProductData) Descriptor() ([]byte, []int) {
	return file_proto_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewProductData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewProductData) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *NewProductData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewProductData) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *NewProductData) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *NewProductData) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *NewProductData) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type ProductData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category    string            `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Description string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       int32             `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"` //last two digits represent paisa
	ImageURL    string            `protobuf:"bytes,6,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	ShopId      string            `protobuf:"bytes,7,opt,name=shopId,proto3" json:"shopId,omitempty"`
	Attributes  map[string]string `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProductData) Reset() {
	*x = ProductData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductData) ProtoMessage() {}

func (x *ProductData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductData.ProtoReflect.Descriptor instead.
func (*ProductData) Descriptor() ([]byte, []int) {
	return file_proto_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProductData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductData) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProductData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductData) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductData) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *ProductData) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *ProductData) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type ProductSearchParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MaxPrice int32  `protobuf:"varint,2,opt,name=maxPrice,proto3" json:"maxPrice,omitempty"`
	MinPrice int32  `protobuf:"varint,3,opt,name=minPrice,proto3" json:"minPrice,omitempty"`
	Category string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *ProductSearchParam) Reset() {
	*x = ProductSearchParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchParam) ProtoMessage() {}

func (x *ProductSearchParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchParam.ProtoReflect.Descriptor instead.
func (*ProductSearchParam) Descriptor() ([]byte, []int) {
	return file_proto_product_service_proto_rawDescGZIP(), []int{3}
}

func (x *ProductSearchParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductSearchParam) GetMaxPrice() int32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *ProductSearchParam) GetMinPrice() int32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *ProductSearchParam) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type MultipleProductData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductData []*ProductData `protobuf:"bytes,1,rep,name=productData,proto3" json:"productData,omitempty"`
}

func (x *MultipleProductData) Reset() {
	*x = MultipleProductData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleProductData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleProductData) ProtoMessage() {}

func (x *MultipleProductData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleProductData.ProtoReflect.Descriptor instead.
func (*MultipleProductData) Descriptor() ([]byte, []int) {
	return file_proto_product_service_proto_rawDescGZIP(), []int{4}
}

func (x *MultipleProductData) GetProductData() []*ProductData {
	if x != nil {
		return x.ProductData
	}
	return nil
}

var File_proto_product_service_proto protoreflect.FileDescriptor

var file_proto_product_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22,
	0xb2, 0x02, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x45,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xbc, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64,
	0x12, 0x42, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x7c, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x22, 0x4b, 0x0a, 0x13, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x32, 0x8d,
	0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x64, 0x6a,
	0x36, 0x38, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_service_proto_rawDescOnce sync.Once
	file_proto_product_service_proto_rawDescData = file_proto_product_service_proto_rawDesc
)

func file_proto_product_service_proto_rawDescGZIP() []byte {
	file_proto_product_service_proto_rawDescOnce.Do(func() {
		file_proto_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_service_proto_rawDescData)
	})
	return file_proto_product_service_proto_rawDescData
}

var file_proto_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_product_service_proto_goTypes = []interface{}{
	(*ProductId)(nil),           // 0: proto.ProductId
	(*NewProductData)(nil),      // 1: proto.NewProductData
	(*ProductData)(nil),         // 2: proto.ProductData
	(*ProductSearchParam)(nil),  // 3: proto.ProductSearchParam
	(*MultipleProductData)(nil), // 4: proto.MultipleProductData
	nil,                         // 5: proto.NewProductData.AttributesEntry
	nil,                         // 6: proto.ProductData.AttributesEntry
}
var file_proto_product_service_proto_depIdxs = []int32{
	5, // 0: proto.NewProductData.attributes:type_name -> proto.NewProductData.AttributesEntry
	6, // 1: proto.ProductData.attributes:type_name -> proto.ProductData.AttributesEntry
	2, // 2: proto.MultipleProductData.productData:type_name -> proto.ProductData
	1, // 3: proto.ProductService.AddProduct:input_type -> proto.NewProductData
	3, // 4: proto.ProductService.FindProduct:input_type -> proto.ProductSearchParam
	0, // 5: proto.ProductService.AddProduct:output_type -> proto.ProductId
	4, // 6: proto.ProductService.FindProduct:output_type -> proto.MultipleProductData
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_product_service_proto_init() }
func file_proto_product_service_proto_init() {
	if File_proto_product_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductId); i {
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
		file_proto_product_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProductData); i {
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
		file_proto_product_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductData); i {
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
		file_proto_product_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchParam); i {
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
		file_proto_product_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleProductData); i {
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
			RawDescriptor: file_proto_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_service_proto_goTypes,
		DependencyIndexes: file_proto_product_service_proto_depIdxs,
		MessageInfos:      file_proto_product_service_proto_msgTypes,
	}.Build()
	File_proto_product_service_proto = out.File
	file_proto_product_service_proto_rawDesc = nil
	file_proto_product_service_proto_goTypes = nil
	file_proto_product_service_proto_depIdxs = nil
}
