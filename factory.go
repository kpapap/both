package both

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (defaultInterval = "1m")
var (typeStr = component.MustNewType("both"))
const (defaultNbcmrInterval = "1m")
const (defaultConfigMapName = "nbcmr-cm")

// NewFactory creates a new receiver factory.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithLogs(createWhitelistLogsReceiver, component.StabilityLevelUndefined),
		receiver.WithLogs(createNbcmrLogsReceiver, component.StabilityLevelUndefined),
	)
}

// createDefaultConfig returns the default configuration for the  receiver.
// This function is used when creating a new factory to provide a default configuration
// for the receiver.
func createWhitelistDefaultConfig() component.Config {
	return &Config{
		Interval: defaultInterval,
	}
}

// createLogsReceiver creates a new instance of the logs receiver.
// createLogsReceiver creates a log receiver based on provided config.
func createWhitelistLogsReceiver(_ context.Context, settings receiver.Settings, cfg component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	// Create the new receiver
	rCfg := cfg.(*Config)
	return newWhitelistReceiver(rCfg, consumer, settings)
}

// createLogsReceiver creates a new instance of the nbcmr receiver.
func createNbcmrLogsReceiver(_ context.Context, settings receiver.Settings, cfg component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	// Create the new receiver
	rCfg := cfg.(*Config)
	return newNbcmrReceiver(rCfg, consumer, settings)
}

// createDefaultConfig returns the default configuration for the nbcmr receiver.
func createDefaultConfig() component.Config {
	return &Config{
		Interval: defaultInterval,
		ConfigMapName: defaultConfigMapName,
	}
}


