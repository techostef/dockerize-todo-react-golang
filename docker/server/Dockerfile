FROM golang:1.21.3

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./server/go.mod .
COPY ./server/go.sum .
RUN go mod download && go mod verify

COPY ./server/ .
RUN go build -o app .

# Expose the port on which the application will run
EXPOSE 8080

CMD ["/app/app"]