include .env

fmt:
	go fmt .

test:
	API_KEY=${API_KEY} go test