package models

type Subscribe struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	ChannelID int                  `json:"channel_id"`
	Channel   ChannelVideoResponse `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SubscribeResponse struct {
	ID        int                  `json:"id"`
	ChannelID int                  `json:"channel_id"`
	Channel   ChannelVideoResponse `json:"-"`
}

type SubscribeChannelResponse struct {
	ID        int                  `json:"id"`
	ChannelID int                  `json:"channel_id"`
	Channel   ChannelVideoResponse `json:"-"`
}

func (SubscribeResponse) TableName() string {
	return "subscribers"
}

func (SubscribeChannelResponse) TableName() string {
	return "subscribers"
}