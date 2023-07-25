// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsBindRateDataPoint(ts, 1, AttributeBindTypeServer)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsLdapBindLastSuccessfulTimeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsLdapBindRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsLdapClientSessionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsLdapSearchRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsNameCacheHitRateDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsNotificationQueuedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsOperationRateDataPoint(ts, 1, AttributeOperationTypeRead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsReplicationNetworkIoDataPoint(ts, 1, AttributeDirectionSent, AttributeNetworkDataTypeCompressed)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsReplicationObjectRateDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsReplicationOperationPendingDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsReplicationPropertyRateDataPoint(ts, 1, AttributeDirectionSent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsReplicationSyncObjectPendingDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsReplicationSyncRequestCountDataPoint(ts, 1, AttributeSyncResultSuccess)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsReplicationValueRateDataPoint(ts, 1, AttributeDirectionSent, AttributeValueTypeDistingushedNames)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsSecurityDescriptorPropagationsEventQueuedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsSuboperationRateDataPoint(ts, 1, AttributeSuboperationTypeSecurityDescriptorPropagationsEvent)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordActiveDirectoryDsThreadCountDataPoint(ts, 1)

			metrics := mb.Emit()

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 0)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "active_directory.ds.bind.rate":
					assert.False(t, validatedMetrics["active_directory.ds.bind.rate"], "Found a duplicate in the metrics slice: active_directory.ds.bind.rate")
					validatedMetrics["active_directory.ds.bind.rate"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of binds per second serviced by this domain controller.", ms.At(i).Description())
					assert.Equal(t, "{binds}/s", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "server", attrVal.Str())
				case "active_directory.ds.ldap.bind.last_successful.time":
					assert.False(t, validatedMetrics["active_directory.ds.ldap.bind.last_successful.time"], "Found a duplicate in the metrics slice: active_directory.ds.ldap.bind.last_successful.time")
					validatedMetrics["active_directory.ds.ldap.bind.last_successful.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The amount of time taken for the last successful LDAP bind.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "active_directory.ds.ldap.bind.rate":
					assert.False(t, validatedMetrics["active_directory.ds.ldap.bind.rate"], "Found a duplicate in the metrics slice: active_directory.ds.ldap.bind.rate")
					validatedMetrics["active_directory.ds.ldap.bind.rate"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of successful LDAP binds per second.", ms.At(i).Description())
					assert.Equal(t, "{binds}/s", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "active_directory.ds.ldap.client.session.count":
					assert.False(t, validatedMetrics["active_directory.ds.ldap.client.session.count"], "Found a duplicate in the metrics slice: active_directory.ds.ldap.client.session.count")
					validatedMetrics["active_directory.ds.ldap.client.session.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of connected LDAP client sessions.", ms.At(i).Description())
					assert.Equal(t, "{sessions}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "active_directory.ds.ldap.search.rate":
					assert.False(t, validatedMetrics["active_directory.ds.ldap.search.rate"], "Found a duplicate in the metrics slice: active_directory.ds.ldap.search.rate")
					validatedMetrics["active_directory.ds.ldap.search.rate"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of LDAP searches per second.", ms.At(i).Description())
					assert.Equal(t, "{searches}/s", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "active_directory.ds.name_cache.hit_rate":
					assert.False(t, validatedMetrics["active_directory.ds.name_cache.hit_rate"], "Found a duplicate in the metrics slice: active_directory.ds.name_cache.hit_rate")
					validatedMetrics["active_directory.ds.name_cache.hit_rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The percentage of directory object name component lookups that are satisfied by the Directory System Agent's name cache.", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "active_directory.ds.notification.queued":
					assert.False(t, validatedMetrics["active_directory.ds.notification.queued"], "Found a duplicate in the metrics slice: active_directory.ds.notification.queued")
					validatedMetrics["active_directory.ds.notification.queued"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of pending update notifications that have been queued to push to clients.", ms.At(i).Description())
					assert.Equal(t, "{notifications}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "active_directory.ds.operation.rate":
					assert.False(t, validatedMetrics["active_directory.ds.operation.rate"], "Found a duplicate in the metrics slice: active_directory.ds.operation.rate")
					validatedMetrics["active_directory.ds.operation.rate"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of operations performed per second.", ms.At(i).Description())
					assert.Equal(t, "{operations}/s", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "active_directory.ds.replication.network.io":
					assert.False(t, validatedMetrics["active_directory.ds.replication.network.io"], "Found a duplicate in the metrics slice: active_directory.ds.replication.network.io")
					validatedMetrics["active_directory.ds.replication.network.io"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of network data transmitted by the Directory Replication Agent.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "compressed", attrVal.Str())
				case "active_directory.ds.replication.object.rate":
					assert.False(t, validatedMetrics["active_directory.ds.replication.object.rate"], "Found a duplicate in the metrics slice: active_directory.ds.replication.object.rate")
					validatedMetrics["active_directory.ds.replication.object.rate"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of objects transmitted by the Directory Replication Agent per second.", ms.At(i).Description())
					assert.Equal(t, "{objects}/s", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "active_directory.ds.replication.operation.pending":
					assert.False(t, validatedMetrics["active_directory.ds.replication.operation.pending"], "Found a duplicate in the metrics slice: active_directory.ds.replication.operation.pending")
					validatedMetrics["active_directory.ds.replication.operation.pending"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of pending replication operations for the Directory Replication Agent.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "active_directory.ds.replication.property.rate":
					assert.False(t, validatedMetrics["active_directory.ds.replication.property.rate"], "Found a duplicate in the metrics slice: active_directory.ds.replication.property.rate")
					validatedMetrics["active_directory.ds.replication.property.rate"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of properties transmitted by the Directory Replication Agent per second.", ms.At(i).Description())
					assert.Equal(t, "{properties}/s", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
				case "active_directory.ds.replication.sync.object.pending":
					assert.False(t, validatedMetrics["active_directory.ds.replication.sync.object.pending"], "Found a duplicate in the metrics slice: active_directory.ds.replication.sync.object.pending")
					validatedMetrics["active_directory.ds.replication.sync.object.pending"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of objects remaining until the full sync completes for the Directory Replication Agent.", ms.At(i).Description())
					assert.Equal(t, "{objects}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "active_directory.ds.replication.sync.request.count":
					assert.False(t, validatedMetrics["active_directory.ds.replication.sync.request.count"], "Found a duplicate in the metrics slice: active_directory.ds.replication.sync.request.count")
					validatedMetrics["active_directory.ds.replication.sync.request.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of sync requests made by the Directory Replication Agent.", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("result")
					assert.True(t, ok)
					assert.EqualValues(t, "success", attrVal.Str())
				case "active_directory.ds.replication.value.rate":
					assert.False(t, validatedMetrics["active_directory.ds.replication.value.rate"], "Found a duplicate in the metrics slice: active_directory.ds.replication.value.rate")
					validatedMetrics["active_directory.ds.replication.value.rate"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of values transmitted by the Directory Replication Agent per second.", ms.At(i).Description())
					assert.Equal(t, "{values}/s", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "sent", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "distingushed_names", attrVal.Str())
				case "active_directory.ds.security_descriptor_propagations_event.queued":
					assert.False(t, validatedMetrics["active_directory.ds.security_descriptor_propagations_event.queued"], "Found a duplicate in the metrics slice: active_directory.ds.security_descriptor_propagations_event.queued")
					validatedMetrics["active_directory.ds.security_descriptor_propagations_event.queued"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of security descriptor propagation events that are queued for processing.", ms.At(i).Description())
					assert.Equal(t, "{events}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "active_directory.ds.suboperation.rate":
					assert.False(t, validatedMetrics["active_directory.ds.suboperation.rate"], "Found a duplicate in the metrics slice: active_directory.ds.suboperation.rate")
					validatedMetrics["active_directory.ds.suboperation.rate"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The rate of sub-operations performed.", ms.At(i).Description())
					assert.Equal(t, "{suboperations}/s", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "security_descriptor_propagations_event", attrVal.Str())
				case "active_directory.ds.thread.count":
					assert.False(t, validatedMetrics["active_directory.ds.thread.count"], "Found a duplicate in the metrics slice: active_directory.ds.thread.count")
					validatedMetrics["active_directory.ds.thread.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of threads in use by the directory service.", ms.At(i).Description())
					assert.Equal(t, "{threads}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
