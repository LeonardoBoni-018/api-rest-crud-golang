package model

import "time"

type BusinessDomain struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	OwnerID     string    `bson:"owner_id" json:"owner_id"`
	Description string    `bson:"description" json:"description"`
	Timezone    string    `bson:"timezone" json:"timezone"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
