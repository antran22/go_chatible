package api

import (
	"log"
	"os"
	"strconv"

	"github.com/go-pg/pg"
)

var MessengerSender *MessageSenderPool
var UserFetcher *UserFetcherPool
var DB *pg.DB

func InitAPIs() {
	workerCount, err := strconv.Atoi(os.Getenv("MESSENGER_WORKERS"))
	if err != nil || workerCount < 1 {
		log.Fatalln("MESSENGER_WORKERS should be a integer larger than 0")
	}
	log.Println("Making new message sender pool")
	MessengerSender = NewMessageSenderPool()
	log.Println("Connecting to database")
	ConnectDB()
	log.Println("Creating schema")
	if err := CreateSchema(); err != nil {
		log.Fatalln("Error while creating schema", err)
		return
	}
	log.Println("Making new user fetcher pool")
	UserFetcher = NewUserFetcherPool()
	log.Println("Finish initializing APIs")
}

func TearDownAPIs() {
	MessengerSender.Close()
	UserFetcher.Close()
	if err := DB.Close(); err != nil {
		panic(err)
	}
}
