FROM golang:latest
WORKDIR /api

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8080

RUN go build -o main .
CMD ./main