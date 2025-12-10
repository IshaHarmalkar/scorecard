package models

import "time"

type Scorecard struct {
	Id         int          `json:"id"`
	UserId     int          `json:"userId"`
	Title      string       `json:"title"`
	Url        string       `json:"url"`
	TotalScore int          `json:"totalScore"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
}

type Section struct {
	Id         		int          	`json:"id"`
	FormId     		int          	`json:"userId"`
	SectionName		string       	`json:"title"`
	FormData   		ScorecardData   `json:"formData"`
	TotalScore 		int          	`json:"totalScore"`
	CreatedAt  		time.Time    	`json:"createdAt"`
	UpdatedAt  		time.Time    	`json:"updatedAt"`
}


type ScorecardData struct {
	Question     string   `json:"question"`
	Options    []string   `json:"options"`
	Score      []int      `json:"score"`
}