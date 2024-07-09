// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: menu.proto

package menu

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

type MealDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId string  `protobuf:"bytes,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price        float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *MealDetails) Reset() {
	*x = MealDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MealDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MealDetails) ProtoMessage() {}

func (x *MealDetails) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MealDetails.ProtoReflect.Descriptor instead.
func (*MealDetails) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{0}
}

func (x *MealDetails) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *MealDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MealDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MealDetails) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type MealInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RestaurantId string  `protobuf:"bytes,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price        float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *MealInfo) Reset() {
	*x = MealInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MealInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MealInfo) ProtoMessage() {}

func (x *MealInfo) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MealInfo.ProtoReflect.Descriptor instead.
func (*MealInfo) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{1}
}

func (x *MealInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MealInfo) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *MealInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MealInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MealInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{3}
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId string `protobuf:"bytes,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Limit        int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset       int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{4}
}

func (x *Filter) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *Filter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Filter) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Meals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meals []*MealInfo `protobuf:"bytes,1,rep,name=meals,proto3" json:"meals,omitempty"`
}

func (x *Meals) Reset() {
	*x = Meals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meals) ProtoMessage() {}

func (x *Meals) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meals.ProtoReflect.Descriptor instead.
func (*Meals) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{5}
}

func (x *Meals) GetMeals() []*MealInfo {
	if x != nil {
		return x.Meals
	}
	return nil
}

var File_menu_proto protoreflect.FileDescriptor

var file_menu_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65,
	0x6e, 0x75, 0x22, 0x7e, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x5b,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x2d, 0x0a, 0x05, 0x4d,
	0x65, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x6d, 0x65, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6d, 0x65, 0x61, 0x6c, 0x73, 0x32, 0xce, 0x01, 0x0a, 0x04, 0x4d,
	0x65, 0x6e, 0x75, 0x12, 0x26, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x61, 0x6c, 0x12, 0x11,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x1a, 0x08, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x49, 0x44, 0x12, 0x27, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x61, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x12, 0x08, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x61, 0x6c, 0x12, 0x0e, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x61, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x22,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x12, 0x08, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x2e, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x61, 0x6c, 0x73,
	0x12, 0x0c, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x0b,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x61, 0x6c, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_proto_rawDescOnce sync.Once
	file_menu_proto_rawDescData = file_menu_proto_rawDesc
)

func file_menu_proto_rawDescGZIP() []byte {
	file_menu_proto_rawDescOnce.Do(func() {
		file_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_proto_rawDescData)
	})
	return file_menu_proto_rawDescData
}

var file_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_menu_proto_goTypes = []any{
	(*MealDetails)(nil), // 0: menu.MealDetails
	(*MealInfo)(nil),    // 1: menu.MealInfo
	(*ID)(nil),          // 2: menu.ID
	(*Void)(nil),        // 3: menu.Void
	(*Filter)(nil),      // 4: menu.Filter
	(*Meals)(nil),       // 5: menu.Meals
}
var file_menu_proto_depIdxs = []int32{
	1, // 0: menu.Meals.meals:type_name -> menu.MealInfo
	0, // 1: menu.Menu.AddMeal:input_type -> menu.MealDetails
	2, // 2: menu.Menu.GetMealByID:input_type -> menu.ID
	1, // 3: menu.Menu.UpdateMeal:input_type -> menu.MealInfo
	2, // 4: menu.Menu.DeleteMeal:input_type -> menu.ID
	4, // 5: menu.Menu.FetchMeals:input_type -> menu.Filter
	2, // 6: menu.Menu.AddMeal:output_type -> menu.ID
	1, // 7: menu.Menu.GetMealByID:output_type -> menu.MealInfo
	3, // 8: menu.Menu.UpdateMeal:output_type -> menu.Void
	3, // 9: menu.Menu.DeleteMeal:output_type -> menu.Void
	5, // 10: menu.Menu.FetchMeals:output_type -> menu.Meals
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_menu_proto_init() }
func file_menu_proto_init() {
	if File_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_menu_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MealDetails); i {
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
		file_menu_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MealInfo); i {
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
		file_menu_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ID); i {
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
		file_menu_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Void); i {
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
		file_menu_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Filter); i {
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
		file_menu_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Meals); i {
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
			RawDescriptor: file_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_proto_goTypes,
		DependencyIndexes: file_menu_proto_depIdxs,
		MessageInfos:      file_menu_proto_msgTypes,
	}.Build()
	File_menu_proto = out.File
	file_menu_proto_rawDesc = nil
	file_menu_proto_goTypes = nil
	file_menu_proto_depIdxs = nil
}
