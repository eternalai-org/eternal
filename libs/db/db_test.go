package db

import (
	"eternal-infer-worker/model_structures"
	"testing"
)

func TestWrite(t *testing.T) {
	data := []model_structures.ContractSyncState{
		{
			ContractAddress:  "ContractAddress",
			Job:              "job",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
		{
			ContractAddress:  "ContractAddress_1",
			Job:              "job_1",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
		{
			ContractAddress:  "ContractAddress_2",
			Job:              "job_2",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
		{
			ContractAddress:  "ContractAddress_3",
			Job:              "job_3",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
	}

	err := Write(model_structures.ContractSyncState{}.CollectionName(), data)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("OK")
}

func TestRead(t *testing.T) {
	data := []model_structures.ContractSyncState{}

	err := Read(model_structures.ContractSyncState{}.CollectionName(), &data)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)
}

func TestUpdate(t *testing.T) {
	data := []model_structures.ContractSyncState{
		{
			ContractAddress:  "ContractAddress_updated",
			Job:              "job_update",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
		{
			ContractAddress:  "ContractAddress_1",
			Job:              "job_1",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
		{
			ContractAddress:  "ContractAddress_updated_2",
			Job:              "job_update_2",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
		{
			ContractAddress:  "ContractAddress_3",
			Job:              "job_3",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
		{
			ContractAddress:  "ContractAddress_3_4",
			Job:              "job_3_4",
			LastSyncedBlock:  5,
			ResyncToBlock:    3,
			ResyncFromBlock:  4,
			ClearDataAndSync: true,
		},
	}

	err := Update(model_structures.ContractSyncState{}.CollectionName(), &data)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)
}
