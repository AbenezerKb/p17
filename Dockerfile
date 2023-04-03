FROM golang:alpine AS builder
WORKDIR /
ADD . .
RUN go build -o bin/sms-gateway /cmd/rest/main.go

FROM alpine
WORKDIR /
COPY --from=builder /bin/sms-gateway .


ENTRYPOINT [ "./sms-gateway" ]
