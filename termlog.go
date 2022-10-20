package termlog

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

const (
	defaultFormat     = `2006-01-02 15:04:05.000000` // Pythonesque Timestamp instead of 2006-01-02 15:04:05 Z0700
	SystemUnsupported = "Unsupported"
	Version           = "0.1.0"
)

const (
	DebugLevel loggingLevel = 5 // Debugging messages
	ErrorLevel loggingLevel = 2 // Error messages (default)
	FatalLevel loggingLevel = 1 // Fatal message with exit code 1
	InfoLevel  loggingLevel = 4 // Informational messages
	PanicLevel loggingLevel = 0 // Panic message with exit code 1 and stack trace
	WarnLevel  loggingLevel = 3 // Warning messages
)

const (
	FormatDebugColor       = "\033[0;37m%s\033[0m" // 0;37 White-Normal (Silver)
	FormatErrorColor       = "\033[1;31m%s\033[0m" // 1;31 Red-Bold
	FormatFatalColor       = "\033[1;31m%s\033[0m" // 1;31 Red-Bold
	FormatInformationColor = "\033[1;36m%s\033[0m" // 1;36 Cyan-Bold
	FormatPanicColor       = "\033[1;31m%s\033[0m" // 1;31 Red-Bold
	FormatWarningColor     = "\033[1;33m%s\033[0m" // 1;33 Yellow-Bold
)

const (
	ForceColorUndefined colorForce = "undefined"
	ForceColorOff       colorForce = "off"
	ForceColorOn        colorForce = "on"
)

const (
	ANSIPrefix             = "\033["   // Not yet utilized
	ANSIColorSuffix        = "m"       // Not yet utilized
	ANSIReset              = "\033[0m" // Not yet utilized
	DebugLabel             = "DEBUG"
	DebugLabelANSI         = "\033[0;37mDEBUG\033[0m"    // 0;37 White-Normal (Silver)
	DebugLabelBgANSI       = "\033[0;30;47mDEBUG\033[0m" // 0;30;47 FG: Black, BG: White-Normal (Silver)
	InformationLabel       = "INFO "
	InformationLabelANSI   = "\033[1;36mINFO \033[0m"    // 1;36 Cyan-Bold
	InformationLabelBgANSI = "\033[1;30;46mINFO \033[0m" // 1;30;46 FG: Black, BG: Cyan-Bold
	WarningLabel           = "WARN "
	WarningLabelANSI       = "\033[1;33mWARN \033[0m"     // 1;33 Yellow-Bold
	WarningLabelBgANSI     = "\033[1;30;103mWARN \033[0m" // 1;30;103 FG: Black, BG: Yellow-Bold
	ErrorLabel             = "ERROR"
	ErrorLabelANSI         = "\033[1;31mERROR\033[0m"     // 1;31 Red-Bold
	ErrorLabelBgANSI       = "\033[1;30;101mERROR\033[0m" // 1;30;101 FG: Black, BG: Red-Bold
	FatalLabel             = "FATAL"
	FatalLabelANSI         = "\033[1;31mFATAL\033[0m"     // 1;31 Red-Bold
	FatalLabelBgANSI       = "\033[1;30;101mFATAL\033[0m" // 1;30;101 FG: Black, BG: Red-Bold
	PanicLabel             = "PANIC"
	PanicLabelANSI         = "\033[1;31mPANIC\033[0m"     // 1;31 Red-Bold
	PanicLabelBgANSI       = "\033[1;30;101mPANIC\033[0m" // 1;30;101 FG: Black, BG: Red-Bold
)

// colorForce used to define a trinary value type
type colorForce string

// loggingLevel is used to define logging levels
type loggingLevel uint8

// Logger is the base structure for logging
type Logger struct {
	ForceColor    colorForce
	Level         loggingLevel
	Namespace     string
	system        string
	FatalExitCode int
	PanicExitCode int
	TimeFormat    string
}

// Debug outputs a debugging level message
func (l *Logger) Debug(format string, msg ...any) {
	if l.Level > InfoLevel {
		l.printLog(DebugLevel, format, msg...)
	}
}

// Error outputs a error level message
func (l *Logger) Error(format string, msg ...any) {
	if l.Level > FatalLevel {
		l.printLog(ErrorLevel, format, msg...)
	}
}

// Panic outputs a panic level message with exit code
func (l *Logger) Fatal(msg ...any) {
	var format string

	if len(msg) > 0 {
		format = msg[0].(string)
		if len(msg) > 1 {
			msg = msg[1:]
		} else {
			msg = []any{}
		}
	}

	l.printLog(FatalLevel, format, msg...)
	os.Exit(l.FatalExitCode)
}

func (l *Logger) getLabel(level loggingLevel) (result string) {
	/*
		ANSI Color Off
			- If the system is supported and Logger.ForceColor is ForceColorOff
			- If the system is unsupported and Logger.ForceColor is not ForceColorOn
		ANSI Color On
			- If the system is supported and Logger.ForceColor is not ForceColorOff
			- If the system is unsupported but Logger.ForceColor is ForceColorOn
	*/

	enableANSI := false
	if l.ForceColor != ForceColorUndefined {
		if l.ForceColor == ForceColorOn {
			enableANSI = true
		}
	} else if l.system != SystemUnsupported {
		enableANSI = true
	}

	switch level {
	case DebugLevel:
		result = DebugLabel
		if enableANSI {
			result = DebugLabelANSI
		}
	case InfoLevel:
		result = InformationLabel
		if enableANSI {
			result = InformationLabelANSI
			// result = fmt.Sprintf(FormatInformationColor, InformationLabel)
		}
	case WarnLevel:
		result = WarningLabel
		if enableANSI {
			result = WarningLabelANSI
			// result = fmt.Sprintf(FormatWarningColor, WarningLabel)
		}
	case ErrorLevel:
		result = ErrorLabel
		if enableANSI {
			result = ErrorLabelANSI
			// result = fmt.Sprintf(FormatErrorColor, ErrorLabel)
		}
	case FatalLevel:
		result = FatalLabel
		if enableANSI {
			result = FatalLabelANSI
			// result = fmt.Sprintf(FormatFatalColor, FatalLabel)
		}
	case PanicLevel:
		result = PanicLabel
		if enableANSI {
			result = PanicLabelANSI
		}
	}

	if len(l.Namespace) > 0 {
		result = fmt.Sprintf("%s %s", l.Namespace, result)
	}

	return result
}

// Info outputs a information level message
func (l *Logger) Info(format string, msg ...any) {
	if l.Level > WarnLevel {
		// printLog(InfoLevel, l.system, l.Namespace, format, msg...)
		l.printLog(InfoLevel, format, msg...)
	}
}

// Panic outputs a panic level message with exit code and stack trace
func (l *Logger) Panic(msg ...any) {
	var format string

	if len(msg) > 0 {
		format = msg[0].(string)
		if len(msg) > 1 {
			msg = msg[1:]
		} else {
			msg = []any{}
		}
	}

	// log.Panic(format)
	l.printLog(PanicLevel, format, msg...)
	fmt.Fprintln(os.Stderr)
	debug.PrintStack()
	// fmt.Fprintln(os.Stderr)
	os.Exit(l.PanicExitCode)
}

func (l *Logger) printLog(level loggingLevel, format string, msg ...any) {
	fmt.Fprintf(os.Stderr, "%s %s %s\n", time.Now().Format(l.TimeFormat), l.getLabel(level), fmt.Sprintf(format, msg...))
}

func (l *Logger) String() string {
	return fmt.Sprintf("%T{level: %d, namespace: %q}", l, l.Level, l.Namespace)
}

// Warn outputs a warning level message
func (l *Logger) Warn(format string, msg ...any) {
	if l.Level > ErrorLevel {
		l.printLog(WarnLevel, format, msg...)
	}
}

// New returns a Logger
func New(args ...string) *Logger {
	var namespace string
	if len(args) > 0 {
		namespace = args[0]
	}

	isMSTerminalSession := false
	isMSYS := false
	system := runtime.GOOS

	if os.Getenv("WT_SESSION") != "" {
		isMSTerminalSession = true
	}
	if os.Getenv("MSYSTEM") != "" {
		isMSYS = true
	}
	// fmt.Printf("termlog.New() | Windows System: %q\n", system)
	/* NOTE: Test for unsupported systems goes here
	It works fine in MS Terminal, Terminal/PowerShell 5 Desktop, Terminal/Command Prompt,
	and PowerShell 7.2.5 Core on macOS Big Sur.
	*/
	if system == "windows" {
		// Test for Windows Terminal as it has ANSI support turned on by default
		if isMSTerminalSession != true && isMSYS != true { // If Windows Terminal is not hosting the session
			system = SystemUnsupported
		}
		// fmt.Printf("termlog.New() | Windows System: %q\n", system)
	}

	// fmt.Printf("termlog.New() | isMSTerminalSession: %t\n", isMSTerminalSession)
	// fmt.Printf("termlog.New() | shellOut: %s\n", shellOut)
	// if shellErr != nil {
	// 	fmt.Printf("termlog.New() | shellErr: %s\n", shellErr.Error())
	// }

	return &Logger{
		ForceColor:    ForceColorUndefined,
		Level:         DebugLevel,
		Namespace:     namespace,
		system:        system,
		FatalExitCode: 1,
		PanicExitCode: 2,
		TimeFormat:    defaultFormat,
	}
}
