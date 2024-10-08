// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: budgets.proto

package budgeting_service

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

type CreateBudget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CategoryId string  `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Amount     float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Period     string  `protobuf:"bytes,4,opt,name=period,proto3" json:"period,omitempty"`
	StartDate  string  `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate    string  `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *CreateBudget) Reset() {
	*x = CreateBudget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budgets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBudget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBudget) ProtoMessage() {}

func (x *CreateBudget) ProtoReflect() protoreflect.Message {
	mi := &file_budgets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBudget.ProtoReflect.Descriptor instead.
func (*CreateBudget) Descriptor() ([]byte, []int) {
	return file_budgets_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBudget) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateBudget) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *CreateBudget) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateBudget) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *CreateBudget) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CreateBudget) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CategoryId string  `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Amount     float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Period     string  `protobuf:"bytes,5,opt,name=period,proto3" json:"period,omitempty"`
	StartDate  string  `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate    string  `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	CreatedAt  string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budgets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_budgets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_budgets_proto_rawDescGZIP(), []int{1}
}

func (x *Budget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Budget) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Budget) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Budget) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Budget) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Budget) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Budget) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Budget) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Budget) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type BudgetFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	UserId     string  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CategoryId string  `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Amount     float64 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Period     string  `protobuf:"bytes,6,opt,name=period,proto3" json:"period,omitempty"`
	StartDate  string  `protobuf:"bytes,7,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate    string  `protobuf:"bytes,8,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *BudgetFilter) Reset() {
	*x = BudgetFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budgets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetFilter) ProtoMessage() {}

func (x *BudgetFilter) ProtoReflect() protoreflect.Message {
	mi := &file_budgets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetFilter.ProtoReflect.Descriptor instead.
func (*BudgetFilter) Descriptor() ([]byte, []int) {
	return file_budgets_proto_rawDescGZIP(), []int{2}
}

func (x *BudgetFilter) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *BudgetFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BudgetFilter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BudgetFilter) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *BudgetFilter) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BudgetFilter) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *BudgetFilter) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *BudgetFilter) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type Budgets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Budgets []*Budget `protobuf:"bytes,1,rep,name=budgets,proto3" json:"budgets,omitempty"`
}

func (x *Budgets) Reset() {
	*x = Budgets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budgets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budgets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budgets) ProtoMessage() {}

func (x *Budgets) ProtoReflect() protoreflect.Message {
	mi := &file_budgets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budgets.ProtoReflect.Descriptor instead.
func (*Budgets) Descriptor() ([]byte, []int) {
	return file_budgets_proto_rawDescGZIP(), []int{3}
}

func (x *Budgets) GetBudgets() []*Budget {
	if x != nil {
		return x.Budgets
	}
	return nil
}

var File_budgets_proto protoreflect.FileDescriptor

var file_budgets_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x17, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x22, 0xfa, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdc, 0x01,
	0x0a, 0x0c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3e, 0x0a, 0x07,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x07, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x32, 0xe3, 0x02, 0x0a,
	0x0d, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x1d, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x19,
	0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x1f, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x17, 0x2e, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_budgets_proto_rawDescOnce sync.Once
	file_budgets_proto_rawDescData = file_budgets_proto_rawDesc
)

func file_budgets_proto_rawDescGZIP() []byte {
	file_budgets_proto_rawDescOnce.Do(func() {
		file_budgets_proto_rawDescData = protoimpl.X.CompressGZIP(file_budgets_proto_rawDescData)
	})
	return file_budgets_proto_rawDescData
}

var file_budgets_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_budgets_proto_goTypes = []interface{}{
	(*CreateBudget)(nil), // 0: budgeting_service.CreateBudget
	(*Budget)(nil),       // 1: budgeting_service.Budget
	(*BudgetFilter)(nil), // 2: budgeting_service.BudgetFilter
	(*Budgets)(nil),      // 3: budgeting_service.Budgets
	(*PrimaryKey)(nil),   // 4: budgeting_service.PrimaryKey
	(*Void)(nil),         // 5: budgeting_service.Void
}
var file_budgets_proto_depIdxs = []int32{
	1, // 0: budgeting_service.Budgets.budgets:type_name -> budgeting_service.Budget
	0, // 1: budgeting_service.BudgetService.Create:input_type -> budgeting_service.CreateBudget
	4, // 2: budgeting_service.BudgetService.GetById:input_type -> budgeting_service.PrimaryKey
	2, // 3: budgeting_service.BudgetService.GetAll:input_type -> budgeting_service.BudgetFilter
	1, // 4: budgeting_service.BudgetService.Update:input_type -> budgeting_service.Budget
	4, // 5: budgeting_service.BudgetService.Delete:input_type -> budgeting_service.PrimaryKey
	1, // 6: budgeting_service.BudgetService.Create:output_type -> budgeting_service.Budget
	1, // 7: budgeting_service.BudgetService.GetById:output_type -> budgeting_service.Budget
	3, // 8: budgeting_service.BudgetService.GetAll:output_type -> budgeting_service.Budgets
	1, // 9: budgeting_service.BudgetService.Update:output_type -> budgeting_service.Budget
	5, // 10: budgeting_service.BudgetService.Delete:output_type -> budgeting_service.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_budgets_proto_init() }
func file_budgets_proto_init() {
	if File_budgets_proto != nil {
		return
	}
	file_budgeting_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_budgets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBudget); i {
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
		file_budgets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Budget); i {
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
		file_budgets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetFilter); i {
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
		file_budgets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Budgets); i {
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
			RawDescriptor: file_budgets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_budgets_proto_goTypes,
		DependencyIndexes: file_budgets_proto_depIdxs,
		MessageInfos:      file_budgets_proto_msgTypes,
	}.Build()
	File_budgets_proto = out.File
	file_budgets_proto_rawDesc = nil
	file_budgets_proto_goTypes = nil
	file_budgets_proto_depIdxs = nil
}
