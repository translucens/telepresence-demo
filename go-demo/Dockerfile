FROM golang:1.12.7-buster AS builder

ENV GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . ./
RUN go vet
RUN go build -o demoapp

FROM debian:buster-slim

COPY --from=builder /build/demoapp /demoapp
ENV PORT=8080
EXPOSE ${PORT}

CMD ["/demoapp"]
