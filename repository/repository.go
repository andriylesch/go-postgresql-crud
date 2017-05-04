package repository

import (
	"errors"

	"github.com/go-postgresql-crud/database"
	"github.com/go-postgresql-crud/logger"
)

//Data ...
type Repository struct {
	DB database.Database
}

// New ...
func New(db database.Database) Repository {
	return Repository{DB: db}
}

// GetData ...
func (r Repository) GetAutoBrands() []AutoBrand {

	brands := []AutoBrand{}
	if err := r.DB.Select(&brands, "SELECT id, name FROM AutoBrands;"); err != nil {
		logger.Error(err, "Error during select data")
	}

	return brands
}

// GetAutoBrandsByID ...
func (r Repository) GetAutoBrandsByID(brandID int) AutoBrand {

	brand := AutoBrand{}
	if err := r.DB.Get(&brand, "SELECT id, name FROM AutoBrands WHERE id = $1;", brandID); err != nil {
		logger.Error(err, "Error during select data")
	}
	return brand
}

// CreateTable ...
func (r Repository) CreateTable() {
	sqlQuery := `CREATE TABLE IF NOT EXISTS AutoBrands (
					id int NOT NULL,
					name varchar(120) );`
	r.DB.MustExec(sqlQuery)
}

// DropTable ...
func (r Repository) DropTable() {
	sqlQuery := `DROP TABLE IF EXISTS AutoBrands;`
	r.DB.MustExec(sqlQuery)
}

// CreateAutoBrands ...
func (r Repository) CreateAutoBrands(array []AutoBrand) {
	for _, item := range array {
		_, err := r.DB.Exec("INSERT INTO AutoBrands VALUES ($1,$2);", item.ID, item.Name)
		if err != nil {
			logger.Error(err, "Error during insert data")
		}
	}
}

// CreateAutoBrand ...
func (r Repository) CreateAutoBrand(brand AutoBrand) (int, error) {

	if brand.ID == 0 || brand.Name == "" {
		return 0, errors.New("brand parameter is not correct")
	}

	var id int
	err := r.DB.Get(&id, "INSERT INTO AutoBrands VALUES ($1,$2) RETURNING id;", brand.ID, brand.Name)
	if err != nil {
		logger.Error(err, "Error during insert data :")
		return 0, err
	}

	return id, nil
}

// CreateAutoBrand ...
// func (r Repository) UpdateAutoBrandsByID(id int, name string) (int, error) {

// 	if id == 0 || name == "" {
// 		return 0, errors.New("id or name parameters are not correct")
// 	}

// 	err := r.DB.Exec("UPDATE AutoBrands SET name=$1 RETURNING id;", brand.ID, brand.Name)
// 	if err != nil {
// 		logger.Error(err, "Error during insert data :")
// 		return 0, err
// 	}

// 	return id, nil
// }

// var DB *sqlx.DB

// func createTable() {
// 	sqlQuery := `CREATE TABLE IF NOT EXISTS Makes (
// 					makeid int NOT NULL,
// 					name varchar(120)
// 				);`
// 	// id SERIAL PRIMARY KEY,

// 	DB.MustExec(sqlQuery)
// }

// func dropTable() {
// 	sqlQuery := `DROP TABLE IF EXISTS Makes;`
// 	DB.MustExec(sqlQuery)
// }

// /* ----------------------------------------------------- */

// func initDb() {

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err := sqlx.Connect("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Error during connection :", err)
// 		return
// 	}

// 	if err = db.Ping(); err != nil {
// 		fmt.Println("Error during ping :", err)
// 		return
// 	}

// 	DB = db
// }

// func DoOperations(array []models.Make) {

// 	if len(array) == 0 {
// 		fmt.Println("array parameter is empty")
// 		return
// 	}

// 	fmt.Println(" 1) Init database")
// 	initDb()

// 	fmt.Println(" 2) Create 'Makes' table")
// 	createTable()

// 	fmt.Println(" 3) Insert data to 'Makes' table")
// 	tx := DB.MustBegin()
// 	for _, item := range array {
// 		tx.MustExec("INSERT INTO Makes VALUES ($1,$2);", item.ID, item.Name)
// 	}
// 	tx.Commit()

// 	// get the data back
// 	fmt.Println(" 4) get data from 'Makes' table")
// 	makes := []makeDAO{}
// 	if err := DB.Select(&makes, "SELECT makeid, name FROM Makes;"); err != nil {
// 		fmt.Println("Error during select data from DB :", err)
// 	}

// 	for _, item := range makes {
// 		fmt.Printf("makeId = %d ; Name = '%v' ;\n", item.MakeID, item.Name)
// 	}

// 	fmt.Println(" 3) Drop 'Makes' table")
// 	dropTable()
// }
