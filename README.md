# WriterTo Utils

[![Go Reference](https://pkg.go.dev/badge/github.com/RealA10N/writer-to-utils.svg)](https://pkg.go.dev/github.com/RealA10N/writer-to-utils)
[![codecov](https://codecov.io/gh/RealA10N/writer-to-utils/graph/badge.svg?token=SDfCQpOpHn)](https://codecov.io/gh/RealA10N/writer-to-utils)
[![CI](https://github.com/RealA10N/writer-to-utils/actions/workflows/ci.yml/badge.svg)](https://github.com/RealA10N/writer-to-utils/actions/workflows/ci.yml)

Useful functions & types that extend the functionality of the `io.WriterTo` interface in Golang.

## MultiWriterTo

A simple utility that concatenates multiple `io.WriterTo` objects, into a single
writing stream.

```go
func main() {
    r1 := strings.NewReader("hello, ")  // implements io.WriterTo
    r2 := strings.NewReader("world!") // implements io.WriterTo

    multiWriter := multiwriterto.MultiWriterTo(r1, r2)
    buffer := new(bytes.Buffer)

    // writes "hello, world!" to the buffer!
    n, err := multiWriter.WriteTo(buffer)
}
```

## BinaryMarshalerAdapter

Convert any `encoding.BinaryMarshaler` interface to a `io.WriterTo` interface.

```go
func main() {
    marshaler := MyMarshaler() // implements: MarshalBinary() (data []byte, err error)
    
    writerTo := BinaryMarshalerAdapter(marshaler)
    buffer := new(bytes.Buffer)

    // calls MarshalBinary() and writes the binary data to the buffer! 
    n, err := writerTo.WriteTo(buffer)
}
```