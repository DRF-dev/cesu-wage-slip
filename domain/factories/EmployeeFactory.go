package factories

import (
	"scrapers-go/domain/models"
	modelsInfra "scrapers-go/infra/models"
)

type EmployeeFactory interface {
	FromAdapterToModel(adapter modelsInfra.EmployeeModelAdapter) models.EmployeeModel
}
