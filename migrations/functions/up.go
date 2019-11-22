package functions

import (
	"fmt"

	"github.com/Maximo-Miranda/example-api-rest/tools"
	usr "github.com/Maximo-Miranda/example-api-rest/internal/user"
)

// UpMigrations ...
func UpMigrations() error {

	conn, err := tools.Connect()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)

	// Create Tables
	migrate := conn.AutoMigrate(
		&usr.User{},
	)
	if migrate.Error != nil {
		fmt.Println("Migration Fail", migrate.Error)
		return migrate.Error
	}

	resultIdxUserDni := conn.Model(&usr.User{}).AddUniqueIndex("idx_users_dni", "dni")
	if resultIdxUserDni.Error != nil {
		fmt.Println("Migration Fail", resultIdxUserDni.Error)
		return migrate.Error
	}

	return nil
}
