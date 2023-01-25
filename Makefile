run:
	go run cmd/rest/main.go

build:
	go build -o app ./cmd/rest

migration-up:
	go run cmd/migration/up/main.go
