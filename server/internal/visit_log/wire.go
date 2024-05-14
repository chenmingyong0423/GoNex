// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject

package visit_log

import (
	"github.com/chenmingyong0423/fnote/server/internal/visit_log/internal/repository"
	"github.com/chenmingyong0423/fnote/server/internal/visit_log/internal/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/visit_log/internal/service"
	"github.com/chenmingyong0423/fnote/server/internal/visit_log/internal/web"
	"github.com/chenmingyong0423/go-eventbus"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var VisitLogProviders = wire.NewSet(web.NewVisitLogHandler, service.NewVisitLogService, repository.NewVisitLogRepository, dao.NewVisitLogDao,
	wire.Bind(new(service.IVisitLogService), new(*service.VisitLogService)),
	wire.Bind(new(repository.IVisitLogRepository), new(*repository.VisitLogRepository)),
	wire.Bind(new(dao.IVisitLogDao), new(*dao.VisitLogDao)))

func InitVisitLogModule(mongoDB *mongo.Database, eventBus *eventbus.EventBus) *Module {
	panic(wire.Build(
		VisitLogProviders,
		wire.Struct(new(Module), "Svc", "Hdl"),
	))
}
