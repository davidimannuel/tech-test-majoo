package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"majoo/soal1/configs"
)

func SetupPostgresDB(conf configs.Postgres) *sql.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Password,
		conf.DB)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	return db
}
