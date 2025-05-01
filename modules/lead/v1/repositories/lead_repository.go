package repositories

import (
	"context"
	"crm-go/config"
	"crm-go/modules/lead/v1/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LeadRepository struct {
	collection *mongo.Collection
}

func NewLeadRepository() *LeadRepository {
	db := config.GetDB()
	return &LeadRepository{
		collection: db.Collection("leads"),
	}
}

func (r *LeadRepository) CreateLead(lead *models.Lead) (*models.Lead, error) {
	lead.CreatedAt = time.Now()
	lead.UpdatedAt = time.Now()
	println("Inserted lead:", lead)
	println("context to be inserted:", context.Background())
	result, err := r.collection.InsertOne(context.Background(), lead)
	println("Inserted lead:")
	if err != nil {
		println("Error inserting lead:", err)
		return nil, err
	}

	println("Lead inserted with ID:", result)

	lead.ID = result.InsertedID.(primitive.ObjectID)
	return lead, nil

}
