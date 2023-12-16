// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mcvl/utility"
)

type (
	IDetail interface {
		Init() error
		Fetch(ctx context.Context) (utility.Protocols, error)
	}
)

var (
	localDetail IDetail
)

func Detail() IDetail {
	if localDetail == nil {
		panic("implement not found for interface IDetail, forgot register?")
	}
	return localDetail
}

func RegisterDetail(i IDetail) {
	localDetail = i
}
