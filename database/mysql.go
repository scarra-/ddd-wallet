package database

import (
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect() error {
	var err error

	dsn, _ := strings.CutPrefix(os.Getenv("MYSQL_DSN"), "mysql://")

	DBConn, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{DisableForeignKeyConstraintWhenMigrating: true},
	)

	return err
}
