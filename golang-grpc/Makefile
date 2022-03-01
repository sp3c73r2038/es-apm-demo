ELASTIC_APM_SERVER_URL := http://127.0.0.1:8201
ELASTIC_APM_API_KEY :=

-include local.mk

gen-proto:
	protoc -I. --go_out=plugins=grpc,paths=source_relative:. proto/*.proto

build-user:
	go build -o bin/user user.go

build-order:
	go build -o bin/order order.go

build-client:
	go build -o bin/client client.go

user: build-user
	ELASTIC_APM_API_KEY=${ELASTIC_APM_API_KEY} \
	ELASTIC_APM_ACTIVE=true ELASTIC_APM_SERVICE_NAME=golang-grpc-demo-user \
		ELASTIC_APM_DISABLE_METRICS="system*,transaction*,golang*" \
		ELASTIC_APM_BREAKDOWN_METRIC=true \
		ELASTIC_APM_ENVIRONMENT=dev \
		ELASTIC_APM_SERVER_URL=${ELASTIC_APM_SERVER_URL} \
		bin/user

order: build-order
	ELASTIC_APM_API_KEY=${ELASTIC_APM_API_KEY} \
	ELASTIC_APM_ACTIVE=true ELASTIC_APM_SERVICE_NAME=golang-grpc-demo-order \
		ELASTIC_APM_DISABLE_METRICS="system*,transaction*,golang*" \
		ELASTIC_APM_BREAKDOWN_METRIC=true \
		ELASTIC_APM_ENVIRONMENT=dev \
		ELASTIC_APM_SERVER_URL=${ELASTIC_APM_SERVER_URL} \
		bin/order

client: build-client
	ELASTIC_APM_API_KEY=${ELASTIC_APM_API_KEY} \
	ELASTIC_APM_ACTIVE=true ELASTIC_APM_SERVICE_NAME=golang-grpc-demo-client \
		ELASTIC_APM_DISABLE_METRICS="system*,transaction*,golang*" \
		ELASTIC_APM_BREAKDOWN_METRIC=true \
		ELASTIC_APM_ENVIRONMENT=dev \
		ELASTIC_APM_SERVER_URL=${ELASTIC_APM_SERVER_URL} \
		bin/client
