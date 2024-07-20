FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN go build -o /docker-gs-ping  ./cmd/main.go 

CMD ["/docker-gs-ping"]