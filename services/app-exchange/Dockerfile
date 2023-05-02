FROM golang:1.17.2-alphine
#setting the working directory and coping source code
WORKDIR /app
COPY . /app/

# building starting the application
RUN go build -o app
CMD ["./app"]

