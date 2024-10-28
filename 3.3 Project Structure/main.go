package main

import (
	"contactApp/app"
	"contactApp/repository"
	"contactApp/utils/log"
	"sync"
)

func main() {
	//logger
	log := log.GetLogger()
	db := app.NewDBConnection(log)
	if db == nil {
		log.Error("DB connection failed")
	}
	defer func() {
		db.Close()
		log.Error("Db closed")
	}()
	var wg sync.WaitGroup
	repository := repository.NewGormRepositoryMySQL()

}
