package dto

type Manufacturer struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    ManufacturerId int `json:"manufacturerId" validate:"required"`
    Name string `json:"name" validate:"required"`
    Address string `json:"address" validate:"required"`
    ContactNumber string `json:"contactNumber" validate:"required"`
    Email string `json:"email" validate:"required"`
    Website string `json:"website" validate:"required"`
    }

