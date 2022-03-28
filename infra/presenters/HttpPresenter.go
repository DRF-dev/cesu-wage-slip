package presenters

import (
	"fmt"
	urlForm "net/url"
	"scrapers-go/domain/enums"
	"scrapers-go/domain/factories"
	"scrapers-go/domain/models"
	"scrapers-go/domain/services"
	modelsInfra "scrapers-go/infra/models"
	"time"
)

type HttpPresenterImpl struct{
	HttpService services.HttpService
	EmployeeFactory factories.EmployeeFactory
	WageSlipFactory factories.WageSlipFactory
}

func (receiver HttpPresenterImpl) PostLogin(credentialsModel models.CredentialsModel) error {
	form := urlForm.Values{}
	form.Set(enums.Username, credentialsModel.Username)
	form.Set(enums.Password, credentialsModel.Password)

	err := receiver.HttpService.PostForm("https://www.cesu.urssaf.fr/info/accueil.login.do", form, nil)
	if err != nil {
		return err
	}

	return nil
}

func (receiver HttpPresenterImpl) GetEmployee() (models.EmployeeModel, error) {
	var employeeModel models.EmployeeModel

	var employeeModelAdapter modelsInfra.EmployeeModelAdapter

	err := receiver.HttpService.Get("https://www.cesu.urssaf.fr/cesuwebdec/status", &employeeModelAdapter)
	if err != nil {
		return employeeModel, err
	}

	employeeModel = receiver.EmployeeFactory.FromAdapterToModel(employeeModelAdapter)
	return employeeModel, nil
}

func (receiver HttpPresenterImpl) GetWageSlips(employeeModel *models.EmployeeModel) error {
	var wageSlipAdapter modelsInfra.WageSlipAdapter

	year, _, _ := time.Now().Date()
	fiveYearsAgo := year - 5
	url := fmt.Sprintf("https://www.cesu.urssaf.fr/cesuwebdec/salaries/%v/bulletinsSalaire?pseudoSiret=&dtDebutRecherche=%v0101&dtFinRecherche=%v1231&numStart=0&nbAffiche=100000&numeroOrdre=0&orderBy=orderByRefDoc", employeeModel.Id, fiveYearsAgo, year)

	err := receiver.HttpService.Get(url, &wageSlipAdapter)
	if err != nil {
		return err
	}

	for _, object := range wageSlipAdapter.ListObjects {
		wageSlip := receiver.WageSlipFactory.CreateModel(object.DocumentaryReference)
		employeeModel.WageSlips = append(employeeModel.WageSlips, wageSlip)
	}

	return nil
}

