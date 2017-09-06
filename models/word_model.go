package models

import "github.com/datake914/lonbutu/common"

type WordModel struct {
	baseModel
	LogicalName     string
	PhysicalName    string
	PhysicalNameAbb string
	Status          common.WordStatus
}
