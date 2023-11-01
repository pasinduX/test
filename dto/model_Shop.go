package dto

type Shop struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    ShopId string `json:"shopId" validate:"required"`
    Name string `json:"name" validate:"required"`
    Address string `json:"address" validate:"required"`
    PhoneNumber string `json:"phoneNumber" validate:"required"`
    Email string `json:"email" validate:"required"`
    Website string `json:"website" validate:"required"`
    CreatedAt string `json:"createdAt" validate:"required"`
    }

