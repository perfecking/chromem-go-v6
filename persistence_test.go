package chromem

import (
	"bytes"
	"encoding/gob"
	"os"
	"slices"
	"testing"
)

func TestPersistence(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "chromem-go")
	if err != nil {
		t.Error("expected nil, got", err)
	}
	t.Cleanup(func() {
		_ = os.RemoveAll(tempDir)
	})

	type s struct {
		Foo string
		Bar []float32
	}
	obj := s{
		Foo: "test",
		Bar: []float32{-0.1, 0.1, 0.2},
	}

	persist(tempDir, obj)

	// Check if the file exists.
	_, err = os.Stat(tempDir + ".gob")
	if err != nil {
		t.Error("expected nil, got", err)
	}
	// Check if the file contains the expected data.
	b, err := os.ReadFile(tempDir + ".gob")
	if err != nil {
		t.Error("expected nil, got", err)
	}
	d := gob.NewDecoder(bytes.NewReader(b))
	res := s{}
	err = d.Decode(&res)
	if err != nil {
		t.Error("expected nil, got", err)
	}
	if res.Foo != obj.Foo {
		t.Error("expected", obj.Foo, "got", res.Foo)
	}
	if slices.Compare[[]float32](res.Bar, obj.Bar) != 0 {
		t.Error("expected", obj.Bar, "got", res.Bar)
	}
}