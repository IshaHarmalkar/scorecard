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