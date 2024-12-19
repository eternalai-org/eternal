package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Snapshot struct {
	Model       `bson:",inline" json:",inline"`
	OrderId     primitive.ObjectID `json:"order_id" bson:"order_id"`
	UserAddress string             `json:"user_address" bson:"user_address"`

	Url string `json:"url" bson:"url"`

	FileName string `json:"file_name" bson:"file_name"`
	FilePath string `json:"file_path" bson:"file_path"`
	FileSize int64  `json:"file_size" bson:"file_size"`

	CrawlingAt  time.Time `json:"crawling_at" bson:"crawling_at"`
	FilecoinCid string    `json:"filecoin_cid" bson:"filecoin_cid"`

	PushToChain    bool           `json:"push_to_chain" bson:"push_to_chain"`
	RawApiResponse []*ApiResponse `json:"raw_api_response" bson:"raw_api_response"`
}

type ApiResponse struct {
	Url  string `json:"url" bson:"url"`
	Data string `json:"data" bson:"data"`
}

func (m Snapshot) CollectionName() string {
	return "snapshot"
}

type ListSnapshotRequest struct {
	Pagination `query:",inline" json:"pagination"`

	UserAddress string               `json:"user_address" query:"-"`
	OrderId     string               `json:"order_id" query:"order_id"`
	OrderIds    []primitive.ObjectID `json:"order_ids" query:"order_ids"`
}

type ListSnapshotResponse struct {
	Snapshots    []*Snapshot `json:"snapshots"`
	TotalRecords int64       `json:"total_records"`
}
