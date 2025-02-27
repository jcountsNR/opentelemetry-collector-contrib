// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkametricsreceiver

import (
	"context"
	"fmt"
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver/internal/metadata"
)

func TestBrokerShutdown(t *testing.T) {
	client := newMockClient()
	client.closed = false
	client.close = nil
	client.Mock.
		On("Close").Return(nil).
		On("Closed").Return(false)
	scraper := brokerScraper{
		client:   client,
		settings: receivertest.NewNopCreateSettings(),
		config:   Config{},
	}
	_ = scraper.shutdown(context.Background())
	client.AssertExpectations(t)
}

func TestBrokerShutdown_closed(t *testing.T) {
	client := newMockClient()
	client.closed = true
	client.Mock.
		On("Closed").Return(true)
	scraper := brokerScraper{
		client:   client,
		settings: receivertest.NewNopCreateSettings(),
		config:   Config{},
	}
	_ = scraper.shutdown(context.Background())
	client.AssertExpectations(t)
}

func TestBrokerScraper_Name(t *testing.T) {
	s := brokerScraper{}
	assert.Equal(t, s.Name(), brokersScraperName)
}

func TestBrokerScraper_createBrokerScraper(t *testing.T) {
	sc := sarama.NewConfig()
	newSaramaClient = mockNewSaramaClient
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
}

func TestBrokerScraperStart(t *testing.T) {
	newSaramaClient = mockNewSaramaClient
	sc := sarama.NewConfig()
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
	assert.NoError(t, bs.Start(context.Background(), nil))
}

func TestBrokerScraper_scrape_handles_client_error(t *testing.T) {
	newSaramaClient = func(addrs []string, conf *sarama.Config) (sarama.Client, error) {
		return nil, fmt.Errorf("new client failed")
	}
	sc := sarama.NewConfig()
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
	_, err = bs.Scrape(context.Background())
	assert.Error(t, err)
}

func TestBrokerScraper_shutdown_handles_nil_client(t *testing.T) {
	newSaramaClient = func(addrs []string, conf *sarama.Config) (sarama.Client, error) {
		return nil, fmt.Errorf("new client failed")
	}
	sc := sarama.NewConfig()
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
	err = bs.Shutdown(context.Background())
	assert.NoError(t, err)
}

func TestBrokerScraper_scrape(t *testing.T) {
	client := newMockClient()
	client.Mock.On("Brokers").Return(testBrokers)
	bs := brokerScraper{
		client:   client,
		settings: receivertest.NewNopCreateSettings(),
		config:   Config{MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig()},
	}
	require.NoError(t, bs.start(context.Background(), componenttest.NewNopHost()))
	md, err := bs.scrape(context.Background())
	assert.NoError(t, err)
	require.Equal(t, 1, md.ResourceMetrics().Len())
	require.Equal(t, 1, md.ResourceMetrics().At(0).ScopeMetrics().Len())
	ms := md.ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics()
	for i := 0; i < ms.Len(); i++ {
		m := ms.At(i)
		switch m.Name() {
		case "kafka.brokers":
			assert.Equal(t, m.Sum().DataPoints().At(0).IntValue(), int64(len(testBrokers)))
		case "kafka.broker.count":
			assert.Equal(t, m.Sum().DataPoints().At(0).IntValue(), int64(len(testBrokers)))
		case "kafka.broker.consumer_fetch_rate":
			assert.Equal(t, m.Gauge().DataPoints().At(0).DoubleValue(), int64(len(testBrokers)))
		case "kafka.broker.incoming_byte_rate":
			assert.Equal(t, m.Gauge().DataPoints().At(0).DoubleValue(), int64(len(testBrokers)))
		case "kafka.broker.outgoing_byte_rate":
			assert.Equal(t, m.Gauge().DataPoints().At(0).DoubleValue(), int64(len(testBrokers)))
		case "kafka.broker.request_latency":
			assert.Equal(t, m.Gauge().DataPoints().At(0).DoubleValue(), int64(len(testBrokers)))
		case "kafka.broker.response_rate":
			assert.Equal(t, m.Gauge().DataPoints().At(0).DoubleValue(), int64(len(testBrokers)))
		case "kafka.broker.response_size":
			assert.Equal(t, m.Gauge().DataPoints().At(0).DoubleValue(), int64(len(testBrokers)))
		case "kafka.broker.request_rate":
			assert.Equal(t, m.Gauge().DataPoints().At(0).DoubleValue(), int64(len(testBrokers)))
		case "kafka.broker.request_size":
			assert.Equal(t, m.Gauge().DataPoints().At(0).DoubleValue(), int64(len(testBrokers)))
		case "kafka.broker.requests_in_flight":
			assert.Equal(t, m.Sum().DataPoints().At(0).IntValue(), int64(len(testBrokers)))
		case "kafka.broker.consumer_fetch_count":
			assert.Equal(t, m.Sum().DataPoints().At(0).IntValue(), int64(len(testBrokers)))
		}
	}
}

func TestBrokersScraper_createBrokerScraper(t *testing.T) {
	sc := sarama.NewConfig()
	newSaramaClient = mockNewSaramaClient
	bs, err := createBrokerScraper(context.Background(), Config{}, sc, receivertest.NewNopCreateSettings())
	assert.NoError(t, err)
	assert.NotNil(t, bs)
}
