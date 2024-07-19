# MultiWriterTo

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
