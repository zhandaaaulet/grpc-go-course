FROM golang

ADD . /go/src/

WORKDIR /app

COPY . /app

RUN go build -o app ./calculatorEndterm/calculator_server

EXPOSE 50051

ENTRYPOINT  /app/app