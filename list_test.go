package main

import (
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	out, err := run("ls")
	if err != nil {
		t.Fatalf("output: %s, error: %v", string(out), err)
	}
	expected := `REPO                TAGS
busybox             glibc, musl
alpine              latest
`
	if !strings.HasSuffix(out, expected) {
		t.Fatalf("expected: %s\ngot: %s", expected, out)
	}
}