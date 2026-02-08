all: compile dockerbuild

compile:
	go build -o bin/stationapi cmd/stationapi/main.go
	go build -o bin/stationscraper cmd/stationscraper/main.go

dockerbuild:
	docker build -t fuelfinder-frontend . -f ./build/frontend.Dockerfile --network host
	docker build -t fuelfinder-stationapi . -f ./build/stationapi.Dockerfile
	docker build -t fuelfinder-scraper . -f ./build/scraper.Dockerfile

compose:
	docker compose -f deployments/docker-compose.yml up -d --build