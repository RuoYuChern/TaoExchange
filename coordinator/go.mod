module tao.exchange.com/coordinator

go 1.20

replace tao.exchange.com/common => ../common

replace tao.exchange.com/grpc => ../grpc

require github.com/uptrace/bunrouter/extra/reqlog v1.0.20

require (
	github.com/fatih/color v1.14.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/rs/cors v1.9.0
	github.com/uptrace/bunrouter v1.0.20 // indirect
	go.opentelemetry.io/otel v1.13.0 // indirect
	go.opentelemetry.io/otel/trace v1.13.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)
