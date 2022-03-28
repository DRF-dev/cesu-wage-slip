package factories

import (
	"scrapers-go/domain/models"
)

type WageSlipFactory interface {
	CreateModel(reference string) models.WageSlip
}
