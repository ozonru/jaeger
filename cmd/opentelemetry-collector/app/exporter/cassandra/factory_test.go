// Copyright (c) 2020 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cassandra

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector/config/configcheck"
	"github.com/open-telemetry/opentelemetry-collector/config/configerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	jConfig "github.com/jaegertracing/jaeger/pkg/config"
	"github.com/jaegertracing/jaeger/plugin/storage/cassandra"
)

func TestCreateTraceExporter(t *testing.T) {
	v, _ := jConfig.Viperize(DefaultOptions().AddFlags)
	opts := DefaultOptions()
	opts.InitFromViper(v)
	factory := Factory{OptionsFactory: func() *cassandra.Options {
		return opts
	}}
	exporter, err := factory.CreateTraceExporter(zap.NewNop(), factory.CreateDefaultConfig())
	require.Nil(t, exporter)
	assert.EqualError(t, err, "gocql: unable to create session: control: unable to connect to initial hosts: dial tcp 127.0.0.1:9042: connect: connection refused")
}

func TestCreateTraceExporter_NilConfig(t *testing.T) {
	factory := Factory{}
	exporter, err := factory.CreateTraceExporter(zap.NewNop(), nil)
	require.Nil(t, exporter)
	assert.EqualError(t, err, "could not cast configuration to jaeger_cassandra")
}

func TestCreateDefaultConfig(t *testing.T) {
	factory := Factory{OptionsFactory: DefaultOptions}
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, configcheck.ValidateConfig(cfg))
}

func TestCreateMetricsExporter(t *testing.T) {
	f := Factory{OptionsFactory: DefaultOptions}
	mReceiver, err := f.CreateMetricsExporter(zap.NewNop(), f.CreateDefaultConfig())
	assert.Equal(t, err, configerror.ErrDataTypeIsNotSupported)
	assert.Nil(t, mReceiver)
}

func TestType(t *testing.T) {
	factory := Factory{}
	assert.Equal(t, TypeStr, factory.Type())
}
