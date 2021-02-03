# HTTP Server

Simple HTTP File server with upload support for pen-testing CTFs

## Features

- Simple
- Easy file uploads
- Custom uploads directory  
- HTTP & HTTPS
- Auto generation of SSL certs

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
sudo cp http /usr/bin/http
```

if you wish to use ports lower than 1000 you can add the bind capability 

```bash
sudo setcap 'cap_net_bind_service=+ep' /usr/bin/http
```

### Usage

Navigate to the directory you wish you server via HTTP and execute the binary

if you wish to set a custom port you can do so with -p

`http -p 80`

if you wish to use https you can use -tls

`http -tls`

## Upload

You can upload a file with curl

#### HTTP

`curl -F file=@test.txt http://10.10.10.10:8080/upload`

#### HTTPS

`curl -F file=@test.txt https://10.10.10.10:8080/upload`


This will upload `test.txt` to the directory that is being served.

### TODO

- Custom uploads directory
- Custom index
