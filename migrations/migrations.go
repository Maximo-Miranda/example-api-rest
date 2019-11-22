package main

import (
	"os"
	"fmt"
	"flag"
	"strings"

	"github.com/Maximo-Miranda/example-api-rest/migrations/functions"
)

func main() {

		migration := flag.String("migration", "", "Define run command up or down")
		flag.Parse()

		switch strings.ToUpper(*migration)  {
		case "UP":
			_ = functions.UpMigrations()
		case "DOWN":
			_ = functions.DownMigration()
		default:
			fmt.Println("The migration flag is requered, exmaple: -migration=up")
			os.Exit(1)
		}

}
