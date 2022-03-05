package main

type Class Object

type ClassBlock struct {
	state             U1
	flags             U2
	accessFlags       U2
	declaringClass    U2
	enclosingClass    U2
	enclosingMethod   U2
	innerAccessFlags  U2
	fieldsCount       U2
	methodsCount      U2
	interfacesCount   U2
	innerClassCount   U2
	constantPoolCount U2
	objectSize        int
	methodTableSize   int
	iMethodTableSize  int
	initingTid        int
	name              string
	signature         string
	sourceFileName    string
	super             *Class
	fields            *FieldBlock
	methods           *MethodBlock
	interfaces        **Class
	methodTable       **MethodBlock
	iMethodTable      *ITableEntry
	classLoader       *Object
	innerClasses      *U2
	bootstrapMethods  string
	extraAttributes   *ExtraAttributes
	constantPool      ConstantPool
}

type FieldBlock struct{}

type MethodBlock struct{}

type ITableEntry struct{}

type ExtraAttributes struct{}

// type ConstantPool struct{}

/*
func ParseClass(classname string, data bytes.Buffer, offset int, len int, classLoader *Object) *Class {
}

func DefineClass(classname string, data bytes.Buffer, offset int, len int, classLoader *Object) *Class {
}

func LinkClass(class *Class) {
}

func InitClass(class *Class) {
}

func FindSystemClass0(name string) *Class {
}

func LoadSystemClass(name string) *Class {
}

func FindPrimitiveClass(name rune) *Class {
}

func FindPrimitiveClassByName(name string) *Class {
}

func FindHashedClass(name string, loader *Object) *Class {
}

func FindClassFromClassLoader(name string, loader *Object) *Class {
}

func FindArrayClassFromClassLoader(name string, loader *Object) *Class {
}

func GetSystemClassLoader() *Object {
}

func BootClassPathSize() int {
}

func GetBootClassPathEntry(index int) string {
}

func BootClassPathResource(filename string, index int) *Object {
}

func FreeClassLoaderData(classLoader *Object) {
}

func FreeClassData(class *Class) {
}

func GetClassPath() string {
}

func GetBootClassPath() string {
}

func GetEndorsedDirs() string {
}

func MarkBootClasses() {
}

func MarkLoaderClasses(loader *Object, mark int) {
}

func ThreadBootClasses() {
}

func ThreadLoaderClasses(classLoader *Object) {
}

func NewLibraryUnloader(classLoader *Object, entry interface{}) {
}

func InitialiseClassStage1(args *InitArgs) int {
}

func InitialiseClassStage2() int {
}

func BootPackage(packageName string) *Object {
}

func BootPackages() *Object {
}

func HideFieldFromGC(hidden *FieldBlock) int {
}
*/
