all: download_stock

linux: cmd/download_stock.go
	go build -o bin/download_stock cmd/download_stock.go

windows: cmd/download_stock.go
	GOOS=windows GOARCH=386 go build -o bin/download_stock cmd/download_stock.go

test: godeploy
	./godeploy sth
