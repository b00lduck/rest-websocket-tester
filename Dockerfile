FROM golang:alpine AS builder
RUN apk add --no-cache git gcc musl-dev
ADD . /src
WORKDIR /src
RUN go get ./... \
 && go vet ./... \
 && go test ./...\
 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app cmd/app/main.go

FROM alpine
RUN chmod 777 /tmp
COPY --from=builder /app /app
USER 100:100
ENTRYPOINT ["/app"]
