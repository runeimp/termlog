package termlog

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

const (
	testingDateFormat     = `2006-01-02`
	testingDateTimeFormat = `2006-01-02 15:04:05`
)

// TestTermLogConditionalLevelFalse type test
func TestTermLogConditionalLevelFalse(t *testing.T) {
	var (
		condition  = false
		output     bytes.Buffer
		want       = ""
		wantMsg    = "This is conditional: false"
		wantFormat = "%s \x1b[1;36mINFO \x1b[0m %s\n"
	) // "2023-02-06 \x1b[1;33mWARN \x1b[0m Unit Testing Output\n"

	want = fmt.Sprintf(wantFormat, time.Now().Format(testingDateFormat), wantMsg)

	logger := New()
	logger.Output = &output // Redirecting the output from os.Stderr for testing
	logger.TimeFormat = testingDateFormat

	logger.ConditionalLevel(condition, WarnLevel, InfoLevel, wantMsg)

	got := output.String()
	if got != want {
		t.Fatalf(`got = %q, want match for %q`, got, want)
	}
}

// TestTermLogConditionalLevelTrue type test
func TestTermLogConditionalLevelTrue(t *testing.T) {
	var (
		condition  = true
		output     bytes.Buffer
		want       = ""
		wantMsg    = "This is conditional: true"
		wantFormat = "%s \x1b[1;33mWARN \x1b[0m %s\n"
	) // "2023-02-06 \x1b[1;33mWARN \x1b[0m Unit Testing Output\n"

	want = fmt.Sprintf(wantFormat, time.Now().Format(testingDateFormat), wantMsg)

	logger := New()
	logger.Output = &output // Redirecting the output from os.Stderr for testing
	logger.TimeFormat = testingDateFormat

	logger.ConditionalLevel(condition, WarnLevel, InfoLevel, wantMsg)

	got := output.String()
	if got != want {
		t.Fatalf(`got = %q, want match for %q`, got, want)
	}
}

// TestTermLogOutput test
func TestTermLogOutput(t *testing.T) {
	var (
		output     bytes.Buffer
		want       = ""
		wantFormat = "%s \x1b[1;33mWARN \x1b[0m %s\n"
		wantMsg    = "Unit Testing Output"
	)

	want = fmt.Sprintf(wantFormat, time.Now().Format(testingDateFormat), wantMsg)

	logger := New()
	logger.Output = &output // Redirecting the output from os.Stderr for testing
	logger.TimeFormat = testingDateFormat
	logger.Warn(wantMsg)

	got := output.String()
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

// TestTermLogType type test
func TestTermLogType(t *testing.T) {
	logger := New()
	got := fmt.Sprintf("%T", logger)
	want := "*termlog.Logger"
	if got != want {
		t.Fatalf(`got = %q, want match for %q`, got, want)
	}
}
