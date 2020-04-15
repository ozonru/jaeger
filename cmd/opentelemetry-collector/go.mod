module github.com/jaegertracing/jaeger/cmd/opentelemetry-collector

go 1.13

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab

replace github.com/jaegertracing/jaeger => ./../../

require (
	github.com/Shopify/sarama v1.22.2-0.20190604114437-cd910a683f9f
	github.com/bombsimon/wsl v1.2.5 // indirect
	github.com/census-instrumentation/opencensus-proto v0.2.1
	github.com/jaegertracing/jaeger v1.17.0
	github.com/magiconair/properties v1.8.1
	github.com/open-telemetry/opentelemetry-collector v0.3.1-0.20200408203355-0e1b2e323d39
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.0
	github.com/uber/jaeger-lib v2.2.0+incompatible
	go.uber.org/zap v1.13.0
	k8s.io/client-go v12.0.0+incompatible // indirect
)
