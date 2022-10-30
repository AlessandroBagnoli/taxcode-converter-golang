cd ../service || exit

go vet .

go fmt .

go install golang.org/x/lint/golint@latest
golint .

go install github.com/securego/gosec/v2/cmd/gosec@latest
gosec .
