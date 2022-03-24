.PHONY: dev
dev:
	go install github.com/makiuchi-d/arelo@latest
	arelo -p '**/*.go' -i '**/.*' -i '**/*_test.go' -- go run .

.PHONY: test
test:
	GO_ENV=test go test ./...
