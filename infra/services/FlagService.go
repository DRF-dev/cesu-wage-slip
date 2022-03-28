package services

import (
	"flag"
	"scrapers-go/domain/factories"
	"scrapers-go/domain/models"
)

type FlagServiceImpl struct{
	Cf factories.CredentialsFactory
	Pf factories.PathFactory
}

func (receiver FlagServiceImpl) GetAll() (models.CredentialsModel, models.PathModel) {
	username := flag.String("u", "username", "Username used for login")
	password := flag.String("p", "password", "Password used for login")
	folder := flag.String("f", "/cesu", "Folder where download all wage slips")

	flag.Parse()

	credentialsModel := receiver.Cf.CreateModel(*username, *password)
	pathModel := receiver.Pf.CreateModel(*folder)

	return credentialsModel, pathModel
}
