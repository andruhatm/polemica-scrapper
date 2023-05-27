FROM golang:1.19.7-alpine AS build

RUN apk add --no-cache git

WORKDIR /polemica-scrapper

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app .

EXPOSE 8081

CMD ["./out/app"]