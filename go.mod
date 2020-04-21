module github.com/aibotsoft/forted-service

go 1.14

require (
	github.com/aibotsoft/gen v0.0.0-00010101000000-000000000000
	github.com/aibotsoft/micro v0.0.0-20200403074803-75ca4377a55d
	github.com/dgraph-io/ristretto v0.0.2
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.14.1
	google.golang.org/grpc v1.28.0
)

replace github.com/aibotsoft/micro => ../micro

replace github.com/aibotsoft/gen => ../gen
