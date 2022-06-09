package storage

import (
	"ae86/internal/model"
	"ae86/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryStorage struct {
	db *gorm.DB
}

func NewCategoryStorage(db *gorm.DB) *CategoryStorage {
	return &CategoryStorage{db: db}
}

func (s *CategoryStorage) GetByID(id uint) (result model.Category, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("CategoryStorage.GetByID failed")
		}
	}()

	err = s.db.
		Model(&model.Category{}).
		Where("id = ? AND is_deleted = ?", id, false).
		First(&result).
		Error
	return
}

func (s *CategoryStorage) GetAllByStoreID(storeID uint) (result []model.Category, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":   err,
				"storeID": storeID,
			}).Error("CategoryStorage.GetAllByStoreID failed")
		}
	}()

	err = s.db.
		Model(&model.Category{}).
		Where("store_id = ? AND is_deleted = ?", storeID, false).
		Find(&result).
		Error
	return
}

func (s *CategoryStorage) Create(category model.Category) (id uint, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":    err,
				"category": category,
			}).Error("CategoryStorage.Create failed")
		}
	}()

	err = s.db.
		Model(&model.Category{}).
		Create(&category).
		Error
	id = category.ID
	return
}

func (s *CategoryStorage) Update(id uint, category model.Category) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":    err,
				"id":       id,
				"category": category,
			}).Error("CategoryStorage.Update failed")
		}
	}()

	err = s.db.
		Model(&model.Category{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Updates(&category).
		Error
	return
}

func (s *CategoryStorage) Delete(id uint) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("CategoryStorage.Delete failed")
		}
	}()

	err = s.db.
		Model(&model.Category{}).
		Where("id = ?", id).
		UpdateColumn("is_deleted", true).
		Error
	return
}
