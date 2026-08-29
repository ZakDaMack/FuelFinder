FROM golang:alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/stationscraper ./cmd/stationscraper

FROM alpine:latest AS fuelfinder-stationscraper
WORKDIR /app
COPY --from=go-build /app/bin/stationscraper .

ENV GOV_API_URL="https://www.fuel-finder.service.gov.uk"
ENV GOV_API_CLIENT_ID=""
ENV GOV_API_CLIENT_SECRET=""
ENV SCRAPER_INTERVAL=15
ENV POSTGIS_DSN="postgres://localhost:5432/fuelfinder?sslmode=disable"
ENV DEBUG_MODE=false
ENV IMMEDIATE=true

CMD ["/app/stationscraper"]
