package main

import (
	"log"
	"scrapers-go/domain/factories"
	"scrapers-go/domain/presenters"
	"scrapers-go/domain/services"
	factoriesImpl "scrapers-go/infra/factories"
	presentersImpl "scrapers-go/infra/presenters"
	servicesImpl "scrapers-go/infra/services"
)

func main() {
	// Dependencies
	var credentialsFactory factories.CredentialsFactory = factoriesImpl.CredentialsFactoryImpl{}
	var pathFactory factories.PathFactory = factoriesImpl.PathFactoryImpl{}
	var employeeFactory factories.EmployeeFactory = factoriesImpl.EmployeeFactoryImpl{}

	var flagService services.FlagService = servicesImpl.FlagServiceImpl{
		Cf: credentialsFactory,
		Pf: pathFactory,
	}

	var httpService services.HttpService = &servicesImpl.HttpServiceImpl{}
	httpService.SetClient(true)

	var httpPresenter presenters.HttpPresenter = presentersImpl.HttpPresenterImpl{
		HttpService: httpService,
		EmployeeFactory: employeeFactory,
	}

	// MAIN
	// credentialsModel, pathModel := flagService.GetAll()
	credentialsModel, _ := flagService.GetAll()

	err := httpPresenter.PostLogin(credentialsModel)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	employeeModel, err := httpPresenter.GetEmployee()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	err = httpPresenter.GetWageSlips(&employeeModel)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	// fmt.Println(employeeModel.WageSlips)

	// var wg sync.WaitGroup
	// for _, wageSlip := range employeeModel.WageSlips {
	// 	wg.Add(1)
	// 	go func(s string) {
	// 		defer wg.Done()
	// 		a := fmt.Sprintf("https://www.cesu.urssaf.fr/cesuwebdec/salaries/%v/editions/bulletinSalairePE?refDoc=%v", httpData.Object.Numero, s)
	//
	// 		res, _ = client.Get(a)
	// 		if err != nil {
	// 			log.Fatalf("Error %v\n", err)
	// 		}
	//
	// 		b := fmt.Sprintf("%v/%v.pdf", *folder, s)
	// 		out, _ := os.Create(b)
	//
	// 		io.Copy(out, res.Body)
	//
	// 		res.Body.Close()
	// 		out.Close()
	// 	}(wageSlip.ReferencePdf)
	// }
	//
	// wg.Wait()
}
