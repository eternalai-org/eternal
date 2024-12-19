package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSession struct {
	Model `bson:",inline" json:",inline"`

	UserId      primitive.ObjectID `json:"user_id" bson:"user_id"`
	UserAddress string             `json:"user_address" bson:"user_address"`
	IP          string             `json:"ip" bson:"ip"`
	SessionId   string             `json:"session_id" bson:"session_id"`
	AuthToken   string             `json:"auth_token" bson:"auth_token"`
	Exp         int64              `json:"exp" bson:"exp"`
	Metadata    string             `json:"metadata" bson:"metadata"`
	IssuedAt    time.Time          `json:"iat" bson:"iat"`
}

func (UserSession) CollectionName() string {
	return "user_session"
}
