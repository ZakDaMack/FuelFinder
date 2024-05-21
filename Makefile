all: compile

compile:
	go build -o bin/apigateway cmd/apigateway/main.go
	go build -o bin/fuelfinderserver cmd/fuelfinderserver/main.go
	go build -o bin/stationscraper cmd/stationscraper/main.go

protogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/fuelfinder/fuelfinder.proto 

dockerbuild:
	docker build -t fuelfinder-frontend . -f ./build/frontend.Dockerfile --network host
	docker build -t fuelfinder-server . -f ./build/server.Dockerfile
	docker build -t fuelfinder-scraper . -f ./build/scraper.Dockerfile
	docker build -t fuelfinder-gateway . -f ./build/gateway.Dockerfile