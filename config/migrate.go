package config

import (
	"log"

	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@localhost:"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_NAME")+"?sslmode=disable",
	)

	if err != nil {
		log.Fatal("Migration init error:", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration success")
}
