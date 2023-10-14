package utility

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/finance?parseTime=true")
	if err != nil {
		log.Panicln("Failed connect to db", err)
	}

	return db
}

func MigrationDB(db *sql.DB) error {
	log.Printf("Migrasi database sedang dijalankan...")

	dbConnect := ConnectDB()
	migrations := &migrate.FileMigrationSource{
		Dir: "migrate",
	}

	_, err := migrate.Exec(dbConnect, "mysql", migrations, migrate.Up)

	return err
}
