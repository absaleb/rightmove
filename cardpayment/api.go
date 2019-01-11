//go:generate protoc --go_out=plugins=grpc,paths=source_relative:. -I .:$GOPATH/src/gitlab.okta-solutions.com/mashroom/backend/common/protobuf api.proto
package cardpayment