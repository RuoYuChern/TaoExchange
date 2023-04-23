module tao.exchange.com/infra

go 1.20

replace tao.exchange.com/common => ../common

require (
	github.com/uptrace/bun v1.1.12
	github.com/uptrace/bun/dialect/pgdialect v1.1.12
	github.com/uptrace/bun/driver/pgdriver v1.1.12
	golang.org/x/exp v0.0.0-20230420155640-133eef4313cb
	tao.exchange.com/common v0.0.0-00010101000000-000000000000
)

require (
	github.com/agrison/go-commons-lang v0.0.0-20200208220349-58e9fcb95174 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	mellium.im/sasl v0.3.1 // indirect
)
