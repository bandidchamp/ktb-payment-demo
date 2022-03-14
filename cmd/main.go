package main

import (
	"ktb-payment/configs"
	"ktb-payment/internal/server"
	"os"
)

func main() {
	err := configs.ReadEnvironment()
	if err != nil {
		panic(err)
	}

	config := &configs.Config{}
	config.Posrgress.Host = os.Getenv("POSTGRESS_HOST")
	config.Posrgress.Port = os.Getenv("POSTGRESS_DBNAME")
	config.Posrgress.Username = os.Getenv("POSTGRESS_USERNAME")
	config.Posrgress.Password = os.Getenv("POSTGRESS_PASSWORD")
	config.Posrgress.DBname = os.Getenv("POSTGRESS_DBNAME")

	config.Fiber.Host = os.Getenv("APP_SERVER_HOST")
	config.Fiber.Port = os.Getenv("APP_SERVER_PORT")
	config.Fiber.Timeout = os.Getenv("APP_SERVER_TIMEOUT")

	Server := server.NewServer(config)
	serr := Server.StartServer()
	if serr != nil {
		panic("Server can't Start.")
	}

}
