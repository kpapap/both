package both

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (typeStr = component.MustNewType("both"))

const (defaultInterval = "1m")
const (defaultHttpInterval = "5m")
const (defaultConfigMapName = "nbcmr-cm")

// NewFactory creates a new receiver factory.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithLogs(createBothLogsReceiver, component.StabilityLevelUndefined),
	)
}

// createLogsReceiver creates a new instance of the logs receiver.
func createBothLogsReceiver(_ context.Context, settings receiver.Settings, cfg component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	// Create the new receiver
	rCfg := cfg.(*Config)
	return newBothReceiver(rCfg, consumer, settings)
}

// createDefaultConfig returns the default configuration for the nbcmr receiver.
func createDefaultConfig() component.Config {
	return &Config{
		Interval: defaultInterval,
		HttpInterval: defaultHttpInterval,
		ConfigMapName: defaultConfigMapName,
	}
}


