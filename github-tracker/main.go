package main

import (
	"context"
	"fmt"
	"github-tracker/github-tracker/models"
	"github-tracker/github-tracker/repository"

	// "github-tracker/github-tracker/repository"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received POST request!")
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the request")
		return
	}

	fmt.Printf(string(body))
}

func insertGitHubWebHook(ctx context.Context, repo repository.Commit, webhook models.WebhookPayload, BODY string, createdAt time.Time) {
	// commit := entity.Commit{
	// 	RepoName:       webhook.Repository.FullName,
	// 	CommitID:       webhook.HeadCommit.ID,
	// 	CommitMessage:  webhook.HeadCommit.Message,
	// 	AuthorUsername: webhook.HeadCommit.Author.Username,
	// 	AuthorMessage:  webhook.HeadCommit.Message,
	// 	AuthorEmail:    webhook.HeadCommit.Author.Email,
	// 	Payload:        BODY,
	// 	CreatedAt:      createdAt,
	// 	UpdatedAt:      createdAt,
	// }
	// if commit != nil {
	// 	fmt.Printf("NO NULL")
	// } else {
	// 	fmt.Printf("NULL")
	// }
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", postHandler).Methods("POST")

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
