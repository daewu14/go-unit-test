package repository

import (
	"belajar-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repo *CategoryRepositoryMock) FindById(id string) *entity.Category {
	argument := repo.Mock.Called(id)
	if argument.Get(0) == nil {
		return nil
	} else {
		category := argument.Get(0).(entity.Category)
		return &category
	}
}
