FROM golang:alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/stationscraper cmd/stationscraper/main.go

FROM alpine:latest AS fuelfinder-stationscraper
WORKDIR /app
COPY --from=go-build /app/bin/stationscraper .
COPY scripts/stationscraper_cron.sh .

ENV SCRAPER_URL="https://www.gov.uk/guidance/access-fuel-price-data"
ENV SCRAPER_INTERVAL=15
ENV POSTGIS_DSN="postgres://localhost:5432/fuelfinder?sslmode=disable"
ENV DEBUG_MODE=false
ENV IMMEDIATE=true

CMD ["/app/stationscraper"]