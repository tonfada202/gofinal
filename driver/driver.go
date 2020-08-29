package driver

import (
	"database/sql"
	"log"
)

var db *sql.DB
var url = "postgres://gdipzyef:cUya0IVsMv02m2LPbFMZDVM8_pOa0Agw@lallah.db.elephantsql.com:5432/gdipzyef"

func ConnectSQL() *sql.DB {
	var err error
	db, err = sql.Open("postgres", url)
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // ตอนสอบต้อง
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTable() {
	db, err := sql.Open("postgres", url)
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // ตอนสอบต้อง
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS customer (
		id SERIAL PRIMARY KEY, 
		name TEXT,
		email TEXT,
		status TEXT
		);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create ", err)
	}
	//log.Println("Okay")
}
