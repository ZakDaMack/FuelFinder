FROM golang:alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/stationapi cmd/stationapi/main.go

FROM alpine:latest AS fuelfinder-stationapi
WORKDIR /app
COPY --from=go-build /app/bin/stationapi .

ENV PORT=8080
ENV DEBUG_MODE=false
ENV POSTGIS_DSN="postgres://database:5432/fuelfinder?sslmode=disable"

CMD ["/app/stationapi"]