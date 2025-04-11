FROM golang:1.20-alpine

RUN apk add --no-cache \
    bash \
    curl \
    git \
    gcc \
    g++ \

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/app/main.go

EXPOSE 5432

CMD ["./app"]

