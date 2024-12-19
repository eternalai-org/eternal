package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event string

const (
	EventMint               Event = "mint"
	EventMakeOffer          Event = "make_offer"
	EventCancelOffer        Event = "cancel_offer"
	EventAcceptOffer        Event = "accept_offer"
	EventUpdateOfferPrice   Event = "update_offer_price"
	EventUpdateListingPrice Event = "update_listing_price"
	EventListing            Event = "listing"
	EventTransfer           Event = "transfer"

	EventListingDone   Event = "listed"
	EventMintDone      Event = "mint_done"
	EventPurchase      Event = "purchase"
	EventCancelListing Event = "cancel_listing"
)

type Activity struct {
	Model       `bson:",inline" json:",inline"`
	UserAddress string             `json:"user_address" bson:"user_address"`
	SnapshotId  primitive.ObjectID `json:"snapshot_id" bson:"snapshot_id"`
	OrderId     primitive.ObjectID `json:"order_id" bson:"order_id"`

	Price   string `json:"price" bson:"price"`
	OfferID string `json:"offer_id" bson:"offer_id"`
	Event   Event  `json:"event" bson:"event"`

	BuyerAddress  string `json:"buyer_address" bson:"buyer_address"`
	SellerAddress string `json:"seller_address" bson:"seller_address"`

	Deadline    int64  `json:"deadline" bson:"deadline"`
	AcceptAsset string `json:"accept_asset" bson:"accept_asset"`
	IsInvalid   bool   `json:"is_invalid" bson:"is_invalid"`
	AtBlock     uint64 `json:"at_block" bson:"at_block"`
}

func (m Activity) CollectionName() string {
	return "activity"
}

type ActivityResponse struct {
	Activity `bson:",inline" json:",inline"`
	Order    *Order    `bson:"order" json:"order"`
	Snapshot *Snapshot `json:"snapshot" bson:"snapshot"`
}

type ListActivityRequest struct {
	Pagination  `query:",inline" json:",inline"`
	SnapshotId  string `json:"snapshot_id" query:"snapshot_id"`
	OrderId     string `json:"order_id" query:"order_id"`
	UserAddress string `json:"user_address" query:"user_address"`
	FetchOrder  *bool  `json:"fetch_order" query:"fetch_order"`
}

type ListActivityResponse struct {
	Activities   []*ActivityResponse `json:"activities"`
	TotalRecords int64               `json:"total_records"`
}
