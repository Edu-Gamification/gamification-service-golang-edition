package repository

import (
	"GamificationEducation/internal/domain"
	errors2 "GamificationEducation/internal/errors"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) FindById(id int64) (domain.User, error) {
	var user domain.User

	if err := userRepository.db.QueryRow("select * from users where id = $1", id).Scan(&user.Id, &user.Name, &user.Surname,
		&user.Patronymic, &user.Password, &user.Email, &user.Active, &user.ClanPoints, &user.Coins, &user.Clan); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors2.ErrNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (userRepository *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User

	rows, err := userRepository.db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name, &user.Surname,
			&user.Patronymic, &user.Password, &user.Email, &user.Active, &user.ClanPoints, &user.Coins, &user.Clan)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
