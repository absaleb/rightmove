FROM registry.gitlab.okta-solutions.com/mashroom/backend/common/grpc:1.0 as builder
WORKDIR /go/src/gitlab.okta-solutions.com/mashroom/backend/cardpayment
COPY . .
RUN export CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go get -v -d ./...  && \
    go generate -v && \
    go install -tags netgo -ldflags '-w -extldflags "-static"' -v ./cmd/...
FROM scratch
COPY --from=builder /go/bin/mashroom-cardpayment /
ENTRYPOINT ["/mashroom-cardpayment"]
ENV ADDR ":10000"
ENV MONGO_URL "mongodb:27017"
ENV MONGO_DATABASE "cardpayment-data"
ENV MONGO_USERNAME "cardpayment-data"
ENV MONGO_PASSWORD "cardpayment-data"
ENV ELASTIC_URL ""
CMD -addr $ADDR