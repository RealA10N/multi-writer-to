package writertoutils_test

import (
	"bytes"
	"testing"

	"alon.kr/x/writertoutils"
	"github.com/stretchr/testify/assert"
)

func TestBufferWriterToHappyFlow(t *testing.T) {
	expected := []byte{1, 2, 3}
	writerTo := writertoutils.BufferWriterTo(expected)

	// loop two times to ensure that we can call WriteTo multiple times and
	// expect the same result.
	for i := 0; i < 2; i++ {
		buffer := bytes.Buffer{}
		n, err := writerTo.WriteTo(&buffer)

		assert.NoError(t, err)
		assert.EqualValues(t, len(expected), n)
		assert.Equal(t, expected, buffer.Bytes())
	}
}
