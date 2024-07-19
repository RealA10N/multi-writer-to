package multiwriterto_test

import (
	"strings"
	"testing"

	multiwriterto "github.com/RealA10N/multi-writer-to"
	"github.com/stretchr/testify/assert"
)

func TestSingleWriterTo(t *testing.T) {
	reader := strings.NewReader("hello, world!")

	multiWriter := multiwriterto.MultiWriterTo(reader)
	output := new(strings.Builder)

	multiWriter.WriteTo(output)
	assert.Equal(t, "hello, world!", output.String())
}

func TestMultipleWriteTo(t *testing.T) {
	r1 := strings.NewReader("hello,")
	r2 := strings.NewReader(" world!")

	multiWriter := multiwriterto.MultiWriterTo(r1, r2)
	output := new(strings.Builder)

	multiWriter.WriteTo(output)
	assert.Equal(t, "hello, world!", output.String())
}
