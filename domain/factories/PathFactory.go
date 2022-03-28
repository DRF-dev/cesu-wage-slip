package factories

import "scrapers-go/domain/models"

type PathFactory interface {
	CreateModel(path string) models.PathModel
}
