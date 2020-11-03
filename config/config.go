package config

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/nakagami/firebirdsql"
)

var DB *sqlx.DB
var Dir string
// SetParametrs установка параметров приложения
func SetParametrs() {
	var n int
	db, err := sqlx.Open("firebirdsql", "sysdba:masterkey@localhost:3050/C:/obmen/FIRST.fdb?auth_plugin_name=Legacy_Auth&wire_auth=true&column_name_to_lower=false")
	db.QueryRow("SELECT Count(*) FROM rdb$relations").Scan(&n)
	fmt.Println("Relations count=", n)
	DB = db
	if err != nil {
		fmt.Printf("error : %v", err.Error())
		log.Fatal(err)
	}
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	Dir=dir
}
