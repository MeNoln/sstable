package sstable

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const headerLen int = 8

type Header struct {
	FileLen     uint32 // len = 4
	IndexOffset uint32 // len = 4
}

func NewHeader(fl, io uint32) Header {
	return Header{
		FileLen:     fl,
		IndexOffset: io,
	}
}

func ReadHeader(r io.Reader) (Header, error) {
	h := Header{}

	bhead := make([]byte, headerLen)
	if _, err := r.Read(bhead); err != nil {
		return h, err
	}

	h, err := UnmarshalHeader(bhead)
	if err != nil {
		return h, err
	}

	return h, nil
}

func (h *Header) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	if err := binary.Write(buf, binary.BigEndian, h); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func UnmarshalHeader(data []byte) (Header, error) {
	h := Header{}

	if len(data) < 8 {
		return h, fmt.Errorf("header data is less then required")
	}

	toff := binary.BigEndian.Uint32(data[:4])
	ioff := binary.BigEndian.Uint32(data[4:8])

	h.FileLen = toff
	h.IndexOffset = ioff

	return h, nil
}
