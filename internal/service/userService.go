package service

import (
	"avito-third/internal/repository"
	"avito-third/internal/user"
	"time"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CRUDOperation(elem *user.UserSegment) ( error) {
	return s.repo.CRUDOperation(elem)
}

func (s *UserService) GetActive(userId int) ([]string, error){
	return s.repo.GetActive(userId)
}

func (s *UserService) GetReport(period time.Time) ([]user.ReportUsers, error){
	return s.repo.GetReport(period)
}