package service

import (
	"avito-third/internal/repository"
	_"context"
	"time"
)
func Start(storage repository.Repository) error {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	if err := storage.DeleteExpiredUser(); err != nil {
		return err
	}

	for {
		select {
		case <-ticker.C:
			if err := storage.DeleteExpiredUser(); err != nil {
				return err
			}
		}
	}
}
