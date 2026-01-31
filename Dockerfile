FROM golang:1.25-alpine3.22
LABEL authors="forkbomb.net"
WORKDIR /heracles-api
COPY . ./
RUN go mod download

EXPOSE 8080

RUN go build -o heracles-api .

CMD ["/heracles-api/heracles-api"]
