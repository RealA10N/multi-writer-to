package writertoutils_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	writertoutils "github.com/RealA10N/writer-to-utils"
	"github.com/stretchr/testify/assert"
)

func TestSingleWriterTo(t *testing.T) {
	reader := strings.NewReader("hello, world!")

	multiWriter := writertoutils.MultiWriterTo(reader)
	output := new(strings.Builder)

	n, err := multiWriter.WriteTo(output)
	const expected = "hello, world!"

	assert.NoError(t, err)
	assert.EqualValues(t, n, len(expected))
	assert.Equal(t, expected, output.String())
}

func TestMultipleWriteTo(t *testing.T) {
	r1 := strings.NewReader("hello,")
	r2 := strings.NewReader(" world!")

	multiWriter := writertoutils.MultiWriterTo(r1, r2)
	output := new(strings.Builder)

	n, err := multiWriter.WriteTo(output)
	const expected = "hello, world!"

	assert.NoError(t, err)
	assert.EqualValues(t, len(expected), n)
	assert.Equal(t, expected, output.String())
}

type CorruptedWriterTo struct {
	str string
	err error
}

// This function simulates a WriteTo that fails for some
// external error, while being able to only write some
// of the intended message.
func (corrupted CorruptedWriterTo) WriteTo(w io.Writer) (int64, error) {

	n, err := io.WriteString(w, corrupted.str)

	if err != nil || n != len(corrupted.str) {
		panic("unexpected error")
	}

	return int64(n), corrupted.err
}
func TestErrorInMiddleOfWrite(t *testing.T) {
	expectedErr := errors.New("network error!")

	r1 := strings.NewReader("hello, ")
	r2 := CorruptedWriterTo{str: "wor", err: expectedErr}

	multiWriter := writertoutils.MultiWriterTo(r1, r2)
	output := new(strings.Builder)

	n, err := multiWriter.WriteTo(output)
	const expected = "hello, wor"

	assert.ErrorIs(t, err, expectedErr)
	assert.EqualValues(t, len(expected), n)
	assert.Equal(t, expected, output.String())
}
