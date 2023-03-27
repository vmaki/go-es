
swag:
	swag init

mUp:
	go run main.go migrate up

docker-build:
	docker build -t go-es .
docker-env:
	docker-compose -f docker-compose-env.yml up -d
docker-prod:
	docker-compose up -d
