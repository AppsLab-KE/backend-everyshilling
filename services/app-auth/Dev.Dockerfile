# # Initial stage: download modules
# FROM golang:1.19-alpine as golang-builder

# RUN apk add  --update build-base git ca-certificates
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;
# FROM golang-builder AS app-builder
# # Specify working directory
# WORKDIR /usr/src/app
# # Copy go mod files
# COPY  go.mod go.sum ./
# # Pull packages
# RUN go mod download && go mod verify
# # Copy all files
# COPY . .
# # Create binary
# # RUN mkdir -p "build/dev" && go build -o build/dev/auth-app

# RUN go build -o dev/auth-app
# FROM app-builder as entrypoint
# ENTRYPOINT ["/usr/src/app/build/dev/auth-app"]


#Building OCI-compliant image to keep the container image small
FROM golang:1.19.0-alpine3.16 as builder
WORKDIR /usr/src/app

# Download module in a separate layer to allow caching for the Docker build
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# Copy source code
COPY . .
# COPY cmd ./cmd
# COPY internal ./internal

RUN CGO_ENABLED=0 go build -o dev/app-auth

FROM alpine:3.16.2
WORKDIR /bin
COPY --from=builder /usr/src/app/build/dev/app-auth /bin/app-auth
ENV GIN_MODE=release
CMD /bin/service/app-auth
