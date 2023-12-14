setup:
	cp .env.example .env

build: setup
	go build -o ./out/blog

run-server: build
	./out/blog server
