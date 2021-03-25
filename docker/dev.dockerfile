FROM golang:1.16.2-alpine3.13

RUN apk add --update tzdata \
    git;

# Install fresh that brings reload functionality
RUN go get github.com/pilu/fresh
RUN go get -u -v github.com/dgrijalva/jwt-go
RUN go get -v gopkg.in/rethinkdb/rethinkdb-go.v6

ARG PORT
ENV PORT=$PORT
EXPOSE $PORT

CMD ["fresh"]