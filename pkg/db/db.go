// db package provides connectivity for Postgresql 13
// This package contains functions and methods to interact with postgresql database.
// Written by Nursultan Kuandyk (github.com/lunarnuts)
// date modified: 10/30/2021

package db

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/jackc/pgx/v4"
)

// Config is the settings used to establish a connection to a PostgreSQL server.
// Used to pass database configuration
type Setup struct {
	User     string // database username
	Password string // database user password
	Host     string // database host address
	Port     int    // database port
	Name     string // database database name
	Type     string // database type
}

// string is used to parse connection parameters into a string that is used by pgx.Connect(ctx,string)
func (dbs *Setup) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		dbs.Type, dbs.User, dbs.Password, dbs.Host, dbs.Port, dbs.Name)
}

// New accepts Setup struct and returns a pointer to Database connection
func New(dbs *Setup) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbs.String())
	if err != nil {
		return nil, fmt.Errorf("unable to connect database: %v", err)
	}
	return conn, err
}

func DeleteAllData(dbs *Setup) (*pgx.Conn, error) {
	c, err := ioutil.ReadFile("pkg/db/delete.sql")
	if err != nil {
		return nil, fmt.Errorf("unable read file: %v", err)
	}
	sql := string(c)
	conn, err := New(dbs)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("unable to delete data from database: %v", err)
	}
	conn.Close(context.Background())
	conn, err = New(dbs)
	return conn, err
}
