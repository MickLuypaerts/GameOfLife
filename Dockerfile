FROM golang:latest as builder
# GO111MODULE=on
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/MickLuypaerts/GameOfLife
COPY /src/ .

RUN go build -ldflags="-w -s" -o /go/bin/GameOfLife

FROM scratch
EXPOSE 8080
WORKDIR /app/
COPY --from=builder /go/bin/GameOfLife .
COPY --from=builder /go/src/github.com/MickLuypaerts/GameOfLife/static /app/static
ENTRYPOINT ["./GameOfLife" ]