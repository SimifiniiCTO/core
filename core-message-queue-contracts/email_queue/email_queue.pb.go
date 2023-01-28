// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: email_queue.proto

package email_queue

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

type EmailType int32

const (
	EmailType_PASSWORD_RESET_EMAIL EmailType = 0
	EmailType_WELCOME_EMAIL        EmailType = 1
	// reset email account email type
	EmailType_RESET_EMAIL EmailType = 2
	// invite code email type
	EmailType_INVITE_CODE EmailType = 3
	// system maintenance email type
	EmailType_SYSTEM_MAINTENANCE EmailType = 4
	// promotional email type
	EmailType_PROMOTIONAL EmailType = 5
)

// Enum value maps for EmailType.
var (
	EmailType_name = map[int32]string{
		0: "PASSWORD_RESET_EMAIL",
		1: "WELCOME_EMAIL",
		2: "RESET_EMAIL",
		3: "INVITE_CODE",
		4: "SYSTEM_MAINTENANCE",
		5: "PROMOTIONAL",
	}
	EmailType_value = map[string]int32{
		"PASSWORD_RESET_EMAIL": 0,
		"WELCOME_EMAIL":        1,
		"RESET_EMAIL":          2,
		"INVITE_CODE":          3,
		"SYSTEM_MAINTENANCE":   4,
		"PROMOTIONAL":          5,
	}
)

func (x EmailType) Enum() *EmailType {
	p := new(EmailType)
	*p = x
	return p
}

func (x EmailType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailType) Descriptor() protoreflect.EnumDescriptor {
	return file_email_queue_proto_enumTypes[0].Descriptor()
}

func (EmailType) Type() protoreflect.EnumType {
	return &file_email_queue_proto_enumTypes[0]
}

func (x EmailType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailType.Descriptor instead.
func (EmailType) EnumDescriptor() ([]byte, []int) {
	return file_email_queue_proto_rawDescGZIP(), []int{0}
}

//
//EmailMessage: represents an email message
type EmailMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From                 string    `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string    `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Subject              string    `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	UserName             string    `protobuf:"bytes,4,opt,name=userName,proto3" json:"userName,omitempty"`
	EmailType            EmailType `protobuf:"varint,5,opt,name=emailType,proto3,enum=email_queue.EmailType" json:"emailType,omitempty"`
	FirstName            string    `protobuf:"bytes,6,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string    `protobuf:"bytes,7,opt,name=lastName,proto3" json:"lastName,omitempty"`
	VerificationAddress  string    `protobuf:"bytes,8,opt,name=verificationAddress,proto3" json:"verificationAddress,omitempty"`
	UserID               string    `protobuf:"bytes,9,opt,name=userID,proto3" json:"userID,omitempty"`
	PasswordResetAuthnID int32     `protobuf:"varint,10,opt,name=passwordResetAuthnID,proto3" json:"passwordResetAuthnID,omitempty"`
	PasswordResetToken   string    `protobuf:"bytes,11,opt,name=passwordResetToken,proto3" json:"passwordResetToken,omitempty"`
}

func (x *EmailMessage) Reset() {
	*x = EmailMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailMessage) ProtoMessage() {}

func (x *EmailMessage) ProtoReflect() protoreflect.Message {
	mi := &file_email_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailMessage.ProtoReflect.Descriptor instead.
func (*EmailMessage) Descriptor() ([]byte, []int) {
	return file_email_queue_proto_rawDescGZIP(), []int{0}
}

func (x *EmailMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EmailMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EmailMessage) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailMessage) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *EmailMessage) GetEmailType() EmailType {
	if x != nil {
		return x.EmailType
	}
	return EmailType_PASSWORD_RESET_EMAIL
}

func (x *EmailMessage) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *EmailMessage) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *EmailMessage) GetVerificationAddress() string {
	if x != nil {
		return x.VerificationAddress
	}
	return ""
}

func (x *EmailMessage) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *EmailMessage) GetPasswordResetAuthnID() int32 {
	if x != nil {
		return x.PasswordResetAuthnID
	}
	return 0
}

func (x *EmailMessage) GetPasswordResetToken() string {
	if x != nil {
		return x.PasswordResetToken
	}
	return ""
}

type EmailContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string    `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string    `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Subject   string    `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Message   string    `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	EmailType EmailType `protobuf:"varint,5,opt,name=emailType,proto3,enum=email_queue.EmailType" json:"emailType,omitempty"`
	UserName  string    `protobuf:"bytes,6,opt,name=userName,proto3" json:"userName,omitempty"`
	FirstName string    `protobuf:"bytes,7,opt,name=firstName,proto3" json:"firstName,omitempty"`
	Metadata  *Metadata `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Token     *Tokens   `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *EmailContract) Reset() {
	*x = EmailContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailContract) ProtoMessage() {}

func (x *EmailContract) ProtoReflect() protoreflect.Message {
	mi := &file_email_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailContract.ProtoReflect.Descriptor instead.
func (*EmailContract) Descriptor() ([]byte, []int) {
	return file_email_queue_proto_rawDescGZIP(), []int{1}
}

func (x *EmailContract) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EmailContract) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EmailContract) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailContract) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EmailContract) GetEmailType() EmailType {
	if x != nil {
		return x.EmailType
	}
	return EmailType_PASSWORD_RESET_EMAIL
}

func (x *EmailContract) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *EmailContract) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *EmailContract) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EmailContract) GetToken() *Tokens {
	if x != nil {
		return x.Token
	}
	return nil
}

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountActivationToken string `protobuf:"bytes,1,opt,name=accountActivationToken,proto3" json:"accountActivationToken,omitempty"`
	PasswordResetToken     string `protobuf:"bytes,2,opt,name=passwordResetToken,proto3" json:"passwordResetToken,omitempty"`
	InviteCodeToken        string `protobuf:"bytes,3,opt,name=inviteCodeToken,proto3" json:"inviteCodeToken,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_email_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_email_queue_proto_rawDescGZIP(), []int{2}
}

func (x *Tokens) GetAccountActivationToken() string {
	if x != nil {
		return x.AccountActivationToken
	}
	return ""
}

func (x *Tokens) GetPasswordResetToken() string {
	if x != nil {
		return x.PasswordResetToken
	}
	return ""
}

func (x *Tokens) GetInviteCodeToken() string {
	if x != nil {
		return x.InviteCodeToken
	}
	return ""
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId       uint32 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SourceService string `protobuf:"bytes,2,opt,name=source_service,json=sourceService,proto3" json:"source_service,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_queue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_email_queue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_email_queue_proto_rawDescGZIP(), []int{3}
}

func (x *Metadata) GetTraceId() uint32 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (x *Metadata) GetSourceService() string {
	if x != nil {
		return x.SourceService
	}
	return ""
}

var File_email_queue_proto protoreflect.FileDescriptor

var file_email_queue_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x22, 0x86, 0x03, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x14, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb5, 0x02, 0x0a, 0x0d, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x16,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4c,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2a, 0x83, 0x01, 0x0a,
	0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x41,
	0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x45, 0x4d, 0x41,
	0x49, 0x4c, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x45, 0x4c, 0x43, 0x4f, 0x4d, 0x45, 0x5f,
	0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x53, 0x45, 0x54,
	0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x56, 0x49,
	0x54, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x59, 0x53,
	0x54, 0x45, 0x4d, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10,
	0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c,
	0x10, 0x05, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x69, 0x6d, 0x69, 0x66, 0x69, 0x6e, 0x69, 0x69, 0x43, 0x54, 0x4f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_email_queue_proto_rawDescOnce sync.Once
	file_email_queue_proto_rawDescData = file_email_queue_proto_rawDesc
)

func file_email_queue_proto_rawDescGZIP() []byte {
	file_email_queue_proto_rawDescOnce.Do(func() {
		file_email_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_queue_proto_rawDescData)
	})
	return file_email_queue_proto_rawDescData
}

var file_email_queue_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_email_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_email_queue_proto_goTypes = []interface{}{
	(EmailType)(0),        // 0: email_queue.EmailType
	(*EmailMessage)(nil),  // 1: email_queue.EmailMessage
	(*EmailContract)(nil), // 2: email_queue.EmailContract
	(*Tokens)(nil),        // 3: email_queue.Tokens
	(*Metadata)(nil),      // 4: email_queue.Metadata
}
var file_email_queue_proto_depIdxs = []int32{
	0, // 0: email_queue.EmailMessage.emailType:type_name -> email_queue.EmailType
	0, // 1: email_queue.EmailContract.emailType:type_name -> email_queue.EmailType
	4, // 2: email_queue.EmailContract.metadata:type_name -> email_queue.Metadata
	3, // 3: email_queue.EmailContract.token:type_name -> email_queue.Tokens
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_email_queue_proto_init() }
func file_email_queue_proto_init() {
	if File_email_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailMessage); i {
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
		file_email_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailContract); i {
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
		file_email_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
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
		file_email_queue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_email_queue_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_email_queue_proto_goTypes,
		DependencyIndexes: file_email_queue_proto_depIdxs,
		EnumInfos:         file_email_queue_proto_enumTypes,
		MessageInfos:      file_email_queue_proto_msgTypes,
	}.Build()
	File_email_queue_proto = out.File
	file_email_queue_proto_rawDesc = nil
	file_email_queue_proto_goTypes = nil
	file_email_queue_proto_depIdxs = nil
}
