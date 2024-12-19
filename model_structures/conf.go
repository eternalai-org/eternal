package model_structures

import "strings"

const JOIN_FOR_MINTINT_KEY = "join_for_minting"

type Conf struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (Conf) CollectionName() string {
	return "configs"
}

func (o *Conf) Query(key string, v interface{}) interface{} {
	input := v.(*[]Conf)
	for _, i := range *input {
		if strings.EqualFold(i.Key, key) {
			return i
		}
	}
	return nil
}
