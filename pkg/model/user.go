package model

import (
	"time"
)

type User struct {
	Model `bson:",inline" json:",inline"`

	UserAddress    string    `json:"usre_address" bson:"user_address"`
	VerifierId     string    `json:"verifier_id" bson:"verifier_id"`
	Name           string    `json:"name" bson:"name"`
	Email          string    `json:"email" bson:"email"`
	ProfileImage   string    `json:"profile_image" bson:"profile_image"`
	VerifiedFollow bool      `json:"verified_follow" bson:"verified_follow"`
	LastSeen       time.Time `json:"last_seen" bson:"last_seen"`

	AuthChallenge string `json:"-" bson:"auth_challenge"`
}

func (User) CollectionName() string {
	return "user"
}

type ListUserRequest struct {
	Pagination `query:",inline" json:",inline"`

	UserAddress string `json:"user_address" query:"user_address"`
}

type ListUserResponse struct {
	Users        []*User `json:"users"`
	TotalRecords int64   `json:"total_records"`
}
