FROM golang:1.25-alpine3.22
LABEL authors="forkbomb.net"
WORKDIR /heracles-api
COPY . ./
RUN go mod download

RUN go build -o heracles-api .

EXPOSE 8080

CMD ["/heracles-api/heracles-api"]
