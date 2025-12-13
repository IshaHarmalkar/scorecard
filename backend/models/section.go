package models

import "time"

type Section struct {
	Id          		int 						`json:"id"`
	ScorecardId 		int 						`json:"scorecardId"`
	SectionTitle  		string  					`json:"sectionTitle"`
	SectionData   		[]QuestionOptionsScore   	`json:"sectionData"`
	TotalScore    		int   						`json:"total_score"`
	CreatedAt    		time.Time   				`json:"created_at"`
	UpdatedAt   		time.Time    				`json:"updated_at"`
}


type QuestionOptionsScore struct {
	Question 			string  		`json:"question"`
	Options  			[]string  		`json:"options"`
	Score    			[]int	 		`json:"score"`  //float?
}



