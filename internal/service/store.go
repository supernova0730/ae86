package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	"ae86/pkg/logger"
	"github.com/sirupsen/logrus"
)

type StoreService struct {
	storage adapter.StorageContainer
}

func NewStoreService(storage adapter.StorageContainer) *StoreService {
	return &StoreService{storage: storage}
}

func (s *StoreService) GetByID(id uint) (result model.Store, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("StoreService.GetByID failed")
		}
	}()

	return s.storage.Store().GetByID(id)
}
