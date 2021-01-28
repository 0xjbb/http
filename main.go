package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main(){
	dir,err  := os.Getwd()

	if err != nil{
		log.Fatal(err)
	}

	port := flag.String("p", "8080", "Listening port.")
	serDir := flag.String("d", dir, "Serve directory.")
	//upload := flag.String("u", dir + "/uploads", "Directory for uploads.")
	flag.Parse()

	fmt.Println("[+] Listening on port", *port, "...")
	fmt.Println("[+] Serving directory:",*serDir)
	fmt.Println("[+] Uploads directory:")//todo

	http.HandleFunc("/upload", uploadHandler)

	http.Handle("/", http.FileServer(http.Dir(*serDir)))
	host := fmt.Sprintf(":%s", *port)
	log.Fatal(http.ListenAndServe(host, logRequest(http.DefaultServeMux)))
}

func uploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		fmt.Fprintf(w, "Send it as post.")
		//todo add an upload form, for an easier time in RDP Sessions.
		return
	}

	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file")

	if err != nil {
		log.Println(err)
		return
	}


	fmt.Println(handler.Filename)
	fmt.Println(file)
	//curl -F file=@test.txt http://localhost:8080/upload
	fmt.Println("File Uploaded! ")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// maybe ignore /uploads?
		ipAddr := strings.Split(r.RemoteAddr, ":")[0]//remote port.
		fmt.Printf("%s %s %s\n", ipAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}