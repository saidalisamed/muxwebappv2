package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/saidalisamed/muxwebappv2/app"
)

func main() {
	app.Run()
}
