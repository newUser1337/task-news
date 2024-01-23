package migration

import (
	"fmt"
	"log"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/newUser1337/task-news/internal/config"
)

func Migrate(cfg *config.Mongo) error {
	mongConnection := strings.Join([]string{cfg.Address, cfg.DbName}, "/")
	migrationDriver, err := database.Open(mongConnection)
	if err != nil {
		return fmt.Errorf("failed to create instance %s", err)
	}

	defer func() {
		if err := migrationDriver.Close(); err != nil {
			log.Fatalf("failed to close migration driver %s", err)
		}
	}()

	m, err := migrate.NewWithDatabaseInstance("file://migration", cfg.DbName, migrationDriver)
	if err != nil {
		return fmt.Errorf("failed to execute migration scripts: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
