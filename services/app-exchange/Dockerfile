# Initial stage: download modules
FROM golang:1.18-alpine as golang-builder

RUN apk add build-base
RUN apk --update add git ca-certificates
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder
WORKDIR /app/exchange

# Copy go mod files
COPY go.mod go.sum \
     /app/exchange/

RUN go mod download

COPY .  /app/exchange


RUN go build -o /tmp/app-exchange

FROM app-builder AS prepare-bin

COPY --from=app-builder /tmp/app-exchange /usr/bin/exchange-service

# Create the target folder and copy the contents

ENTRYPOINT ["/usr/bin/exchange-service"]