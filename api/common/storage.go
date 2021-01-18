package common

import (
	"github.com/go-pg/pg/v10"
	"github.com/marioarizaj/serverless-example/models"
)

type StorageIFace interface {
	CreateAuthor(author *models.Author) (*models.Author, error)
	GetAllAuthors() ([]models.Author, error)
	UpdateAuthor(author *models.Author) (*models.Author, error)
	DeleteAuthor(author *models.Author) error

	CreateArticle(author *models.Article) (*models.Article, error)
	GetAllArticles() ([]models.Article, error)
	UpdateArticle(author *models.Article) (*models.Article, error)
	DeleteArticle(author *models.Article) error
}

type Storage struct {
	db *pg.DB
}

var sto StorageIFace

func ConnectToDB() StorageIFace {
	if sto != nil {
		return sto
	}
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "user",
		Password: "pass",
		Database: "db_name",
	})

	sto = &Storage{db: db}

	return sto
}
