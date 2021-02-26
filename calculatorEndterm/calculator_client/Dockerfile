FROM golang

ADD . /go/src/

WORKDIR /app

COPY . /app

RUN go build -o app ./

EXPOSE 8080

ENTRYPOINT  /app/app