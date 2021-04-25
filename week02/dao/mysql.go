package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	USER_NAME = "root"
	PASS_WORD = "root"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "test"
	CHARSET   = "utf8"
)

var (
	MysqlDB    *sql.DB
	MysqlDBErr error
)

func init() {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	MysqlDB, MysqlDBErr = sql.Open("mysql", dbURL)
	if MysqlDBErr != nil {
		log.Println("dbURL: " + dbURL)
		panic("MYSQL配置不正确: " + MysqlDBErr.Error())
	}

	MysqlDB.SetMaxOpenConns(100)
	MysqlDB.SetMaxIdleConns(20)
	MysqlDB.SetConnMaxLifetime(100 * time.Second)

	if MysqlDBErr = MysqlDB.Ping(); nil != MysqlDBErr {
		panic("连接mysql失败: " + MysqlDBErr.Error())
	}
}
