package factories

import "scrapers-go/domain/models"

type CredentialsFactoryImpl struct{}

func (CredentialsFactoryImpl) CreateModel(username string, password string) models.CredentialsModel {
	return models.CredentialsModel{
		Username: username,
		Password: password,
	}
}
