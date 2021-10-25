FROM golang:1.16-alpine
ENV DB_PATH=.
EXPOSE 8004
WORKDIR /app
COPY . .
WORKDIR /app/cmd/lbc-test
RUN go build -o lbc-test main.go
ENTRYPOINT DB_PATH=$(DB_PATH) ./lbc-test 