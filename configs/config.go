package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IConfig interface {
	Get(key string) string
}

type Configs struct {
}

func (cfg *Configs) Get(key string) string {
	return os.Getenv(key)
}

func New(fileName ...string) Configs {
	err := godotenv.Load(fileName...)
	if err != nil {
		log.Panic(err)
	}

	return Configs{}
}
