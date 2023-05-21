# Initial stage: download modules
FROM golang:1.18-alpine as golang-builder

RUN apk add build-base openssl
RUN apk --update add git ca-certificates
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder
WORKDIR /app/auth

# # Copy go mod files
# COPY go.mod go.sum \
#      /app/auth/

COPY  . /app/auth/

# RUN go mod download


RUN go build -o /tmp/app-auth

# Generate private and public keys
RUN mkdir -p /etc/auth-service

RUN if [ ! -e "/etc/auth-service/public.pem" ]; then \
       openssl genrsa -out /etc/auth-service/private.pem 2048; \
       openssl rsa -in /etc/auth-service/private.pem -pubout -out /etc/auth-service/public.pem;  \
    fi;

FROM app-builder AS prepare-bin

COPY --from=app-builder /tmp/app-auth /usr/bin/auth-service

ENTRYPOINT ["/usr/bin/auth-service"]

