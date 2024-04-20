FROM golang:alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/fueldataserver cmd/fueldataserver/main.go

FROM alpine:latest AS ofd-fueldataserver
WORKDIR /app
COPY --from=go-build /app/bin/fueldataserver .
ENV PORT=50051
ENV MONGO_URI="mongodb://database"
CMD ["/app/fueldataserver"]