package services

import (
	"errors"
	"final-project-hacktiv8/models"

	"gorm.io/gorm"
)

type SocialMediaService struct {
	DB *gorm.DB
}

func (sms *SocialMediaService) CreateSocialMedia(socmed models.SocialMedia) (models.SocialMedia, error) {
	result := sms.DB.Create(&socmed)

	if result.Error != nil {
		return models.SocialMedia{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.SocialMedia{}, errors.New("error, try again")
	}

	return socmed, nil
}

func (sms *SocialMediaService) GetAllSocialMedia(userId uint) ([]models.SocialMedia, error) {
	var accounts []models.SocialMedia

	if err := sms.DB.Where("user_id = ?", userId).Find(&accounts).Error; err != nil {
		return []models.SocialMedia{}, err
	}

	return accounts, nil
}