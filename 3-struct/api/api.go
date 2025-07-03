package api

import "struct/config"

type Api struct {
	Key string
}

func NewApi(cfg config.Config) *Api {
	return &Api{
		Key: cfg.Key,
	}
}
