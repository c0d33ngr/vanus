module github.com/linkall-labs/vanus/raft

go 1.17

require (
	github.com/cockroachdb/datadriven v0.0.0-20200714090401-bf6692d28da5
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	go.etcd.io/etcd/client/pkg/v3 v3.5.7
	go.opentelemetry.io/otel/trace v1.11.2
)

require (
	github.com/certifi/gocertifi v0.0.0-20200922220541-2c3bb06c6054 // indirect
	github.com/cockroachdb/errors v1.2.4 // indirect
	github.com/cockroachdb/logtags v0.0.0-20190617123548-eb05cc24525f // indirect
	github.com/getsentry/raven-go v0.2.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

// Bad imports are sometimes causing attempts to pull that code.
// This makes the error more explicit.
replace go.etcd.io/etcd => ./FORBIDDEN_DEPENDENCY

replace go.etcd.io/etcd/v3 => ./FORBIDDEN_DEPENDENCY

replace github.com/linkall-labs/vanus/observability => ../observability
