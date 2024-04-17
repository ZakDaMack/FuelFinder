all: compile

compile:
	go build -o bin/apigateway cmd/apigateway/main.go
	go build -o bin/fueldataserver cmd/fueldataserver/main.go
	go build -o bin/stationscraper cmd/stationscraper/main.go

protogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/fueldata/fueldata.proto 
