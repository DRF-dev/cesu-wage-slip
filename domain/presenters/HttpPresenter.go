package presenters

import (
	"scrapers-go/domain/models"
)

type HttpPresenter interface {
	PostLogin(credentials models.CredentialsModel) error
	GetEmployee() (models.EmployeeModel, error)
	GetWageSlips(employeeModel *models.EmployeeModel) error
}
