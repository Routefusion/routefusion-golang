package client

import (
	"bytes"
	"io"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestOffsetReaderRead(t *testing.T) {
	buf := []byte("testData")
	reader := &offsetReader{buf: bytes.NewReader(buf)}

	tempBuf := make([]byte, len(buf))

	n, err := reader.Read(tempBuf)

	if err != nil {
		t.Fatal(err)
	}

	if n != len(buf) {
		t.Errorf("lengths don't match. expected: %d, actual: %d",
			n, len(buf))
	}

	if !reflect.DeepEqual(buf, tempBuf) {
		t.Errorf("bufs not equal. expected: %v, actual: %v", buf, tempBuf)
	}
}

func TestOffsetReaderSeek(t *testing.T) {
	buf := []byte("testData")
	reader := newOffsetReader(bytes.NewReader(buf), 0)

	n, err := reader.Seek(0, 2)
	if err != nil {
		t.Fatal(err)
	}
	if int64(len(buf)) != n {
		t.Error("seek does not do so on offset")
	}
}

func TestOffsetReaderClose(t *testing.T) {
	buf := []byte("testData")
	reader := &offsetReader{buf: bytes.NewReader(buf)}

	err := reader.Close()
	if err != nil {
		t.Fatal(err)
	}

	tempBuf := make([]byte, len(buf))
	n, err := reader.Read(tempBuf)
	if n != 0 {
		t.Error("bytes read should be 0 for EOF")
	}
	if !reflect.DeepEqual(err, io.EOF) {
		t.Error("EOF error is not seen")
	}
}

func TestOffsetReaderRace(t *testing.T) {
	wg := sync.WaitGroup{}

	f := func(reader *offsetReader) {
		defer wg.Done()
		var err error
		buf := make([]byte, 1)
		_, err = reader.Read(buf)
		for err != io.EOF {
			_, err = reader.Read(buf)
		}

	}

	closeFn := func(reader *offsetReader) {
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(20)+1) * time.Millisecond)
		reader.Close()
	}
	for i := 0; i < 50; i++ {
		reader := &offsetReader{buf: bytes.NewReader(make([]byte, 1024*1024))}
		wg.Add(1)
		go f(reader)
		wg.Add(1)
		go closeFn(reader)
	}
	wg.Wait()
}

func BenchmarkOffsetReader(b *testing.B) {
	bufSize := 1024 * 1024 * 100
	buf := make([]byte, bufSize)
	reader := &offsetReader{buf: bytes.NewReader(buf)}

	tempBuf := make([]byte, 1024)

	for i := 0; i < b.N; i++ {
		reader.Read(tempBuf)
	}
}

func BenchmarkBytesReader(b *testing.B) {
	bufSize := 1024 * 1024 * 100
	buf := make([]byte, bufSize)
	reader := bytes.NewReader(buf)

	tempBuf := make([]byte, 1024)

	for i := 0; i < b.N; i++ {
		reader.Read(tempBuf)
	}
}
