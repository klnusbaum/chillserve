test: testHandlers testConfig

testHandlers:
	go test ./handlers/... -coverprofile=coverage.out
	go tool cover -func=coverage.out

testConfig:
	go test ./config/... -coverprofile=coverage.out
	go tool cover -func=coverage.out
