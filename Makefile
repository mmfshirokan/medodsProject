mock gen service:
	mockery --dir ./internal/service --all --output ./internal/service/mock --with-expecter
mock gen handlers:
	mockery --dir ./internal/handlers --all --output ./internal/handlers/mock --with-expecter
lint:
	golangci-lint run