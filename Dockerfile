# build databse and setup index
FROM mongo:latest as ofd-database
WORKDIR /app
COPY scripts/mongo_setup ./setup
RUN mongosh < setup

# build stage
FROM node:lts-alpine as web-build
WORKDIR /app
COPY web/package*.json ./
RUN npm install
COPY web .
RUN export REACT_APP_API_URL="/api/"
RUN npm run build

# production stage
# this ensures that only the relevant files are copied over
FROM nginx:stable-alpine as ofd-web
COPY --from=web-build /app/build /usr/share/nginx/html
EXPOSE 80
HEALTHCHECK CMD curl --fail http://localhost || exit 1
CMD ["nginx", "-g", "daemon off;"]

FROM golang:alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/apigateway cmd/apigateway/main.go
RUN go build -o bin/fueldataserver cmd/fueldataserver/main.go
RUN go build -o bin/stationscraper cmd/stationscraper/main.go

FROM alpine:latest AS ofd-fueldataserver
WORKDIR /app
COPY --from=go-build /app/bin/fueldataserver .
ENV PORT=50051
ENV MONGO_URI="mongodb://database"
CMD ["/app/fueldataserver"]

FROM alpine:latest AS ofd-apigateway
WORKDIR /app
COPY --from=go-build /app/bin/apigateway .
ENV GRPC_HOST="fueldataserver:50051"
ENV PORT=8080
EXPOSE 8080
CMD ["/app/apigateway"]

FROM alpine:latest AS ofd-stationscraper
WORKDIR /app
COPY --from=go-build /app/bin/stationscraper .
COPY scripts/stationscraper_cron.sh .
ENV GRPC_HOST="fueldataserver:50051"
CMD crontab /app/stationscraper_cron.sh && exec crond -f