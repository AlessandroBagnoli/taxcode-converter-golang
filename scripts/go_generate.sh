cd .. || exit
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/vektra/mockery/v2@latest
go generate ./...
