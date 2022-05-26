package main

import (
	"github.com/Dshepett/task-manager/internal/config"
	"log"

	"github.com/Dshepett/task-manager/internal/app"
)

func main() {
	conf := config.New()
	log.Println(conf)
	a := app.New(conf)
	a.Run()
}
