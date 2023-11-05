.PHONY: prod test run_tests run

BINARY_NAME=lastdisco-backend

prod:
	cd cmd/main && GOARCH=amd64 GOOS=windows go build -o ../../bin/${BINARY_NAME}-windows-testing -tags prod
	cd cmd/main && GOARCH=amd64 GOOS=darwin  go build -o ../../bin/${BINARY_NAME}-darwin-testing  -tags prod
	cd cmd/main && GOARCH=amd64 GOOS=linux   go build -o ../../bin/${BINARY_NAME}-linux-testing   -tags prod

test:
	cd cmd/main && GOARCH=amd64 GOOS=windows go build -o ../../bin/${BINARY_NAME}-windows-testing -tags test
	cd cmd/main && GOARCH=amd64 GOOS=darwin  go build -o ../../bin/${BINARY_NAME}-darwin-testing  -tags test
	cd cmd/main && GOARCH=amd64 GOOS=linux   go build -o ../../bin/${BINARY_NAME}-linux-testing   -tags test

clean:
	cd cmd/main && go clean
	rm bin/${BINARY_NAME}*

vet:
	cd cmd/main && go vet -tags test

dep:
	go mod download

run_tests:
	@go test ./test/

run:
	@./_bin/lastdisco-backend
