FROM golang:1.14
RUN go get github.com/google/wire/cmd/wire
WORKDIR /src
COPY . ./
RUN make

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /srv
COPY --from=0 /src/query .
CMD ["./query"]
