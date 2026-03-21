package service

import "awesomeProject/model"

type HeroClient interface {
	GetHeroPowerStats(heroID string, token string) (*model.HeroPowerStatsResponse, error)
}

type HeroService interface {
	GetHeroPowerStats(heroID string, token string) (*model.HeroPowerStatsResponse, error)
}

type heroService struct {
	client HeroClient
}

func NewHeroService(client HeroClient) HeroService {
	return &heroService{
		client: client,
	}
}

func (s *heroService) GetHeroPowerStats(heroID string, token string) (*model.HeroPowerStatsResponse, error) {
	return s.client.GetHeroPowerStats(heroID, token)
}
