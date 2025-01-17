package main

import (
	"log"

	"github.com/Timber868/french-league/cmd/api"
	"github.com/Timber868/french-league/config"
	"github.com/Timber868/french-league/db"
	"github.com/go-sql-driver/mysql"
)

// Server info
const serverAddress string = ":8080"

func main() {
	db, _ := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp", //Network protocol, tcp is standard for IP or hostname connections
		AllowNativePasswords: true,  //Allows the use of MySQL’s native password authentication method. This is often needed with more recent versions of MySQL.
		ParseTime:            true,  //Date columns wont be left as string but as Time.time instead
	})

	//Initialize the api server without our own custom method
	server := api.NewApiServer(serverAddress, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
