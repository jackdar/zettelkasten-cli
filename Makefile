build:
	go build -o bin/zk main.go

run:
	./bin/zk

clean:
	rm -rf bin/*
