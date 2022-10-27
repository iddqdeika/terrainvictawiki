package project

import (
	"encoding/json"
	"os"
)

type Metadata struct {
	FriendlyName              string
	DataName                  string
	TechCategory              string
	AITechRole                string
	AICriticalTech            *bool
	AIProjectRole             string
	ResearchCost              int
	Prereqs                   []string
	AltPrereq0                string
	RequiredObjectiveNames    []string
	AltRequiredObjective0Name string
	RequiredMilestone         string
	Effects                   []string
	OneTimeGlobally           bool
	Repeatable                bool
	RequiresNation            string
	RequiredControlPoint      []string
	FactionPrereq             []string
	FactionAvailableChance    int
	FactionAlways             string
	InitialUnlockChance       int
	DeltaUnlockChance         int
	MaxUnlockChance           int
	OrgGranted                string
	ResourcesGranted          []*Resource
	IconResource              string
	CompletedIllustrationPath string
}

type templateObj struct {
	FriendlyName              string      `json:"friendlyName"`
	DataName                  string      `json:"dataName"`
	TechCategory              string      `json:"techCategory"`
	AITechRole                string      `json:"AI_techRole"`
	AICriticalTech            *bool       `json:"AI_criticalTech"`
	AIProjectRole             string      `json:"AI_projectRole"`
	ResearchCost              int         `json:"researchCost"`
	Prereqs                   []string    `json:"prereqs"`
	AltPrereq0                string      `json:"altPrereq0"`
	RequiredObjectiveNames    []string    `json:"requiredObjectiveNames"`
	AltRequiredObjective0Name string      `json:"altRequiredObjective0Name"`
	RequiredMilestone         string      `json:"requiredMilestone"`
	Effects                   []string    `json:"effects"`
	OneTimeGlobally           bool        `json:"oneTimeGlobally"`
	Repeatable                bool        `json:"repeatable"`
	RequiresNation            string      `json:"requiresNation"`
	RequiredControlPoint      []string    `json:"requiredControlPoint"`
	FactionPrereq             []string    `json:"factionPrereq"`
	FactionAvailableChance    int         `json:"factionAvailableChance"`
	FactionAlways             string      `json:"factionAlways"`
	InitialUnlockChance       int         `json:"initialUnlockChance"`
	DeltaUnlockChance         int         `json:"deltaUnlockChance"`
	MaxUnlockChance           int         `json:"maxUnlockChance"`
	OrgGranted                string      `json:"orgGranted"`
	ResourcesGranted          []*Resource `json:"resourcesGranted"`
	IconResource              string      `json:"iconResource"`
	CompletedIllustrationPath string      `json:"completedIllustrationPath"`
}

func (t *templateObj) ToMetadata() *Metadata {
	return &Metadata{
		FriendlyName:              t.FriendlyName,
		DataName:                  t.DataName,
		TechCategory:              t.TechCategory,
		AITechRole:                t.AITechRole,
		AICriticalTech:            t.AICriticalTech,
		AIProjectRole:             t.AIProjectRole,
		ResearchCost:              t.ResearchCost,
		Prereqs:                   t.Prereqs,
		AltPrereq0:                t.AltPrereq0,
		RequiredObjectiveNames:    t.RequiredObjectiveNames,
		AltRequiredObjective0Name: t.AltRequiredObjective0Name,
		RequiredMilestone:         t.RequiredMilestone,
		Effects:                   t.Effects,
		OneTimeGlobally:           t.OneTimeGlobally,
		Repeatable:                t.Repeatable,
		RequiresNation:            t.RequiresNation,
		RequiredControlPoint:      t.RequiredControlPoint,
		FactionPrereq:             t.FactionPrereq,
		FactionAvailableChance:    t.FactionAvailableChance,
		FactionAlways:             t.FactionAlways,
		InitialUnlockChance:       t.InitialUnlockChance,
		DeltaUnlockChance:         t.DeltaUnlockChance,
		MaxUnlockChance:           t.MaxUnlockChance,
		OrgGranted:                t.OrgGranted,
		ResourcesGranted:          t.ResourcesGranted,
		IconResource:              t.IconResource,
		CompletedIllustrationPath: t.CompletedIllustrationPath,
	}
}

type Resource struct {
	Resource string      `json:"resource"`
	Value    interface{} `json:"value"`
}

func getMetadataFromFile(filename string) ([]*Metadata, error) {
	list, err := readTemplateFile(filename)
	if err != nil {
		return nil, err
	}
	res := make([]*Metadata, 0)
	for _, o := range list {
		res = append(res, o.ToMetadata())
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
