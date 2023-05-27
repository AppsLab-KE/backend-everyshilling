# Initial stage: download modules
FROM golang:1.18-alpine as golang-builder

RUN apk add build-base
RUN apk --update add git ca-certificates
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder
WORKDIR /app/db

# Copy go mod files
COPY go.mod go.sum \
     /app/db/

RUN go mod download

COPY .  /app/db


RUN go build -o /tmp/app-db

FROM app-builder AS prepare-bin

COPY --from=app-builder /tmp/app-db /usr/bin/database-service

# Create the target folder and copy the contents

ENTRYPOINT ["/usr/bin/database-service"]