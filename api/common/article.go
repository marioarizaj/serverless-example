package common

import "github.com/marioarizaj/serverless-example/models"

func (s *Storage) CreateArticle(article *models.Article) (*models.Article, error) {
	_, err := s.db.Model(article).Insert(article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *Storage) GetAllArticles() ([]models.Article, error) {
	var authors []models.Article
	err := s.db.Model(&models.Article{}).Select(&authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}
