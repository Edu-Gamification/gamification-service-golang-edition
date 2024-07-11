package service

import (
	"GamificationEducation/internal/domain"
	"math"
)

type ClanRepository interface {
	FindAll() ([]domain.Clan, error)
	FindByName(name string) (domain.Clan, error)
	FindById(id int64) (domain.Clan, error)
	FindClanMembers(clanId int64) ([]domain.User, error)
}

type ClanService struct {
	clanRepository ClanRepository
}

func NewClanService(repository ClanRepository) *ClanService {
	return &ClanService{clanRepository: repository}
}

func (clanService *ClanService) GetAll() ([]domain.Clan, error) {
	clans, err := clanService.clanRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return clans, nil
}

func (clanService *ClanService) GetByName(name string) (domain.Clan, error) {
	clan, err := clanService.clanRepository.FindByName(name)
	if err != nil {
		return domain.Clan{}, err
	}
	return clan, nil
}

func (clanService *ClanService) GetMembers(id int64) ([]domain.User, error) {
	members, err := clanService.clanRepository.FindClanMembers(id)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (clanService *ClanService) GetById(id int64) (domain.Clan, error) {
	clan, err := clanService.clanRepository.FindById(id)
	if err != nil {
		return domain.Clan{}, err
	}
	return clan, nil
}

// GetMinClan возвращает клан с минимальным кол-вом участников
func (clanService *ClanService) GetMinClan() (domain.Clan, error) {
	minCount := math.MaxInt
	var minClanId int64
	clans, err := clanService.GetAll()
	if err != nil {
		return domain.Clan{}, nil
	}
	for _, clan := range clans {
		members, err := clanService.GetMembers(clan.Id)
		if err != nil {
			return domain.Clan{}, nil
		}
		if len(members) < minCount {
			minCount = len(members)
			minClanId = clan.Id
		}
	}
	res, err := clanService.GetById(minClanId)
	if err != nil {
		return domain.Clan{}, err
	}
	return res, nil
}
