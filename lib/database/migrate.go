package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func Migrate(db *sql.DB, dbName string, path string) {
	fmt.Println("Start migration process")
	version, ok := doMigrate(db, dbName, path)
	fmt.Println("Migration Complete (Version, Status): ", version, ok)
}

func doMigrate(db *sql.DB, dbName string, path string) (int, bool) {
	var err error

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if nil != err {
		log.Fatal("error when creating database driver: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", path), dbName, driver)
	if nil != err {
		log.Fatal("error when creating database instance: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error when migrate up: ", err)
	}

	version, dirty, err := m.Version()
	if nil != err {
		log.Fatal(err)
	}

	return int(version), dirty
}
