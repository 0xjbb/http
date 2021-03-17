package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

//curl -F file=@test.txt http://localhost:8080/upload
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {

		w.Write([]byte(`
			<html>
			<head>
				<title>Upload</title>
			</head>
			<body>
			<form enctype="multipart/form-data" action="/upload" method="POST">
				<input type="file" name="file" />
				<input type="submit" value="upload" />
			</form>
			</body>
			</html>
		`))
		return
	}

	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file")

	if err != nil {
		log.Println(err)
		return
	}

	fileName := path.Join(*uploadDirectory, path.Base(handler.Filename))
	fh, err := os.Create(fileName)
	defer fh.Close()

	if err != nil {
		log.Println(err) //?fatal???
		return
	}

	io.Copy(fh, file)
	ipAddr := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Printf("%s %s %s %s\n", ipAddr, "UPLOAD", handler.Filename, fileName)
	w.Write([]byte("File uploaded."))
}
