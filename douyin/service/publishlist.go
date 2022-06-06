package service

import (
	"douyin/model"
	"log"

	"github.com/jinzhu/gorm"
)

func GetPublishList(db *gorm.DB) (*[30]model.Videofeed, error) {
	list, err := model.PublishList(db)
	if err != nil {
		log.Println(err)
	}
	return list, nil
}
