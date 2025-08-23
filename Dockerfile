FROM golang:1.25-alpine3.22
LABEL authors="forkbomb.net"
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o /heracles-api

EXPOSE 8080

CMD ["/heracles-api"]
