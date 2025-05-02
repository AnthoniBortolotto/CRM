package repositories

import (
	"context"
	"crm-go/config"
	"crm-go/modules/lead/v1/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	result, err := r.collection.InsertOne(context.Background(), lead)
	if err != nil {
		println("Error inserting lead:", err)
		return nil, err
	}
	lead.ID = result.InsertedID.(primitive.ObjectID)
	return lead, nil
}

func (r *LeadRepository) GetLeadByID(id primitive.ObjectID) (*models.Lead, error) {
	var lead models.Lead
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&lead)
	if err != nil {
		println("Error finding lead:", err)
		return nil, err
	}
	return &lead, nil
}
