package main

import (
	"fmt"

	"github.com/runeimp/termlog"
)

func main() {
	fmt.Printf("==> Testing Logger Up to Error Level With Info Level Set and ForceColor is ForceColorOff\n\n")
	testInfoLogger()
	fmt.Printf("\n==> Testing Namespaced Logger with Custom TimeFormat Up to Error Level With Default Settings\n\n")
	testNamespacedLogger()
	fmt.Printf("\n==> Testing Logger with the ConditionalLevel method\n\n")
	testConditionalLevelLogger()
	fmt.Printf("\n==> Testing Logger Up to Panic Level With Custom Exit Codes and ForceColor is ForceColorOn\n\n")
	testCustomExitCodesLogger()
	fmt.Println()
}

func testInfoLogger() {
	log := termlog.New()
	log.ForceColor = termlog.ForceColorOff
	log.Level = termlog.InfoLevel
	log.Debug("Debugging message")
	log.Info("Informational message")
	log.Warn("Warning message")
	log.Error("Error message")
}

func testCustomExitCodesLogger() {
	log := termlog.New()
	log.ForceColor = termlog.ForceColorOn
	log.FatalExitCode = 13
	log.PanicExitCode = 42
	log.Debug("Debugging message")
	log.Info("Informational message")
	log.Warn("Warning message")
	log.Error("Error message")
	log.Fatal("Fatal message") // The default exit code 1
	log.Panic("Panic message") // The default exit code 2 preceded by a stack trace
}

func testNamespacedLogger() {
	log := termlog.New("my-namespace")
	log.TimeFormat = "2006-01-02 15:04:05"
	log.Debug("Debugging message")
	log.Info("Informational message")
	log.Warn("Warning message")
	log.Error("Error message")
	// log.Fatal("Fatal message") // Message then default exit code 1
	// log.Panic("Panic message") // Message then default exit code 2 and stack trace
}

func testConditionalLevelLogger() {
	log := termlog.New()
	log.Level = termlog.WarnLevel
	condition := 0 < 1
	log.ConditionalLevel(condition, termlog.InfoLevel, termlog.WarnLevel, "This is condition: %-5t # Should be Info Log Level in this test", condition)
	condition = false
	log.ConditionalLevel(condition, termlog.InfoLevel, termlog.WarnLevel, "This is condition: %5t # Should be Warn Log Level in this test", condition)
}
