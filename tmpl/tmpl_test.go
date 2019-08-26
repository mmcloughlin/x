package tmpl

import (
	"bytes"
	"path/filepath"
	"testing"
)

func Corpus(t *testing.T) []string {
	t.Helper()
	filenames, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	return filenames
}

func TestRoundTrip(t *testing.T) {
	for _, filename := range Corpus(t) {
		src, err := FileSystemLoader.Load(filename)
		if err != nil {
			t.Fatal(err)
		}

		tpl, err := ParseFile(filename, src)
		if err != nil {
			t.Fatal(err)
		}

		got, err := tpl.Bytes()
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(got, src) {
			t.Logf("src:\n%s", src)
			t.Logf("got:\n%s", got)
			t.Fatalf("%s: roundtrip error", filename)
		}
	}
}
