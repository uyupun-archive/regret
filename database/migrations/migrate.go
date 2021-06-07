package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/uyupun/regret/database"
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
		Dir: "database/migrations",
	}
	return migrations
}

func connectDB() *sql.DB {
	dsn, err := database.GetDSN()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
