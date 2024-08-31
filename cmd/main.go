package main

import (
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/cmd/app"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/config"
)

func main() {
	cfg := config.Load()
	app.StartApp(cfg)
}
