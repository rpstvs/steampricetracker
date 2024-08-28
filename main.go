package main

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/rpstvs/steamprice-api/internals/api"
)

func main() {
	godotenv.Load(".env")

	steamClient := api.NewClient(30 * time.Second)

	//server := server.ReturnServer()

	//c := cron.New()
	/*
		c.AddFunc("0 0 * * * *", func() {
			fmt.Println("starting job")

		})
		c.Start()
	*/
	steamClient.UpdateDB(0)
	//server.Start()
}
