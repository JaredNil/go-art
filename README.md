# go-art
go mod init github.com/jarednil/go-art
go build && ./go-art
go get github.com/joho/godotenv
go mod vendor
go mod tidy

 go get github.com/go-chi/cors
  go get github.com/go-chi/chi