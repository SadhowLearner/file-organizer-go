# File Organizer (Go)

A small, configurable file organizer written in Go. Use the built binary (organizer or organizer.exe) to sort, move, or copy files from build outputs, downloads, or any folder into organized directories by extension, date, or custom rules.

## Features
- Move or copy files into target folders
- Filter by extensions or patterns
- Recursive directory processing
- Dry-run and verbose modes for safe testing
- Config file support (YAML/JSON) for repeatable rules

## Build (Windows)
From the project root:

- Native build on Windows:
    ```
    go build -o organizer.exe main.go
    ```

- Cross-compile on Linux/macOS for Windows (64-bit):
    ```
    GOOS=windows GOARCH=amd64 go build -o organizer.exe main.go
    ```
- In Linux:
    ```
    go build -o organizer main.go
    ```
Place `organizer.exe` anywhere on your PATH or run it from the project folder.

## Basic usage

Just move output file `(organizer/orgaizer.exe)` to folder that you want to organize the folder and double click the output file example in `Download` folder.

You can modif this project as you want.