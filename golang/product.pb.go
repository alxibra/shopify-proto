// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: product.proto

package shopify

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title             string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	BodyHtml          string             `protobuf:"bytes,3,opt,name=body_html,json=bodyHtml,proto3" json:"body_html,omitempty"`
	Vendor            string             `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor,omitempty"`
	ProductType       string             `protobuf:"bytes,5,opt,name=product_type,json=productType,proto3" json:"product_type,omitempty"`
	CreatedAt         string             `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Handle            string             `protobuf:"bytes,7,opt,name=handle,proto3" json:"handle,omitempty"`
	UpdatedAt         string             `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PublishedAt       string             `protobuf:"bytes,9,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	TemplateSuffix    string             `protobuf:"bytes,10,opt,name=template_suffix,json=templateSuffix,proto3" json:"template_suffix,omitempty"`
	Status            string             `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	PublishedScope    string             `protobuf:"bytes,12,opt,name=published_scope,json=publishedScope,proto3" json:"published_scope,omitempty"`
	Tags              string             `protobuf:"bytes,13,opt,name=tags,proto3" json:"tags,omitempty"`
	AdminGraphqlApiId string             `protobuf:"bytes,14,opt,name=admin_graphql_api_id,json=adminGraphqlApiId,proto3" json:"admin_graphql_api_id,omitempty"`
	Variants          []*Product_Variant `protobuf:"bytes,15,rep,name=variants,proto3" json:"variants,omitempty"`
	Options           []*Product_Option  `protobuf:"bytes,16,rep,name=options,proto3" json:"options,omitempty"`
	Images            []*Product_Image   `protobuf:"bytes,17,rep,name=images,proto3" json:"images,omitempty"`
	Image             *Product_Image     `protobuf:"bytes,18,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Product) GetBodyHtml() string {
	if x != nil {
		return x.BodyHtml
	}
	return ""
}

func (x *Product) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *Product) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

func (x *Product) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Product) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

func (x *Product) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Product) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

func (x *Product) GetTemplateSuffix() string {
	if x != nil {
		return x.TemplateSuffix
	}
	return ""
}

func (x *Product) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Product) GetPublishedScope() string {
	if x != nil {
		return x.PublishedScope
	}
	return ""
}

func (x *Product) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Product) GetAdminGraphqlApiId() string {
	if x != nil {
		return x.AdminGraphqlApiId
	}
	return ""
}

func (x *Product) GetVariants() []*Product_Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *Product) GetOptions() []*Product_Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Product) GetImages() []*Product_Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Product) GetImage() *Product_Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type Product_Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price                string  `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Sku                  string  `protobuf:"bytes,4,opt,name=sku,proto3" json:"sku,omitempty"`
	Position             uint32  `protobuf:"varint,5,opt,name=position,proto3" json:"position,omitempty"`
	InventoryPolicy      string  `protobuf:"bytes,6,opt,name=inventory_policy,json=inventoryPolicy,proto3" json:"inventory_policy,omitempty"`
	CompareAtPrice       string  `protobuf:"bytes,7,opt,name=compare_at_price,json=compareAtPrice,proto3" json:"compare_at_price,omitempty"`
	FulfillmentService   string  `protobuf:"bytes,8,opt,name=fulfillment_service,json=fulfillmentService,proto3" json:"fulfillment_service,omitempty"`
	InventoryManagement  string  `protobuf:"bytes,9,opt,name=inventory_management,json=inventoryManagement,proto3" json:"inventory_management,omitempty"`
	Option1              string  `protobuf:"bytes,10,opt,name=option1,proto3" json:"option1,omitempty"`
	Option2              string  `protobuf:"bytes,11,opt,name=option2,proto3" json:"option2,omitempty"`
	Option3              string  `protobuf:"bytes,12,opt,name=option3,proto3" json:"option3,omitempty"`
	CreatedAt            string  `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string  `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Taxable              bool    `protobuf:"varint,15,opt,name=taxable,proto3" json:"taxable,omitempty"`
	Barcode              string  `protobuf:"bytes,16,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Grams                uint32  `protobuf:"varint,17,opt,name=grams,proto3" json:"grams,omitempty"`
	ImageId              int64   `protobuf:"varint,18,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Weight               float64 `protobuf:"fixed64,19,opt,name=weight,proto3" json:"weight,omitempty"`
	WeightUnit           string  `protobuf:"bytes,20,opt,name=weight_unit,json=weightUnit,proto3" json:"weight_unit,omitempty"`
	InventoryItemId      int64   `protobuf:"varint,21,opt,name=inventory_item_id,json=inventoryItemId,proto3" json:"inventory_item_id,omitempty"`
	InventoryQuantity    uint32  `protobuf:"varint,22,opt,name=inventory_quantity,json=inventoryQuantity,proto3" json:"inventory_quantity,omitempty"`
	OldInventoryQuantity uint32  `protobuf:"varint,23,opt,name=old_inventory_quantity,json=oldInventoryQuantity,proto3" json:"old_inventory_quantity,omitempty"`
	RequiresShipping     bool    `protobuf:"varint,24,opt,name=requires_shipping,json=requiresShipping,proto3" json:"requires_shipping,omitempty"`
	AdminGraphqlApiId    string  `protobuf:"bytes,25,opt,name=admin_graphql_api_id,json=adminGraphqlApiId,proto3" json:"admin_graphql_api_id,omitempty"`
}

func (x *Product_Variant) Reset() {
	*x = Product_Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_Variant) ProtoMessage() {}

func (x *Product_Variant) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_Variant.ProtoReflect.Descriptor instead.
func (*Product_Variant) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Product_Variant) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product_Variant) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Product_Variant) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Product_Variant) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Product_Variant) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Product_Variant) GetInventoryPolicy() string {
	if x != nil {
		return x.InventoryPolicy
	}
	return ""
}

func (x *Product_Variant) GetCompareAtPrice() string {
	if x != nil {
		return x.CompareAtPrice
	}
	return ""
}

func (x *Product_Variant) GetFulfillmentService() string {
	if x != nil {
		return x.FulfillmentService
	}
	return ""
}

func (x *Product_Variant) GetInventoryManagement() string {
	if x != nil {
		return x.InventoryManagement
	}
	return ""
}

func (x *Product_Variant) GetOption1() string {
	if x != nil {
		return x.Option1
	}
	return ""
}

func (x *Product_Variant) GetOption2() string {
	if x != nil {
		return x.Option2
	}
	return ""
}

func (x *Product_Variant) GetOption3() string {
	if x != nil {
		return x.Option3
	}
	return ""
}

func (x *Product_Variant) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Product_Variant) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Product_Variant) GetTaxable() bool {
	if x != nil {
		return x.Taxable
	}
	return false
}

func (x *Product_Variant) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *Product_Variant) GetGrams() uint32 {
	if x != nil {
		return x.Grams
	}
	return 0
}

func (x *Product_Variant) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *Product_Variant) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Product_Variant) GetWeightUnit() string {
	if x != nil {
		return x.WeightUnit
	}
	return ""
}

func (x *Product_Variant) GetInventoryItemId() int64 {
	if x != nil {
		return x.InventoryItemId
	}
	return 0
}

func (x *Product_Variant) GetInventoryQuantity() uint32 {
	if x != nil {
		return x.InventoryQuantity
	}
	return 0
}

func (x *Product_Variant) GetOldInventoryQuantity() uint32 {
	if x != nil {
		return x.OldInventoryQuantity
	}
	return 0
}

func (x *Product_Variant) GetRequiresShipping() bool {
	if x != nil {
		return x.RequiresShipping
	}
	return false
}

func (x *Product_Variant) GetAdminGraphqlApiId() string {
	if x != nil {
		return x.AdminGraphqlApiId
	}
	return ""
}

type Product_Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId uint64   `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Position  uint32   `protobuf:"varint,4,opt,name=position,proto3" json:"position,omitempty"`
	Values    []string `protobuf:"bytes,5,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Product_Option) Reset() {
	*x = Product_Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_Option) ProtoMessage() {}

func (x *Product_Option) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_Option.ProtoReflect.Descriptor instead.
func (*Product_Option) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Product_Option) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product_Option) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Product_Option) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product_Option) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Product_Option) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Product_Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Position          uint32   `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
	CreatedAt         string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt         string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Alt               string   `protobuf:"bytes,5,opt,name=alt,proto3" json:"alt,omitempty"`
	Width             uint32   `protobuf:"varint,6,opt,name=width,proto3" json:"width,omitempty"`
	Height            uint32   `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	Src               string   `protobuf:"bytes,8,opt,name=src,proto3" json:"src,omitempty"`
	VariantIds        []uint32 `protobuf:"varint,9,rep,packed,name=variant_ids,json=variantIds,proto3" json:"variant_ids,omitempty"`
	AdminGraphqlApiId string   `protobuf:"bytes,10,opt,name=admin_graphql_api_id,json=adminGraphqlApiId,proto3" json:"admin_graphql_api_id,omitempty"`
}

func (x *Product_Image) Reset() {
	*x = Product_Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_Image) ProtoMessage() {}

func (x *Product_Image) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_Image.ProtoReflect.Descriptor instead.
func (*Product_Image) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Product_Image) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product_Image) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Product_Image) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Product_Image) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Product_Image) GetAlt() string {
	if x != nil {
		return x.Alt
	}
	return ""
}

func (x *Product_Image) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Product_Image) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Product_Image) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *Product_Image) GetVariantIds() []uint32 {
	if x != nil {
		return x.VariantIds
	}
	return nil
}

func (x *Product_Image) GetAdminGraphqlApiId() string {
	if x != nil {
		return x.AdminGraphqlApiId
	}
	return ""
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x68, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x22, 0xd7, 0x0e, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f,
	0x64, 0x79, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x6f, 0x64, 0x79, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x66, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x69,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x71, 0x6c, 0x41, 0x70, 0x69, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73,
	0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x11, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x1a, 0xc5, 0x06, 0x0a, 0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x61, 0x74,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x65, 0x41, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x13,
	0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a,
	0x14, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74,
	0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x11, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x16, 0x6f, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x6f, 0x6c, 0x64, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x71, 0x6c, 0x41, 0x70, 0x69, 0x49, 0x64, 0x1a, 0x7f, 0x0a, 0x06, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x95, 0x02, 0x0a, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63,
	0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x73, 0x12, 0x2f, 0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x41, 0x70, 0x69,
	0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_product_proto_goTypes = []interface{}{
	(*Product)(nil),         // 0: shopify.Product
	(*Product_Variant)(nil), // 1: shopify.Product.Variant
	(*Product_Option)(nil),  // 2: shopify.Product.Option
	(*Product_Image)(nil),   // 3: shopify.Product.Image
}
var file_product_proto_depIdxs = []int32{
	1, // 0: shopify.Product.variants:type_name -> shopify.Product.Variant
	2, // 1: shopify.Product.options:type_name -> shopify.Product.Option
	3, // 2: shopify.Product.images:type_name -> shopify.Product.Image
	3, // 3: shopify.Product.image:type_name -> shopify.Product.Image
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_Variant); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_Option); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_Image); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
