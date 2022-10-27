package project

import (
	"fmt"
	"terrainvictawiki/pkg/localization"
)

const (
	defaultTemplatePath     = "./res/data/Templates/TIProjectTemplate.json"
	defaultLocalizationPath = "./res/data/Localization/en/TIProjectTemplate.en"
	defaultLocalizationRoot = "TIProjectTemplate"
)

func NewRepo() (*Repository, error) {
	r, err := NewRepositoryFromFile(defaultTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("cannot init tech handler: %v\r\n", err)
	}
	TryLocalizeRepo(defaultLocalizationPath, r)

	return r, nil
}

func TryLocalizeRepo(filename string, r *Repository) {
	l, err := localization.NewStorage(filename)
	if err != nil {
		fmt.Printf("cant load localization file %v: %v\r\n", filename, err)
		return
	}
	f := localization.NewStorageFacade(l, defaultLocalizationRoot)
	r.Localize(f)
}
