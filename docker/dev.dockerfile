FROM golang:1.15.7-alpine3.13

RUN apk add --update tzdata \
    git;

# Install fresh that brings reload functionality
RUN go get github.com/pilu/fresh

ARG PORT
ENV PORT=$PORT
EXPOSE $PORT

CMD ["fresh"]