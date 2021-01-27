package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	dir,err  := os.Getwd()

	if err != nil{
		log.Fatal(err)
	}

	port := flag.String("p", "8080", "Listening port.")
	serDir := flag.String("d", dir, "Serve directory")
	flag.Parse()

	fmt.Println("[+] Listening on port", *port, "...")
	fmt.Println("[+] Serving directory:",*serDir)

	http.Handle("/", http.FileServer(http.Dir(*serDir)))
	host := fmt.Sprintf(":%s", *port)
	log.Fatal(http.ListenAndServe(host, nil))
}
