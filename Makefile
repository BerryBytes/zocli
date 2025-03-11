default: run

run:
	go run main.go

test:
	go test -v ./...

lint:
	golangci-lint run

debug:clean build-all
	go run main.go -vvv

build:
	CGO_ENABLED=0 go build -gcflags "all=-N -l"	--ldflags "-s -w" -o bin/zocli main.go

build-all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l"	--ldflags "-s -w" -o bin/zocli-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -gcflags "all=-N -l"	--ldflags "-s -w" -o bin/zocli-linux-arm64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -gcflags "all=-N -l"	--ldflags "-s -w" -o bin/zocli-windows-amd64.exe main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -gcflags "all=-N -l"	--ldflags "-s -w" -o bin/zocli-darwin-amd64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -gcflags "all=-N -l"	--ldflags "-s -w" -o bin/zocli-darwin-arm64 main.go

test-cover:
	go test --coverprofile=cover.out -v ./...
	go tool cover -func=cover.out

clean:
	rm bin -rf

source=printer/debug.go
dest=printer/mock/mock.go
mock-gen:
	mockgen -source=pkg/utils/${source} -destination=pkg/utils/${dest}
