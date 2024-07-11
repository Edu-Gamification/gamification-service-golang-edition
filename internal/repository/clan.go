package repository

import (
	"GamificationEducation/internal/domain"
	errors2 "GamificationEducation/internal/errors"
	"database/sql"
	"errors"
)

type ClanRepository struct {
	db *sql.DB
}

func NewClanRepository(db *sql.DB) *ClanRepository {
	return &ClanRepository{db: db}
}

func (clanRepository *ClanRepository) FindAll() ([]domain.Clan, error) {
	var clans []domain.Clan

	rows, err := clanRepository.db.Query("select * from clans")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var clan domain.Clan
		err := rows.Scan(&clan.Id, &clan.Name, &clan.PointsAmount)
		if err != nil {
			return nil, err
		}
		members, err := clanRepository.FindClanMembers(clan.Id)
		if err != nil {
			return nil, err
		}
		clan.Members = members
		clans = append(clans, clan)
	}
	return clans, nil
}

func (clanRepository *ClanRepository) FindByName(name string) (domain.Clan, error) {
	var clan domain.Clan
	if err := clanRepository.db.QueryRow("select * from clans where name = $1", name).Scan(&clan.Id,
		&clan.Name, &clan.PointsAmount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Clan{}, errors2.ErrNotFound
		}
		return domain.Clan{}, err
	}
	members, err := clanRepository.FindClanMembers(clan.Id)
	if err != nil {
		return domain.Clan{}, err
	}
	clan.Members = members
	return clan, nil
}

func (clanRepository *ClanRepository) FindById(id int64) (domain.Clan, error) {
	var clan domain.Clan
	if err := clanRepository.db.QueryRow("select * from clans where id = $1", id).Scan(&clan.Id, &clan.Name,
		&clan.PointsAmount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Clan{}, errors2.ErrNotFound
		}
		return domain.Clan{}, err
	}
	members, err := clanRepository.FindClanMembers(clan.Id)
	if err != nil {
		return domain.Clan{}, err
	}
	clan.Members = members
	return clan, nil
}

func (clanRepository *ClanRepository) FindClanMembers(clanId int64) ([]domain.User, error) {
	var members []domain.User
	rows, err := clanRepository.db.Query("select * from users where clan = $1", clanId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var member domain.User
		err := rows.Scan(&member.Id, &member.Name, &member.Surname,
			&member.Patronymic, &member.Password, &member.Email, &member.Active, &member.ClanPoints, &member.Coins, &member.Clan)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}
