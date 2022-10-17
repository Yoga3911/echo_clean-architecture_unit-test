# GO IMAGE
FROM golang:1.19.2-alpine3.16

WORKDIR /app

COPY . .

RUN go build -o main.app .

RUN cp main.app /

WORKDIR /

RUN rm -rf ./app

# ENV JWT_KEY=cobasaja123
# ENV DB_HOST=host.docker.internal
# ENV DB_USERNAME=root
# ENV DB_PASSWORD=123456
# ENV DB_NAME=gorm_crud
# ENV DB_PORT=3307

CMD ["./main.app"]