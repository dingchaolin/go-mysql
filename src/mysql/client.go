package mysql

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

type Client struct{
	Host      string
	Port      int
	User      string
	Password  string
	Database  string
	Charset   string
	ConnectionLimit int
	DriverName string
	Conn *sql.DB
}

func (c * Client) InitConnection(){
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.Charset)

	var err error
	c.Conn, err = sql.Open(c.DriverName, dataSourceName)

	if err != nil {
		log.Panicln(fmt.Sprintf("%v", err ))
	} else {
		c.Conn.SetMaxOpenConns(1)
		c.Conn.SetMaxIdleConns(1)
		c.Conn.Ping()
	}
}


func (c * Client) Fetch(sql string, args ...interface{})(*sql.Rows, error){
	return c.Conn.Query( sql, args... )
}

func (c * Client) FetchNoArgs(sql string)(*sql.Rows, error){
	return c.Conn.Query( sql )
}