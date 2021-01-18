package common

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/marioarizaj/serverless-example/models"
)

type StorageIFace interface {
	CreateAuthor(author *models.Author) (*models.Author, error)
	GetAllAuthors() ([]models.Author, error)

	CreateArticle(article *models.Article) (*models.Article, error)
	GetAllArticles() ([]models.Article, error)
}

type Storage struct {
	db *pg.DB
}

var sto StorageIFace

func ConnectToDB() (StorageIFace, error) {
	if sto != nil {
		return sto, nil
	}
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
	err := db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	sto = &Storage{db: db}

	return sto, nil
}
