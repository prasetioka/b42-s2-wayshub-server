package models

type Channel struct {
	ID          int         `json:"id"`
	Email       string      `json:"email" gorm:"type: varchar(255)"`
	Password    string      `json:"password" gorm:"type: varchar(255)"`
	ChannelName string      `json:"channelName" gorm:"type: varchar(255)"`
	Description string      `json:"description" gorm:"type: varchar(255)"`
	Cover       string      `json:"cover" gorm:"type: varchar(255)"`
	Photo       string      `json:"photo" gorm:"type: varchar(255)"`
	Subscribe   []Subscribe `json:"subscribe" gorm:"many2many:subscription"`
}

type ChannelVideoResponse struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	ChannelName string `json:"channelName"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Photo       string `json:"photo"`
}

func (ChannelVideoResponse) TableName() string {
	return "channels"
}
