package repository

import (
	"context"
	"database/sql"
	"github-tracker/github-tracker/repository/entity"
)

type Commit interface {
	Insert(ctx context.Context, commit *entity.Commit) (err error)
	GetCommitByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error)
}

type commit struct {
	Conn *sql.DB
}

func NewCommit(conn *sql.DB) commit {
	return commit {
		Conn: conn,
	}
}

func (c commit) Insert(ctx context.Context, commit *entity.Commit) (err error) {
	query := `INSERT INTO commits (repo_name, commit_id, commit_message, author_username, author_email, payload, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := c.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRowContext(
		ctx, 
		commit.RepoName, 
		commit.CommitID, 
		commit.CommitMessage, 
		commit.AuthorUsername, 
		commit.AuthorEmail, 
		commit.Payload, 
		commit.CreatedAt, 
		commit.UpdatedAt).Err()

	return err
}

func (c commit) GetCommitByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error) {
	query := `SELECT *
	FROM COMMITS
	WHERE author_email = $1`

	rows, err := c.Conn.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var commit entity.Commit
		err = rows.Scan(
			&commit.ID, 
			&commit.RepoName, 
			&commit.CommitID, 
			&commit.CommitMessage, 
			&commit.AuthorUsername, 
			&commit.AuthorEmail, 
			&commit.Payload, 
			&commit.CreatedAt, 
			&commit.UpdatedAt)
		if err != nil {
			return nil, err
		}

		commits = append(commits, commit)
	}

	return commits, nil
}
