include .env

fmt:
	go fmt .

test:
	CHAPA_API_KEY=${CHAPA_API_KEY} go test