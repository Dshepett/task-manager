package main

import (
	"github.com/dshepett/task-manager/internal/config"
	"log"

	"github.com/dshepett/task-manager/internal/app"
)

func main() {
	conf := config.New()
	log.Println(conf)
	a := app.New(conf)
	a.Run()
}
