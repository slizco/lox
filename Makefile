lint:
	golint

dep:
	dep ensure

vet:	dep
	go vet $(go list ./... | grep -v vendor)

qtest:
	go test -v $(go list ./... | grep -v vendor)

test:	dep qtest
