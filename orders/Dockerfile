FROM golang:1.16-alpine AS builder
USER root
WORKDIR /builder
ADD . /builder
RUN go build -o orders .

FROM alpine:3.14.0
WORKDIR /app
EXPOSE 8081
COPY --from=builder /builder/orders .
ENTRYPOINT ["/app/orders"]
