// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.17.3
// source: proto/order.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type OrderFindReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderFindReq) Reset() {
	*x = OrderFindReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFindReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFindReq) ProtoMessage() {}

func (x *OrderFindReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFindReq.ProtoReflect.Descriptor instead.
func (*OrderFindReq) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderFindReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserOrderFindReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserOrderFindReq) Reset() {
	*x = UserOrderFindReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrderFindReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrderFindReq) ProtoMessage() {}

func (x *UserOrderFindReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrderFindReq.ProtoReflect.Descriptor instead.
func (*UserOrderFindReq) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *UserOrderFindReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Product string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	UserId  int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *Order) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *Orders) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_proto_order_proto protoreflect.FileDescriptor

var file_proto_order_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4a, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x06,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0x5f, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0d, 0x66, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x07,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData = file_proto_order_proto_rawDesc
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_proto_rawDescData)
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_order_proto_goTypes = []interface{}{
	(*OrderFindReq)(nil),     // 0: OrderFindReq
	(*UserOrderFindReq)(nil), // 1: UserOrderFindReq
	(*Order)(nil),            // 2: Order
	(*Orders)(nil),           // 3: Orders
}
var file_proto_order_proto_depIdxs = []int32{
	2, // 0: Orders.orders:type_name -> Order
	0, // 1: OrderService.findOrder:input_type -> OrderFindReq
	1, // 2: OrderService.findUserOrder:input_type -> UserOrderFindReq
	2, // 3: OrderService.findOrder:output_type -> Order
	3, // 4: OrderService.findUserOrder:output_type -> Orders
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderFindReq); i {
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
		file_proto_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrderFindReq); i {
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
		file_proto_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_proto_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
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
			RawDescriptor: file_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_rawDesc = nil
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	FindOrder(ctx context.Context, in *OrderFindReq, opts ...grpc.CallOption) (*Order, error)
	FindUserOrder(ctx context.Context, in *UserOrderFindReq, opts ...grpc.CallOption) (*Orders, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) FindOrder(ctx context.Context, in *OrderFindReq, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/OrderService/findOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) FindUserOrder(ctx context.Context, in *UserOrderFindReq, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/OrderService/findUserOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	FindOrder(context.Context, *OrderFindReq) (*Order, error)
	FindUserOrder(context.Context, *UserOrderFindReq) (*Orders, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) FindOrder(context.Context, *OrderFindReq) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrder not implemented")
}
func (*UnimplementedOrderServiceServer) FindUserOrder(context.Context, *UserOrderFindReq) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserOrder not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_FindOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).FindOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/FindOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).FindOrder(ctx, req.(*OrderFindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_FindUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserOrderFindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).FindUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/FindUserOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).FindUserOrder(ctx, req.(*UserOrderFindReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findOrder",
			Handler:    _OrderService_FindOrder_Handler,
		},
		{
			MethodName: "findUserOrder",
			Handler:    _OrderService_FindUserOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}