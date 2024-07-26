package writertoutils

import (
	"encoding"
	"io"
)

// An adapter from the encoding.BinaryMarshaler interface to the io.WriterTo
// interface.
type binaryMarshalerAdapter struct {
	marshaler encoding.BinaryMarshaler
}

func (adapter binaryMarshalerAdapter) WriteTo(writer io.Writer) (int64, error) {
	data, err := adapter.marshaler.MarshalBinary()
	if err != nil {
		return 0, err
	}

	n, err := writer.Write(data)
	return int64(n), err
}

// An adapter from the encoding.BinaryMarshaler interface to the io.WriterTo
// interface.
func BinaryMarshalerAdapter(marshaler encoding.BinaryMarshaler) io.WriterTo {
	return binaryMarshalerAdapter{marshaler: marshaler}
}
