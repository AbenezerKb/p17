FROM golang:alpine AS builder
WORKDIR /
ADD . .
RUN go build -o bin/sewasew /cmd/rest/main.go

FROM alpine
WORKDIR /
COPY --from=builder /bin/sewasew .


ENTRYPOINT [ "./sewasew" ]
