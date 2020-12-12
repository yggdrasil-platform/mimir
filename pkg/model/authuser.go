package model

import "time"

type AuthUser struct {
  ID int `json:"id" gorm:"primarykey"`
  Password string `json:"password"`
  UserID int `json:"userId" gorm:"unique"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}
