# Initial stage: download modules
FROM golang:1.18-alpine as golang-builder

RUN apk add  --update build-base git ca-certificates
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder

# Specify working directory
WORKDIR /usr/src/app

# Copy go mod files
COPY  go.mod go.sum ./

# Pull packages
RUN go mod download && go mod verify

# Create binary
RUN mkdir -p "build/prod" && go build -o build/prod/auth-app

FROM app-builder as entrypoint

ENTRYPOINT ["/usr/src/app/build/prod/auth-app"]
