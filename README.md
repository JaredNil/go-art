# go-art
go mod init github.com/jarednil/go-art
go build && ./go-art
go get github.com/joho/godotenv
go mod vendor
go mod tidy

go get github.com/go-chi/cors
go get github.com/go-chi/chi
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

IN FOLDER SCHEMA
goose postgres postgres://postgres:user@localhost:5432/goroo up
sqlc generate