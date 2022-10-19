# GO IMAGE
FROM golang:1.19.2 as app

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app .

FROM alpine:latest

COPY --from=app /app/main.app .

CMD ["/main.app"]