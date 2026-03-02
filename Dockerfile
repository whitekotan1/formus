FROM golang:1.24 as build_stage

WORKDIR /app

COPY go.mod ./
COPY go.sum ./


RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/formus ./cmd/


FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=build_stage /app/formus .
COPY static/ ./static/

EXPOSE 5000

CMD ["./formus"]