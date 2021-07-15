FROM golang:1.16-alpine AS builder
USER root
WORKDIR /builder
ADD . /builder
RUN go build -o orders ./cmd/main.go

FROM alpine:3.14.0
ARG AZ_COSMOS_ORDER_CONNECTION_STRING
ENV AZ_COSMOS_ORDER_CONNECTION_STRING=${AZ_COSMOS_ORDER_CONNECTION_STRING}
WORKDIR /app
EXPOSE 8081

COPY --from=builder /builder/orders .
ENTRYPOINT ["/app/orders"]
