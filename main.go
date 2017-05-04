package main

import (
	"fmt"

	"github.com/go-postgresql-crud/config"
	"github.com/go-postgresql-crud/database"
	"github.com/go-postgresql-crud/logger"
	"github.com/go-postgresql-crud/repository"
)

func init() {
	config.InitConfig()
	logger.InitLogger()
	database.Connect()
}

func main() {
	defer database.Close()

	brands := []repository.AutoBrand{}

	brands = append(brands, repository.AutoBrand{ID: 1, Name: "Chevrolet"})
	brands = append(brands, repository.AutoBrand{ID: 2, Name: "BMW"})
	brands = append(brands, repository.AutoBrand{ID: 3, Name: "AUDI"})
	brands = append(brands, repository.AutoBrand{ID: 4, Name: "FIAT"})
	brands = append(brands, repository.AutoBrand{ID: 5, Name: "Mercedes-Benz"})

	repositoryMng := repository.New(database.DbObject)

	logger.Info("-------- Create table in database.")
	repositoryMng.CreateTable()

	logger.Info("-------- Inserting array to table")
	repositoryMng.CreateAutoBrands(brands)

	logger.Info("-------- Inserting 'Volvo' autobrand to table")
	brand := repository.AutoBrand{ID: 6, Name: "Volvo"}
	id, err := repositoryMng.CreateAutoBrand(brand)
	if err != nil {
		logger.Error(err, "")
	}

	logger.Info(fmt.Sprint("new brand was added into table and Id = ", id))

	logger.Info("-------- Get all data from table")
	result := repositoryMng.GetAutoBrands()
	for index, brandItem := range result {
		logger.Info(fmt.Sprintf("%d) Brand Id = '%d'; Name = '%s'; \n", index, brandItem.ID, brandItem.Name))
	}

	logger.Info("-------- Drop table from database.")
	repositoryMng.DropTable()

}
