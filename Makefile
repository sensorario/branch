default:
	go test ./... -v

update_mod:
	GO111MODULE=on go mod tidy

update:
	go get -u all

GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test ./...
.PHONY: test/cover
test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out
