package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type AttributeReader = func(io.Reader, ConstantPool) (AttributeInfo, error)

var AttributeReaders = make(map[string]AttributeReader)

func init() {
	AttributeReaders["ConstantValue"] = readConstantValueAttribute
	AttributeReaders["Code"] = readCodeAttribute
}

type AttributeInfo interface {
	Name() string
}

func ReadAttributes(r io.Reader, pool ConstantPool) (attrs []AttributeInfo, err error) {
	attributeCount, err := ReadU2(r)
	attrs = make([]AttributeInfo, attributeCount)
	for i := range attrs {
		attrs[i], err = ReadAttribute(r, pool)
		if err != nil {
			return
		}
	}
	return
}

func ReadAttribute(r io.Reader, pool ConstantPool) (attr AttributeInfo, err error) {
	nameIndex, err := ReadU2(r)
	length, err := ReadU4(r)

	name := pool[nameIndex].(*UTF8Constant)
	reader, ok := AttributeReaders[name.Str]

	if !ok {
		data := make([]U1, length)
		_, err = r.Read(data)
		attr = &UnknownAttribute{
			name.Str,
			data,
		}
		return
	}

	attr, err = reader(r, pool)
	return
}

type UnknownAttribute struct {
	name string
	data []U1
}

func (ua *UnknownAttribute) Name() string {
	return ua.name
}

func (ua *UnknownAttribute) String() string {
	return fmt.Sprintf("<%s (%d)>", ua.name, len(ua.data))
}

type ConstantValueAttribute struct {
	constantValueIndex U2
}

func (*ConstantValueAttribute) Name() string {
	return "ConstantValue"
}

func readConstantValueAttribute(r io.Reader, _ ConstantPool) (attr AttributeInfo, err error) {
	attr = new(ConstantValueAttribute)
	err = binary.Read(r, binary.BigEndian, attr)
	return
}

type CodeAttribute struct {
	maxStack       U2
	maxLocals      U2
	code           []U1
	exceptionTable []ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   U2
	endPc     U2
	handlerPc U2
	catchType U2
}

func (*CodeAttribute) Name() string {
	return "Code"
}

func readCodeAttribute(r io.Reader, pool ConstantPool) (attr AttributeInfo, err error) {
	maxStack, err := ReadU2(r)
	maxLocals, err := ReadU2(r)

	codeLength, err := ReadU4(r)
	code := make([]U1, codeLength)
	_, err = r.Read(code)

	exceptionTableLength, err := ReadU2(r)
	exceptionTable := make([]ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		err = binary.Read(r, binary.BigEndian, &exceptionTable[i])
		if err != nil {
			return
		}
	}

	attributesCount, err := ReadU2(r)
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i], err = readAttribute(r, pool)
		if err != nil {
			return
		}
	}

	attr = &CodeAttribute{
		maxStack,
		maxLocals,
		code,
		exceptionTable,
		attributes,
	}
	return
}

type StackMapTableAttribute []StackMapFrame

func (StackMapTableAttribute) Name() string {
	return "StackMapTable"
}

type VerificationTypeInfo struct {
	Tag               U1
	PoolIndexOrOffset U2
}

type StackMapFrame struct {
	FrameType U1
	*SameLocals1StackItemFrame
	*SameLocals1StackItemFrameExtended
	*ChopFrame
	*SameFrameExtended
	*AppendFrame
	*FullFrame
}

type SameLocals1StackItemFrame struct {
	StackItem VerificationTypeInfo
}

type SameLocals1StackItemFrameExtended struct {
	OffsetDelta U2
	StackItem   VerificationTypeInfo
}

type ChopFrame struct {
	OffsetDelta U2
}

type SameFrameExtended struct {
	OffsetDelta U2
}

type AppendFrame struct {
	OffsetDelta U2
	Locals      []VerificationTypeInfo
}

type FullFrame struct {
	OffsetDelta U2
	Locals      []VerificationTypeInfo
	Stack       []VerificationTypeInfo
}

func readStackMapTableAttribute(r io.Reader) (attr AttributeInfo, err error) {
	numberOfEntries, err := ReadU1(r)
	table := make(StackMapTableAttribute, numberOfEntries)
	for i := range table {
		var frame StackMapFrame
		frame.FrameType, err = ReadU1(r)
		if frame.FrameType == 255 {
		} else if frame.FrameType > 251 {
		} else if frame.FrameType == 251 {
		} else if frame.FrameType > 247 {
		} else if frame.FrameType == 247 {
		} else if frame.FrameType > 127 {
			// err
		} else if frame.FrameType > 63 {
		} /* else frame is same_frame */
		table[i] = frame
	}
	attr = table
	return
}

type ExceptionsAttribute struct {
}

func (*ExceptionsAttribute) Name() string {
	return "Exceptions"
}

type BootstrapMethodsAttribute struct {
}

func (*BootstrapMethodsAttribute) Name() string {
	return "StackMapTable"
}

type NestHostAttribute struct {
}

func (*NestHostAttribute) Name() string {
	return "NestHost"
}

type NestMembersAttribute struct{}

func (*NestMembersAttribute) Name() string {
	return "NestMembers"
}

type PermittedSubclassesAttribute struct{}

func (*PermittedSubclassesAttribute) Name() string {
	return "PermittedSubclasses"
}

type InnerClassesAttribute struct{}

func (*InnerClassesAttribute) Name() string {
	return "InnerClasses"
}

type EnclosingMethodAttribute struct{}

func (*EnclosingMethodAttribute) Name() string {
	return "EnclosingMethod"
}

type SyntheticAttribute struct{}

func (*SyntheticAttribute) Name() string {
	return "Synthetic"
}

type SignatureAttribute struct{}

func (*SignatureAttribute) Name() string {
	return "Signature"
}

type RecordAttribute struct{}

func (*RecordAttribute) Name() string {
	return "Record"
}

type SourceFileAttribute struct{}

func (*SourceFileAttribute) Name() string {
	return "SourceFile"
}

type LineNumberTableAttribute struct{}

func (*LineNumberTableAttribute) Name() string {
	return "LineNumberTable"
}

type LocalVariableTableAttribute struct{}

func (*LocalVariableTableAttribute) Name() string {
	return "LocalVariableTable"
}

type LocalVariableTypeTableAttribute struct{}

func (*LocalVariableTypeTableAttribute) Name() string {
	return "LocalVariableTypeTable"
}

type SourceDebugExtensionAttribute struct{}

func (*SourceDebugExtensionAttribute) Name() string {
	return "SourceDebugExtension"
}

type DeprecatedAttribute struct{}

func (*DeprecatedAttribute) Name() string {
	return "Deprecated"
}

type RuntimeVisibleAnnotationsAttribute struct{}

func (*RuntimeVisibleAnnotationsAttribute) Name() string {
	return "RuntimeVisibleAnnotations"
}

type RuntimeInvisibleAnnotationsAttribute struct{}

func (*RuntimeInvisibleAnnotationsAttribute) Name() string {
	return "RuntimeInvisibleAnnotations"
}

type RuntimeVisibleParameterAnnotationsAttribute struct{}

func (*RuntimeVisibleParameterAnnotationsAttribute) Name() string {
	return "RuntimeVisibleParameterAnnotations"
}

type RuntimeInvisibleParameterAnnotationsAttribute struct{}

func (*RuntimeInvisibleParameterAnnotationsAttribute) Name() string {
	return "RuntimeInvisibleParameterAnnotations"
}

type RuntimeVisibleTypeAnnotationsAttribute struct{}

func (*RuntimeVisibleTypeAnnotationsAttribute) Name() string {
	return "RuntimeVisibleTypeAnnotations"
}

type RuntimeInvisibleTypeAnnotationsAttribute struct{}

func (*RuntimeInvisibleTypeAnnotationsAttribute) Name() string {
	return "RuntimeInvisibleTypeAnnotations"
}

type AnnotationDefaultAttribute struct{}

func (*AnnotationDefaultAttribute) Name() string {
	return "AnnotationDefault"
}

type MethodParametersAttribute struct{}

func (*MethodParametersAttribute) Name() string {
	return "MethodParameters"
}

type ModuleAttribute struct{}

func (*ModuleAttribute) Name() string {
	return "Module"
}

type ModulePackagesAttribute struct{}

func (*ModulePackagesAttribute) Name() string {
	return "ModulePackages"
}

type ModuleMainClassAttribute struct{}

func (*ModuleMainClassAttribute) Name() string {
	return "ModuleMainClass"
}
