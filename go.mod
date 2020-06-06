module github.com/aibotsoft/forted-service

go 1.14

require (
	github.com/aibotsoft/gen v0.0.0-20200531091936-c4d5d714bf82
	github.com/aibotsoft/micro v0.0.0-20200603095207-2998daaf8ae2
	github.com/dgraph-io/ristretto v0.0.2
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.0
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	google.golang.org/genproto v0.0.0-20200605102947-12044bf5ea91 // indirect
	google.golang.org/grpc v1.29.1
)

replace github.com/aibotsoft/micro => ../micro

replace github.com/aibotsoft/gen => ../gen
