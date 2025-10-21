# File Organizer (Go)

A small *file organizer* written in Go. Use the built binary (organizer or organizer.exe) to  move files from build outputs, downloads, or any folder into organized directories by extension.

## Features
- Move or copy files into target folders
- Filter by extensions or patterns
- Recursive directory processing

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