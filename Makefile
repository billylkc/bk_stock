all: download_stock

industry: cmd/download_industry.go
	go build -o bin/download_industry cmd/download_industry.go

stock: cmd/download_stock.go
	go build -o bin/download_stock cmd/download_stock.go

sector: cmd/download_sector.go
	go build -o bin/download_sector cmd/download_sector.go

windows: cmd/download_stock.go
	GOOS=windows GOARCH=386 go build -o bin/download_stock cmd/download_stock.go

test: godeploy
	./godeploy sth
