package services

import "gaia-api/domain/entities"

type LoggerService struct {
	UserLogs *[]entities.UserLogs
}

func NewCustomLogger() *LoggerService {
	userLogs := make([]entities.UserLogs, 0)
	return &LoggerService{UserLogs: &userLogs}
}
