package repositories

import (
	"wayshub-server/models"

	"gorm.io/gorm"
)

type ChannelRepository interface {
	FindChannels() ([]models.Channel, error)
	GetChannel(ID int) (models.Channel, error)
	UpdateChannel(channel models.Channel) (models.Channel, error)
	DeleteChannel(channel models.Channel) (models.Channel, error)
}

func RepositoryChannel(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindChannels() ([]models.Channel, error) {
	var channels []models.Channel
	err := r.db.Preload("Subscribe").Preload("Video").Find(&channels).Error

	return channels, err
}

func (r *repository) GetChannel(ID int) (models.Channel, error) {
	var channel models.Channel
	err := r.db.Preload("Subscribe").First(&channel, ID).Error

	return channel, err
}

func (r *repository) UpdateChannel(channel models.Channel) (models.Channel, error) {
	err := r.db.Save(&channel).Error

	return channel, err
}

func (r *repository) DeleteChannel(channel models.Channel) (models.Channel, error) {
	err := r.db.Delete(&channel).Error

	return channel, err
}
