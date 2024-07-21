FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /docker-gs-ping  ./cmd/main.go 

CMD ["/docker-gs-ping"]