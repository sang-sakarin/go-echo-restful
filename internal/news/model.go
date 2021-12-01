package news

import "go-echo-restful/internal/common"

type News struct {
	common.SoftControlModel `pg:"override"`
	Title                   string `json:"title"`
	Description             string `json:"description"`
	IsPublish               bool   `json:"is_publish"`
}
