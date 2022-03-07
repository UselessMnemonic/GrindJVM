package main

import (
	"encoding/binary"
	"errors"
	"io"
)

type U1 = uint8
type U2 = uint16
type U4 = uint32
type U8 = uint64

type JByte = int8
type JShort = int16
type JInt = int32
type JLong = int64

func ReadU1(r io.Reader) (v U1, err error) {
	bytes := make([]byte, 1)
	num, err := r.Read(bytes)
	if err != nil {
		return
	}
	if num != 1 {
		err = errors.New("did not read enough bytes")
	}
	v = bytes[0]
	return
}

func ReadU2(r io.Reader) (v U2, err error) {
	bytes := make([]byte, 2)
	num, err := r.Read(bytes)
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

func ReadU4(r io.Reader) (v U4, err error) {
	bytes := make([]byte, 4)
	num, err := r.Read(bytes)
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
