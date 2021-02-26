FROM golang

ADD . /go/src/

WORKDIR /app

COPY . /app

RUN go build -o app ./calculatorEndterm/calculator_client

EXPOSE 8080

ENTRYPOINT  /app/app