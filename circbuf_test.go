package circbuf

import (
	"bytes"
	"io"
	"testing"
)

func TestBuffer_implWriter(t *testing.T) {
	buf, err := NewBuffer(1)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	
	var _ io.Writer = buf
}

func TestBuffer_ShortWrite(t *testing.T) {
	buf, err := NewBuffer(1024)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	inp := []byte("hello world")

	n, err := buf.Write(inp)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if n != len(inp) {
		t.Fatalf("bad: %v", n)
	}

	if !bytes.Equal(buf.Bytes(), inp) {
		t.Fatalf("bad: %v", buf.Bytes())
	}
}

func TestBuffer_FullWrite(t *testing.T) {
	inp := []byte("hello world")

	buf, err := NewBuffer(int64(len(inp)))
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	n, err := buf.Write(inp)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if n != len(inp) {
		t.Fatalf("bad: %v", n)
	}

	if !bytes.Equal(buf.Bytes(), inp) {
		t.Fatalf("bad: %v", buf.Bytes())
	}
}

func TestBuffer_LongWrite(t *testing.T) {
	inp := []byte("hello world")

	buf, err := NewBuffer(6)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	n, err := buf.Write(inp)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if n != len(inp) {
		t.Fatalf("bad: %v", n)
	}

	expect := []byte(" world")
	if !bytes.Equal(buf.Bytes(), expect) {
		t.Fatalf("bad: %s", buf.Bytes())
	}
}

func TestBuffer_HugeWrite(t *testing.T) {
	inp := []byte("hello world")

	buf, err := NewBuffer(3)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	n, err := buf.Write(inp)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if n != len(inp) {
		t.Fatalf("bad: %v", n)
	}

	expect := []byte("rld")
	if !bytes.Equal(buf.Bytes(), expect) {
		t.Fatalf("bad: %s", buf.Bytes())
	}
}

func TestBuffer_ManySmall(t *testing.T) {
	inp := []byte("hello world")

	buf, err := NewBuffer(3)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	for _, b := range inp {
		n, err := buf.Write([]byte{b})
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if n != 1 {
			t.Fatalf("bad: %v", n)
		}
	}

	expect := []byte("rld")
	if !bytes.Equal(buf.Bytes(), expect) {
		t.Fatalf("bad: %v", buf.Bytes())
	}
}

func TestBuffer_MultiPart(t *testing.T) {
	inputs := [][]byte{
		[]byte("hello world\n"),
		[]byte("this is a test\n"),
		[]byte("my cool input\n"),
	}
	total := 0

	buf, err := NewBuffer(16)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	for _, b := range inputs {
		total += len(b)
		n, err := buf.Write(b)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if n != len(b) {
			t.Fatalf("bad: %v", n)
		}
	}

	if int64(total) != buf.TotalWritten() {
		t.Fatalf("bad total")
	}

	expect := []byte("t\nmy cool input\n")
	if !bytes.Equal(buf.Bytes(), expect) {
		t.Fatalf("bad: %v", buf.Bytes())
	}
}
