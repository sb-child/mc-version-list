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
	IVerFetcher interface {
		Init() error
		Fetch(ctx context.Context) (utility.Versions, error)
	}
)

var (
	localVerFetcher IVerFetcher
)

func VerFetcher() IVerFetcher {
	if localVerFetcher == nil {
		panic("implement not found for interface IVerFetcher, forgot register?")
	}
	return localVerFetcher
}

func RegisterVerFetcher(i IVerFetcher) {
	localVerFetcher = i
}
