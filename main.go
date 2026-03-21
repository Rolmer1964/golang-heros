package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Isso é o token da superHero
const (
	myToken = "f82cfc46bc540e8d69d147443076dcb2"
)

type MarvelResponse struct {
	Response     string `json:"response"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

func getMarvelCharacters(w http.ResponseWriter, r *http.Request) {

	var heroId = r.URL.Query().Get("hero_id")
	if heroId == "" {
		http.Error(w, "Parâmetro 'id do herói' é obrigatório", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf(
		"https://superheroapi.com/api/%s/%s/powerstats",
		myToken,
		heroId,
	)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erro ao acessar a API Super Hero", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro ao ler resposta da Marvel", http.StatusInternalServerError)
		return
	}

	var marvelResp MarvelResponse
	if err := json.Unmarshal(body, &marvelResp); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(marvelResp)
}

func main() {
	http.HandleFunc("/personagens", getMarvelCharacters)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor rodando em http://localhost:%s/personagens", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
