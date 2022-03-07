package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type ClassFile struct {
	Magic        U4
	MinorVersion U2
	MajorVersion U2
	CPInfo       ConstantPool
	AccessFlags  U2
	ThisClass    U2
	SuperClass   U2
	Interfaces   []U2
	Fields       []FieldInfo
	Methods      []MethodInfo
	Attributes   []AttributeInfo
}

func readClassFile(r io.Reader) (classFile *ClassFile, err error) {
	classFile = new(ClassFile)
	classFile.Magic, err = ReadU4(r)
	classFile.MinorVersion, err = ReadU2(r)
	classFile.MajorVersion, err = ReadU2(r)

	cpCount, err := ReadU2(r)
	classFile.CPInfo = make(ConstantPool, cpCount)
	for i := 1; i < len(classFile.CPInfo); i++ {
		classFile.CPInfo[i], err = readConstantPool(r)
		if err != nil {
			return
		}
	}

	classFile.AccessFlags, err = ReadU2(r)
	classFile.ThisClass, err = ReadU2(r)
	classFile.SuperClass, err = ReadU2(r)

	interfacesCount, err := ReadU2(r)
	classFile.Interfaces = make([]U2, interfacesCount)
	for i := range classFile.Interfaces {
		classFile.Interfaces[i], err = ReadU2(r)
		if err != nil {
			return
		}
	}

	fieldsCount, err := ReadU2(r)
	classFile.Fields = make([]FieldInfo, fieldsCount)
	for i := range classFile.Fields {
		classFile.Fields[i], err = readFieldInfo(r, classFile.CPInfo)
		if err != nil {
			return
		}
	}

	methodsCount, err := ReadU2(r)
	classFile.Methods = make([]MethodInfo, methodsCount)
	for i := range classFile.Methods {
		classFile.Methods[i], err = readMethodInfo(r, classFile.CPInfo)
		if err != nil {
			return
		}
	}

	attributesCount, err := ReadU2(r)
	classFile.Attributes = make([]AttributeInfo, attributesCount)
	for i := range classFile.Attributes {
		classFile.Attributes[i], err = readAttribute(r, classFile.CPInfo)
		if err != nil {
			return
		}
	}
	return
}

func readConstantPool(r io.Reader) (cpInfo PoolConstant, err error) {
	tag, err := ReadU1(r)
	switch tag {
	case CONSTANT_Class:
		cpInfo = new(ClassConstant)
		err = binary.Read(r, binary.BigEndian, cpInfo)

	case CONSTANT_Fieldref:
		cpInfo = new(FieldRefConstant)
		err = binary.Read(r, binary.BigEndian, cpInfo)

	case CONSTANT_Methodref:
		cpInfo = new(MethodRefConstant)
		err = binary.Read(r, binary.BigEndian, cpInfo)

	case CONSTANT_InterfaceMethodref:
		cpInfo = new(InterfaceMethodRefConstant)
		err = binary.Read(r, binary.BigEndian, cpInfo)

	case CONSTANT_String:
		cpInfo = new(StringConstant)
		err = binary.Read(r, binary.BigEndian, cpInfo)

	case CONSTANT_Integer:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_Float:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_Long:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_Double:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_NameAndType:
		cpInfo = new(NameAndTypeConstant)
		err = binary.Read(r, binary.BigEndian, cpInfo)

	case CONSTANT_Utf8:
		length, _ := ReadU2(r)
		bytes := make([]U1, length)
		_, err = r.Read(bytes)
		cpInfo = &UTF8Constant{
			string(bytes),
		}

	case CONSTANT_MethodHandle:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_MethodType:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_Dynamic:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_InvokeDynamic:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_Module:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	case CONSTANT_Package:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))

	default:
		err = errors.New(fmt.Sprintf("unsupported constant type %d\n", tag))
	}
	return
}

func readMethodOrFieldInfo(r io.Reader, pool ConstantPool) (accessFlags, nameIndex, descriptorIndex U2, attributes []AttributeInfo, err error) {
	accessFlags, err = ReadU2(r)
	nameIndex, err = ReadU2(r)
	descriptorIndex, err = ReadU2(r)
	attributes, err = ReadAttributes(r, pool)
	return
}

func readMethodInfo(r io.Reader, pool ConstantPool) (info MethodInfo, err error) {
	info.AccessFlags, info.NameIndex, info.DescriptorIndex, info.Attributes, err = readMethodOrFieldInfo(r, pool)
	info.Name = pool[info.NameIndex].(*UTF8Constant).Str
	info.Descriptor = pool[info.DescriptorIndex].(*UTF8Constant).Str
	for _, attr := range info.Attributes {
		if ptr, ok := attr.(*CodeAttribute); ok {
			info.Code = ptr
			break
		}
	}
	return
}

func readFieldInfo(r io.Reader, pool ConstantPool) (info FieldInfo, err error) {
	info.AccessFlags, info.NameIndex, info.DescriptorIndex, info.Attributes, err = readMethodOrFieldInfo(r, pool)
	return
}

const (
	ACC_PUBLIC    = 0x0001
	ACC_PRIVATE   = 0x0002
	ACC_PROTECTED = 0x0004
	ACC_STATIC    = 0x0008
	ACC_FINAL     = 0x0010
	ACC_VOLATILE  = 0x0040
	ACC_TRANSIENT = 0x0080
	ACC_SYNTHETIC = 0x1000
	ACC_ENUM      = 0x4000
)

type FieldInfo struct {
	AccessFlags     U2
	NameIndex       U2
	DescriptorIndex U2
	Attributes      []AttributeInfo
}

const (
	ACC_SYNCHRONIZED = 0x0020
	ACC_BRIDGE       = 0x0040
	ACC_VARARGS      = 0x0080
	ACC_NATIVE       = 0x0100
	ACC_ABSTRACT     = 0x0400
	ACC_STRICT       = 0x0800
)

type MethodInfo struct {
	AccessFlags     U2
	NameIndex       U2
	DescriptorIndex U2
	Attributes      []AttributeInfo

	Name       string
	Descriptor string
	Code       *CodeAttribute
}
