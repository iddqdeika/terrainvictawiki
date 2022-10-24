package tech

import (
	"encoding/json"
	"os"
)

type Metadata struct {
	FriendlyName              string
	DataName                  string
	TechCategory              string
	AITechRole                string
	AICriticalTech            bool
	EndGameTech               bool
	ResearchCost              int
	Prereqs                   []string
	AltPrereq0                string
	Year                      string
	Effects                   []string
	IconResource              string
	CompletedIllustrationPath string
	VoiceoverPath             string
}

type templateObj struct {
	FriendlyName              string   `json:"friendlyName"`
	DataName                  string   `json:"dataName"`
	TechCategory              string   `json:"techCategory"`
	AITechRole                string   `json:"AI_techRole"`
	AICriticalTech            *bool    `json:"AI_criticalTech"`
	EndGameTech               *bool    `json:"endGameTech"`
	ResearchCost              int      `json:"researchCost"`
	Prereqs                   []string `json:"prereqs"`
	AltPrereq0                string   `json:"altPrereq0"`
	Year                      string   `json:"year"`
	Effects                   []string `json:"effects"`
	IconResource              string   `json:"iconResource"`
	CompletedIllustrationPath string   `json:"completedIllustrationPath"`
	VoiceoverPath             string   `json:"voiceoverPath"`
}

func (o *templateObj) toMetadata() *Metadata {
	res := &Metadata{
		FriendlyName:              o.FriendlyName,
		DataName:                  o.DataName,
		TechCategory:              o.TechCategory,
		AITechRole:                o.AITechRole,
		AICriticalTech:            o.AICriticalTech != nil && *o.AICriticalTech,
		EndGameTech:               o.EndGameTech != nil && *o.EndGameTech,
		ResearchCost:              o.ResearchCost,
		Prereqs:                   o.Prereqs,
		AltPrereq0:                o.AltPrereq0,
		Year:                      o.Year,
		Effects:                   o.Effects,
		IconResource:              o.IconResource,
		CompletedIllustrationPath: o.CompletedIllustrationPath,
		VoiceoverPath:             o.VoiceoverPath,
	}
	return res
}

func getMetadataFromFile(filename string) ([]*Metadata, error) {
	list, err := readTemplateFile(filename)
	if err != nil {
		return nil, err
	}
	res := make([]*Metadata, len(list))
	for i, obj := range list {
		res[i] = obj.toMetadata()
	}
	return res, nil
}

func readTemplateFile(filename string) ([]*templateObj, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var list []*templateObj
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
