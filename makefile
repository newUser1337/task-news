all:
	go build -o ./build/news ./cmd/main.go 

clean:
	rm ./build/news