package dto

type Certificates struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    CertificatesId string `json:"certificatesId" validate:"required"`
    Name string `json:"name" validate:"required"`
    IssueDate string `json:"issueDate" validate:"required"`
    ExpireDate string `json:"expireDate" validate:"required"`
    Issuer string `json:"issuer" validate:"required"`
    CertificateNumber string `json:"certificateNumber" validate:"required"`
    Description string `json:"description" validate:"required"`
    File string `json:"file" validate:"required"`
    }

