# MultiWriterTo

[![Go Reference](https://pkg.go.dev/badge/github.com/RealA10N/multi-writer-to.svg)](https://pkg.go.dev/github.com/RealA10N/multi-writer-to)
[![codecov](https://codecov.io/gh/RealA10N/multi-writer-to/graph/badge.svg?token=SDfCQpOpHn)](https://codecov.io/gh/RealA10N/multi-writer-to)
[![CI](https://github.com/RealA10N/multi-writer-to/actions/workflows/ci.yml/badge.svg)](https://github.com/RealA10N/multi-writer-to/actions/workflows/ci.yml)

A simple utility that concatenates multiple `WriterTo` objects, into a single
writing stream.

```go
func main() {
    r1 := strings.NewReader("hello, ")  // implements io.WriterTo
    r2 := strings.NewReader("world!") // implements io.WriterTo

    multiWriter := multiwriterto.MultiWriterTo(r1, r2)
    buffer := new(strings.Builder)

    // writes "hello, world!" to the buffer!
    n, err := multiWriter.WriteTo(buffer)
}
```
