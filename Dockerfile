FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /go-web-app

EXPOSE 9000

CMD [ "/go-web-app" ]