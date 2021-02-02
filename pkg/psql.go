package pkg

import (
	"database/sql"
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
	Driver         string
	Username       string
	Password       string
	Host           string
	Port           int
	Database       string
	Timeout        time.Duration //Total timeout for try
	ConnectTimeout time.Duration //Postgres connection timeout
	Retries        int
	FailOnPG       bool //defaults to true. A error is thrown if the client can establish a connectoin to the database, but the call fails because of an postgres error. (e.g. wrhong credentials, non existing database...=
}

func (s *SQLConnection) Connect() error {

	db, err := sql.Open(s.Driver, getPGConnection(*s))
	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil && !s.FailOnPG {
		_, ok := err.(pq.PGError)
		if ok {
			//dont handle if type is from pgError
			return nil
		}
	}

	return err
}

func OpenSQL(connection SQLConnection) error {
	return ConnectWithRetries(&connection, connection.Retries, connection.ConnectTimeout, connection.Timeout)
}

func getMySQLConnection(c SQLConnection) string {
	return fmt.Sprintf("%s:%s@%s:%d/%s?timeout=%ds", c.Username, c.Password, c.Host, c.Port, c.Database, int(c.Timeout.Seconds()))
}

func getPGConnection(c SQLConnection) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s connect_timeout=%d sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Database, int(c.Timeout.Seconds()))
}
