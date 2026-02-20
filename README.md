# Shell-Go

A simple shell implementation in Go (not golang btw).

## Overview

Shell-Go is an interactive command-line shell written in Go that provides a REPL (Read-Eval-Print Loop) interface with various builtin commands for file system operations, I/O, system management, and network utilities.

## Features

- **Interactive REPL**: Command-line interface with custom prompt showing current directory
- **Signal Handling**: Graceful handling of SIGINT and SIGTERM signals
- **Custom Prompt**: Path formatting that displays home directory as `~`
- **Environment Variable Management**: Built-in support for setting and viewing environment variables
- **Modular Command Architecture**: Commands organized by category (filesystem, I/O, network, system)
- **Flag Support**: Many commands support flags using the cobra library

## Requirements

- Go 1.23.2 or later
- Optional: `just` for build automation (or use `go` commands directly)

## Installation & Building

### Using Just

```bash
# Build the binary
just build

# Run directly
just run
```

The build process creates a binary named `shell-go` in the current directory.

## Usage

Start the shell by running the binary:

```bash
./shell-go
```

Type commands and press Enter to execute them. To exit the shell, use the `exit` command or press Ctrl+C (handled gracefully via signal handling).

## Available Commands

### File System Commands

- **`cat [file]`** - Read and display file contents
  - `-n`, `--number` - Number the output lines, starting at 1
  - `-e` - Display non-printing characters and a dollar sign (`$`) at the end of each line

- **`cd [directory]`** - Change directory
  - Supports `~` as shorthand for home directory
  - Supports both absolute and relative paths

- **`ls`** - List files and directories
  - `-a`, `--all` - View all files and directories (including hidden files)
  - `-d`, `--dirs` - View only directories
  - `-f`, `--files` - View only files

- **`mkdir [directory]`** - Create a new directory

- **`pwd`** - Print the current working directory

- **`rm [file]`** - Remove a file
  - `-R`, `-r` - Attempt to remove the file hierarchy rooted in each file argument (recursive)

- **`rmdir [directory]`** - Remove an empty directory

- **`touch [file]`** - Create an empty file or update file timestamp

- **`ln [target] [link]`** or **`link [target] [link]`** - Create a symbolic link

### I/O Commands

- **`clear`** - Clear the terminal screen

- **`echo [text]`** - Print text to standard output

### System Commands

- **`exit`** - Exit the shell gracefully

- **`export [key]=[value]`** - Set an environment variable
  - Example: `export PATH=/usr/bin`

- **`env`** - Print all environment variables

- **`hostname`** - Print the name of the current host system

- **`whoami`** - Print the current username

### Network Commands

- **`nslookup [domain]`** - Perform DNS lookup for a domain name
  - Displays DNS server information and resolved IP addresses

## Project Structure

```
shell-go/
├── cmd/
│   └── main.go          # Main entry point with REPL loop
├── internal/
│   └── shell/
│       ├── exit/         # Exit functionality
│       ├── prompt/       # Prompt formatting (home directory as ~)
│       └── signals/      # Signal handling (SIGINT, SIGTERM)
├── pkg/
│   └── builtins/
│       ├── fs/           # File system commands
│       │   ├── cat.go
│       │   ├── cd.go
│       │   ├── link.go
│       │   ├── ls.go
│       │   ├── mkdir.go
│       │   ├── pwd.go
│       │   ├── rm.go
│       │   ├── rmdir.go
│       │   └── touch.go
│       ├── io/           # I/O commands
│       │   ├── clear.go
│       │   └── echo.go
│       ├── network/      # Network commands
│       │   └── nslookup.go
│       └── system/       # System commands
│           ├── env.go
│           ├── exit.go
│           ├── export.go
│           ├── hostname.go
│           └── whoami.go
├── go.mod
├── justfile
└── README.md
```

## Development Notes

- The shell uses environment variables (specifically `PWD`) for state management
- Signal handling runs in a separate goroutine to allow graceful shutdown
- Many commands use the [cobra](https://github.com/spf13/cobra) library for flag parsing and command structure
- The prompt formatter automatically converts absolute paths starting with `$HOME` to use `~` notation
- Unknown commands display "Unknown command" and continue the REPL loop
