package dto

type GroceryItems struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    GroceryitemsId string `json:"groceryitemsId" validate:"required"`
    Name string `json:"name" validate:"required"`
    Category string `json:"category" validate:"required"`
    Price float64 `json:"price" validate:"required"`
    Quantity int `json:"quantity" validate:"required"`
    }

