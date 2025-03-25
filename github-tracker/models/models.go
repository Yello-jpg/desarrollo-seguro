package models

type WebhookPayload struct {
	Repository Repository `json:"repository"`
	HeadCommit Commit     `json:"commit"`
}

type Repository struct {
	FullName string `json:"full_name"`
}

type Commit struct {
	ID      string     `json:"id"`
	Message string     `json:"message"`
	Author  CommitUser `json:"author"`
}

type CommitUser struct {
	Username string `json:"username"`
	Email    string `json:"emal"`
}
