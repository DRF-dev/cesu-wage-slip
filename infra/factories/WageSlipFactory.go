package factories

import (
	"scrapers-go/domain/models"
)

type WageSlipFactoryImpl struct{}

func (WageSlipFactoryImpl) CreateModel(reference string) models.WageSlip {
	return models.WageSlip{
		ReferencePdf: reference,
	}
}
