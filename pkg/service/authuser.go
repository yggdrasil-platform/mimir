package service

import (
  "github.com/kieranroneill/mimir/pkg/logger"
  "github.com/kieranroneill/mimir/pkg/model"
  "gorm.io/gorm"
)

type AuthUserService struct {
	database *gorm.DB
}

func(s *AuthUserService) Create(m model.AuthUser) (*model.AuthUser, error) {
	result := s.database.Create(&m)
	if result.Error != nil {
		return nil, result.Error
	}

	return &m, nil
}

func(s *AuthUserService) GetById(id int) (*model.AuthUser, error) {
	var m model.AuthUser

	result := s.database.First(&m, id)
	if result.Error != nil {
		logger.Error.Printf(result.Error.Error())
		return nil, nil
	}

	return &m, nil
}

func(s *AuthUserService) GetByUserId(uid int) *model.AuthUser {
  var m model.AuthUser

  result := s.database.Where("user_id = ?", uid).First(&m)
  if result.Error != nil {
    logger.Error.Printf(result.Error.Error())
    return nil
  }

  return &m
}

func NewAuthUserService(db *gorm.DB) *AuthUserService {
	return &AuthUserService{database: db}
}
