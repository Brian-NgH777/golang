package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bodycategories struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang
	BcName        string             `json:"bcName" bson:"bcName,omitempty"`
	BcThumbnail   string             `json:"bcThumbnail" bson:"bcThumbnail"`
	BcDescription string             `json:"bcDescription" bson:"bcDescription"`
}
