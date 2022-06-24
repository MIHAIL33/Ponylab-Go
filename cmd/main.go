package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"github.com/MIHAIL33/Ponylab-Go/pkg/cache"
	"github.com/MIHAIL33/Ponylab-Go/pkg/handler"
	"github.com/MIHAIL33/Ponylab-Go/pkg/mqtt"
	"github.com/MIHAIL33/Ponylab-Go/pkg/receiver"
	"github.com/MIHAIL33/Ponylab-Go/pkg/repository"
	"github.com/MIHAIL33/Ponylab-Go/pkg/service"
	"github.com/MIHAIL33/Ponylab-Go/server.go"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	//load env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln("error loading env variables: %s", err.Error())
	}

	//connect to postgres
	db, err := dbConnect()
	if err != nil {
		log.Fatalln(err.Error())
	}

	repos := repository.NewRepository(db)

	cache := cache.NewCache()

	service := service.NewService(repos, cache)
	handler.NewHandler(service)
	//loading data to cache
	err = service.AddAllInCache()
	if err != nil {
		log.Println("failed loading data to cache from postgres")
	}

	//connect to mqtt
	mqttClient := mqtt.NewMQTTClient(os.Getenv("Address"), os.Getenv("PORT"), os.Getenv("ClientID"))
	err = mqttClient.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer mqttClient.Disconnect(100)

	//listen devices-state
	deviceStateChannel := make(chan mqtt.MQTTPayload)
	go mqttClient.Listen(os.Getenv("TOPIC_STATE"), deviceStateChannel)

	receiver := receiver.NewReceiver(service, deviceStateChannel)
	go receiver.GetMQTTPayload()

	//run http server
	runServer()
}

func dbConnect() (*sql.DB, error) {
	var wait time.Duration = 1
	var count int = 0

	for {
		if count == 10 {
			return nil, errors.New("no connection to postgres")
		}
		count++
		db, err := repository.NewPostgresDB(repository.Config{
			Host: "localhost",
			Port: "5432",
			Username: "postgres",
			Password: "postgres",
			DBName: "devices",
			SSLMode: "disable",
		})
		if err != nil {
			log.Println("No connection to Postgres, repeat...")
			wait *= 2
			time.Sleep(wait * time.Second)
			continue
		}
		log.Println("Connected to postgres")
		return db, nil
	}
	
}

func runServer() {
	srv := new(server.Server)
	err := srv.Run("8000")
	if err != nil {
		log.Fatalln(err.Error())
	}
}