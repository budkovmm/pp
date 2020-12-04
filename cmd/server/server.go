package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"pp/api/server"
	"pp/api/utils"
)

func main()  {
	if err := serve(); err != nil {
		log.Fatal(err)
	}
}

func serve() error{
	utils.LoadEnvs()
	db := utils.GetDbConnection()
	r, err := server.NewServer(db,"/api/v1")
	if err != nil {
		return err
	}
	h := http.Server{
		Addr:              ":8080",
		Handler:           r,
	}
	if err = h.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
