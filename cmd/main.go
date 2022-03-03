package main

import (
	"CourseWork/internal/apichi"
	"CourseWork/internal/apichi/openapichi"
	"CourseWork/internal/config"
	"CourseWork/internal/database"
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/server"
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	//Generate random seed
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	//Load Config
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	cfg, err := config.LoadConfig(path + "/config/config.env")
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	udf, err := database.NewFullDataFile("shorturl.db", "adminurl.db", "data.db", "ip.db")
	if err != nil {
		log.Fatal(err)
	}
	dbbe := dbbackend.NewDataStorage(udf)
	hs := apichi.NewHandlers(dbbe)
	//rt := apichi.NewRouter(hs)
	rt := openapichi.NewOpenApiRouter(hs)
	srv := server.NewServer(":8000", rt, cfg)

	srv.Start(dbbe)

	log.Println("Hello url shortener!")

	<-ctx.Done()

	srv.Stop()
	cancel()
	udf.Close()

	log.Print("Exit")
}
