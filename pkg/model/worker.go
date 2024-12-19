package model

import "time"

type WorkerStatus int64

const (
	WorkerStatusDisable WorkerStatus = iota + 1
	WorkerStatusEnable
)

type Worker struct {
	Model    `bson:",inline" json:",inline"`
	Name     string       `json:"name" bson:"name"`
	Code     string       `json:"code" bson:"code"`
	Status   WorkerStatus `json:"status" bson:"status"`
	Crontab  string       `json:"crontab" bson:"crontab"`
	Interval int          `json:"interval" bson:"interval"`
	State    *State       `json:"state" bson:"state"`
}

type State struct {
	LastSyncedBlock uint64    `json:"last_synced_block" bson:"last_synced_block"`
	ResyncFromBlock uint64    `json:"resync_from_block" bson:"resync_from_block"`
	LastRunDatetime time.Time `json:"last_run_datetime" bson:"last_run_datetime"`
	IsRunning       bool      `json:"is_running" bson:"is_running"`
}

func (m Worker) CollectionName() string {
	return "worker"
}
