test:
	@go test -v -cover -coverprofile=coverage.out github.com/priyakdey/trie

cov: test
	@go tool cover -html=coverage.out -o coverage.html
	@python3 -m http.server 7001 --bind 127.0.0.1 --directory .

fmt:
	@go fmt
