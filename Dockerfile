FROM golang:1.20-alpine3.17 as builder

ENV GO111MODULE on
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app
COPY go.mod .
COPY go.sum .

COPY . ./

RUN go build -o go-hexagonal

FROM alpine:3.17 as release
WORKDIR /app

COPY --from=builder /app/go-hexagonal /app/cmd/
COPY --from=builder /app/config /app/config
RUN chmod +x /app/cmd/go-hexagonal

CMD ["cmd/go-hexagonal"]