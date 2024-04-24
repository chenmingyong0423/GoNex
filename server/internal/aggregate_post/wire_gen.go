// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package aggregate_post

import (
	"github.com/chenmingyong0423/fnote/server/internal/aggregate_post/internal/web"
	"github.com/chenmingyong0423/fnote/server/internal/post/service"
	"github.com/chenmingyong0423/fnote/server/internal/post_draft"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitAggregatePostModule(postServ service.IPostService, postDraftModel *post_draft.Model) *Model {
	iPostDraftService := postDraftModel.Svc
	aggregatePostHandler := web.NewAggregatePostHandler(postServ, iPostDraftService)
	model := &Model{
		Hdl: aggregatePostHandler,
	}
	return model
}

// wire.go:

var AggregatePostProviders = wire.NewSet(web.NewAggregatePostHandler)
