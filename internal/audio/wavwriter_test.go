package audio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewWAVCreatesOutDirectory(t *testing.T) {
	tmp := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer func() { _ = os.Chdir(oldWd) }()

	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir temp: %v", err)
	}

	w, err := NewWAV("test", 44100, 1)
	if err != nil {
		t.Fatalf("NewWAV returned error: %v", err)
	}
	defer func() { _ = w.Close() }()

	outPath := filepath.Join("out", "test.wav")
	if _, err := os.Stat(outPath); err != nil {
		t.Fatalf("expected wav file %q to exist: %v", outPath, err)
	}
}
