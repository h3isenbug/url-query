all: build

build:
	wire /src/cmd/url-query
	go build -o query /src/cmd/url-query/
