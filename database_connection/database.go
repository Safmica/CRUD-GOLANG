package database_connection

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db,err := sql.Open("mysql", "root:@tcp(localhost:3306)/gdsc")
	if err != nil{
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleTime(5 * time.Minute)
	db.SetMaxLifeTime(60 * time.Minute)

	return db
}
