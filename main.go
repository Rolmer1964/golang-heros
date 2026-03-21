package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"awesomeProject/client"
	"awesomeProject/handler"
	"awesomeProject/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	apiClient := client.NewSuperheroClient(httpClient)
	heroService := service.NewHeroService(apiClient)
	heroHandler := handler.NewHeroHandler(heroService)

	mux := http.NewServeMux()
	mux.HandleFunc("/personagens", heroHandler.GetHeroPowerStats)

	log.Printf("Servidor rodando em http://localhost:%s/personagens", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("erro ao iniciar servidor: %v", err)
	}
}
