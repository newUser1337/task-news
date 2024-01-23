up:
	docker-compose up -d 

down:
	docker-compose down 

dev_build:
	go build -o ./build/news ./cmd/main.go 
	docker build --tag news-go . 

clean:
	rm ./build/news