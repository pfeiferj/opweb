// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type CustomReserved0 capnp.Struct

// CustomReserved0_TypeID is the unique identifier for the type CustomReserved0.
const CustomReserved0_TypeID = 0x81c2f05a394cf4af

func NewCustomReserved0(s *capnp.Segment) (CustomReserved0, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved0(st), err
}

func NewRootCustomReserved0(s *capnp.Segment) (CustomReserved0, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved0(st), err
}

func ReadRootCustomReserved0(msg *capnp.Message) (CustomReserved0, error) {
	root, err := msg.Root()
	return CustomReserved0(root.Struct()), err
}

func (s CustomReserved0) String() string {
	str, _ := text.Marshal(0x81c2f05a394cf4af, capnp.Struct(s))
	return str
}

func (s CustomReserved0) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved0) DecodeFromPtr(p capnp.Ptr) CustomReserved0 {
	return CustomReserved0(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved0) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved0) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved0) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved0) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved0_List is a list of CustomReserved0.
type CustomReserved0_List = capnp.StructList[CustomReserved0]

// NewCustomReserved0 creates a new list of CustomReserved0.
func NewCustomReserved0_List(s *capnp.Segment, sz int32) (CustomReserved0_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved0](l), err
}

// CustomReserved0_Future is a wrapper for a CustomReserved0 promised by a client call.
type CustomReserved0_Future struct{ *capnp.Future }

func (f CustomReserved0_Future) Struct() (CustomReserved0, error) {
	p, err := f.Future.Ptr()
	return CustomReserved0(p.Struct()), err
}

type CustomReserved1 capnp.Struct

// CustomReserved1_TypeID is the unique identifier for the type CustomReserved1.
const CustomReserved1_TypeID = 0xaedffd8f31e7b55d

func NewCustomReserved1(s *capnp.Segment) (CustomReserved1, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved1(st), err
}

func NewRootCustomReserved1(s *capnp.Segment) (CustomReserved1, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved1(st), err
}

func ReadRootCustomReserved1(msg *capnp.Message) (CustomReserved1, error) {
	root, err := msg.Root()
	return CustomReserved1(root.Struct()), err
}

func (s CustomReserved1) String() string {
	str, _ := text.Marshal(0xaedffd8f31e7b55d, capnp.Struct(s))
	return str
}

func (s CustomReserved1) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved1) DecodeFromPtr(p capnp.Ptr) CustomReserved1 {
	return CustomReserved1(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved1) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved1) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved1) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved1) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved1_List is a list of CustomReserved1.
type CustomReserved1_List = capnp.StructList[CustomReserved1]

// NewCustomReserved1 creates a new list of CustomReserved1.
func NewCustomReserved1_List(s *capnp.Segment, sz int32) (CustomReserved1_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved1](l), err
}

// CustomReserved1_Future is a wrapper for a CustomReserved1 promised by a client call.
type CustomReserved1_Future struct{ *capnp.Future }

func (f CustomReserved1_Future) Struct() (CustomReserved1, error) {
	p, err := f.Future.Ptr()
	return CustomReserved1(p.Struct()), err
}

type CustomReserved2 capnp.Struct

// CustomReserved2_TypeID is the unique identifier for the type CustomReserved2.
const CustomReserved2_TypeID = 0xf35cc4560bbf6ec2

func NewCustomReserved2(s *capnp.Segment) (CustomReserved2, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved2(st), err
}

func NewRootCustomReserved2(s *capnp.Segment) (CustomReserved2, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved2(st), err
}

func ReadRootCustomReserved2(msg *capnp.Message) (CustomReserved2, error) {
	root, err := msg.Root()
	return CustomReserved2(root.Struct()), err
}

func (s CustomReserved2) String() string {
	str, _ := text.Marshal(0xf35cc4560bbf6ec2, capnp.Struct(s))
	return str
}

func (s CustomReserved2) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved2) DecodeFromPtr(p capnp.Ptr) CustomReserved2 {
	return CustomReserved2(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved2) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved2) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved2) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved2) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved2_List is a list of CustomReserved2.
type CustomReserved2_List = capnp.StructList[CustomReserved2]

// NewCustomReserved2 creates a new list of CustomReserved2.
func NewCustomReserved2_List(s *capnp.Segment, sz int32) (CustomReserved2_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved2](l), err
}

// CustomReserved2_Future is a wrapper for a CustomReserved2 promised by a client call.
type CustomReserved2_Future struct{ *capnp.Future }

func (f CustomReserved2_Future) Struct() (CustomReserved2, error) {
	p, err := f.Future.Ptr()
	return CustomReserved2(p.Struct()), err
}

type CustomReserved3 capnp.Struct

// CustomReserved3_TypeID is the unique identifier for the type CustomReserved3.
const CustomReserved3_TypeID = 0xda96579883444c35

func NewCustomReserved3(s *capnp.Segment) (CustomReserved3, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved3(st), err
}

func NewRootCustomReserved3(s *capnp.Segment) (CustomReserved3, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved3(st), err
}

func ReadRootCustomReserved3(msg *capnp.Message) (CustomReserved3, error) {
	root, err := msg.Root()
	return CustomReserved3(root.Struct()), err
}

func (s CustomReserved3) String() string {
	str, _ := text.Marshal(0xda96579883444c35, capnp.Struct(s))
	return str
}

func (s CustomReserved3) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved3) DecodeFromPtr(p capnp.Ptr) CustomReserved3 {
	return CustomReserved3(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved3) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved3) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved3) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved3) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved3_List is a list of CustomReserved3.
type CustomReserved3_List = capnp.StructList[CustomReserved3]

// NewCustomReserved3 creates a new list of CustomReserved3.
func NewCustomReserved3_List(s *capnp.Segment, sz int32) (CustomReserved3_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved3](l), err
}

// CustomReserved3_Future is a wrapper for a CustomReserved3 promised by a client call.
type CustomReserved3_Future struct{ *capnp.Future }

func (f CustomReserved3_Future) Struct() (CustomReserved3, error) {
	p, err := f.Future.Ptr()
	return CustomReserved3(p.Struct()), err
}

type CustomReserved4 capnp.Struct

// CustomReserved4_TypeID is the unique identifier for the type CustomReserved4.
const CustomReserved4_TypeID = 0x80ae746ee2596b11

func NewCustomReserved4(s *capnp.Segment) (CustomReserved4, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved4(st), err
}

func NewRootCustomReserved4(s *capnp.Segment) (CustomReserved4, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved4(st), err
}

func ReadRootCustomReserved4(msg *capnp.Message) (CustomReserved4, error) {
	root, err := msg.Root()
	return CustomReserved4(root.Struct()), err
}

func (s CustomReserved4) String() string {
	str, _ := text.Marshal(0x80ae746ee2596b11, capnp.Struct(s))
	return str
}

func (s CustomReserved4) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved4) DecodeFromPtr(p capnp.Ptr) CustomReserved4 {
	return CustomReserved4(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved4) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved4) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved4) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved4) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved4_List is a list of CustomReserved4.
type CustomReserved4_List = capnp.StructList[CustomReserved4]

// NewCustomReserved4 creates a new list of CustomReserved4.
func NewCustomReserved4_List(s *capnp.Segment, sz int32) (CustomReserved4_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved4](l), err
}

// CustomReserved4_Future is a wrapper for a CustomReserved4 promised by a client call.
type CustomReserved4_Future struct{ *capnp.Future }

func (f CustomReserved4_Future) Struct() (CustomReserved4, error) {
	p, err := f.Future.Ptr()
	return CustomReserved4(p.Struct()), err
}

type CustomReserved5 capnp.Struct

// CustomReserved5_TypeID is the unique identifier for the type CustomReserved5.
const CustomReserved5_TypeID = 0xa5cd762cd951a455

func NewCustomReserved5(s *capnp.Segment) (CustomReserved5, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved5(st), err
}

func NewRootCustomReserved5(s *capnp.Segment) (CustomReserved5, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved5(st), err
}

func ReadRootCustomReserved5(msg *capnp.Message) (CustomReserved5, error) {
	root, err := msg.Root()
	return CustomReserved5(root.Struct()), err
}

func (s CustomReserved5) String() string {
	str, _ := text.Marshal(0xa5cd762cd951a455, capnp.Struct(s))
	return str
}

func (s CustomReserved5) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved5) DecodeFromPtr(p capnp.Ptr) CustomReserved5 {
	return CustomReserved5(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved5) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved5) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved5) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved5) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved5_List is a list of CustomReserved5.
type CustomReserved5_List = capnp.StructList[CustomReserved5]

// NewCustomReserved5 creates a new list of CustomReserved5.
func NewCustomReserved5_List(s *capnp.Segment, sz int32) (CustomReserved5_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved5](l), err
}

// CustomReserved5_Future is a wrapper for a CustomReserved5 promised by a client call.
type CustomReserved5_Future struct{ *capnp.Future }

func (f CustomReserved5_Future) Struct() (CustomReserved5, error) {
	p, err := f.Future.Ptr()
	return CustomReserved5(p.Struct()), err
}

type CustomReserved6 capnp.Struct

// CustomReserved6_TypeID is the unique identifier for the type CustomReserved6.
const CustomReserved6_TypeID = 0xf98d843bfd7004a3

func NewCustomReserved6(s *capnp.Segment) (CustomReserved6, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved6(st), err
}

func NewRootCustomReserved6(s *capnp.Segment) (CustomReserved6, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved6(st), err
}

func ReadRootCustomReserved6(msg *capnp.Message) (CustomReserved6, error) {
	root, err := msg.Root()
	return CustomReserved6(root.Struct()), err
}

func (s CustomReserved6) String() string {
	str, _ := text.Marshal(0xf98d843bfd7004a3, capnp.Struct(s))
	return str
}

func (s CustomReserved6) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved6) DecodeFromPtr(p capnp.Ptr) CustomReserved6 {
	return CustomReserved6(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved6) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved6) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved6) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved6) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved6_List is a list of CustomReserved6.
type CustomReserved6_List = capnp.StructList[CustomReserved6]

// NewCustomReserved6 creates a new list of CustomReserved6.
func NewCustomReserved6_List(s *capnp.Segment, sz int32) (CustomReserved6_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved6](l), err
}

// CustomReserved6_Future is a wrapper for a CustomReserved6 promised by a client call.
type CustomReserved6_Future struct{ *capnp.Future }

func (f CustomReserved6_Future) Struct() (CustomReserved6, error) {
	p, err := f.Future.Ptr()
	return CustomReserved6(p.Struct()), err
}

type CustomReserved7 capnp.Struct

// CustomReserved7_TypeID is the unique identifier for the type CustomReserved7.
const CustomReserved7_TypeID = 0xb86e6369214c01c8

func NewCustomReserved7(s *capnp.Segment) (CustomReserved7, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved7(st), err
}

func NewRootCustomReserved7(s *capnp.Segment) (CustomReserved7, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved7(st), err
}

func ReadRootCustomReserved7(msg *capnp.Message) (CustomReserved7, error) {
	root, err := msg.Root()
	return CustomReserved7(root.Struct()), err
}

func (s CustomReserved7) String() string {
	str, _ := text.Marshal(0xb86e6369214c01c8, capnp.Struct(s))
	return str
}

func (s CustomReserved7) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved7) DecodeFromPtr(p capnp.Ptr) CustomReserved7 {
	return CustomReserved7(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved7) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved7) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved7) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved7) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved7_List is a list of CustomReserved7.
type CustomReserved7_List = capnp.StructList[CustomReserved7]

// NewCustomReserved7 creates a new list of CustomReserved7.
func NewCustomReserved7_List(s *capnp.Segment, sz int32) (CustomReserved7_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved7](l), err
}

// CustomReserved7_Future is a wrapper for a CustomReserved7 promised by a client call.
type CustomReserved7_Future struct{ *capnp.Future }

func (f CustomReserved7_Future) Struct() (CustomReserved7, error) {
	p, err := f.Future.Ptr()
	return CustomReserved7(p.Struct()), err
}

type CustomReserved8 capnp.Struct

// CustomReserved8_TypeID is the unique identifier for the type CustomReserved8.
const CustomReserved8_TypeID = 0xf416ec09499d9d19

func NewCustomReserved8(s *capnp.Segment) (CustomReserved8, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved8(st), err
}

func NewRootCustomReserved8(s *capnp.Segment) (CustomReserved8, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved8(st), err
}

func ReadRootCustomReserved8(msg *capnp.Message) (CustomReserved8, error) {
	root, err := msg.Root()
	return CustomReserved8(root.Struct()), err
}

func (s CustomReserved8) String() string {
	str, _ := text.Marshal(0xf416ec09499d9d19, capnp.Struct(s))
	return str
}

func (s CustomReserved8) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved8) DecodeFromPtr(p capnp.Ptr) CustomReserved8 {
	return CustomReserved8(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved8) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved8) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved8) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved8) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved8_List is a list of CustomReserved8.
type CustomReserved8_List = capnp.StructList[CustomReserved8]

// NewCustomReserved8 creates a new list of CustomReserved8.
func NewCustomReserved8_List(s *capnp.Segment, sz int32) (CustomReserved8_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved8](l), err
}

// CustomReserved8_Future is a wrapper for a CustomReserved8 promised by a client call.
type CustomReserved8_Future struct{ *capnp.Future }

func (f CustomReserved8_Future) Struct() (CustomReserved8, error) {
	p, err := f.Future.Ptr()
	return CustomReserved8(p.Struct()), err
}

type CustomReserved9 capnp.Struct

// CustomReserved9_TypeID is the unique identifier for the type CustomReserved9.
const CustomReserved9_TypeID = 0xa1680744031fdb2d

func NewCustomReserved9(s *capnp.Segment) (CustomReserved9, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved9(st), err
}

func NewRootCustomReserved9(s *capnp.Segment) (CustomReserved9, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CustomReserved9(st), err
}

func ReadRootCustomReserved9(msg *capnp.Message) (CustomReserved9, error) {
	root, err := msg.Root()
	return CustomReserved9(root.Struct()), err
}

func (s CustomReserved9) String() string {
	str, _ := text.Marshal(0xa1680744031fdb2d, capnp.Struct(s))
	return str
}

func (s CustomReserved9) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CustomReserved9) DecodeFromPtr(p capnp.Ptr) CustomReserved9 {
	return CustomReserved9(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CustomReserved9) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CustomReserved9) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CustomReserved9) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CustomReserved9) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// CustomReserved9_List is a list of CustomReserved9.
type CustomReserved9_List = capnp.StructList[CustomReserved9]

// NewCustomReserved9 creates a new list of CustomReserved9.
func NewCustomReserved9_List(s *capnp.Segment, sz int32) (CustomReserved9_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CustomReserved9](l), err
}

// CustomReserved9_Future is a wrapper for a CustomReserved9 promised by a client call.
type CustomReserved9_Future struct{ *capnp.Future }

func (f CustomReserved9_Future) Struct() (CustomReserved9, error) {
	p, err := f.Future.Ptr()
	return CustomReserved9(p.Struct()), err
}

const schema_b526ba661d550a59 = "x\xda2H\xe1v`1\xe4M\x8fd`\x0a\x16`" +
	"ae\xfb/\x98\x1d\xf9(\xafd]\x03\x83 /\xe3" +
	"\xffH\xaeP\xd9\xb4]j[\x19X\xd8\x19\x18\x84g" +
	"\xea\xbc\x12^\xaa\xc3\xce\xc0\xfc\x7f\xfd\x17\x1f\xcb\xa8\x0f" +
	"\x87\x1a\xb1(j\xd4y%\xdc\x0bV\xa4{[\x9e\xd9" +
	"\x85=c!\x16E\x99:\xaf\x84K\xc1\x8aB\x97\x04" +
	"\xde\xd4);\xbb\x14\x8b\xa2@\x9dW\xc2\xb1`E\xb1" +
	"[\x9f\x1b\xf6\xff\xbd\xbf\x0e\x8b\"K\x9dW\xc2\xae`" +
	"E'\x18}\x143\x93\xf3v`Q\xa4\xa8\xf3JX" +
	"\x17\xac\xc8\xd4\xc7\xa5yF\xf8\xb4[X\x14q\xea\xbc" +
	"\x12\x16\x05+:\x94\xb7\x9f;\xecH\xccg,\x8a>" +
	"j\xbf\x12\xfe\xab\x0dR$9w\xae'\xe7\x1b\xb1/" +
	"X\x14\xdd\xd4~%\xfc\x14\xach1K\xc1_\xeb\x96" +
	"\xde\x9fX\x14\x1d\xd4~%|V\x9b\x9d\xe1?\x16x" +
	"\xf0\x7friqI~\xae^2SbA^\x81\x95" +
	"3\x98\x17\x94Z\x9cZ\xc4_\x96\x9ab\x12\xc0\xc8\x88" +
	"_\x85\x01A\x15\x96\x04U\x98\x12TaHP\x859" +
	"A\x15\xc6\x04U\x18\x11TaAP\x85Y\x00##" +
	" \x00\x00\xff\xff\"f\xfa\xbf"

func RegisterCustomSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_b526ba661d550a59,
		Nodes: []uint64{
			0x80ae746ee2596b11,
			0x81c2f05a394cf4af,
			0xa1680744031fdb2d,
			0xa5cd762cd951a455,
			0xaedffd8f31e7b55d,
			0xb86e6369214c01c8,
			0xda96579883444c35,
			0xf35cc4560bbf6ec2,
			0xf416ec09499d9d19,
			0xf98d843bfd7004a3,
		},
		Compressed: true,
	})
}