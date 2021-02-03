package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var(
	serverDirectory *string
	uploadDirectory *string

)

func main(){
	dir,err  := os.Getwd()

	if err != nil{
		log.Fatal(err)
	}

	port := flag.String("p", "8080", "Listening port.")
	serverDirectory = flag.String("d", dir, "Serve directory.")
	uploadDirectory = flag.String("u", dir, "Custom uploads directory ( Default is CWD )")
	isTLS := flag.Bool("tls", false, "Enable HTTPS.")

	flag.Parse()

	fmt.Println("[+] Listening on port", *port, "...")
	fmt.Println("[+] Serving directory:", *serverDirectory)
	fmt.Println("[+] Uploads directory:", *uploadDirectory)

	http.HandleFunc("/upload", uploadHandler)

	http.Handle("/", http.FileServer(http.Dir(*serverDirectory)))
	host := fmt.Sprintf(":%s", *port)

	// TODO clean this bit up
	if *isTLS {
		cert, key, err := GenerateCert()
		if err != nil{
			log.Fatal(err)
		}

		fcert, err := tls.X509KeyPair(cert, key)

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{fcert},
		}

		server := http.Server{
			Addr: host,
			Handler: logRequest(http.DefaultServeMux),
			TLSConfig: tlsConfig,
		}
		log.Fatal(server.ListenAndServeTLS("",""))
	}else{
		log.Fatal(http.ListenAndServe(host, logRequest(http.DefaultServeMux)))
	}
}

//curl -F file=@test.txt http://localhost:8080/upload
func uploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{

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

	if err != nil{
		log.Println(err)//?fatal???
		return
	}

	io.Copy(fh, file)
	ipAddr := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Printf("%s %s %s %s", ipAddr, "UPLOAD", handler.Filename, fileName)
	w.Write([]byte("File uploaded."))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// maybe ignore /uploads?
		ipAddr := strings.Split(r.RemoteAddr, ":")[0]//remove remote port.
		fmt.Printf("%s %s %s\n", ipAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}