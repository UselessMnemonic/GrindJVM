package main

import (
	"encoding/binary"
	"errors"
	"os"
)

type U1 = uint8
type U2 = uint16
type U4 = uint32
type U8 = uint64

type JByte = int8
type JShort = int16
type JInt = int32
type JLong = int64

func ReadU1(file *os.File) (v U1, err error) {
	bytes := make([]byte, 1)
	num, err := file.Read(bytes)
	if err != nil {
		return
	}
	if num != 1 {
		err = errors.New("did not read enough bytes")
	}
	v = bytes[0]
	return
}

func ReadU2(file *os.File) (v U2, err error) {
	bytes := make([]byte, 2)
	num, err := file.Read(bytes)
	if err != nil {
		return
	}
	if num != 2 {
		err = errors.New("did not read enough bytes")
		return
	}
	v = binary.BigEndian.Uint16(bytes)
	return
}

func ReadU4(file *os.File) (v U4, err error) {
	bytes := make([]byte, 4)
	num, err := file.Read(bytes)
	if err != nil {
		return
	}
	if num != 4 {
		err = errors.New("did not read enough bytes")
		return
	}
	v = binary.BigEndian.Uint32(bytes)
	return
}
