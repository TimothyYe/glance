package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateReaderLoadsTxtFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sample.txt")
	if err := os.WriteFile(path, []byte("hello\nworld\n"), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	r, err := createReader(path)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if r == nil {
		t.Fatal("expected a reader, got nil")
	}
	if got := r.Current(); got != "hello" {
		t.Fatalf("expected loaded content to start with %q, got %q", "hello", got)
	}
}

func TestCreateReaderRejectsUnsupportedFormat(t *testing.T) {
	r, err := createReader("book.pdf")
	if err == nil {
		t.Fatal("expected an error for unsupported format, got nil")
	}
	if r != nil {
		t.Fatalf("expected nil reader on error, got %v", r)
	}
}

func TestResolveVersionFallsBackToDev(t *testing.T) {
	if got := resolveVersion(""); got != "dev" {
		t.Fatalf("expected %q for an empty injected version, got %q", "dev", got)
	}
}

func TestResolveVersionUsesInjectedValue(t *testing.T) {
	if got := resolveVersion("1.2.3"); got != "1.2.3" {
		t.Fatalf("expected the injected version %q, got %q", "1.2.3", got)
	}
}

func TestCreateReaderPropagatesLoadError(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "does-not-exist.txt")

	r, err := createReader(missing)
	if err == nil {
		t.Fatal("expected an error for a missing file, got nil")
	}
	if r != nil {
		t.Fatalf("expected nil reader on error, got %v", r)
	}
}
