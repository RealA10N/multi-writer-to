package multiwriterto

import "io"

type multiWriterTo struct {
	writerTos []io.WriterTo
}

func MultiWriterTo(writerTos ...io.WriterTo) io.WriterTo {
	return multiWriterTo{writerTos: writerTos}
}

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
