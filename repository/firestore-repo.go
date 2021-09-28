package repository

import (
	"../entity"
)

type repo struct{}

//NewFirestoreRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	// implement save method for Firestore
	return &entity.Post{}, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	// implement get all method for Firestore
	return []entity.Post{}, nil
}
