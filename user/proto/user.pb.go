// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Birthday             int64    `protobuf:"varint,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Birthday             int64    `protobuf:"varint,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Nickname             string   `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginResponse) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

func (m *LoginResponse) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

// 返回用户id
type ResponseUserId struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseUserId) Reset()         { *m = ResponseUserId{} }
func (m *ResponseUserId) String() string { return proto.CompactTextString(m) }
func (*ResponseUserId) ProtoMessage()    {}
func (*ResponseUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{3}
}

func (m *ResponseUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseUserId.Unmarshal(m, b)
}
func (m *ResponseUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseUserId.Marshal(b, m, deterministic)
}
func (m *ResponseUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseUserId.Merge(m, src)
}
func (m *ResponseUserId) XXX_Size() int {
	return xxx_messageInfo_ResponseUserId.Size(m)
}
func (m *ResponseUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseUserId.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseUserId proto.InternalMessageInfo

func (m *ResponseUserId) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	// 响应是否成功
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// 响应信息
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "com.weiguo.srv.user.User")
	proto.RegisterType((*LoginRequest)(nil), "com.weiguo.srv.user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "com.weiguo.srv.user.LoginResponse")
	proto.RegisterType((*ResponseUserId)(nil), "com.weiguo.srv.user.ResponseUserId")
	proto.RegisterType((*Response)(nil), "com.weiguo.srv.user.Response")
}

func init() {
	proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74)
}

var fileDescriptor_ed89022014131a74 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x75, 0x9b, 0x36, 0xb6, 0xe3, 0x27, 0xab, 0x48, 0x2c, 0x08, 0x75, 0x4f, 0x3d, 0x45, 0x50,
	0xf0, 0x0f, 0x08, 0xa2, 0xe0, 0x29, 0xda, 0x1f, 0x90, 0xa6, 0x43, 0x1b, 0x34, 0xd9, 0xb8, 0x93,
	0xb4, 0xf8, 0x77, 0xbc, 0x78, 0xf2, 0x3f, 0xca, 0x6c, 0x3e, 0x30, 0x90, 0xf6, 0x12, 0xf2, 0xf6,
	0xbd, 0xbc, 0xbc, 0x37, 0x3b, 0x70, 0x52, 0x10, 0x9a, 0x1b, 0x7e, 0xf8, 0x99, 0xd1, 0xb9, 0x96,
	0x67, 0x91, 0x4e, 0xfc, 0x0d, 0xc6, 0xcb, 0x42, 0xfb, 0x64, 0xd6, 0x3e, 0x53, 0xea, 0x5b, 0x40,
	0x7f, 0x46, 0x68, 0xe4, 0x18, 0x86, 0x69, 0x1c, 0xbd, 0xa7, 0x61, 0x82, 0x9e, 0x98, 0x88, 0xe9,
	0x28, 0x68, 0x30, 0x73, 0x2c, 0xb6, 0x5c, 0xaf, 0xe4, 0x6a, 0xcc, 0x5c, 0x16, 0x12, 0x6d, 0xb4,
	0x59, 0x78, 0x4e, 0xc9, 0xd5, 0x58, 0x9e, 0xc3, 0x20, 0x5b, 0xe9, 0x14, 0xbd, 0xbe, 0x25, 0x4a,
	0xc0, 0xa7, 0x98, 0x84, 0xf1, 0x87, 0x37, 0x28, 0x4f, 0x2d, 0x60, 0x9f, 0x79, 0x6c, 0xf2, 0xd5,
	0x22, 0xfc, 0xf2, 0xdc, 0x89, 0x98, 0x3a, 0x41, 0x83, 0xd5, 0x23, 0x1c, 0xbe, 0xe8, 0x65, 0x9c,
	0x06, 0xf8, 0x59, 0x20, 0xe5, 0xad, 0x3c, 0x62, 0x47, 0x9e, 0x5e, 0x3b, 0x8f, 0xfa, 0x11, 0x70,
	0x54, 0x19, 0x51, 0xa6, 0x53, 0x42, 0x79, 0x01, 0x2e, 0x7f, 0xf9, 0xbc, 0xa8, 0x7c, 0x2a, 0xb4,
	0xb3, 0x71, 0xd3, 0xca, 0xe9, 0x6c, 0xd5, 0xdf, 0xd6, 0x6a, 0xd0, 0x6e, 0xd5, 0x9a, 0xb8, 0xdb,
	0x9e, 0xb8, 0x9a, 0xc2, 0x71, 0x9d, 0x71, 0x56, 0x26, 0xda, 0x92, 0x54, 0xdd, 0xc3, 0xb0, 0x69,
	0xe3, 0xc1, 0x3e, 0x15, 0x51, 0x84, 0x44, 0x56, 0x34, 0x0c, 0x6a, 0x28, 0x4f, 0xc1, 0x49, 0x68,
	0x59, 0x55, 0xe1, 0xd7, 0xdb, 0x5f, 0x01, 0x07, 0x6c, 0xfd, 0x8a, 0x66, 0x1d, 0x47, 0x28, 0x9f,
	0x00, 0x1e, 0x0c, 0x86, 0xb9, 0xfd, 0x9f, 0xbc, 0xf4, 0x3b, 0x96, 0xc5, 0x67, 0x6a, 0x7c, 0xd5,
	0x49, 0xd5, 0x19, 0xd4, 0x9e, 0x7c, 0x83, 0x11, 0x0b, 0xed, 0xa0, 0xe5, 0x75, 0xa7, 0xfa, 0xff,
	0x6d, 0x8e, 0xd5, 0x2e, 0x49, 0xed, 0x3a, 0x77, 0xed, 0x12, 0xdf, 0xfd, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x38, 0xae, 0x32, 0xcc, 0xd7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	// 创建一个用户
	CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 用户登录
	UserLogin(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.weiguo.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.CreateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.UserLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	// 创建一个用户
	CreateUser(context.Context, *User, *Response) error
	// 用户登录
	UserLogin(context.Context, *LoginRequest, *LoginResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) CreateUser(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.CreateUser(ctx, in, out)
}

func (h *UserService) UserLogin(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserServiceHandler.UserLogin(ctx, in, out)
}
