package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lead struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email       string             `bson:"email" json:"email"`
	Phone       string             `bson:"phone" json:"phone"`
	ClientName  string             `bson:"client_name" json:"client_name"`
	ProductName string             `bson:"product_name" json:"product_name"`
	LeadSource  string             `bson:"lead_source" json:"lead_source"`
	LeadStatus  string             `bson:"lead_status" json:"lead_status"`
	LeadOwnerID primitive.ObjectID `bson:"lead_owner_id" json:"lead_owner_id"`
	TotalPrice  float64            `bson:"total_price" json:"total_price"`
	Discount    float64            `bson:"discount" json:"discount"`
	Notes       string             `bson:"notes" json:"notes"`

	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type CreateLeadRequest struct {
	Email               string  `json:"email" binding:"required,email"`
	Phone               string  `json:"phone" binding:"required"`
	ClientName          string  `json:"client_name" binding:"required"`
	ProductName         string  `json:"product_name" binding:"required"`
	LeadSource          string  `json:"lead_source" binding:"required"`
	TotalPrice          float64 `json:"total_price" binding:"required"`
	Discount            float64 `json:"discount"`
	Notes               string  `json:"notes"`
	AuthenticationToken string  `json:"authentication_token" binding:"required"`
}
