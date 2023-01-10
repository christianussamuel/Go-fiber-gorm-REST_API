package models

type Master struct {
	ProfileCode       int64  `json:"profileCode,omitempty"`
	Photo             string `json:"photo,omitempty"`
	WorkingExperience string `json:"workingExperience,omitempty"`
	Employment        string `json:"employment,omitempty"`
	Education         string `json:"education,omitempty"`
	Skill             string `json:"skill,omitempty"`
}
