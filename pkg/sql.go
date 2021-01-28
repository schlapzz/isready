package pkg

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	_ "gorm.io/driver/postgres"
	"time"
)

//mongodb
//redis
//elasticsearch
//kafka
//curl

type SQLConnection struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Timeout  time.Duration
	Retries  int
}

func OpenSQL(connection SQLConnection) error {

	var err error
	var connString string

	switch connection.Driver {
	case "mysql":
		connString = getMySQLConnection(connection)
	case "postgres":
		connString = getPGConnection(connection)
	default:
		return errors.New("no matching database driver found for: " + connection.Driver)
	}

	var db *sql.DB

	for i := 0; i < connection.Retries; i++ {
		db, err = sql.Open(connection.Driver, connString)
		if err == nil {

			err = db.Ping()

			if err == nil {
				fmt.Println("connected")
				continue
			} else {
				if _, ok := err.(*pq.Error); ok {
					fmt.Printf("PG Ping failed: %s\n", err.Error())
					return err
				}
			}
		}
		fmt.Println("...")
		time.Sleep(10 * time.Second)
	}

	if err != nil {
		fmt.Println("error : " + err.Error())
	}

	return err
}

func getMySQLConnection(c SQLConnection) string {
	return fmt.Sprintf("%s:%s@%s:%d/%s?timeout=%ds", c.Username, c.Password, c.Host, c.Port, c.Database, int(c.Timeout.Seconds()))
}

func getPGConnection(c SQLConnection) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s connect_timeout=%d sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Database, int(c.Timeout.Seconds()))
}
