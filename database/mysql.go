package database

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect() error {
	var err error

	dsn, _ := strings.CutPrefix(os.Getenv("MYSQL_DSN"), "mysql://")

	maxWaitTime := 15 * time.Second

	startTime := time.Now()

	for {
		DBConn, err = gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{DisableForeignKeyConstraintWhenMigrating: true},
		)

		if err == nil {
			break
		}

		if time.Since(startTime) >= maxWaitTime {
			return fmt.Errorf("timed out waiting for database connection")
		}

		time.Sleep(1 * time.Second)
	}

	return err
}
