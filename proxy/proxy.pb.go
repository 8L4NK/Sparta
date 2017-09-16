// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy.proto

/*
Package proxy is a generated protocol buffer package.

It is generated from these files:
	proxy.proto

It has these top-level messages:
	CognitoIdentity
	LambdaContext
	ProxyRequest
	ProxyResponse
*/
package proxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// CognitoIdentity is the identity information from a Cognito auth
// See http://docs.aws.amazon.com/lambda/latest/dg/nodejs-prog-model-context.html
type CognitoIdentity struct {
	CognitoIdentityId     string `protobuf:"bytes,1,opt,name=cognito_identity_id,json=cognitoIdentityId" json:"cognito_identity_id,omitempty"`
	CognitoIdentityPoolId string `protobuf:"bytes,2,opt,name=cognito_identity_pool_id,json=cognitoIdentityPoolId" json:"cognito_identity_pool_id,omitempty"`
}

func (m *CognitoIdentity) Reset()                    { *m = CognitoIdentity{} }
func (m *CognitoIdentity) String() string            { return proto.CompactTextString(m) }
func (*CognitoIdentity) ProtoMessage()               {}
func (*CognitoIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CognitoIdentity) GetCognitoIdentityId() string {
	if m != nil {
		return m.CognitoIdentityId
	}
	return ""
}

func (m *CognitoIdentity) GetCognitoIdentityPoolId() string {
	if m != nil {
		return m.CognitoIdentityPoolId
	}
	return ""
}

// LambdaContext defines the AWS Lambda Context object provided by the AWS Lambda runtime.
// See http://docs.aws.amazon.com/lambda/latest/dg/nodejs-prog-model-context.html
// for more information on field values.  Note that the golang version doesn't functions
// defined on the Context object.
type LambdaContext struct {
	FunctionName       string           `protobuf:"bytes,1,opt,name=function_name,json=functionName" json:"function_name,omitempty"`
	FunctionVersion    string           `protobuf:"bytes,2,opt,name=function_version,json=functionVersion" json:"function_version,omitempty"`
	InvokedFunctionArn string           `protobuf:"bytes,3,opt,name=invoked_function_arn,json=invokedFunctionArn" json:"invoked_function_arn,omitempty"`
	MemoryLimitInMb    string           `protobuf:"bytes,4,opt,name=memory_limit_in_mb,json=memoryLimitInMb" json:"memory_limit_in_mb,omitempty"`
	AwsRequestId       string           `protobuf:"bytes,5,opt,name=aws_request_id,json=awsRequestId" json:"aws_request_id,omitempty"`
	LogGroupName       string           `protobuf:"bytes,6,opt,name=log_group_name,json=logGroupName" json:"log_group_name,omitempty"`
	LogStreamName      string           `protobuf:"bytes,7,opt,name=log_stream_name,json=logStreamName" json:"log_stream_name,omitempty"`
	Identity           *CognitoIdentity `protobuf:"bytes,8,opt,name=identity" json:"identity,omitempty"`
}

func (m *LambdaContext) Reset()                    { *m = LambdaContext{} }
func (m *LambdaContext) String() string            { return proto.CompactTextString(m) }
func (*LambdaContext) ProtoMessage()               {}
func (*LambdaContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LambdaContext) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *LambdaContext) GetFunctionVersion() string {
	if m != nil {
		return m.FunctionVersion
	}
	return ""
}

func (m *LambdaContext) GetInvokedFunctionArn() string {
	if m != nil {
		return m.InvokedFunctionArn
	}
	return ""
}

func (m *LambdaContext) GetMemoryLimitInMb() string {
	if m != nil {
		return m.MemoryLimitInMb
	}
	return ""
}

func (m *LambdaContext) GetAwsRequestId() string {
	if m != nil {
		return m.AwsRequestId
	}
	return ""
}

func (m *LambdaContext) GetLogGroupName() string {
	if m != nil {
		return m.LogGroupName
	}
	return ""
}

func (m *LambdaContext) GetLogStreamName() string {
	if m != nil {
		return m.LogStreamName
	}
	return ""
}

func (m *LambdaContext) GetIdentity() *CognitoIdentity {
	if m != nil {
		return m.Identity
	}
	return nil
}

// ProxyRequest is the request made by the NodeJS layer to the
// golang process
type ProxyRequest struct {
	Context *LambdaContext `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Event   []byte         `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (m *ProxyRequest) Reset()                    { *m = ProxyRequest{} }
func (m *ProxyRequest) String() string            { return proto.CompactTextString(m) }
func (*ProxyRequest) ProtoMessage()               {}
func (*ProxyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProxyRequest) GetContext() *LambdaContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *ProxyRequest) GetEvent() []byte {
	if m != nil {
		return m.Event
	}
	return nil
}

// ProxyResponse is what the Go process sends back
type ProxyResponse struct {
	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *ProxyResponse) Reset()                    { *m = ProxyResponse{} }
func (m *ProxyResponse) String() string            { return proto.CompactTextString(m) }
func (*ProxyResponse) ProtoMessage()               {}
func (*ProxyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProxyResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*CognitoIdentity)(nil), "proxy.CognitoIdentity")
	proto.RegisterType((*LambdaContext)(nil), "proxy.LambdaContext")
	proto.RegisterType((*ProxyRequest)(nil), "proxy.ProxyRequest")
	proto.RegisterType((*ProxyResponse)(nil), "proxy.ProxyResponse")
}

func init() { proto.RegisterFile("proxy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x5f, 0x8b, 0xd4, 0x30,
	0x10, 0xc0, 0xd9, 0xf3, 0xf6, 0xee, 0x9c, 0x6d, 0x5d, 0x8d, 0xab, 0xf4, 0xf1, 0xe8, 0x89, 0xac,
	0x08, 0x45, 0xd6, 0x07, 0x9f, 0x65, 0x41, 0x29, 0xac, 0xb2, 0x54, 0xf1, 0x35, 0xa4, 0x4d, 0x2c,
	0xc1, 0x26, 0x53, 0xd3, 0xec, 0x9f, 0xfa, 0x81, 0xfd, 0x1c, 0xd2, 0x24, 0x5d, 0xd8, 0xf5, 0xad,
	0xf9, 0xcd, 0x6f, 0x26, 0x99, 0x99, 0xc2, 0xac, 0x35, 0x78, 0xec, 0xb3, 0xd6, 0xa0, 0x45, 0x32,
	0x75, 0x87, 0xf4, 0x0f, 0xcc, 0xd7, 0x58, 0x6b, 0x69, 0x31, 0xe7, 0x42, 0x5b, 0x69, 0x7b, 0x92,
	0xc1, 0xf3, 0xca, 0x23, 0x2a, 0x03, 0xa3, 0x92, 0x27, 0x93, 0xfb, 0xc9, 0xf2, 0x71, 0xf1, 0xac,
	0x3a, 0xb7, 0x73, 0x4e, 0x3e, 0x40, 0xf2, 0x9f, 0xdf, 0x22, 0x36, 0x43, 0xd2, 0x95, 0x4b, 0x7a,
	0x71, 0x91, 0xb4, 0x45, 0x6c, 0x72, 0x9e, 0xfe, 0xbd, 0x82, 0x78, 0xc3, 0x54, 0xc9, 0xd9, 0x1a,
	0xb5, 0x15, 0x47, 0x4b, 0x1e, 0x20, 0xfe, 0xb9, 0xd3, 0x95, 0x95, 0xa8, 0xa9, 0x66, 0x4a, 0x84,
	0x4b, 0xa3, 0x11, 0x7e, 0x65, 0x4a, 0x90, 0x37, 0xf0, 0xf4, 0x24, 0xed, 0x85, 0xe9, 0x24, 0xea,
	0x70, 0xcf, 0x7c, 0xe4, 0x3f, 0x3c, 0x26, 0xef, 0x60, 0x21, 0xf5, 0x1e, 0x7f, 0x09, 0x4e, 0x4f,
	0x29, 0xcc, 0xe8, 0xe4, 0x91, 0xd3, 0x49, 0x88, 0x7d, 0x0a, 0xa1, 0x8f, 0x46, 0x93, 0xb7, 0x40,
	0x94, 0x50, 0x68, 0x7a, 0xda, 0x48, 0x25, 0x2d, 0x95, 0x9a, 0xaa, 0x32, 0xb9, 0xf6, 0xe5, 0x7d,
	0x64, 0x33, 0x04, 0x72, 0xfd, 0xa5, 0x24, 0xaf, 0xe0, 0x09, 0x3b, 0x74, 0xd4, 0x88, 0xdf, 0x3b,
	0xd1, 0xd9, 0xa1, 0xdf, 0xa9, 0x7f, 0x2f, 0x3b, 0x74, 0x85, 0x87, 0x39, 0x1f, 0xac, 0x06, 0x6b,
	0x5a, 0x1b, 0xdc, 0xb5, 0xbe, 0xab, 0x1b, 0x6f, 0x35, 0x58, 0x7f, 0x1e, 0xa0, 0xeb, 0xea, 0x35,
	0xcc, 0x07, 0xab, 0xb3, 0x46, 0x30, 0xe5, 0xb5, 0x5b, 0xa7, 0xc5, 0x0d, 0xd6, 0xdf, 0x1c, 0x75,
	0xde, 0x0a, 0xee, 0xc6, 0x29, 0x27, 0x77, 0xf7, 0x93, 0xe5, 0x6c, 0xf5, 0x32, 0xf3, 0x7b, 0xbd,
	0xd8, 0x63, 0x71, 0xf2, 0xd2, 0xef, 0x10, 0x6d, 0x07, 0x25, 0xbc, 0x89, 0x64, 0x70, 0x5b, 0xf9,
	0x89, 0xbb, 0x01, 0xcf, 0x56, 0x8b, 0x50, 0xe2, 0x6c, 0x1b, 0xc5, 0x28, 0x91, 0x05, 0x4c, 0xc5,
	0x5e, 0x68, 0xeb, 0xc6, 0x1c, 0x15, 0xfe, 0x90, 0x3e, 0x40, 0x1c, 0xaa, 0x76, 0x2d, 0xea, 0x4e,
	0x10, 0x02, 0xd7, 0x25, 0xf2, 0xde, 0xd5, 0x8c, 0x0a, 0xf7, 0x5d, 0xde, 0xb8, 0xbf, 0xed, 0xfd,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xdc, 0x17, 0x68, 0x7c, 0x02, 0x00, 0x00,
}
