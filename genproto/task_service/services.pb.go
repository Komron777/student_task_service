// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task_service/services.proto

package task

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("task_service/services.proto", fileDescriptor_5fc1ae02c8860a9c) }

var fileDescriptor_5fc1ae02c8860a9c = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x49, 0x2c, 0xce,
	0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x87, 0xd2, 0xc5, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x2c, 0x20, 0x49, 0x29, 0x71, 0x14, 0x25, 0x20, 0x0e, 0x44, 0xda, 0x68, 0x1d,
	0x13, 0x17, 0x77, 0x48, 0x62, 0x71, 0x76, 0x30, 0x44, 0x4a, 0xc8, 0x8c, 0x8b, 0xcb, 0xb9, 0x28,
	0x35, 0xb1, 0x24, 0x15, 0x24, 0x28, 0x24, 0xac, 0x07, 0x56, 0x8a, 0x10, 0x09, 0x4a, 0x2d, 0x94,
	0xc2, 0x22, 0x58, 0x2c, 0xa4, 0xce, 0xc5, 0xee, 0x9e, 0x5a, 0x02, 0xd6, 0xc4, 0x05, 0x91, 0x77,
	0xaa, 0xf4, 0x4c, 0x91, 0x12, 0x80, 0xb0, 0xa1, 0x52, 0x20, 0x85, 0x06, 0x5c, 0x5c, 0xa1, 0x05,
	0x29, 0x68, 0x16, 0x20, 0x44, 0x40, 0x16, 0xf0, 0x42, 0x04, 0x83, 0x4b, 0x93, 0x93, 0x53, 0x8b,
	0x41, 0x46, 0x73, 0xb9, 0xa4, 0xe6, 0xa4, 0x42, 0x75, 0x20, 0x9b, 0x8e, 0xa6, 0x50, 0x8f, 0x8b,
	0xdb, 0x27, 0xb3, 0xb8, 0xc4, 0xbf, 0x2c, 0xb5, 0xc8, 0xa5, 0x34, 0x55, 0x88, 0x1b, 0x22, 0xeb,
	0x9a, 0x5b, 0x50, 0x52, 0x29, 0x25, 0x04, 0xe1, 0x80, 0xe4, 0x41, 0xc6, 0x14, 0x83, 0x9c, 0x62,
	0xcc, 0xc5, 0x09, 0xe7, 0x0b, 0x61, 0x2a, 0x28, 0xc4, 0xa6, 0xc9, 0x49, 0xe0, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d,
	0x1c, 0x92, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x8e, 0xa4, 0x77, 0x87, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskRes, error)
	GetTask(ctx context.Context, in *ById, opts ...grpc.CallOption) (*GetTaskRes, error)
	UpdateTask(ctx context.Context, in *UpdateTaskReq, opts ...grpc.CallOption) (*Success, error)
	DeleteTask(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Success, error)
	ListOverDue(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTasksRes, error)
	ListTasks(ctx context.Context, in *ListTasksReq, opts ...grpc.CallOption) (*ListTasksRes, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskRes, error) {
	out := new(CreateTaskRes)
	err := c.cc.Invoke(ctx, "/task.TaskService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *ById, opts ...grpc.CallOption) (*GetTaskRes, error) {
	out := new(GetTaskRes)
	err := c.cc.Invoke(ctx, "/task.TaskService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskReq, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/task.TaskService/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteTask(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/task.TaskService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListOverDue(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTasksRes, error) {
	out := new(ListTasksRes)
	err := c.cc.Invoke(ctx, "/task.TaskService/ListOverDue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListTasks(ctx context.Context, in *ListTasksReq, opts ...grpc.CallOption) (*ListTasksRes, error) {
	out := new(ListTasksRes)
	err := c.cc.Invoke(ctx, "/task.TaskService/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	CreateTask(context.Context, *CreateTaskReq) (*CreateTaskRes, error)
	GetTask(context.Context, *ById) (*GetTaskRes, error)
	UpdateTask(context.Context, *UpdateTaskReq) (*Success, error)
	DeleteTask(context.Context, *ById) (*Success, error)
	ListOverDue(context.Context, *Empty) (*ListTasksRes, error)
	ListTasks(context.Context, *ListTasksReq) (*ListTasksRes, error)
}

// UnimplementedTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) CreateTask(ctx context.Context, req *CreateTaskReq) (*CreateTaskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (*UnimplementedTaskServiceServer) GetTask(ctx context.Context, req *ById) (*GetTaskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedTaskServiceServer) UpdateTask(ctx context.Context, req *UpdateTaskReq) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (*UnimplementedTaskServiceServer) DeleteTask(ctx context.Context, req *ById) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (*UnimplementedTaskServiceServer) ListOverDue(ctx context.Context, req *Empty) (*ListTasksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOverDue not implemented")
}
func (*UnimplementedTaskServiceServer) ListTasks(ctx context.Context, req *ListTasksReq) (*ListTasksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*CreateTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateTask(ctx, req.(*UpdateTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteTask(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListOverDue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListOverDue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/ListOverDue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListOverDue(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListTasks(ctx, req.(*ListTasksReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskService_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TaskService_DeleteTask_Handler,
		},
		{
			MethodName: "ListOverDue",
			Handler:    _TaskService_ListOverDue_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _TaskService_ListTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task_service/services.proto",
}
