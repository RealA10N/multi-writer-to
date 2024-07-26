package writertoutils

import "io"

type multiWriterTo struct {
	writerTos []io.WriterTo
}

// Implements the io.WriterTo interface.
//
// WriteTo writes data from all of the provided writers to w, one after another
// until all data has been written or an error has occurred. The return value n
// is the number of total bytes written. If an error has been encountered during
// the write, it is also returned.
func MultiWriterTo(writerTos ...io.WriterTo) io.WriterTo {
	return multiWriterTo{writerTos: writerTos}
}

// writes data from all of the provided writers to w, one after another
// until all data has been written or an error has occurred.
// The return value n is the number of total bytes written.
// If an error has been encountered during the write, it is also returned.
func (multiWriter multiWriterTo) WriteTo(destWriter io.Writer) (n int64, err error) {
	var k int64

	for _, srcWriter := range multiWriter.writerTos {
		k, err = srcWriter.WriteTo(destWriter)
		n += k
		if err != nil {
			return
		}
	}

	return
}
