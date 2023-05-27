# Initial stage: download modules
FROM ubuntu:latest as golang-builder


RUN apt-get update && apt-get install -y git cron golang
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64;


FROM golang-builder AS app-builder
WORKDIR /app/ratescron

# Copy go mod files
COPY go.mod go.sum \
     /app/ratescron/

RUN go mod download

COPY .  /app/ratescron


RUN go build -o /tmp/app-ratescron

FROM app-builder AS prepare-bin

COPY --from=app-builder /tmp/app-ratescron /usr/bin/ratescron-service

RUN ls /etc/cron.d/
# Create a cron job file
RUN echo "0 0 * * * /usr/bin/ratescron-service >> /var/log/cron.log 2>&1" > /etc/cron.d/rates-cron-job

# Give execution rights to the cron job
RUN chmod 0644 /etc/cron.d/rates-cron-job

# Apply the cron job
RUN crontab /etc/cron.d/rates-cron-job

RUN touch /var/log/cron.log 
# Run the command to start the cron service
CMD ["cron", "-f", "-L", "/var/log/cron.log"]