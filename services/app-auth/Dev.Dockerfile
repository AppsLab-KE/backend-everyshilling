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


# Stage 1: Build the application
FROM golang:1.19-alpine  AS builder
WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app-auth .
RUN go build -o dev/auth-app


# Stage 2: Run the application
FROM gcr.io/distroless/base-debian10

#path to the file to be copied
COPY --from=builder /usr/src/app .
ENTRYPOINT ["/app-auth"]
