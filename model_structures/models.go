package model_structures

type IModel interface {
	CollectionName() string
}

type TracerData struct {
	m map[string]interface{}
}

func (d *TracerData) Add(key string, data interface{}) *TracerData {
	d.m[key] = data
	return d
}

func (d TracerData) Data() interface{} {
	return d.m
}

func NewTracerData() *TracerData {
	return &TracerData{m: map[string]interface{}{}}
}
