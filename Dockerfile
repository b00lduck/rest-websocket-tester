FROM golang:alpine AS builder
RUN apk add --no-cache git gcc musl-dev
ADD . /src
WORKDIR /src
RUN go get ./... \
 && go vet ./... \
 && go test ./...\
 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app cmd/app/main.go

FROM scratch
USER 100:100
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
