FROM golang:alpine as base-builder

LABEL maintainer='@ctrose17 <>'

WORKDIR /app

EXPOSE 8081

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build

FROM alpine:latest

WORKDIR /app

COPY --from=base-builder /app/SimFBA .

COPY --from=base-builder /app/data /app/data



ENV PORT 8081
ENV ROOT=/app
ENV GOPATH /go
EXPOSE 8081
EXPOSE 443

CMD ["./SimFBA"]