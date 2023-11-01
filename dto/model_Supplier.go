package dto

type Supplier struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    SupplierId string `json:"supplierId" validate:"required"`
    Address string `json:"address" validate:"required"`
    Phone string `json:"phone" validate:"required"`
    Email string `json:"email" validate:"required"`
    SupplierName string `json:"supplierName" validate:"required"`
    }

