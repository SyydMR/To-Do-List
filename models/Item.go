package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Item struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status" gorm:"default:'To Do'"`
    UserID      int64  `json:"user_id"`
}

func (u *User) AddItem(item *Item) error {
    if item == nil {
        return errors.New("item cannot be nil")
    }

    item.UserID = int64(u.ID)
    if err := db.Save(item).Error; err != nil {
        log.Printf("Error saving item for user %d: %v", u.ID, err)
        return fmt.Errorf("error saving item for user %d: %v", u.ID, err)
    }

    return nil
}

func GetUserAllItem(id int64) []Item {
	var items []Item
	db.Where("user_id = ?", id).Find(&items)
	return items
}

func GetAllItem() []Item {
	var items []Item
	db.Find(&items)
	return items
}

func GetItemById(id int64) (*Item, *gorm.DB) {
	var getItem Item
	db := db.Where("ID=?", id).Find(&getItem)
	return &getItem, db
}

func RemoveItem(ID int64) error {
	var item Item
	if err := db.Where("ID=?", ID).Delete(&item).Error; err != nil {
		return err
	}
	return nil
}