package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"awesomeProject/service"
)

type HeroHandler struct {
	service service.HeroService
}

func NewHeroHandler(service service.HeroService) *HeroHandler {
	return &HeroHandler{
		service: service,
	}
}

func (h *HeroHandler) GetHeroPowerStats(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// 🔐 pegar token do header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header é obrigatório", http.StatusUnauthorized)
		return
	}

	// espera formato: Bearer TOKEN
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Formato inválido de Authorization", http.StatusUnauthorized)
		return
	}

	token := parts[1]

	// 🔍 pegar hero_id
	heroID := r.URL.Query().Get("hero_id")
	if heroID == "" {
		http.Error(w, "Parâmetro 'hero_id' é obrigatório", http.StatusBadRequest)
		return
	}

	// 🔥 agora passa o token pra frente
	result, err := h.service.GetHeroPowerStats(heroID, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
