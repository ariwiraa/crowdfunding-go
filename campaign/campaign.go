package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID               uint32   `gorm:"primary_key;auto_increment" json:"id"`
	UserID           uint32   `json:"user_id"`
	Name             string `gorm:"type:varchar(100);not null" json:"name"`
	ShortDescription string `gorm:"type:varchar(100);not null" json:"short_description"`
	Description      string `gorm:"type:longtext;not null" json:"description"`
	Perks            string `gorm:"type:longtext;not null" json:"perks"`
	BackerCount      int    `gorm:"not null" json:"backer_count"`
	GoalAmount       int    `gorm:"not null" json:"goal_amount"`
	CurrentAmount    int    `gorm:"not null" json:"current_amount"`
	Slug             string `json:"slug"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	CampaignImages []CampaignImage
	User 		user.User
}

type CampaignImage struct {
	ID         uint32 `gorm:"primary_key;auto_increment" json:"id"`
	CampaignID uint32 `json:"campaign_id"`
	FileName   string `gorm:"type:varchar(255);not null" json:"name"`
	IsPrimary  int `json:"is_primary"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}