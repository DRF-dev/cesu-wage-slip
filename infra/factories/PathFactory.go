package factories

import (
	"scrapers-go/domain/models"
)

type PathFactoryImpl struct{}

func (PathFactoryImpl) CreateModel(path string) models.PathModel {
	return models.PathModel{
		LocalSaves: path,
	}
}
