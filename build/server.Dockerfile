FROM golang:alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/fuelfinderserver cmd/fuelfinderserver/main.go

FROM alpine:latest AS fuelfinder-fueldataserver
WORKDIR /app
COPY --from=go-build /app/bin/fuelfinderserver .

ENV PORT=50051
ENV DEBUG_MODE=false
ENV MONGO_URI="mongodb://database"

CMD ["/app/fuelfinderserver"]