// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema/proto/contracts/queue_contracts.proto

package contracts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EmailType int32

const (
	// welcome email type
	EmailType_welcome EmailType = 0
	// reset password email type
	EmailType_reset_password EmailType = 1
	// reset email account email type
	EmailType_reset_email EmailType = 2
	// invite code email type
	EmailType_invite_code EmailType = 3
	// system maintenance email type
	EmailType_system_maintenance EmailType = 4
	// promotional email type
	EmailType_promotional EmailType = 5
)

var EmailType_name = map[int32]string{
	0: "welcome",
	1: "reset_password",
	2: "reset_email",
	3: "invite_code",
	4: "system_maintenance",
	5: "promotional",
}

var EmailType_value = map[string]int32{
	"welcome":            0,
	"reset_password":     1,
	"reset_email":        2,
	"invite_code":        3,
	"system_maintenance": 4,
	"promotional":        5,
}

func (x EmailType) String() string {
	return proto.EnumName(EmailType_name, int32(x))
}

func (EmailType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b7c0a1623e62d66c, []int{0}
}

type EmailContract struct {
	Sender               string    `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Target               string    `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Subject              string    `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Message              string    `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Type                 EmailType `protobuf:"varint,5,opt,name=type,proto3,enum=EmailType" json:"type,omitempty"`
	Firstname            string    `protobuf:"bytes,6,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string    `protobuf:"bytes,7,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Metadata             *Metadata `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Token                *Tokens   `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EmailContract) Reset()         { *m = EmailContract{} }
func (m *EmailContract) String() string { return proto.CompactTextString(m) }
func (*EmailContract) ProtoMessage()    {}
func (*EmailContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7c0a1623e62d66c, []int{0}
}

func (m *EmailContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailContract.Unmarshal(m, b)
}
func (m *EmailContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailContract.Marshal(b, m, deterministic)
}
func (m *EmailContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailContract.Merge(m, src)
}
func (m *EmailContract) XXX_Size() int {
	return xxx_messageInfo_EmailContract.Size(m)
}
func (m *EmailContract) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailContract.DiscardUnknown(m)
}

var xxx_messageInfo_EmailContract proto.InternalMessageInfo

func (m *EmailContract) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EmailContract) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *EmailContract) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailContract) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EmailContract) GetType() EmailType {
	if m != nil {
		return m.Type
	}
	return EmailType_welcome
}

func (m *EmailContract) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *EmailContract) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *EmailContract) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EmailContract) GetToken() *Tokens {
	if m != nil {
		return m.Token
	}
	return nil
}

type Tokens struct {
	AccountActivationToken string   `protobuf:"bytes,1,opt,name=accountActivationToken,proto3" json:"accountActivationToken,omitempty"`
	PasswordResetToken     string   `protobuf:"bytes,2,opt,name=passwordResetToken,proto3" json:"passwordResetToken,omitempty"`
	InviteCodeToken        string   `protobuf:"bytes,3,opt,name=inviteCodeToken,proto3" json:"inviteCodeToken,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Tokens) Reset()         { *m = Tokens{} }
func (m *Tokens) String() string { return proto.CompactTextString(m) }
func (*Tokens) ProtoMessage()    {}
func (*Tokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7c0a1623e62d66c, []int{1}
}

func (m *Tokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tokens.Unmarshal(m, b)
}
func (m *Tokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tokens.Marshal(b, m, deterministic)
}
func (m *Tokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tokens.Merge(m, src)
}
func (m *Tokens) XXX_Size() int {
	return xxx_messageInfo_Tokens.Size(m)
}
func (m *Tokens) XXX_DiscardUnknown() {
	xxx_messageInfo_Tokens.DiscardUnknown(m)
}

var xxx_messageInfo_Tokens proto.InternalMessageInfo

func (m *Tokens) GetAccountActivationToken() string {
	if m != nil {
		return m.AccountActivationToken
	}
	return ""
}

func (m *Tokens) GetPasswordResetToken() string {
	if m != nil {
		return m.PasswordResetToken
	}
	return ""
}

func (m *Tokens) GetInviteCodeToken() string {
	if m != nil {
		return m.InviteCodeToken
	}
	return ""
}

type Metadata struct {
	TraceId              uint32   `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SourceService        string   `protobuf:"bytes,2,opt,name=source_service,json=sourceService,proto3" json:"source_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7c0a1623e62d66c, []int{2}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetTraceId() uint32 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *Metadata) GetSourceService() string {
	if m != nil {
		return m.SourceService
	}
	return ""
}

func init() {
	proto.RegisterEnum("EmailType", EmailType_name, EmailType_value)
	proto.RegisterType((*EmailContract)(nil), "EmailContract")
	proto.RegisterType((*Tokens)(nil), "Tokens")
	proto.RegisterType((*Metadata)(nil), "Metadata")
}

func init() {
	proto.RegisterFile("schema/proto/contracts/queue_contracts.proto", fileDescriptor_b7c0a1623e62d66c)
}

var fileDescriptor_b7c0a1623e62d66c = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0xcd, 0x6e, 0xdb, 0x24, 0xaf, 0xb4, 0x1b, 0xde, 0xa1, 0x44, 0x51, 0x29, 0x85, 0x85,
	0x22, 0x92, 0xc2, 0x0a, 0xde, 0x75, 0xf1, 0x20, 0xe8, 0x25, 0xee, 0xc9, 0x4b, 0x98, 0x9d, 0x3c,
	0xd7, 0x68, 0x33, 0x13, 0x67, 0x5e, 0x5a, 0xfa, 0x39, 0xbc, 0xf9, 0x69, 0x65, 0x66, 0xd2, 0x08,
	0xa2, 0xc7, 0xff, 0xef, 0xf7, 0x1e, 0x7d, 0xfd, 0x67, 0xe0, 0xa5, 0x95, 0x5f, 0xa9, 0x15, 0xbb,
	0xce, 0x68, 0xd6, 0x3b, 0xa9, 0x15, 0x1b, 0x21, 0xd9, 0xee, 0x7e, 0xf4, 0xd4, 0x53, 0x35, 0xe6,
	0xc2, 0xfb, 0xcd, 0xcf, 0x0b, 0x58, 0xbc, 0x6b, 0x45, 0xb3, 0xbf, 0x1d, 0x04, 0xae, 0x60, 0x66,
	0x49, 0xd5, 0x64, 0xf2, 0x68, 0x1d, 0x6d, 0xd3, 0x72, 0x48, 0x8e, 0xb3, 0x30, 0x0f, 0xc4, 0xf9,
	0x45, 0xe0, 0x21, 0x61, 0x0e, 0xb1, 0xed, 0xef, 0xbf, 0x91, 0xe4, 0xfc, 0xd2, 0x8b, 0x73, 0x74,
	0xa6, 0x25, 0x6b, 0xc5, 0x03, 0xe5, 0x93, 0x60, 0x86, 0x88, 0xcf, 0x61, 0xc2, 0xa7, 0x8e, 0xf2,
	0xe9, 0x3a, 0xda, 0x2e, 0x6f, 0xa0, 0xf0, 0x17, 0xdc, 0x9d, 0x3a, 0x2a, 0x3d, 0xc7, 0xa7, 0x90,
	0x7e, 0x69, 0x8c, 0x65, 0x25, 0x5a, 0xca, 0x67, 0x7e, 0xf7, 0x0f, 0xc0, 0x27, 0x90, 0xec, 0xc5,
	0x20, 0x63, 0x2f, 0xc7, 0x8c, 0xd7, 0x90, 0xb4, 0xc4, 0xa2, 0x16, 0x2c, 0xf2, 0x64, 0x1d, 0x6d,
	0xe7, 0x37, 0x69, 0xf1, 0x71, 0x00, 0xe5, 0xa8, 0xf0, 0x19, 0x4c, 0x59, 0x7f, 0x27, 0x95, 0xa7,
	0x7e, 0x26, 0x2e, 0xee, 0x5c, 0xb2, 0x65, 0xa0, 0x9b, 0x5f, 0x11, 0xcc, 0x02, 0xc1, 0xd7, 0xb0,
	0x12, 0x52, 0xea, 0x5e, 0xf1, 0x1b, 0xc9, 0xcd, 0x41, 0x70, 0xa3, 0x95, 0x57, 0x43, 0x3d, 0xff,
	0xb1, 0x58, 0x00, 0x76, 0xc2, 0xda, 0xa3, 0x36, 0x75, 0x49, 0x96, 0x38, 0xec, 0x84, 0xea, 0xfe,
	0x61, 0x70, 0x0b, 0x57, 0x8d, 0x3a, 0x34, 0x4c, 0xb7, 0xba, 0xa6, 0x30, 0x1c, 0xea, 0xfc, 0x1b,
	0x6f, 0x3e, 0x40, 0x72, 0xfe, 0x47, 0xf8, 0x18, 0x12, 0xf7, 0xd5, 0xa8, 0x6a, 0x6a, 0x7f, 0xcf,
	0xa2, 0x8c, 0x7d, 0x7e, 0x5f, 0xe3, 0x35, 0x2c, 0xad, 0xee, 0x8d, 0xa4, 0xca, 0x92, 0x39, 0x34,
	0x92, 0x86, 0x1f, 0x5f, 0x04, 0xfa, 0x29, 0xc0, 0x17, 0x47, 0x48, 0xc7, 0xf6, 0x71, 0x0e, 0xf1,
	0x91, 0xf6, 0x52, 0xb7, 0x94, 0x3d, 0x42, 0x84, 0xa5, 0x71, 0xf7, 0x55, 0xe7, 0x6b, 0xb3, 0x08,
	0xaf, 0x60, 0x1e, 0x18, 0xb9, 0x9d, 0xec, 0xc2, 0x81, 0x70, 0x5f, 0x25, 0x75, 0x4d, 0xd9, 0x25,
	0xae, 0x00, 0xed, 0xc9, 0x32, 0xb5, 0x55, 0x2b, 0x1a, 0xc5, 0xa4, 0x84, 0x92, 0x94, 0x4d, 0xdc,
	0x60, 0x67, 0x74, 0xab, 0x5d, 0x43, 0x62, 0x9f, 0x4d, 0xdf, 0xce, 0x3f, 0xa7, 0xe3, 0x63, 0xbc,
	0x9f, 0xf9, 0xd7, 0xf8, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x25, 0x1f, 0x78, 0xbd,
	0x02, 0x00, 0x00,
}
