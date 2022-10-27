package tech

import (
	"fmt"
	"terrainvictawiki/pkg/localization"
)

const (
	defaultTemplatePath     = "./res/data/Templates/TITechTemplate.json"
	defaultLocalizationPath = "./res/data/Localization/en/TITechTemplate.en"
	defaultLocalizationRoot = "TITechTemplate"
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
