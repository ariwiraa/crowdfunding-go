package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID         	uint32	`gorm:"primary_key;auto_increment" json:"id"`
	CampaignID 	uint32    `json:"campaign_id"`
	UserID     	uint32    `json:"user_id"`
	Amount     	int    	`gorm:"not null" json:"amount"`
	Status     	string 	`gorm:"type:varchar(100);not null" json:"status"`
	Code       	string 	`gorm:"type:varchar(255);not null" json:"code"`
	PaymentURL 	string 	`gorm:"type:varchar(255);not null" json:"payment_url"`
	User       	user.User
	Campaign  	campaign.Campaign
	CreatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}