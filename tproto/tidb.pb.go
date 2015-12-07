// Code generated by protoc-gen-go.
// source: tidb.proto
// DO NOT EDIT!

/*
Package tproto is a generated protocol buffer package.

It is generated from these files:
	tidb.proto

It has these top-level messages:
	TableName
	CharCollation
	JoinColumn
	Join
	TIndexColumn
	TIndex
	TColumn
	FieldType
	TTable
	Property
*/
package tproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JoinType int32

const (
	JoinType_LEFT_OUTER_JOIN  JoinType = 0
	JoinType_RIGHT_OUTER_JOIN JoinType = 1
	JoinType_FULL_OUTER_JOIN  JoinType = 2
	JoinType_INNER_JOIN       JoinType = 3
)

var JoinType_name = map[int32]string{
	0: "LEFT_OUTER_JOIN",
	1: "RIGHT_OUTER_JOIN",
	2: "FULL_OUTER_JOIN",
	3: "INNER_JOIN",
}
var JoinType_value = map[string]int32{
	"LEFT_OUTER_JOIN":  0,
	"RIGHT_OUTER_JOIN": 1,
	"FULL_OUTER_JOIN":  2,
	"INNER_JOIN":       3,
}

func (x JoinType) Enum() *JoinType {
	p := new(JoinType)
	*p = x
	return p
}
func (x JoinType) String() string {
	return proto.EnumName(JoinType_name, int32(x))
}
func (x *JoinType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(JoinType_value, data, "JoinType")
	if err != nil {
		return err
	}
	*x = JoinType(value)
	return nil
}
func (JoinType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TableName struct {
	SchemaName       *string `protobuf:"bytes,1,req,name=schemaName" json:"schemaName,omitempty"`
	TableName        *string `protobuf:"bytes,2,req,name=tableName" json:"tableName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TableName) Reset()                    { *m = TableName{} }
func (m *TableName) String() string            { return proto.CompactTextString(m) }
func (*TableName) ProtoMessage()               {}
func (*TableName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TableName) GetSchemaName() string {
	if m != nil && m.SchemaName != nil {
		return *m.SchemaName
	}
	return ""
}

func (m *TableName) GetTableName() string {
	if m != nil && m.TableName != nil {
		return *m.TableName
	}
	return ""
}

type CharCollation struct {
	CharacterSetName   *string `protobuf:"bytes,1,opt,name=characterSetName" json:"characterSetName,omitempty"`
	CollationOrderName *string `protobuf:"bytes,2,opt,name=collationOrderName" json:"collationOrderName,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *CharCollation) Reset()                    { *m = CharCollation{} }
func (m *CharCollation) String() string            { return proto.CompactTextString(m) }
func (*CharCollation) ProtoMessage()               {}
func (*CharCollation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CharCollation) GetCharacterSetName() string {
	if m != nil && m.CharacterSetName != nil {
		return *m.CharacterSetName
	}
	return ""
}

func (m *CharCollation) GetCollationOrderName() string {
	if m != nil && m.CollationOrderName != nil {
		return *m.CollationOrderName
	}
	return ""
}

// One of the column pairs which define the Join.
type JoinColumn struct {
	ParentColumn     *string `protobuf:"bytes,1,opt,name=parentColumn" json:"parentColumn,omitempty"`
	ChildColumn      *string `protobuf:"bytes,2,opt,name=childColumn" json:"childColumn,omitempty"`
	Position         *int32  `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *JoinColumn) Reset()                    { *m = JoinColumn{} }
func (m *JoinColumn) String() string            { return proto.CompactTextString(m) }
func (*JoinColumn) ProtoMessage()               {}
func (*JoinColumn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JoinColumn) GetParentColumn() string {
	if m != nil && m.ParentColumn != nil {
		return *m.ParentColumn
	}
	return ""
}

func (m *JoinColumn) GetChildColumn() string {
	if m != nil && m.ChildColumn != nil {
		return *m.ChildColumn
	}
	return ""
}

func (m *JoinColumn) GetPosition() int32 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

type Join struct {
	Columns          []*JoinColumn `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	ConstraintName   *TableName    `protobuf:"bytes,3,opt,name=constraintName" json:"constraintName,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Join) Reset()                    { *m = Join{} }
func (m *Join) String() string            { return proto.CompactTextString(m) }
func (*Join) ProtoMessage()               {}
func (*Join) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Join) GetColumns() []*JoinColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Join) GetConstraintName() *TableName {
	if m != nil {
		return m.ConstraintName
	}
	return nil
}

type TIndexColumn struct {
	Name             *string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Offset           *int32  `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	Length           *int32  `protobuf:"varint,3,opt,name=Length" json:"Length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TIndexColumn) Reset()                    { *m = TIndexColumn{} }
func (m *TIndexColumn) String() string            { return proto.CompactTextString(m) }
func (*TIndexColumn) ProtoMessage()               {}
func (*TIndexColumn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TIndexColumn) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TIndexColumn) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *TIndexColumn) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

// Column info in Table.
// Only contains publiced indices.
type TIndex struct {
	ID               *int32          `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name             *string         `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	IsPK             *bool           `protobuf:"varint,3,opt,name=IsPK" json:"IsPK,omitempty"`
	IsUnique         *bool           `protobuf:"varint,4,opt,name=IsUnique" json:"IsUnique,omitempty"`
	Columns          []*TIndexColumn `protobuf:"bytes,5,rep,name=Columns" json:"Columns,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *TIndex) Reset()                    { *m = TIndex{} }
func (m *TIndex) String() string            { return proto.CompactTextString(m) }
func (*TIndex) ProtoMessage()               {}
func (*TIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TIndex) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *TIndex) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TIndex) GetIsPK() bool {
	if m != nil && m.IsPK != nil {
		return *m.IsPK
	}
	return false
}

func (m *TIndex) GetIsUnique() bool {
	if m != nil && m.IsUnique != nil {
		return *m.IsUnique
	}
	return false
}

func (m *TIndex) GetColumns() []*TIndexColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

// Column info in Table.
// Only contains publiced columns.
type TColumn struct {
	ID               *int64     `protobuf:"varint,1,req,name=ID" json:"ID,omitempty"`
	Name             *string    `protobuf:"bytes,2,req,name=Name" json:"Name,omitempty"`
	Type             *FieldType `protobuf:"bytes,3,opt,name=Type" json:"Type,omitempty"`
	Offset           *int32     `protobuf:"varint,4,opt,name=Offset" json:"Offset,omitempty"`
	DefaultValue     *string    `protobuf:"bytes,5,opt,name=defaultValue" json:"defaultValue,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *TColumn) Reset()                    { *m = TColumn{} }
func (m *TColumn) String() string            { return proto.CompactTextString(m) }
func (*TColumn) ProtoMessage()               {}
func (*TColumn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TColumn) GetID() int64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *TColumn) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TColumn) GetType() *FieldType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *TColumn) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *TColumn) GetDefaultValue() string {
	if m != nil && m.DefaultValue != nil {
		return *m.DefaultValue
	}
	return ""
}

type FieldType struct {
	Tp               *int32         `protobuf:"varint,1,req,name=Tp" json:"Tp,omitempty"`
	Flag             *uint32        `protobuf:"varint,2,req,name=Flag" json:"Flag,omitempty"`
	Flen             *int32         `protobuf:"varint,3,req,name=Flen" json:"Flen,omitempty"`
	Decimal          *int32         `protobuf:"varint,4,req,name=Decimal" json:"Decimal,omitempty"`
	CharsetInfo      *CharCollation `protobuf:"bytes,5,opt,name=CharsetInfo" json:"CharsetInfo,omitempty"`
	Elems            []string       `protobuf:"bytes,6,rep,name=Elems" json:"Elems,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FieldType) Reset()                    { *m = FieldType{} }
func (m *FieldType) String() string            { return proto.CompactTextString(m) }
func (*FieldType) ProtoMessage()               {}
func (*FieldType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FieldType) GetTp() int32 {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return 0
}

func (m *FieldType) GetFlag() uint32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *FieldType) GetFlen() int32 {
	if m != nil && m.Flen != nil {
		return *m.Flen
	}
	return 0
}

func (m *FieldType) GetDecimal() int32 {
	if m != nil && m.Decimal != nil {
		return *m.Decimal
	}
	return 0
}

func (m *FieldType) GetCharsetInfo() *CharCollation {
	if m != nil {
		return m.CharsetInfo
	}
	return nil
}

func (m *FieldType) GetElems() []string {
	if m != nil {
		return m.Elems
	}
	return nil
}

type TTable struct {
	ID               *int64         `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Columns          []*TColumn     `protobuf:"bytes,3,rep,name=columns" json:"columns,omitempty"`
	Indexes          []*TIndex      `protobuf:"bytes,4,rep,name=indexes" json:"indexes,omitempty"`
	CharColl         *CharCollation `protobuf:"bytes,5,opt,name=charColl" json:"charColl,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *TTable) Reset()                    { *m = TTable{} }
func (m *TTable) String() string            { return proto.CompactTextString(m) }
func (*TTable) ProtoMessage()               {}
func (*TTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TTable) GetID() int64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *TTable) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TTable) GetColumns() []*TColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *TTable) GetIndexes() []*TIndex {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *TTable) GetCharColl() *CharCollation {
	if m != nil {
		return m.CharColl
	}
	return nil
}

type Property struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Property) Reset()                    { *m = Property{} }
func (m *Property) String() string            { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()               {}
func (*Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Property) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Property) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*TableName)(nil), "tproto.TableName")
	proto.RegisterType((*CharCollation)(nil), "tproto.CharCollation")
	proto.RegisterType((*JoinColumn)(nil), "tproto.JoinColumn")
	proto.RegisterType((*Join)(nil), "tproto.Join")
	proto.RegisterType((*TIndexColumn)(nil), "tproto.TIndexColumn")
	proto.RegisterType((*TIndex)(nil), "tproto.TIndex")
	proto.RegisterType((*TColumn)(nil), "tproto.TColumn")
	proto.RegisterType((*FieldType)(nil), "tproto.FieldType")
	proto.RegisterType((*TTable)(nil), "tproto.TTable")
	proto.RegisterType((*Property)(nil), "tproto.Property")
	proto.RegisterEnum("tproto.JoinType", JoinType_name, JoinType_value)
}

var fileDescriptor0 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0x9c, 0xc4, 0xe3, 0x7e, 0xa4, 0xdb, 0x22, 0x59, 0x5c, 0x5a, 0x19, 0x01, 0x6d,
	0x0f, 0x3e, 0xe4, 0x8c, 0x38, 0x34, 0x1f, 0xd4, 0x25, 0x4a, 0xa2, 0xe2, 0xf4, 0xc0, 0xa5, 0xda,
	0xd8, 0x9b, 0x78, 0xc1, 0xf1, 0x1a, 0x7b, 0x83, 0xc8, 0x89, 0x2b, 0x27, 0x7e, 0x03, 0x3f, 0x95,
	0xdd, 0xb5, 0x9d, 0x3a, 0x20, 0x71, 0xb2, 0x66, 0x76, 0xde, 0xcc, 0x7b, 0x6f, 0xc6, 0x00, 0x9c,
	0x86, 0x0b, 0x37, 0xcd, 0x18, 0x67, 0xa8, 0xc5, 0xd5, 0xd7, 0xe9, 0x81, 0xe9, 0xe3, 0x45, 0x4c,
	0x26, 0x78, 0x4d, 0x10, 0x02, 0xc8, 0x83, 0x88, 0xac, 0xb1, 0x8c, 0x6c, 0xed, 0xa2, 0x71, 0x69,
	0xa2, 0x13, 0x30, 0x79, 0x55, 0x60, 0x37, 0x64, 0xca, 0x19, 0xc2, 0x61, 0x3f, 0xc2, 0x59, 0x9f,
	0xc5, 0x31, 0xe6, 0x94, 0x25, 0xc8, 0x86, 0x6e, 0x20, 0x12, 0x38, 0xe0, 0x24, 0xfb, 0x48, 0x78,
	0x89, 0xd6, 0x04, 0xfa, 0x05, 0xa0, 0xa0, 0x2a, 0x9b, 0x66, 0x21, 0xc9, 0xca, 0x36, 0xe2, 0xcd,
	0xf1, 0x00, 0xee, 0x18, 0x4d, 0x44, 0x9b, 0xcd, 0x3a, 0x41, 0x67, 0x70, 0x90, 0xe2, 0x8c, 0x24,
	0xbc, 0x88, 0x4b, 0xfc, 0x29, 0x58, 0x41, 0x44, 0xe3, 0xb0, 0x4c, 0x2a, 0x20, 0xea, 0x42, 0x27,
	0x65, 0x39, 0x95, 0x3d, 0x6d, 0x5d, 0x64, 0x0c, 0xe7, 0x01, 0x9a, 0xb2, 0x15, 0x7a, 0x09, 0xed,
	0x40, 0x55, 0xe6, 0xa2, 0x54, 0xbf, 0xb4, 0x7a, 0xc8, 0x2d, 0x74, 0xba, 0xb5, 0x49, 0x57, 0x70,
	0x14, 0xb0, 0x24, 0xe7, 0x19, 0xa6, 0x49, 0xc1, 0x55, 0x36, 0xb1, 0x7a, 0x27, 0x55, 0xed, 0xce,
	0x10, 0xe7, 0x2d, 0x1c, 0xf8, 0x5e, 0x12, 0x92, 0xef, 0x25, 0xf4, 0x00, 0x9a, 0x35, 0x71, 0x47,
	0xd0, 0x9a, 0x2e, 0x97, 0x39, 0xe1, 0x8a, 0x97, 0x21, 0xe3, 0x31, 0x49, 0x56, 0x3c, 0x2a, 0x59,
	0x51, 0x68, 0x15, 0x68, 0xe1, 0x6b, 0xc3, 0x1b, 0x28, 0x94, 0xb1, 0xeb, 0x51, 0x68, 0x11, 0x91,
	0x97, 0xcf, 0x3e, 0x28, 0x44, 0x47, 0x2a, 0xf3, 0xf2, 0x79, 0x42, 0xbf, 0x6e, 0x88, 0xdd, 0x54,
	0x99, 0x57, 0xd0, 0xee, 0x97, 0x8a, 0x0c, 0xa5, 0xe8, 0x6c, 0xc7, 0xb2, 0x46, 0xcc, 0xf9, 0x0c,
	0x6d, 0xbf, 0xe4, 0x58, 0xcd, 0x6a, 0x5c, 0xea, 0xb5, 0x59, 0x72, 0x95, 0xe7, 0xd0, 0xf4, 0xb7,
	0xe9, 0x3f, 0x72, 0x47, 0x94, 0xc4, 0xa1, 0x7c, 0xa8, 0x09, 0x6a, 0x2a, 0xaa, 0x62, 0x27, 0x21,
	0x59, 0xe2, 0x4d, 0xcc, 0x1f, 0x70, 0x2c, 0x28, 0x19, 0x6a, 0x6f, 0x3f, 0xc0, 0x7c, 0x82, 0x88,
	0x69, 0x7e, 0xaa, 0xa6, 0x29, 0x65, 0xa3, 0x18, 0xaf, 0xd4, 0xb4, 0xc3, 0x22, 0x22, 0x72, 0x43,
	0xf2, 0xed, 0x18, 0xda, 0x03, 0x12, 0xd0, 0x35, 0x8e, 0x45, 0x6f, 0x99, 0xb8, 0x06, 0x4b, 0x1e,
	0x91, 0x18, 0xe6, 0x25, 0x4b, 0xa6, 0x5a, 0x5b, 0xbd, 0xe7, 0x15, 0xa7, 0xfd, 0xfb, 0x3a, 0x04,
	0x63, 0x18, 0x93, 0x75, 0x6e, 0xb7, 0x84, 0x05, 0xa6, 0xf3, 0x4b, 0x13, 0xc6, 0xaa, 0x25, 0xd5,
	0x8c, 0xd5, 0xff, 0x32, 0xf6, 0xe2, 0xe9, 0x14, 0x74, 0x65, 0xdc, 0xf1, 0xce, 0xb8, 0xd2, 0xa8,
	0x73, 0x68, 0x53, 0x69, 0x21, 0xc9, 0x05, 0x25, 0x59, 0x71, 0xb4, 0x6f, 0x2d, 0x7a, 0x03, 0x9d,
	0xa0, 0xe4, 0xf1, 0x5f, 0x7e, 0xce, 0x6b, 0xe8, 0xcc, 0x32, 0x96, 0x92, 0x8c, 0x6f, 0x91, 0x05,
	0xfa, 0x17, 0xb2, 0x2d, 0x2f, 0x44, 0x10, 0xff, 0xa6, 0x9c, 0x53, 0x9c, 0xae, 0x3f, 0x41, 0x47,
	0xde, 0xa1, 0x32, 0xee, 0x14, 0x8e, 0xc7, 0xc3, 0x91, 0xff, 0x38, 0x9d, 0xfb, 0xc3, 0xfb, 0xc7,
	0xbb, 0xa9, 0x37, 0xe9, 0x3e, 0x13, 0x86, 0x77, 0xef, 0xbd, 0xf7, 0xb7, 0x7b, 0x59, 0x4d, 0x96,
	0x8e, 0xe6, 0xe3, 0x71, 0x3d, 0xd9, 0x10, 0xbb, 0x02, 0x6f, 0x32, 0xa9, 0x62, 0xfd, 0xe6, 0x1d,
	0x5c, 0xb1, 0x6c, 0xe5, 0xe2, 0x14, 0x8b, 0x3f, 0xd8, 0x8d, 0x70, 0xc8, 0x58, 0xea, 0x46, 0x0b,
	0x9c, 0x13, 0x57, 0xfd, 0xf6, 0x41, 0xea, 0xae, 0x48, 0x42, 0x32, 0xcc, 0x49, 0x78, 0x63, 0xfa,
	0x74, 0x70, 0x33, 0x93, 0x4a, 0x6e, 0xb5, 0x9f, 0x9a, 0xf6, 0x5b, 0xd3, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x15, 0x32, 0x65, 0x14, 0x1d, 0x04, 0x00, 0x00,
}