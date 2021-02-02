package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	/*"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"*/)

func Connect() *sql.DB {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_DATABASE")

	URL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", HOST, PORT, USER, PASS, DBNAME)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect postgre")
}

/*func Connect() {
  db, err := gorm.Open("postgres", "host=localhost port=5432 user=lakuin dbname=lakuin password=lakuin123")
  defer db.Close()
	if err != nil{
		log.Fatal(err)
	}
  fmt.Println("sukses")
}*/
