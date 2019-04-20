package handlers

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"
)

func TestRerun(t *testing.T) {
	saved := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = w

	go func() {
		NewRerun("sleep", []string{"1"}, 2).Run()
	}()

	select {
	case <-time.After(5 * time.Second):
	}

	_ = w.Close()
	os.Stdout = saved

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	expected := "sleep [1]\nsleep [1]\n"
	if buf.String() != expected {
		t.Errorf("actual=%v expected=%v", buf.String(), expected)
	}
}
