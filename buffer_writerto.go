package writertoutils

import "io"

type bufferWriterTo struct {
	buffer []byte
}

func (writerTo bufferWriterTo) WriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write(writerTo.buffer)
	return int64(n), err
}

// BufferWriterTo writes the same buffer each time the WriteTo metho is called.
func BufferWriterTo(buffer []byte) io.WriterTo {
	return bufferWriterTo{buffer}
}
