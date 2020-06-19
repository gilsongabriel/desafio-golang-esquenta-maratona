FROM golang:1.14.4-alpine3.12 as builder
MAINTAINER "Gilson Gabriel <gilson@codemastersolucoes.com>"

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-w -s" -installsuffix cgo -o main .

CMD ["./main"]