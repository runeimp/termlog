#
# TermLog Project
#

MAIN_CODE := 'termlog.go'

set windows-shell := ["powershell", "-c"] # To use PowerShell Desktop instead of Core on Windows
set shell := ["pwsh", "-c"] # PowerShell Core (Multi-Platform)

alias ver := version

shebang := if os() == 'windows' { 'powershell' } else { '/bin/sh' }
powershell := if os() == 'windows' { 'powershell' } else { '/usr/bin/env pwsh' } # Shebang for PS Desktop on Windows and PS Core everywhere else
termwiper := if os() == 'windows' { '_termwipe-ps' } else { '_termwipe-sh' }
testcoverage := if os() == 'windows' { 'just _test-coverage-ps' } else { 'just _test-coverage-sh' }


@_default: _termwipe
	just --list

# Build tester for Linux, macOS, and Windows
build: _termwipe
	#!{{shebang}}
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/tester cmd/tester/tester.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/tester cmd/tester/tester.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows/tester.exe cmd/tester/tester.go

# Run test code, coverage, or units (unit tests)
@test cmd="code": _termwipe
	just test-{{cmd}}

# Run test code
@test-code:
	go run cmd/tester/tester.go; echo "Exit Code: $LastExitCode"

# Run Go test coverage
@test-coverage:
	# echo "You need to run:"
	# echo "go test -coverprofile=c.out"
	# echo "go tool cover -func=c.out"
	{{testcoverage}}

_test-coverage-sh:
	#!/bin/sh
	echo "$ go test -coverprofile=c.out"
	GO111MODULE=auto go test -coverprofile=c.out
	echo "$ go tool cover -func=c.out"
	GO111MODULE=auto go tool cover -func=c.out
	rm c.out

_test-coverage-ps:
	#!{{powershell}}
	# NOTE: not tested so may have issues
	go test -coverprofile=c.out
	go tool cover -func=c.out
	Remove-Item c.out

# Run Go unit tests
test-units:
	go test


# Display version of app
@version:
	# cat {{MAIN_CODE}} | grep -E '^	Version' | head -1 | cut -d'"' -f2 # Probably don't need head -1
	((Get-Content {{MAIN_CODE}} | Select-String '^	Version') -Split '"')[1]


@_termwipe:
	just {{termwiper}}

@_termwipe-ps:
	Clear-Host

_termwipe-sh:
	#!/bin/sh
	# set -exo pipefail
	if [ ${#VISUAL_STUDIO_CODE} -gt 0 ]; then
		clear
	elif [ ${KITTY_WINDOW_ID} -gt 0 ] || [ ${#TMUX} -gt 0 ] || [ "${TERM_PROGRAM}" = 'vscode' ]; then
		printf '\033c'
	elif [ "${TERM_PROGRAM}" = 'Apple_Terminal' ] || [ "${TERM_PROGRAM}" = 'iTerm.app' ]; then
		osascript -e 'tell application "System Events" to keystroke "k" using command down'
	elif [ -x "$(which tput)" ]; then
		tput reset
	elif [ -x "$(which tcap)" ]; then
		tcap rs
	elif [ -x "$(which reset)" ]; then
		reset
	else
		clear
	fi

