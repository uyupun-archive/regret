package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("Too few arguments")
	}

	command := args[1]
	fmt.Println(command)
	if command == "up" {
		migrateUp()
	} else if command == "down" {
		migrateDown()
	} else {
		panic("Command not match")
	}
}

func migrateUp() {
	migrations := getMigrationsRef()

	db := connectDB()
	defer db.Close()

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func migrateDown() {
	migrations := getMigrationsRef()

	db := connectDB()
	defer db.Close()

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Down)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func getMigrationsRef() *migrate.FileMigrationSource {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	return migrations
}

func connectDB() *sql.DB {
	dsn := getDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func getDSN() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	parseTime := "true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s", userName, password, host, port, name, parseTime)

	return dsn
}
