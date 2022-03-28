package services

import "scrapers-go/domain/models"

type FlagService interface {
	GetAll() (models.CredentialsModel, models.PathModel)
}
