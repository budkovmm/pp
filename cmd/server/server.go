package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"pp/pkg/utils"
	"pp/pkg/server"
)

func main()  {
	if err := serve(); err != nil {
		log.Fatal(err)
	}
}

func serve() error{
	utils.LoadEnvs()
	apiPort := os.Getenv(utils.HttpApiPort)
	log.Printf("API PORT IS %s", apiPort)
	db := utils.GetPgDbConnection()
	r, err := server.NewServer("/api/v1", db)
	if err != nil {
		return err
	}
	h := http.Server{
		Addr:    apiPort,
		Handler: r,
	}
	if err = h.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
