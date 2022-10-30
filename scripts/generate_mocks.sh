cd .. || exit
docker run --rm -v "$PWD":/src -w /src vektra/mockery --all --output ./service/mocks
