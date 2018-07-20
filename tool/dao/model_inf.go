package dao

import (
	"cloud/tool/dao/impl"
)

type ModelInf interface {
	CreateModel(url, table string) (string, error)
	ShowTables(url string) ([]string, error)
}

var (
	Model ModelInf = impl.Model{}
)
