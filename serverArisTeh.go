package main

import (
	"fmt"
	egui "пппппп/server/externalserver"

	"github.com/jmoiron/sqlx"
	_ "github.com/nakagami/firebirdsql"
)

// Config структура подключения к БД
type Config struct {
	DB     *sqlx.DB
	Pathdb string
}

// Parametrs переменная хранящая данные о БД
var Parametrs Config

// OpenDB подключение к БД
func OpenDB() error {
	Parametrs.Pathdb = "C:/obmen/FIRST.fdb"
	db, err := sqlx.Open("firebirdsql", "sysdba:masterkey@localhost:3050/C:/obmen/FIRST.fdb?auth_plugin_name=Legacy_Auth&wire_auth=true&column_name_to_lower=false")
	Parametrs.DB = db

	egui.RegFunc("login", egui.Login)
	return err
}

func parse() {
	for {
		select {
		case res, ok := <-egui.СН:
			if !ok {
				break
			}
			println(" 32 " + res.Message.Action)
			switch res.Message.Action {
			case "login":

				println(" 33 " + res.Message.Action)
				res.Client.Sendout(egui.Login(res.Message.Parameters))

			case "runproc":
				go egui.Runproc(res)
			case "logincheck":
				go egui.Runproc(res)
			case "getform":
				// egui.Login(res.Message.Parameters)
				// res.Client.Sendout()
				go egui.Runproc(res)
			case "getdata":
				go egui.Runproc(res)
			case "formclose":

			case "sendoutAndReturn":
				// запишим в канал полученное значение от клиента
				egui.VСН <- res
			case "endapp":
				// клиент закрыл приложение
				egui.Manager.Unregister <- res.Client
			}
		}
	}
}

func main() {
	err := OpenDB()
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	go parse()
	egui.Init(":8080")
}
