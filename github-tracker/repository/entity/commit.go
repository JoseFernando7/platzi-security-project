package entity

import "time"

type Commit struct {
	ID      		int     `json:"id"`
	RepoName 		string  `json:"repo_name"`
	CommitID 		string 	`json:"commit_id"`
	CommitMessage 	string 	`json:"commit_message"`
	AuthorUsername 	string 	`json:"author_username"`
	AuthorEmail 	string 	`json:"author_email"`
	Payload 		string 	`json:"payload"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
}
