package functions

import (
	"fmt"

	"github.com/Maximo-Miranda/example-api-rest/tools"
	usr "github.com/Maximo-Miranda/example-api-rest/internal/user"
)

// DownMigration ...
func DownMigration() error {

	conn, err := tools.Connect()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)

	// Delete tables
	migrate := conn.DropTableIfExists(
		&usr.User{},
	)
	if migrate.Error != nil {
		fmt.Println("Rollback Migration Fail", migrate.Error)
		return migrate.Error
	}

	return nil
}
