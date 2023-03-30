# Initial stage: download modules
FROM golang:1.19-alpine as golang-builder

RUN apk add  --update build-base git ca-certificates
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder

# Specify working directory
WORKDIR /usr/src/app

# Copy go mod files
COPY  go.mod go.sum ./

# Pull packages
RUN go mod download && go mod verify

# Copy all files
COPY . .

# Create binary
# RUN mkdir -p "build/dev" && go build -o build/dev/auth-app

RUN go build -o dev/auth-app

FROM app-builder as entrypoint

ENTRYPOINT ["/usr/src/app/build/dev/auth-app"]
