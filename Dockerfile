FROM  golang:1.20-alpine3.18
RUN apk update && apk add --no-cache gcc musl-dev
WORKDIR /myapp
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app
EXPOSE 8080
CMD ["./app"]