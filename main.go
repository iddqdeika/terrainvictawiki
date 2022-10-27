package main

import (
	"fmt"
	"io"
	"net/http"
	"terrainvictawiki/pkg/api/rest"
	"time"
)

func main() {

	fmt.Println(time.Unix(1666724060, 0))

	fmt.Printf("starting app...\r\n")

	rest.Init()
	http.Handle("/", http.FileServer(http.Dir("./res/web")))
	fmt.Printf("starting server...\r\n")
	http.ListenAndServe(":8080", nil)
}

func fromWriter(writer io.ReadWriter) string {
	data, err := io.ReadAll(writer)
	if err != nil {
		return ""
	}
	return string(data)
}
