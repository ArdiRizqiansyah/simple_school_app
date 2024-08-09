package database

import (
	"be-classroom/app/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.AppConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DbHost,
		appConfig.DbPort,
		appConfig.DbUser,
		appConfig.DbPassword,
		appConfig.DbName,
	)

	db, err = sql.Open(appConfig.DbDialect, dsn)

	if err != nil {
		log.Panicln("error occured while trying to validate database arguments: ", err.Error())
		return
	}

	if err := db.Ping(); err != nil {
		log.Panicln("error occured while trying to connect to database: ", err.Error())
		return
	}
}

func InitializeDatabase() {
	handleDatabaseConnection()
}

func GetInstanceDatabaseConnection() *sql.DB {
	return db
}

// perintah database postgres
// postgres://postgres:postgres@localhost:5432/nama_database?sslmode=disable

// perintah membuat migrasi
//  migrate create -ext postgres -dir db/migrations create_users_table(nama file migrasi)

// perintah migrasi up
// migrate -database "postgres://postgres:postgres@localhost:5432/nama_database?sslmode=disable" -path db/migrations up

// perintah migrasi down
// migrate -database "postgres://postgres:postgres@localhost:5432/nama_database?sslmode=disable" -path db/migrations down
