package main

import (
	"api-template-basic-golang/internal/app/container"
	"api-template-basic-golang/internal/config"
	"api-template-basic-golang/internal/infra/db"
	"api-template-basic-golang/internal/presentation/http/router"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := config.Load()

	// Conecta no banco de dados
	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados: %v", err)
	}
	defer conn.Close()

	authContainer := container.NewAuthContainer(conn, os.Getenv("JWT_SECRET"))
	userContainer := container.NewUserContainer(conn)

	routes := router.SetupRouter(authContainer, userContainer)

	log.Printf("Servidor rodando em http://localhost:%s", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, routes)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
