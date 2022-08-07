// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package commands

import (
	"github.com/google/wire"
	"github.com/omiga-group/omiga/src/order/order-api/graphql"
	"github.com/omiga-group/omiga/src/order/order-api/http"
	"github.com/omiga-group/omiga/src/order/order-api/services"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"go.uber.org/zap"
)

func NewHttpServer(
	logger *zap.SugaredLogger,
	appSettings configuration.AppSettings,
	postgresSettings postgres.PostgresSettings) (http.HttpServer, error) {
	wire.Build(postgres.NewPostgres, repositories.NewEntgoClient, http.NewHttpServer, graphql.NewGraphQLServer, services.NewOrderService)

	return nil, nil
}