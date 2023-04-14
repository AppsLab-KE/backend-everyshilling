# Initial stage: download modules
FROM golang:1.18-alpine as golang-builder

RUN apk add build-base poppler-utils
RUN apk --update add git ca-certificates
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder


WORKDIR /app
# Copy go mod files
COPY go.mod go.sum \
     /app/db/

RUN go mod download

COPY .  /app/db


RUN go build -o app-scrapper-bin /app/services/app-scrapper/cmd/main/main.go

FROM app-builder AS prepare-bin

COPY --from=app-builder /app/app-scrapper-bin /usr/bin/scrapping-service

ENTRYPOINT ["/usr/bin/scrapping-service"]