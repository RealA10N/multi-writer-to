package writertoutils_test

import (
	"bytes"
	"errors"
	"testing"

	"alon.kr/x/writertoutils"
	"github.com/stretchr/testify/assert"
)

type MarshalerMock struct {
	Data []byte
	Err  error
}

func (marshaler MarshalerMock) MarshalBinary() (data []byte, err error) {
	return marshaler.Data, marshaler.Err
}

func TestHappyFlow(t *testing.T) {
	data := []byte{'h', 'i', '!', 0}
	marshaler := MarshalerMock{Data: data}

	writerTo := writertoutils.BinaryMarshalerAdapter(marshaler)
	buffer := new(bytes.Buffer)
	n, err := writerTo.WriteTo(buffer)

	assert.NoError(t, err)
	assert.EqualValues(t, n, len(data))
	assert.Equal(t, data, buffer.Bytes())
}

func TestMarshalerError(t *testing.T) {
	data := []byte("hello, ")
	expectedErr := errors.New("failed converting to binary")
	marshaler := MarshalerMock{Data: data, Err: expectedErr}

	writerTo := writertoutils.BinaryMarshalerAdapter(marshaler)
	buffer := new(bytes.Buffer)
	n, err := writerTo.WriteTo(buffer)

	assert.ErrorIs(t, expectedErr, err)
	assert.EqualValues(t, 0, n)
	assert.Empty(t, buffer)
}
