package cgo

import (
	"testing"
)

func TestApp(t *testing.T) {
	cook, err := NewCook("Joe", false)
	if err != nil {
		t.Fatalf("unexpected error from New: %v", err)
	}
	defer cook.Close()

	got := cook.Greet("Hello, ")
	want := "Hello, Joe"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAppError(t *testing.T) {
	_, err := NewCook("Joe", true)
	if err == nil {
		t.Fatalf("want error, got nil")
	}
}
