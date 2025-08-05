package main

import (
	"github.com/Sa-Leonardo/gopportunities/config"
	"github.com/Sa-Leonardo/gopportunities/router"
)

var (
	logger *config.Logger
)

// sempre que u não der um nome para o import, ele vai ser importado pelo ultimo pathName do modulo (gin	// sempre que u não der um nome para o import, ele vai ser importado pelo ultimo pathName do modulo, nesse caso (gin

func main() {

	logger = config.GETLogger("main")

	// Initialize configs

	err := config.Init()

	if err != nil {
		logger.Errorf("Config Inicialization error: %s", err)
		return
	}

	// Initialize raouter
	router.Initialize()
}
