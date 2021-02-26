package daos

import (
	"database/sql"
	"fibunacci/config"
	"fibunacci/models"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var crudConfig config.Crud

var gConn *sql.DB

type SqlDao struct {
}

func NewSqlDao(fn string) *SqlDao {
	initialize(fn)
	return &SqlDao{}
}

func initialize(fn string) {
	if crudConfig.Ddl.Create == "" {
		crud := config.LoadDdl(fn)
		crudConfig = crud
		log.Print(crudConfig)
		initializeDb()
		createTable()
	}
}

func getConnection() *sql.DB {
	return gConn
}

func initializeDb() error {

	// Real world scenario move this to a secure vault
	dbUser, dbPassword, dbName, dbHostname, dbPort :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHostname == "" || dbPort == "" {
		errorStr := fmt.Sprintf("Invalid database configuration..host=%s port=%s user=%s password=%s dbname=%s", dbHostname, dbPort, dbUser, dbPassword, dbName)
		log.Fatalln(errorStr)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHostname, dbPort, dbUser, dbPassword, dbName)

	log.Println("DSN : ", dsn)

	conn, _ := sql.Open("postgres", dsn)

	gConn = conn

	log.Println("Database connection established")
	return nil
}

func createTable() {
	log.Println("Create Fibonacci table...")
	db := getConnection()
	statement, err := db.Prepare(crudConfig.Ddl.Create)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Fibonacci table created")
}

func (sqlDao *SqlDao) InsertFibonacci(fibonacci *models.Fibonacci, tx *sql.Tx) error {

	log.Println("Persisting Fibonacci run data ...")

	statement, err := tx.Prepare(crudConfig.Ddl.Insert)
	if err != nil {
		return err
	}
	var id int
	err = statement.QueryRow(fibonacci.NextNumber, fibonacci.PreviousNumber).Scan(&id)

	if err != nil {
		return err
	} else {
		fibonacci.Id = int(id)
	}
	return nil
}

func (sqlDao *SqlDao) GetByPk(pk int, tx *sql.Tx) models.Fibonacci {
	row := tx.QueryRow(crudConfig.Ddl.SelectByPk, pk)
	var id int
	var previous uint32
	var next uint32
	row.Scan(&id, &previous, &next)
	byPk := models.Fibonacci{Id: id, PreviousNumber: previous, NextNumber: next}
	return byPk
}

func (sqlDao *SqlDao) Delete() {

}
