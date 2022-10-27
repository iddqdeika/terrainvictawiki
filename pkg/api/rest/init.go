package rest

import (
	"fmt"
	"net/http"
	"terrainvictawiki/pkg/project"
	"terrainvictawiki/pkg/tech"
)

const (
	rootRestPath = "/rest"
)

func Init() {
	TryInitTech(rootRestPath)
	TryInitProject(rootRestPath)
}

func TryInitTech(root string) {
	r, err := tech.NewRepo()
	if err != nil {
		fmt.Printf("cant init tech repo: %v\r\n", err)
		return
	}
	th := tech.NewHTTPHandler(r)
	http.HandleFunc(root+"/tech/all", th.HandleGetAll)
}

func TryInitProject(root string) {
	r, err := project.NewRepo()
	if err != nil {
		fmt.Printf("cant init project repo: %v\r\n", err)
		return
	}
	th := project.NewHTTPHandler(r)
	http.HandleFunc(root+"/project/all", th.HandleGetAll)
}
