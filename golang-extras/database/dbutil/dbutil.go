package dbutil

import (
	"fmt"
	"log"
)

type TestData struct {
	ID   int32
	Name string
}

type City struct {
	ID          int64
	Name        string
	CountryCode string
}

const (
	Driver   = "mysql"
	Username = "root"
	Paswword = "ashokdey"
	Hostname = "127.0.0.1:3306"
	Database = "world_x"
)

var DBSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", Username, Paswword, Hostname, Database)

func Check(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
