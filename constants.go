package main

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_Dynamic            = 17
	CONSTANT_InvokeDynamic      = 18
	CONSTANT_Module             = 19
	CONSTANT_Package            = 20
)

type PoolConstant interface {
	Tag() U1
}

type ConstantPool = []PoolConstant

type ClassConstant struct {
	NameIndex U2
}

func (_ ClassConstant) Tag() U1 {
	return CONSTANT_Class
}

type FieldRefConstant struct {
	ClassIndex       U2
	NameAndTypeIndex U2
}

func (_ FieldRefConstant) Tag() U1 {
	return CONSTANT_Fieldref
}

type MethodRefConstant struct {
	ClassIndex       U2
	NameAndTypeIndex U2
}

func (_ MethodRefConstant) Tag() U1 {
	return CONSTANT_Methodref
}

type InterfaceMethodRefConstant struct {
	ClassIndex       U2
	NameAndTypeIndex U2
}

func (_ InterfaceMethodRefConstant) Tag() U1 {
	return CONSTANT_InterfaceMethodref
}

type StringConstant struct {
	StringIndex U2
}

func (_ StringConstant) Tag() U1 {
	return CONSTANT_String
}

type IntegerConstant struct {
	Bytes U4
}

type FloatConstant struct {
	Bytes U4
}

type LongConstant struct {
	High U4
	Low  U4
}

type DoubleConstant struct {
	High U4
	Low  U4
}

type NameAndTypeConstant struct {
	NameIndex       U2
	DescriptorIndex U2
}

func (_ NameAndTypeConstant) Tag() U1 {
	return CONSTANT_NameAndType
}

type UTF8Constant struct {
	Str string
}

func (_ UTF8Constant) Tag() U1 {
	return CONSTANT_Utf8
}

type MethodHandleConstant struct {
	ReferenceKind  U1
	ReferenceIndex U2
}

type MethodTypeConstant struct {
	DescriptorIndex U2
}

type DynamicConstant struct {
	BootstrapMethodAttributeIndex U2
	NameAndTypeIndex              U2
}

type InvokeDynamicConstant struct {
}

type ModuleConstant struct {
	NameIndex U2
}

type PackageConstant struct {
	NameIndex U2
}
