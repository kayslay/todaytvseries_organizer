package main

import (
	"log"

	"github.com/kayslay/todaytvseries_organizer/config"
	"github.com/kayslay/todaytvseries_organizer/helpers"
)

func main() {

	err := helpers.Start(*config.InitConfig())
	if err != nil {
		log.Println(err)
	}
}
