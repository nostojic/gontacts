.PHONY: up down

up:
	docker-compose up -d

down:
	docker-compose down

start: up
	CompileDaemon -command="./gontacts"

server:
	CompileDaemon -command="./gontacts"