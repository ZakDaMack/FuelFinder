FROM golang:alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/apigateway cmd/apigateway/main.go

FROM alpine:latest AS fuelfinder-apigateway
WORKDIR /app
COPY --from=go-build /app/bin/apigateway .

ENV GRPC_HOST="fueldataserver:50051"
ENV PORT=8080
ENV DEBUG_MODE=false

EXPOSE 8080
CMD ["/app/apigateway"]