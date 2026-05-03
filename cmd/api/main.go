package main

import (
	"fmt"

	"github.com/MatheusMikio/eduGuard_api/config"
	"github.com/MatheusMikio/eduGuard_api/internal/routes"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("Erro ao inicializar config: %s", err.Error())
		return
	}
	db := config.GetDb()
	routes.Init(db)
}
