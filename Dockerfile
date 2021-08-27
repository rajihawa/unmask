# Multi-Stage build for small and secure image
# Step 1: build executable binary
FROM golang:1.16.5-alpine3.13 As builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go get -d -v

RUN go build -o /go/bin/mask-off

# Step 2: run from scratch
FROM scratch

COPY --from=builder /go/bin/mask-off /go/bin/mask-off

ENTRYPOINT [ "/go/bin/mask-off" ]