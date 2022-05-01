package globals

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ibidathoillah/majoo-test/config"
	"github.com/ibidathoillah/majoo-test/lib/database"
	"github.com/jmoiron/sqlx"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

var db *sqlx.DB
var once sync.Once

func DB() *sqlx.DB {
	once.Do(func() {
		var err error
		var conn string

		var driver = config.GetEnv(config.DB_DRIVER)

		if driver == "postgres" {
			conn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				config.GetEnv(config.DB_HOST),
				config.GetEnv(config.DB_PORT),
				config.GetEnv(config.DB_USER),
				config.GetEnv(config.DB_PASS),
				config.GetEnv(config.DB_NAME),
			)

		} else if driver == "mysql" {
			conn = fmt.Sprintf(
				"%s:%s@(%s:%s)/%s?multiStatements=true",
				config.GetEnv(config.DB_USER),
				config.GetEnv(config.DB_PASS),
				config.GetEnv(config.DB_HOST),
				config.GetEnv(config.DB_PORT),
				config.GetEnv(config.DB_NAME),
			)
		}

		db, err = sqlx.Connect(driver, conn)

		if nil != err {
			log.Fatal(err)
		}
	})
	return db
}

func GetDefaultLimit() int64 {
	return int64(10)
}

func GetQuery(ctx context.Context) *database.Queryable {
	q, ok := database.QueryFromContext(ctx)
	if !ok {
		panic("values when get query from context. please using transaction")
	}
	return &q
}
