FROM golang:1.16.2-alpine3.13

RUN apk add --update tzdata \
    git;


RUN go get -u -v github.com/dgrijalva/jwt-go
RUN go get -u -v gopkg.in/rethinkdb/rethinkdb-go.v6
RUN go get -v github.com/rs/xid
RUN go get -u -v golang.org/x/crypto/bcrypt
RUN go get -u -v github.com/gorilla/handlers
RUN go get -v github.com/go-playground/validator

CMD ["go", "test", "./tests"]