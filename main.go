package main

import (
	"github.com/gactocat/snowshoe/api"
	"github.com/gactocat/snowshoe/config"
)

func main() {
	config.Init()
	api.Start()
}
