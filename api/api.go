package api

import (
	"log"
	"os"
	"strconv"

	"github.com/go-pg/pg"
)

var MessengerSender *MessengerSenderPool
var UserFetcher *UserFetcherPool
var DB *pg.DB

func InitAPIs() {
	workerCount, err := strconv.Atoi(os.Getenv("MESSENGER_WORKERS"))
	if err != nil || workerCount < 1 {
		log.Fatalln("MESSENGER_WORKERS should be a integer larger than 0")
	}
	MessengerSender = NewWorkerPool(workerCount)
	ConnectDB()
	if err := CreateSchema(); err != nil {
		log.Fatalln("Error while creating schema", err)
		return
	}
	UserFetcher = NewUserFetcherPool()
	log.Println("Finish initializing APIs")
}
