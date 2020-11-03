package main

import (
	"fmt"
	"пппппп/server/config"
	egui "пппппп/server/externalserver"
	"пппппп/server/model"

	_ "github.com/nakagami/firebirdsql"
)

// OpenDB подключение к БД
func OpenDB() error {
	config.SetParametrs()
	egui.RegFunc("login", model.Login)
	return nil
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
				res.Client.Sendout(model.Login(res.Message.Parameters))

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
