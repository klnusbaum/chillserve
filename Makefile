test: testhandlers

testhandlers:
	go test ./handlers/... -coverprofile=coverage.out
