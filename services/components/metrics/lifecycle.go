package metrics

import (
	"fmt"
	"time"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/forta-network/forta-node/clients"
	"github.com/forta-network/forta-node/config"
)

// Lifecycle creates lifecycle metrics. It is useful in
// understanding what is going on during lifecycle management.
type Lifecycle interface {
	ClientDial(...config.AgentConfig)
	ClientClose(...config.AgentConfig)

	StatusRunning(...config.AgentConfig)
	StatusAttached(...config.AgentConfig)
	StatusInitialized(...config.AgentConfig)
	StatusStopping(...config.AgentConfig)
	StatusActive(...config.AgentConfig)
	StatusInactive(...config.AgentConfig)

	ActionUpdate(...config.AgentConfig)
	ActionRestart(...config.AgentConfig)
	ActionSubscribe([]domain.CombinerBotSubscription)
	ActionUnsubscribe([]domain.CombinerBotSubscription)

	FailurePull(error, ...config.AgentConfig)
	FailureLaunch(error, ...config.AgentConfig)
	FailureStop(error, ...config.AgentConfig)
	FailureDial(error, ...config.AgentConfig)
	FailureInitialize(error, ...config.AgentConfig)
	FailureInitializeResponse(error, ...config.AgentConfig)
	FailureInitializeValidate(error, ...config.AgentConfig)
	FailureTooManyErrs(error, ...config.AgentConfig)

	BotError(metricName string, err error, cfgs ...config.AgentConfig)
	SystemError(metricName string, err error)

	SystemStatus(metricName string, details string)

	HealthCheckAttempt(botConfigs ...config.AgentConfig)
	HealthCheckSuccess(botConfigs ...config.AgentConfig)
	HealthCheckError(err error, botConfigs ...config.AgentConfig)
}

type lifecycle struct {
	msgClient clients.MessageClient
}

// NewLifecycleClient creates a new metrics client.
func NewLifecycleClient(msgClient clients.MessageClient) Lifecycle {
	return &lifecycle{
		msgClient: msgClient,
	}
}

func (lc *lifecycle) ClientDial(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricClientDial, "", botConfigs))
}

func (lc *lifecycle) ClientClose(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricClientClose, "", botConfigs))
}

func (lc *lifecycle) StatusRunning(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricStatusRunning, "", botConfigs))
}

func (lc *lifecycle) StatusAttached(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricStatusAttached, "", botConfigs))
}

func (lc *lifecycle) StatusInitialized(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricStatusInitialized, "", botConfigs))
}

func (lc *lifecycle) StatusStopping(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricStatusStopping, "", botConfigs))
}

func (lc *lifecycle) StatusActive(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricStatusActive, "", botConfigs))
}

func (lc *lifecycle) StatusInactive(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricStatusInactive, "", botConfigs))
}

func (lc *lifecycle) ActionUpdate(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricActionUpdate, "", botConfigs))
}

func (lc *lifecycle) ActionRestart(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricActionRestart, "", botConfigs))
}

func (lc *lifecycle) ActionSubscribe(subscriptions []domain.CombinerBotSubscription) {
	SendAgentMetrics(lc.msgClient, fromBotSubscriptions(domain.MetricActionSubscribe, subscriptions))
}

func (lc *lifecycle) ActionUnsubscribe(subscriptions []domain.CombinerBotSubscription) {
	SendAgentMetrics(lc.msgClient, fromBotSubscriptions(domain.MetricActionUnsubscribe, subscriptions))
}

func (lc *lifecycle) FailurePull(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricFailurePull, err.Error(), botConfigs))
}

func (lc *lifecycle) FailureLaunch(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricFailureLaunch, err.Error(), botConfigs))
}

func (lc *lifecycle) FailureStop(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricFailureStop, err.Error(), botConfigs))
}

func (lc *lifecycle) FailureDial(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricFailureDial, err.Error(), botConfigs))
}

func (lc *lifecycle) FailureInitialize(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricFailureInitialize, err.Error(), botConfigs))
}

func (lc *lifecycle) FailureInitializeResponse(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricFailureInitializeResponse, err.Error(), botConfigs))
}

func (lc *lifecycle) FailureInitializeValidate(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricFailureInitializeValidate, err.Error(), botConfigs))
}

func (lc *lifecycle) FailureTooManyErrs(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricFailureTooManyErrs, err.Error(), botConfigs))
}

func (lc *lifecycle) BotError(metricName string, err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(fmt.Sprintf("agent.error.%s", metricName), err.Error(), botConfigs))
}

func (lc *lifecycle) SystemError(metricName string, err error) {
	SendAgentMetrics(lc.msgClient, systemMetrics(fmt.Sprintf("system.error.%s", metricName), err.Error()))
}

func (lc *lifecycle) SystemStatus(metricName string, details string) {
	SendAgentMetrics(lc.msgClient, systemMetrics(fmt.Sprintf("system.status.%s", metricName), details))
}

func (lc *lifecycle) HealthCheckAttempt(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricHealthCheckAttempt, "", botConfigs))
}

func (lc *lifecycle) HealthCheckSuccess(botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricHealthCheckSuccess, "", botConfigs))
}

func (lc *lifecycle) HealthCheckError(err error, botConfigs ...config.AgentConfig) {
	SendAgentMetrics(lc.msgClient, fromBotConfigs(domain.MetricHealthCheckError, err.Error(), botConfigs))
}

func fromBotSubscriptions(action string, subscriptions []domain.CombinerBotSubscription) (metrics []*protocol.AgentMetric) {
	for _, botSub := range subscriptions {
		metrics = append(metrics, CreateAgentMetricV1(config.AgentConfig{ID: botSub.Subscriber.BotID}, action, 1))
	}
	return
}

func fromBotConfigs(metricName string, details string, botConfigs []config.AgentConfig) (metrics []*protocol.AgentMetric) {
	details = utils.ObfuscateURLs(details)
	for _, botConfig := range botConfigs {
		metric := &protocol.AgentMetric{
			AgentId:   botConfig.ID,
			Timestamp: time.Now().Format(time.RFC3339),
			Name:      metricName,
			Details:   details,
			Value:     1,
			ShardId:   botConfig.ShardID(),
			ChainId:   int64(botConfig.ChainID),
		}
		metrics = append(metrics, metric)
	}
	return
}

func systemMetrics(metricName string, details string) []*protocol.AgentMetric {
	return fromBotConfigs(metricName, details, []config.AgentConfig{{ID: "system"}})
}
