FROM golang:1.25-alpine3.22
LABEL authors="forkbomb.net"
ARG port
WORKDIR /heracles-api
COPY . ./
RUN go mod download

ENV PORT=$port

EXPOSE $PORT

RUN go build -o heracles-api .

CMD ["/heracles-api/heracles-api"]
