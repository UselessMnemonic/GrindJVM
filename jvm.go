package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("HelloWorld.class", os.O_RDONLY, 0)
	if err != nil {
		return
	}

	classFile, err := readClassFile(file)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Magic: %#x \n", classFile.Magic)
	fmt.Printf("Minor Version: %d \n", classFile.MinorVersion)
	fmt.Printf("Major Version: %d \n", classFile.MajorVersion)

	fmt.Printf("Constant Pool:\n")
	for i, info := range classFile.CPInfo {
		fmt.Printf("  Constant #%d %T %v \n", i+1, info, info)
	}

	fmt.Printf("Access Flags: %#x \n", classFile.AccessFlags)
	fmt.Printf("This Class: %d \n", classFile.ThisClass)
	fmt.Printf("Super Class: %d \n", classFile.SuperClass)
	fmt.Printf("Interfaces: %v \n", classFile.Interfaces)

	fmt.Printf("Fields (%d):\n", len(classFile.Fields))
	for _, field := range classFile.Fields {
		fmt.Printf("  Field: %v\n", field)
	}

	fmt.Printf("Methods (%d):\n", len(classFile.Methods))
	for _, method := range classFile.Methods {
		fmt.Printf("  Method: %v\n", method)
	}

	fmt.Printf("Attributes (%d):\n", len(classFile.Attributes))
	for _, attr := range classFile.Attributes {
		fmt.Printf("  Attribute: %v\n", attr)
	}
}
