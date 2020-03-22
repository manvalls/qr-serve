# QR Serve

This is a small windows utility which serves a file over LAN,
showing a QR code which can be scanned by other devices on the
same network to retrieve the file.

## How to use

Drag and drop the file you want to serve to QR Serve, or right
click the file and choose "Open with...", then find and pick
the QR Serve binary.

## Download

[Click here](https://github.com/manvalls/qr-serve/releases/latest/download/qr-serve.exe)

## Building QR Serve

- Install Go and project dependencies
- Get rsrc - `go install github.com/akavel/rsrc`
- Run `release.sh`