package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	serverDirectory *string
	uploadDirectory *string
	port *string
	isTLS *bool
)

func init(){
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	port = flag.String("p", "8080", "Listening port.")
	serverDirectory = flag.String("d", dir, "Serve directory.")
	uploadDirectory = flag.String("u", dir, "Custom uploads directory ( Default is CWD )")
	isTLS = flag.Bool("tls", false, "Enable HTTPS.")
}

func main() {
	flag.Parse()
	banner()

	r := mux.NewRouter()

	r.HandleFunc("/shell", shellHandler)
	r.HandleFunc("/upload", uploadHandler)
	//http.HandleFunc("/shell", shellHandler)
	r.Handle("/", http.FileServer(http.Dir(*serverDirectory)))


	//CustomFileServer(*serverDirectory)
	http.Handle("/", r)
	host := fmt.Sprintf(":%s", *port)

	// TODO clean this bit up
	if *isTLS {
		cert, key, err := GenerateCert()
		if err != nil {
			log.Fatal(err)
		}

		fcert, err := tls.X509KeyPair(cert, key)

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{fcert},
		}

		server := http.Server{
			Addr:      host,
			Handler:   logRequest(http.DefaultServeMux),
			TLSConfig: tlsConfig,
		}
		log.Fatal(server.ListenAndServeTLS("", ""))
	} else {
		log.Fatal(http.ListenAndServe(host, logRequest(http.DefaultServeMux)))
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddr := strings.Split(r.RemoteAddr, ":")[0] //remove remote port.
		fmt.Printf("%s\t%s\t%s\n", ipAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func banner(){

	fmt.Println("[+] Listening on port", *port, "...")
	fmt.Println("[+] Serving directory:", *serverDirectory)
	fmt.Println("[+] Uploads directory:", *uploadDirectory)
}