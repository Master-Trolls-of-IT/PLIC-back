package service

import "gaia-api/domain/entity"

type LoggerService struct {
	UserLogs *[]entity.UserLogs
}

func NewCustomLogger() *LoggerService {
	userLogs := make([]entity.UserLogs, 0)
	return &LoggerService{UserLogs: &userLogs}
}
