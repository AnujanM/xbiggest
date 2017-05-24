all: build

build:
	go build -o xbiggest main.go

lint:
	${GOPATH}/bin/golint .

clean:
	rm xbiggest