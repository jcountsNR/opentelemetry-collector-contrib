// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for redis metrics.
type MetricsConfig struct {
	RedisClientsBlocked                    MetricConfig `mapstructure:"redis.clients.blocked"`
	RedisClientsConnected                  MetricConfig `mapstructure:"redis.clients.connected"`
	RedisClientsMaxInputBuffer             MetricConfig `mapstructure:"redis.clients.max_input_buffer"`
	RedisClientsMaxOutputBuffer            MetricConfig `mapstructure:"redis.clients.max_output_buffer"`
	RedisCmdCalls                          MetricConfig `mapstructure:"redis.cmd.calls"`
	RedisCmdLatency                        MetricConfig `mapstructure:"redis.cmd.latency"`
	RedisCmdUsec                           MetricConfig `mapstructure:"redis.cmd.usec"`
	RedisCommands                          MetricConfig `mapstructure:"redis.commands"`
	RedisCommandsProcessed                 MetricConfig `mapstructure:"redis.commands.processed"`
	RedisConnectionsReceived               MetricConfig `mapstructure:"redis.connections.received"`
	RedisConnectionsRejected               MetricConfig `mapstructure:"redis.connections.rejected"`
	RedisCPUTime                           MetricConfig `mapstructure:"redis.cpu.time"`
	RedisDbAvgTTL                          MetricConfig `mapstructure:"redis.db.avg_ttl"`
	RedisDbExpires                         MetricConfig `mapstructure:"redis.db.expires"`
	RedisDbKeys                            MetricConfig `mapstructure:"redis.db.keys"`
	RedisKeysEvicted                       MetricConfig `mapstructure:"redis.keys.evicted"`
	RedisKeysExpired                       MetricConfig `mapstructure:"redis.keys.expired"`
	RedisKeyspaceHits                      MetricConfig `mapstructure:"redis.keyspace.hits"`
	RedisKeyspaceMisses                    MetricConfig `mapstructure:"redis.keyspace.misses"`
	RedisLatestFork                        MetricConfig `mapstructure:"redis.latest_fork"`
	RedisMaxmemory                         MetricConfig `mapstructure:"redis.maxmemory"`
	RedisMemoryFragmentationRatio          MetricConfig `mapstructure:"redis.memory.fragmentation_ratio"`
	RedisMemoryLua                         MetricConfig `mapstructure:"redis.memory.lua"`
	RedisMemoryPeak                        MetricConfig `mapstructure:"redis.memory.peak"`
	RedisMemoryRss                         MetricConfig `mapstructure:"redis.memory.rss"`
	RedisMemoryUsed                        MetricConfig `mapstructure:"redis.memory.used"`
	RedisNetInput                          MetricConfig `mapstructure:"redis.net.input"`
	RedisNetOutput                         MetricConfig `mapstructure:"redis.net.output"`
	RedisRdbChangesSinceLastSave           MetricConfig `mapstructure:"redis.rdb.changes_since_last_save"`
	RedisReplicationBacklogFirstByteOffset MetricConfig `mapstructure:"redis.replication.backlog_first_byte_offset"`
	RedisReplicationOffset                 MetricConfig `mapstructure:"redis.replication.offset"`
	RedisRole                              MetricConfig `mapstructure:"redis.role"`
	RedisSlavesConnected                   MetricConfig `mapstructure:"redis.slaves.connected"`
	RedisUptime                            MetricConfig `mapstructure:"redis.uptime"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		RedisClientsBlocked: MetricConfig{
			Enabled: true,
		},
		RedisClientsConnected: MetricConfig{
			Enabled: true,
		},
		RedisClientsMaxInputBuffer: MetricConfig{
			Enabled: true,
		},
		RedisClientsMaxOutputBuffer: MetricConfig{
			Enabled: true,
		},
		RedisCmdCalls: MetricConfig{
			Enabled: false,
		},
		RedisCmdLatency: MetricConfig{
			Enabled: false,
		},
		RedisCmdUsec: MetricConfig{
			Enabled: false,
		},
		RedisCommands: MetricConfig{
			Enabled: true,
		},
		RedisCommandsProcessed: MetricConfig{
			Enabled: true,
		},
		RedisConnectionsReceived: MetricConfig{
			Enabled: true,
		},
		RedisConnectionsRejected: MetricConfig{
			Enabled: true,
		},
		RedisCPUTime: MetricConfig{
			Enabled: true,
		},
		RedisDbAvgTTL: MetricConfig{
			Enabled: true,
		},
		RedisDbExpires: MetricConfig{
			Enabled: true,
		},
		RedisDbKeys: MetricConfig{
			Enabled: true,
		},
		RedisKeysEvicted: MetricConfig{
			Enabled: true,
		},
		RedisKeysExpired: MetricConfig{
			Enabled: true,
		},
		RedisKeyspaceHits: MetricConfig{
			Enabled: true,
		},
		RedisKeyspaceMisses: MetricConfig{
			Enabled: true,
		},
		RedisLatestFork: MetricConfig{
			Enabled: true,
		},
		RedisMaxmemory: MetricConfig{
			Enabled: false,
		},
		RedisMemoryFragmentationRatio: MetricConfig{
			Enabled: true,
		},
		RedisMemoryLua: MetricConfig{
			Enabled: true,
		},
		RedisMemoryPeak: MetricConfig{
			Enabled: true,
		},
		RedisMemoryRss: MetricConfig{
			Enabled: true,
		},
		RedisMemoryUsed: MetricConfig{
			Enabled: true,
		},
		RedisNetInput: MetricConfig{
			Enabled: true,
		},
		RedisNetOutput: MetricConfig{
			Enabled: true,
		},
		RedisRdbChangesSinceLastSave: MetricConfig{
			Enabled: true,
		},
		RedisReplicationBacklogFirstByteOffset: MetricConfig{
			Enabled: true,
		},
		RedisReplicationOffset: MetricConfig{
			Enabled: true,
		},
		RedisRole: MetricConfig{
			Enabled: false,
		},
		RedisSlavesConnected: MetricConfig{
			Enabled: true,
		},
		RedisUptime: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for redis resource attributes.
type ResourceAttributesConfig struct {
	RedisVersion ResourceAttributeConfig `mapstructure:"redis.version"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		RedisVersion: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for redis metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
