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
package appsetup

import (
	"github.com/google/wire"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

func NewConfigurationHelper(logger *zap.SugaredLogger) (configuration.ConfigurationHelper, error) {
	wire.Build(
		os.NewOsHelper,
		configuration.NewConfigurationHelper)

	return nil, nil
}

func NewTimeHelper() (time.TimeHelper, error) {
	wire.Build(
		time.NewTimeHelper)

	return nil, nil
}

func NewDatabase(
	logger *zap.SugaredLogger,
	postgresConfig postgres.PostgresConfig) (database.Database, error) {
	wire.Build(postgres.NewPostgres)

	return nil, nil
}

func NewOsHelper() (os.OsHelper, error) {
	wire.Build(os.NewOsHelper)

	return nil, nil
}

func NewPulsarClient(
	logger *zap.SugaredLogger,
	pulsarConfig pulsar.PulsarConfig) (pulsar.PulsarClient, error) {
	wire.Build(
		os.NewOsHelper,
		pulsar.NewPulsarClient)

	return nil, nil
}
