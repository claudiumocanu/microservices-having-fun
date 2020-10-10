package database

import (
	// "database/sql"
	"github.com/jinzhu/gorm"

	// _ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DBConn *gorm.DB

// func init() {
// 	var err error
// 	connStr := "postgres://gouser:gopass@pgdb/godb?sslmode=disable"
// 	db, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
