dev:
	~/go/bin/air ./cmd/http/main.go

dist:
	go build -o ./build/go-web ./cmd/http/main.go

db-migrate:
	go run ./cli/migrate/main.go