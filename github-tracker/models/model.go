package models

// Commit represents a Git commit
type Commit struct {
  ID        string   `json:"id"`
  Message   string   `json:"message"`
  Author    CommitUser    `json:"author"`
}

// Repository represents a GitHub repository
type Repository struct {
  FullName      string `json:"full_name"`
}

// PushEvent represents a GitHub push event
type GithubWebhook struct {
  Repository Repository `json:"repository"`
  HeadCommit Commit     `json:"head_commit"`
}

// CommitUser represents a Git commit author
type CommitUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
