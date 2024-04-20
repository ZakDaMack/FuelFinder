FROM golang:alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/stationscraper cmd/stationscraper/main.go

FROM alpine:latest AS ofd-stationscraper
WORKDIR /app
COPY --from=go-build /app/bin/stationscraper .
COPY scripts/stationscraper_cron.sh .
ENV GRPC_HOST="fueldataserver:50051"
ENV INTERVAL=15
CMD ["/app/stationscraper"]