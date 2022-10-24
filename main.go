package main

import (
	"fmt"
	"io"
	"net/http"
	"terrainvictawiki/pkg/localization"
	"terrainvictawiki/pkg/tech"
)

func main() {
	fmt.Printf("starting app...\r\n")

	tryInitTech()
	http.Handle("/", http.FileServer(http.Dir("./res/web")))
	fmt.Printf("starting server...\r\n")
	http.ListenAndServe(":8080", nil)
}

func tryInitTech() {
	r, err := newTechRepo()
	if err != nil {
		fmt.Printf("cant init tech repo: %v\r\n", err)
		return
	}

	th := tech.NewHTTPHandler(r)
	http.HandleFunc("/tech/all", th.HandleGetAll)
}

func newTechRepo() (*tech.Repository, error) {
	r, err := tech.NewRepositoryFromFile("./res/data/Templates/TITechTemplate.json")
	if err != nil {
		return nil, fmt.Errorf("cannot init tech handler: %v\r\n", err)
	}
	tryLocalizeRepo("./res/data/Localization/en/TITechTemplate.en", r)

	return r, nil
}

func tryLocalizeRepo(filename string, r *tech.Repository) {
	l, err := localization.NewStorage(filename)
	if err != nil{
		fmt.Printf("cant load localization file %v: %v\r\n", filename, err)
		return
	}
	f := localization.NewStorageFacade(l, "TITechTemplate")
	r.Localize(f)
}

func fromWriter(writer io.ReadWriter) string {
	data, err := io.ReadAll(writer)
	if err != nil {
		return ""
	}
	return string(data)
}
