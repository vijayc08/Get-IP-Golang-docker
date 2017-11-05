FROM golang:1.7

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./myip.go

CMD ["./myip"]

EXPOSE 9999
