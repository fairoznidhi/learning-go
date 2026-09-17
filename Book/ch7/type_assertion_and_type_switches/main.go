package main

import (
	"bufio"
	"fmt"
	"io"
)

type FastReader struct{}

func (f FastReader) Read(p []byte) (int, error) { return 0, nil }

func (f FastReader) WriteTo(w io.Writer) (int64, error) { return 0, nil }

func main() {
	var r io.Reader = FastReader{}
	_, ok := r.(io.WriterTo)
	fmt.Println(ok)

	buffered := bufio.NewReader(FastReader{})
	_, ok2 := interface{}(buffered).(io.WriterTo)
	fmt.Println(ok2)
}
