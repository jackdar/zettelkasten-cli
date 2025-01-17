run:
	./bin/zk

build:
	go build -o bin/zk main.go

clean:
	rm -rf bin/*
