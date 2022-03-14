package configs

import "github.com/joho/godotenv"

type Config struct {
	Posrgress    Posrgress
	Fiber        Fiber
	FiberMockKtb FiberMockKtb
}

type Posrgress struct {
	Host     string
	DBname   string
	Username string
	Password string
	Port     string
}

type Fiber struct {
	Host    string
	Port    string
	Timeout string
}

type FiberMockKtb struct {
	Host    string
	Port    string
	Timeout string
}

func ReadEnvironment() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
