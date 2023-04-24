# Initial stage: download modules
FROM golang:1.18-alpine as golang-builder

RUN apk add build-base
RUN apk --update add git ca-certificates
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder
WORKDIR /app/auth

# Copy go mod files
COPY go.mod go.sum \
     /app/auth/

RUN go mod download

COPY .  /app/auth


RUN go build -o /tmp/app-auth

FROM app-builder AS prepare-bin

COPY --from=app-builder /tmp/app-auth /usr/bin/auth-service

ENTRYPOINT ["/usr/bin/auth-service"]