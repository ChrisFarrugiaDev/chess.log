go mod init chess_log/go_backend

go mod tidy

go build -o chesslog ./cmd/web/.

go build -ldflags="-s -w" -o chesslog ./cmd/web/.