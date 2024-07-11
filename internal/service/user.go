package service

import "GamificationEducation/internal/domain"

type UserRepository interface {
	FindById(id int64) (domain.User, error)
	FindAll() ([]domain.User, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

func (userService *UserService) GetById(id int64) (domain.User, error) {
	user, err := userService.userRepository.FindById(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (userService *UserService) GetAll() ([]domain.User, error) {
	users, err := userService.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
