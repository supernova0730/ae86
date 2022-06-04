package storage

import (
	"ae86/internal/model"
	"ae86/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ManagerStorage struct {
	db *gorm.DB
}

func NewManagerStorage(db *gorm.DB) *ManagerStorage {
	return &ManagerStorage{db: db}
}

func (s *ManagerStorage) GetByID(id uint) (result model.Manager, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("ManagerStorage.GetByID failed")
		}
	}()

	err = s.db.
		Model(&model.Manager{}).
		Where("id = ? AND is_deleted = ?", id, false).
		First(&result).
		Error
	return
}

func (s *ManagerStorage) GetByUsername(username string) (result model.Manager, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":    err,
				"username": username,
			}).Error("ManagerStorage.GetByUsername failed")
		}
	}()

	err = s.db.
		Model(&model.Manager{}).
		Where("username = ? AND is_deleted = ?", username, false).
		First(&result).
		Error
	return
}

func (s *ManagerStorage) Create(manager model.Manager) (id uint, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":   err,
				"manager": manager,
			}).Error("ManagerStorage.Create failed")
		}
	}()

	err = s.db.
		Model(&model.Manager{}).
		Create(&manager).
		Error
	id = manager.ID
	return
}

func (s *ManagerStorage) Update(id uint, manager model.Manager) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":   err,
				"id":      id,
				"manager": manager,
			}).Error("ManagerStorage.Update failed")
		}
	}()

	err = s.db.
		Model(&model.Manager{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Updates(&manager).
		Error
	return
}

func (s *ManagerStorage) Delete(id uint) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("ManagerStorage.Delete failed")
		}
	}()

	err = s.db.
		Model(&model.Manager{}).
		Where("id = ?", id).
		UpdateColumn("is_deleted", true).
		Error
	return
}
