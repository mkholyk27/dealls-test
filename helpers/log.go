package helpers

import (
	"dating-app/configs"
	"log"
	"os"
	"time"
)

func Logger(endpoint string, fileName string, data string) {
	cfg := configs.New()

	now := time.Now()
	dateNow := now.Format("2006-01-02")

	path := cfg.Get("LOG_FILE") + dateNow
	pathExists(path)

	fullPath := path + "/" + endpoint
	pathExists(fullPath)

	pathFile := fullPath + "/" + fileName
	f, err := os.OpenFile(pathFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(data)
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err := createPath(path)
			if err != nil {
				return false
			}
		}
	}
	return true
}

func createPath(path string) error {
	err := os.Mkdir(path, 0775)
	if err != nil {
		return err
	}
	return nil
}
