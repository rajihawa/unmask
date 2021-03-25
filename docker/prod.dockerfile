FROM golang:1.16.2-alpine3.13

RUN apk update && apk add build-base

ARG PORT
ENV PORT=$PORT

WORKDIR /app

COPY . .

RUN make build

RUN RUN go get -u -v github.com/dgrijalva/jwt-go

EXPOSE $PORT

CMD ["make", "run"]