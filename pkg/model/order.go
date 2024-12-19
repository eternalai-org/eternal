package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus int64

const (
	OrderStatusActive OrderStatus = iota + 1
	OrderStatusPending
	OrderStatusListing
)

type Order struct {
	Model `bson:",inline" json:",inline"`

	UserId        primitive.ObjectID `json:"user_id" bson:"user_id"`
	UserAddress   string             `json:"user_address" bson:"user_address"`
	WebsiteUrl    string             `json:"website_url" bson:"website_url"`
	Email         string             `json:"email" bson:"email"`
	YearToStore   int                `json:"year_to_store" bson:"year_to_store"`
	MonthToStore  int                `json:"month_to_store" bson:"month_to_store"`
	Status        OrderStatus        `json:"status" bson:"status"`
	TxHash        string             `json:"tx_hash" bson:"tx_hash"`
	NftId         string             `json:"nft_id" bson:"nft_id"`
	Metadata      []*Metadata        `json:"metadata" bson:"metadata"`
	Description   string             `json:"description" bson:"description"`
	Title         string             `json:"title" bson:"title"`
	Favicon       string             `json:"favicon" bson:"favicon"`
	ThumbnailUrl  string             `json:"thumbnail_url" bson:"thumbnail_url"`
	ConfigCrontab string             `json:"config_crontab" bson:"config_crontab"`
	EventStatus   Event              `json:"event_status" bson:"event_status"`
	ListingId     string             `json:"listing_id" bson:"listing_id"`

	DepositAddress string  `json:"deposit_address" bson:"deposit_address"`
	DepositPrivKey string  `json:"-" bson:"deposit_priv_key"`
	DepositTxHash  string  `json:"deposit_tx_hash" bson:"deposit_tx_hash"`
	MintFee        float64 `json:"mint_fee" bson:"mint_fee"`
}

type OrderWithSnapshot struct {
	Order     `bson:",inline" json:",inline"`
	Snapshots []*Snapshot `json:"snapshots" bson:"snapshots"`
}

type Metadata struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

func (m Order) CollectionName() string {
	return "order"
}

type CreateOrderRequest struct {
	WebsiteUrl   string `json:"website_url" query:"website_url"`
	YearToStore  int    `json:"year_to_store" query:"year_to_store"`
	MonthToStore int    `json:"month_to_store" query:"month_to_store"`
	UserAddress  string `json:"user_address" query:"-"`
}

type ListOrderRequest struct {
	Pagination `query:",inline" json:",inline"`
	Keyword    string `json:"keyword" query:"keyword"`

	UserAddress string       `json:"user_address" query:"user_address"`
	Status      *OrderStatus `json:"order" query:"status"`
	EventStatus *Event       `json:"event_status" query:"event_status"`
	NftId       string       `json:"nft_id" query:"nft_id"`
	Ids         []string     `json:"ids" query:"ids"`
	HasSnapshot *bool        `json:"has_snapshot" query:"has_snapshot"`
}

type ListOrderResponse struct {
	Orders       []*OrderResp `json:"orders"`
	TotalRecords int64        `json:"total_records"`
}

type OrderResp struct {
	Order         `json:",inline" bson:"order,inline"`
	SnapshotCount int64 `json:"snapshot_count" bson:"snapshot_count"`
}

type ListOrderByDomainResponse struct {
	Orders       []*ListOrderWithSnapshotByDomain `json:"orders"`
	TotalRecords int64                            `json:"total_records"`
}

type ListOrderWithSnapshotByDomain struct {
	WebsiteUrl         string               `json:"website_url" bson:"website_url"`
	OrderWithSnapshots []*OrderWithSnapshot `json:"order_with_snapshots" bson:"order_with_snapshots"`
	TotalRecords       int64                `json:"total_records" bson:"total_records"`
}

type ListOrderWithSnapshotResponse struct {
	OrderWithSnapshots []*OrderWithSnapshot `json:"order_with_snapshots"`
	TotalRecords       int64                `json:"total_records"`
}
