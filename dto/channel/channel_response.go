package channeldto

import "wayshub-server/models"

type ChannelResponse struct {
	ID          int                `json:"id"`
	Email       string             `json:"email"`
	ChannelName string             `json:"channelName"`
	Description string             `json:"description"`
	Cover   string             `json:"cover"`
	Photo       string             `json:"photo"`
	Subscribe   []models.Subscribe `json:"subscribe" gorm:"many2many:subscription"`
}

type DeleteResponse struct {
	ID int `json:"id"`
}
