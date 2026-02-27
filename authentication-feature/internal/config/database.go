package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	// bisa diganti sesuai local
	// TODO pindahin ke env bro
	// formtanya"root:password@tcp(127.0.0.1:3306)/my_api_db?parseTime=true"
	dsn := "root@tcp(127.0.0.1:3306)/golang_test?parseTime=true"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	// connection pool (belom dibutuhin, tapi ditulis aja biar inget)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
