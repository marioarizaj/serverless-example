package common

import "github.com/marioarizaj/serverless-example/models"

func (s *Storage) CreateAuthor(author *models.Author) (*models.Author, error) {
	_, err := s.db.Model(author).Insert(author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *Storage) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := s.db.Model(&models.Author{}).Select(&authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (s *Storage) UpdateAuthor(author *models.Author) (*models.Author, error) {
	_, err := s.db.Model(author).Update(author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *Storage) DeleteAuthor(author *models.Author) error {
	_, err := s.db.Model(author).Delete(author)
	if err != nil {
		return err
	}
	return nil
}
