cd .. || exit
go install github.com/swaggo/swag/cmd/swag@latest
swag init --pd
