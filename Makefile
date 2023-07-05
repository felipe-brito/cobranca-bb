BINARY_NAME=cobranca-bb

install:
	go env -w CGO_ENABLED="1"
	go env -w GO111MODULE="on"
	go env -w GOBIN=${GOPATH}/bin
	go mod download -x
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@v1.8.7
	swag -v

run:
	go run cmd/http/main.go

dep:
	go get -d -u ./...
	go mod download -x

doc-swag:
	swag init --parseDependency --parseInternal --parseDepth 1 --output internal/http/swagger/docs --generalInfo cmd/http/main.go

build: doc-swag
	go build -race -o bin/${BINARY_NAME} -ldflags="-s -m" ./cmd/http

docker-compose-up:
	docker compose -f ./resource/docker/docker-compose.yml up -d


