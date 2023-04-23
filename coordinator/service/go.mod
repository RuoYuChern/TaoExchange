module tao.exchange.com/coordinator

go 1.20

replace tao.exchange.com/common => ../../common

replace tao.exchange.com/infra/orm => ../../infra/orm

replace tao.exchange.com/infra => ../../infra

replace tao.exchange.com/grpc => ../../grpc

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/rs/cors v1.9.0
	github.com/uptrace/bunrouter v1.0.20
	github.com/uptrace/bunrouter/extra/reqlog v1.0.20
	golang.org/x/exp v0.0.0-20230420155640-133eef4313cb
	tao.exchange.com/common v0.0.0-00010101000000-000000000000
	tao.exchange.com/grpc v0.0.0-00010101000000-000000000000
	tao.exchange.com/infra/orm v0.0.0-00010101000000-000000000000
)

require (
	github.com/agrison/go-commons-lang v0.0.0-20200208220349-58e9fcb95174 // indirect
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/fatih/color v1.14.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	github.com/uptrace/bun v1.1.12 // indirect
	github.com/uptrace/bun/dialect/pgdialect v1.1.12 // indirect
	github.com/uptrace/bun/driver/pgdriver v1.1.12 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.opentelemetry.io/otel v1.13.0 // indirect
	go.opentelemetry.io/otel/trace v1.13.0 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	mellium.im/sasl v0.3.1 // indirect
	tao.exchange.com/infra v0.0.0-00010101000000-000000000000 // indirect
)
