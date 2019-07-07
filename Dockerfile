FROM golang:1.12.6 AS builder
WORKDIR /go/src/github.com/mmcken3/wait-for-psql
COPY . .
RUN CGO_ENABLED=0 go install ./...

FROM alpine
WORKDIR /usr/bin
COPY --from=builder /go/bin/wait-for-psql /usr/bin/wait-for-psql
ENTRYPOINT ["wait-for-psql"]