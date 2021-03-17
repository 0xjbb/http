package main

import (
	"log"
	"net/http"
	"os"
)


func CustomFileServer(w http.ResponseWriter, r *http.Request, dir string){
	files, err := os.ReadDir(dir)

	if err != nil{
		log.Fatal(err)
	}

	for _,v := range files{
		w.Write([]byte(v))
	}


}