package main

import (
	"mysql"
	"fmt"
)

var (
	Host      =   "192.168.64.184"
	Port      =   3306
	User      =   "read"
	Password  =   "read123"
	Database  =   "doc"
	Charset   =   "utf8"
	ConnectionLimit =  20
	DriverName = "mysql"
)

func main(){
	c := mysql.Client{
		Host:      Host,
		Port:      Port,
		User:      User,
		Password:  Password,
		Database:  Database,
		Charset:   Charset,
		ConnectionLimit: ConnectionLimit,
		DriverName: DriverName,
		Conn:     nil,
	}


	c.InitConnection()
	//sql := "select * from process_list where port > 9090 and name = '本地服务' "
	//rows, err := c.Fetch( sql )
	//sql := "select binlog_id,serverIndex,`database` from binlog where serverIndex >= ? and `database` = ? "
	//args:= []interface{}{0,"artronv71"}
	for i := 0; i < 200 ; i ++ {
		sql := "select process_list.ip,process_list.port,consume_list.process_id,consume_list.serverIndex,consume_list.database from consume_list, process_list  where consume_list.process_id = process_list.process_id   order by process_list.process_id"
		rows, err := c.Fetch( sql  )
		if err != nil{
			fmt.Println( err )
			return
		}
		for rows.Next() {
			//var binlog_id = 0
			//var serverIndex = 0
			//var database = ""
			var process_id = 0
			var ip = ""
			var port= 0
			var serverIndex = 0
			var database = ""
			rows.Scan( &ip, &port, &process_id, &serverIndex, &database)
			fmt.Println( "i ==== ", i, ip, port, process_id, serverIndex, database )
		}
	}
}

