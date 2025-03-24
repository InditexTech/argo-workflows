package metrics

import (
	"context"
	"sync"

	"github.com/argoproj/argo-workflows/v3/util/telemetry"

	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

<<<<<<< HEAD
const (
	argoNamespace            = "argo"
	workflowsSubsystem       = "workflows"
	DefaultMetricsServerPort = 9090
	DefaultMetricsServerPath = "/metrics"
)

var (
	maxOperationTimeSeconds            = envutil.LookupEnvDurationOr("MAX_OPERATION_TIME", 30*time.Second).Seconds()
	operationDurationMetricBucketCount = envutil.LookupEnvIntOr("OPERATION_DURATION_METRIC_BUCKET_COUNT", 6)
)

type ServerConfig struct {
	Enabled      bool
	Path         string
	Port         int
	TTL          time.Duration
	IgnoreErrors bool
	Secure       bool
}

func (s ServerConfig) SameServerAs(other ServerConfig) bool {
	return s.Port == other.Port && s.Path == other.Path && s.Enabled && other.Enabled && s.Secure == other.Secure
}

type metric struct {
	metric      prometheus.Metric
	lastUpdated time.Time
	realtime    bool
	completed   bool
}

=======
>>>>>>> draft-3.6.5
type Metrics struct {
	*telemetry.Metrics

	callbacks         Callbacks
	realtimeMutex     sync.Mutex
	realtimeWorkflows map[string][]realtimeTracker
}

func New(ctx context.Context, serviceName, prometheusName string, config *telemetry.Config, callbacks Callbacks, extraOpts ...metricsdk.Option) (*Metrics, error) {
	m, err := telemetry.NewMetrics(ctx, serviceName, prometheusName, config, extraOpts...)
	if err != nil {
		return nil, err
	}

	err = m.Populate(ctx,
		telemetry.AddVersion,
		telemetry.AddDeprecationCounter,
	)
	if err != nil {
		return nil, err
	}

	metrics := &Metrics{
		Metrics:           m,
		callbacks:         callbacks,
		realtimeWorkflows: make(map[string][]realtimeTracker),
	}

	err = metrics.populate(ctx,
		addIsLeader,
		addPodPhaseGauge,
		addPodPhaseCounter,
		addPodMissingCounter,
		addPodPendingCounter,
		addWorkflowPhaseGauge,
		addCronWfTriggerCounter,
		addCronWfPolicyCounter,
		addWorkflowPhaseCounter,
		addWorkflowTemplateCounter,
		addWorkflowTemplateHistogram,
		addOperationDurationHistogram,
		addErrorCounter,
		addLogCounter,
		addK8sRequests,
		addWorkflowConditionGauge,
		addWorkQueueMetrics,
	)
	if err != nil {
		return nil, err
	}

	go metrics.customMetricsGC(ctx, config.TTL)

	return metrics, nil
}

type addMetric func(context.Context, *Metrics) error

func (m *Metrics) populate(ctx context.Context, adders ...addMetric) error {
	for _, adder := range adders {
		if err := adder(ctx, m); err != nil {
			return err
		}
	}
<<<<<<< HEAD
	for _, metric := range m.workflowsByPhase {
		allMetrics = append(allMetrics, metric)
	}
	for _, metric := range m.podsByPhase {
		allMetrics = append(allMetrics, metric)
	}
	for _, metric := range m.errors {
		allMetrics = append(allMetrics, metric)
	}
	for _, metric := range m.workqueueMetrics {
		allMetrics = append(allMetrics, metric)
	}
	for _, metric := range m.workersBusy {
		allMetrics = append(allMetrics, metric)
	}
	for _, metric := range m.customMetrics {
		allMetrics = append(allMetrics, metric.metric)
	}
	return allMetrics
}

func (m *Metrics) StopRealtimeMetricsForKey(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.workflows[key]; !exists {
		return
	}

	realtimeMetrics := m.workflows[key]
	for _, metric := range realtimeMetrics {
		if realtimeMetric, ok := m.customMetrics[metric]; ok {
			realtimeMetric.completed = true
			m.customMetrics[metric] = realtimeMetric
		}
	}
}

func (m *Metrics) DeleteRealtimeMetricsForKey(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.workflows[key]; !exists {
		return
	}

	realtimeMetrics := m.workflows[key]
	for _, metric := range realtimeMetrics {
		delete(m.customMetrics, metric)
	}

	delete(m.workflows, key)
}

func (m *Metrics) OperationCompleted(durationSeconds float64) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.operationDurations.Observe(durationSeconds)
}

func (m *Metrics) GetCustomMetric(key string) prometheus.Metric {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	// It's okay to return nil metrics in this function
	return m.customMetrics[key].metric
}

func (m *Metrics) UpsertCustomMetric(key string, ownerKey string, newMetric prometheus.Metric, realtime bool) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	metricDesc := newMetric.Desc().String()
	if _, inUse := m.defaultMetricDescs[metricDesc]; inUse {
		return fmt.Errorf("metric '%s' is already in use by the system, please use a different name", newMetric.Desc())
	}
	name, help := recoverMetricNameAndHelpFromDesc(metricDesc)
	if existingHelp, inUse := m.metricNameHelps[name]; inUse && help != existingHelp {
		return fmt.Errorf("metric '%s' has help string '%s' but should have '%s' (help strings must be identical for metrics of the same name)", name, help, existingHelp)
	} else {
		m.metricNameHelps[name] = help
	}
	m.customMetrics[key] = metric{metric: newMetric, lastUpdated: time.Now(), realtime: realtime}

	// If this is a realtime metric, track it
	if realtime {
		m.workflows[ownerKey] = append(m.workflows[ownerKey], key)
	}

=======
>>>>>>> draft-3.6.5
	return nil
}
