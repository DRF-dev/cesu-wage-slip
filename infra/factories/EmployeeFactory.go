package factories

import (
	"scrapers-go/domain/models"
	modelsInfra "scrapers-go/infra/models"
)

type EmployeeFactoryImpl struct{}

func (EmployeeFactoryImpl) FromAdapterToModel(adapter modelsInfra.EmployeeModelAdapter) models.EmployeeModel {
	return models.EmployeeModel{
		Id: adapter.Object.Numero,
		WageSlips: make([]models.WageSlip, 1),
	}
}
