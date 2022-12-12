package repositories

import (
	"wayshub-server/models"

	"gorm.io/gorm"
)

type SubscribeRepository interface {
	Subscription() ([]models.Subscribe, error)
	GetSubscription(ID int) (models.Subscribe, error)
	Subscribe(subscribe models.Subscribe) (models.Subscribe, error)
	Unsubscribe(unsubscribe models.Subscribe) (models.Subscribe, error)
}

func RepositorySubscribe(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetSubscription(ID int) (models.Subscribe, error) {
	var subscribe models.Subscribe
	err := r.db.Preload("Channel").First(&subscribe, ID).Error
  
	return subscribe, err
}

func (r *repository) Subscription() ([]models.Subscribe, error) {
	var subscription []models.Subscribe
	err := r.db.Preload("Channel").Find(&subscription).Error

	return subscription, err
}

func (r *repository) Subscribe(subscribe models.Subscribe) (models.Subscribe, error) {
	err := r.db.Preload("Channel").Create(&subscribe).Error

	return subscribe, err
}

func (r *repository) Unsubscribe(unsubscribe models.Subscribe) (models.Subscribe, error) {
	err := r.db.Delete(&unsubscribe).Error

	return unsubscribe, err
}