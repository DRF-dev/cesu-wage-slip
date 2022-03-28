package factories

import "scrapers-go/domain/models"

type CredentialsFactory interface {
	CreateModel(username string, password string) models.CredentialsModel
}
