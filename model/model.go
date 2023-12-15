package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Garage struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson: "_id,omitempty"`
	OwnerName string             `json:"ownername,omitempty"`
	ModalName string             `json:"modalname,omitempty"`
	CarNumber string             `json:"carnumber,omitempty"`
	Repair    bool               `json:"repair,omitempty"`
}
