## Multistage build to reduce the image size

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
# ENTRYPOINT ["/usr/src/app/build/dev/auth-app"]
ENTRYPOINT ["/app-auth"]
