# HTTP Server

Simple HTTP File server with upload support for pen-testing CTFs

## Getting Started

### Installation

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

Build the project.

```bash
go build
```

(Optional, but recommended)

Copy the binary to a directory in your PATH

```bash
sudo cp httpserverr /usr/bin/http
```

### Usage

Navigate to the directory you wish you server via HTTP and execute the binary

## Upload

You can upload a file with curl

`curl -F file=@test.txt http://10.10.10.10:8080/upload`


### TODO

- Add SSL support
- Custom index
- Add upload form for easy uploads when using RDP