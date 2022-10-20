package termlog

import (
	"fmt"
	"testing"
)

// TestTermLogType type test
func TestTermLogType(t *testing.T) {
	logger := New()
	got := fmt.Sprintf("%T", logger)
	want := "*termlog.Logger"
	if got != want {
		t.Fatalf(`got = %q, want match for %q`, got, want)
	}
}

// TestTermLogStringDefault string test
func TestTermLogStringDefault(t *testing.T) {
	logger := New()
	got := fmt.Sprintf("%s", logger)
	want := `*termlog.Logger{level: 5, namespace: ""}`
	if got != want {
		t.Fatalf(`got = %q, want match for %q`, got, want)
	}
}

// TestTermLogStringNamespaced string test
func TestTermLogStringNamespaced(t *testing.T) {
	logger := New("myname")
	got := fmt.Sprintf("%s", logger)
	want := `*termlog.Logger{level: 5, namespace: "myname"}`
	if got != want {
		t.Fatalf(`got = %q, want match for %q`, got, want)
	}
}
