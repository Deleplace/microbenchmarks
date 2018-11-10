package indirectcall

import (
	"bytes"
	"encoding/binary"
	"io"
)

func write1(w io.Writer, a []int32) error {
	return binary.Write(w, binary.LittleEndian, a)
}

func write2(w *bytes.Buffer, a []int32) error {
	return binary.Write(w, binary.LittleEndian, a)
}

func write3(w io.Writer, a []int32) error {
	var buf [4]byte
	for _, v := range a {
		if v%42 != 0 {
			_, err := w.Write(buf[:])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func write4(w *bytes.Buffer, a []int32) error {
	var buf [4]byte
	for _, v := range a {
		if v%42 != 0 {
			_, err := w.Write(buf[:])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func write3or4(w io.Writer, a []int32) error {
	switch w := w.(type) {
	case *bytes.Buffer:
		return write4(w, a)
	default:
		return write3(w, a)
	}
}
