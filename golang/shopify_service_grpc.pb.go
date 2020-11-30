// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShopifyClient is the client API for Shopify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopifyClient interface {
	GetProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*Products, error)
	GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error)
}

type shopifyClient struct {
	cc grpc.ClientConnInterface
}

func NewShopifyClient(cc grpc.ClientConnInterface) ShopifyClient {
	return &shopifyClient{cc}
}

func (c *shopifyClient) GetProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, "/shopify.Shopify/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopifyClient) GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/shopify.Shopify/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopifyServer is the server API for Shopify service.
// All implementations must embed UnimplementedShopifyServer
// for forward compatibility
type ShopifyServer interface {
	GetProducts(context.Context, *ProductsRequest) (*Products, error)
	GetProduct(context.Context, *ProductRequest) (*Product, error)
	mustEmbedUnimplementedShopifyServer()
}

// UnimplementedShopifyServer must be embedded to have forward compatible implementations.
type UnimplementedShopifyServer struct {
}

func (*UnimplementedShopifyServer) GetProducts(context.Context, *ProductsRequest) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (*UnimplementedShopifyServer) GetProduct(context.Context, *ProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (*UnimplementedShopifyServer) mustEmbedUnimplementedShopifyServer() {}

func RegisterShopifyServer(s *grpc.Server, srv ShopifyServer) {
	s.RegisterService(&_Shopify_serviceDesc, srv)
}

func _Shopify_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopifyServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shopify.Shopify/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopifyServer).GetProducts(ctx, req.(*ProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shopify_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopifyServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shopify.Shopify/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopifyServer).GetProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shopify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shopify.Shopify",
	HandlerType: (*ShopifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _Shopify_GetProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Shopify_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shopify_service.proto",
}