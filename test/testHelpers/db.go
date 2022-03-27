package testhelpers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-txdb"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ExecMigrations(postgresURL string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "../../../db/migrations",
	}
	pg, err := sql.Open("postgres", postgresURL)
	if err != nil {
		logrus.Fatal(err)
	}

	appliedCount, err := migrate.Exec(pg, "postgres", migrations, migrate.Up)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	log.Printf("Applied %v migrations", appliedCount)
	return nil
}
func NewTest(name string) (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		// "test-db",
		"localhost",
		25432,
		"postgresql",
		"postgresql",
		"test-api",
	)
	// migrate
	err := ExecMigrations(conn)
	if err != nil {
		panic(err)
	}

	txdb.Register(name, "postgres", conn)
	dialector := postgres.New(
		postgres.Config{
			DriverName: name,
			DSN:        conn,
		},
	)
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return db, nil
}

func Prepare(name string, seeds []interface{}) (*gorm.DB, error) {
	db, err := NewTest(name)
	if err != nil {
		return nil, err
	}

	for _, s := range seeds {
		if err := db.Create(s).Error; err != nil {
			return nil, err
		}
	}

	return db, nil
}
