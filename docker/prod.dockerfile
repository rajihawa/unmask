FROM golang:1.15.7-alpine3.13

RUN apk update && apk add build-base

ARG PORT
ENV PORT=$PORT

WORKDIR /app

COPY . .

RUN make build

EXPOSE $PORT

CMD ["make", "run"]