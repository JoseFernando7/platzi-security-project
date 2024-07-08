package main

import (
	"context"
	"fmt"
	"github-tracker/github-tracker/models"
	"github-tracker/github-tracker/repository"
	"github-tracker/github-tracker/repository/entity"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Another comment
func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received POST request!")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)

		return
	}

	fmt.Println("Body:", string(body))
}

func insertGithubWebhook(ctx context.Context, repo repository.Commit, webhook models.GithubWebhook, body string, createdTime time.Time) error {
	commit := entity.Commit {
		RepoName: 		webhook.Repository.FullName,
		CommitID: 		webhook.HeadCommit.ID,
		CommitMessage: 	webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail: 	webhook.HeadCommit.Author.Email,
		Payload: 		body,
		CreatedAt: 		createdTime,
		UpdatedAt: 		createdTime,
	}

	err := repo.Insert(ctx, &commit)

	return err
}

// Comment
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", postHandler).Methods("POST")

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err.Error())
	}
}
