all: compile

compile:
	go build -o bin/apigateway cmd/apigateway/main.go
	go build -o bin/fueldataserver cmd/fueldataserver/main.go
	go build -o bin/stationscraper cmd/stationscraper/main.go

protogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/fueldata/fueldata.proto 

dockerbuild:
	docker build -t ofd-frontend . -f ./build/frontend.Dockerfile --network host
	docker build -t ofd-server . -f ./build/server.Dockerfile
	docker build -t ofd-scraper . -f ./build/scraper.Dockerfile
	docker build -t ofd-gateway . -f ./build/gateway.Dockerfile