package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"awesomeProject/model"
)

type SuperheroClient struct {
	httpClient *http.Client
}

func NewSuperheroClient(httpClient *http.Client) *SuperheroClient {
	return &SuperheroClient{
		httpClient: httpClient,
	}
}

func (c *SuperheroClient) GetHeroPowerStats(heroID string, token string) (*model.HeroPowerStatsResponse, error) {

	// ⚠️ API exige token na URL
	url := fmt.Sprintf("https://superheroapi.com/api/%s/%s/powerstats", token, heroID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar request: %w", err)
	}

	// 🔥 header (padrão moderno)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar API externa: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API retornou status %d", resp.StatusCode)
	}

	var result model.HeroPowerStatsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta", err)
	}

	return &result, nil
}
