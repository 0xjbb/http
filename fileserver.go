package main

import (
	"net/http"
)

type fileServer struct{

}

func CustomFileServer() *fileServer{
	return &fileServer{

	}
}

func (fs *fileServer) Serve(root http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}