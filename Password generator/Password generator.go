package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var Config Configuration

type Configuration struct {
	Host     string
	DBname   string
	User     string
	Password string
	Port     int
}

func main() {
	// passwordgenerator
	passwordLength := 20
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

	password := make([]byte, passwordLength)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println("Random wachtwoord:", string(password))

	// JSON file --------------------------------------------------------------------------------
	file, err := os.Open("conf.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Database
	connectionstring := fmt.Sprintf("%s:%s@tcp(%s)/%s", Config.User, Config.Password, Config.Host, Config.DBname)
	db, err = sql.Open("mysql", connectionstring)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
