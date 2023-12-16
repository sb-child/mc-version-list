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
	IProtoFetcher interface {
		Init() error
		Fetch(ctx context.Context) (utility.Protocols, error)
	}
)

var (
	localProtoFetcher IProtoFetcher
)

func ProtoFetcher() IProtoFetcher {
	if localProtoFetcher == nil {
		panic("implement not found for interface IProtoFetcher, forgot register?")
	}
	return localProtoFetcher
}

func RegisterProtoFetcher(i IProtoFetcher) {
	localProtoFetcher = i
}
