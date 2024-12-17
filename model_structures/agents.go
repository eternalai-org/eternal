package model_structures

type ExternalDataRequestType string

var ExternalDataRequestTypeBatch ExternalDataRequestType = "batch"

type AgentInferExternalData struct {
	RoomId             string                   `json:"room_id,omitempty"`
	AgentID            string                   `json:"agent_id,omitempty"`
	OutputMaxCharacter *uint                    `json:"output_max_character,omitempty"`
	Type               *ExternalDataRequestType `json:"type,omitempty"`
}
