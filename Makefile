test: testHandlers testConfig

testHandlers:
	go test ./handlers/... -coverprofile=coverage.out

testConfig:
	go test ./config/... -coverprofile=coverage.out
